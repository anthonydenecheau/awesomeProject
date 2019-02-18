package usecase

import "github.com/anthonydenecheau/gopocservice/breeder"
import "github.com/anthonydenecheau/gopocservice/breeder/repository"

type BreederUsecase interface {
	GetByID(id int64) (*breeder.Breeder, error)
}

type breederUsecase struct {
	breederRepos repository.BreederRepository
}

func NewBreederUsecase(a repository.BreederRepository) BreederUsecase {
	return &breederUsecase{
		breederRepos: a,
	}
}

func (a *breederUsecase) GetByID(id int64) (*breeder.Breeder, error) {

	res, err := a.breederRepos.GetById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
