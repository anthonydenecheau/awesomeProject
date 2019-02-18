package repository

import (
	models "github.com/anthonydenecheau/gopocservice/breeder"
	"github.com/go-pg/pg"
)

type pgBreederRepository struct {
	Db *pg.DB
}

func NewPgBreederRepository(Db *pg.DB) BreederRepository {
	return &pgBreederRepository{Db}
}

func (pgdb *pgBreederRepository) GetById(id int64) (*models.Breeder, error) {

	user := models.Breeder{Id: id}
	error := pgdb.Db.Select(&user)
	/*
		if len(user) > 0 {
			a = list[0]
		} else {
			return nil, models.NOT_FOUND_ERROR
		}
	*/
	return &user, error
}
