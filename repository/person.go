package repository

import (
	"github.com/anthonydenecheau/gopocservice/model"
	"github.com/go-pg/pg"
)

// Our store will have two methods, to add a new bird,
// and to get all existing birds
// Each method returns an error, in case something goes wrong
type Person interface {
	GetPeople() ([]*model.Ws_dog_eleveur, error)
	GetPeopleById(id int64) (*model.Ws_dog_eleveur, error)
}

// The `dbStore` struct will implement the `Store` interface
// It also takes the sql DB connection object, which represents
// the database connection.
type DbPerson struct {
	Db *pg.DB
}

func (pgdb *DbPerson) GetPeopleById(id int64) (p *model.Ws_dog_eleveur, e error) {
	user := model.Ws_dog_eleveur{Id: id}
	error := pgdb.Db.Select(&user)
	return &user, error
}
func (pgdb *DbPerson) GetPeople() ([]*model.Ws_dog_eleveur, error) {
	// Create the data structure that is returned from the function.
	// By default, this will be an empty array of birds
	//people:= []*model.Person{}
	return nil, nil
}

// The store variable is a package level variable that will be available for
// use throughout our application code
var person Person

/*
We will need to call the InitStore method to initialize the store. This will
typically be done at the beginning of our application (in this case, when the server starts up)
This can also be used to set up the store as a mock, which we will be observing
later on
*/
func InitPerson(p Person) {
	person = p
}
func GetPerson() Person {
	return person
}
