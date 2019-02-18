package repository

import model "github.com/anthonydenecheau/gopocservice/breeder"

type BreederRepository interface {
	GetById(id int64) (*model.Breeder, error)
}
