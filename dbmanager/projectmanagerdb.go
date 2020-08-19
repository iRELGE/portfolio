package dbmanager

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"rabie.com/portfolio/errorrepo"
	"rabie.com/portfolio/models"
)

// Createproject insert project on db
func Createproject(project models.Project) error {
	db, err := Connectdb()
	errorrepo.DieIf(err)
	err = project.Insert(context.Background(), db, boil.Infer())
	return err

}
