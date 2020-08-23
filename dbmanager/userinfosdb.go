package dbmanager

import (
	"context"
	"sync"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"rabie.com/portfolio/createjwt"
	"rabie.com/portfolio/errorrepo"
	"rabie.com/portfolio/hash"
	"rabie.com/portfolio/models"
)

var wg sync.WaitGroup

// RegisterUserdb use this func to register user in db
func RegisterUserdb(u models.User) (bool, models.Userprofile) {
	db, err := Connectdb()
	errorrepo.DieIf(err)
	//user profile
	up := &models.Userprofile{Email: u.Email}

	// check if user exist

	check := checkifuserexist(u.Email, db)

	if check == false {
		// hash PASSWORD sha256
		u.Password = hash.Hashpassword(u.Password)
		// exists, err := models.UserExists(context.Background(), db, u.)
		u.Status = true
		err = u.Insert(context.Background(), db, boil.Infer())
		errorrepo.DieIf(err)
		//use relation to insert in user prfile
		err = u.AddUserprofiles(context.Background(), db, true, up)
		errorrepo.DieIf(err)

		ur := &models.UserRole{RolesID: 2}
		_ = u.AddUserRoles(context.Background(), db, true, ur)

		return check, *up
	}
	return check, *up
}

// LoginExist return a boolean value true if exist and false if not
func LoginExist(Email, Password string) (bool, createjwt.JwtCustomClaims) {
	db, err := Connectdb()
	errorrepo.DieIf(err)
	// retriev all user info from db
	logged, err := models.Users(qm.Where(models.UserColumns.Email+"= ?", Email)).One(context.Background(), db)
	errorrepo.DieIf(err)
	//hash password for db
	Password = hash.Hashpassword(Password)
	claim := new(createjwt.JwtCustomClaims)
	//retriev user role
	role, err := logged.UserRoles().One(context.Background(), db)
	//alluser clailm
	errorrepo.DieIf(err)
	claim.Email = logged.Email
	claim.ID = logged.ID
	claim.RolesID = role.RolesID
	if logged.Password == Password && logged.Status == true {
		return true, *claim
	}
	return false, *claim

}

// Userprofiledb : get user profile frome db
func Userprofiledb(up int) models.Userprofile {
	db, err := Connectdb()
	errorrepo.DieIf(err)
	ul, err := models.FindUser(context.Background(), db, up)
	uf, err := ul.Userprofiles().One(context.Background(), db)
	errorrepo.DieIf(err)
	return *uf
}

// Updateuserprofiledb use this function to update user profile
func Updateuserprofiledb(up models.Userprofile) error {
	db, err := Connectdb()
	errorrepo.DieIf(err)
	uf, err := models.FindUserprofile(context.Background(), db, up.ID)
	uf.Adress = up.Adress
	uf.LastName = up.LastName
	uf.Name = up.Name
	uf.Phone = up.Phone
	uf.Photo = up.Photo
	_, err = uf.Update(context.Background(), db, boil.Infer())
	return err

}
