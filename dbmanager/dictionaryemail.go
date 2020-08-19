package dbmanager

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"rabie.com/portfolio/errorrepo"
	"rabie.com/portfolio/models"
)

func checkifuserexist(email string, db *sql.DB) bool {
	exists, err := models.Users(qm.Where(models.UserColumns.Email+"= ?", email)).Exists(context.Background(), db)
	errorrepo.DieIf(err)

	if exists == true {
		return true
	}
	return false
}
