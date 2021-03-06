package service

import (
	"crudexample/db"
	"crudexample/models/entity"
	"errors"
	"github.com/goonode/mogo"
	"gopkg.in/mgo.v2/bson"
)

//Userservice is to handle user relation db query
type Userservice struct{}

//Create is to register new user
func (userservice Userservice) Create(user *(entity.User)) error {
	conn := db.GetConnection()
	defer conn.Session.Close()

	doc := mogo.NewDoc(entity.User{}).(*(entity.User))
	err := doc.FindOne(bson.M{"email": user.Email}, doc)
	if err == nil {
		return errors.New("already exist")
	}
	userModel := mogo.NewDoc(user).(*(entity.User))
	err = mogo.Save(userModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return vErr
	}
	return err
}

// Delete a user from DB
func (userservice Userservice) Delete(email string) error {
	user, _ := userservice.FindByEmail(email)
	conn := db.GetConnection()
	defer conn.Session.Close()
	err := user.Remove()
	return err
}

//Find user
func (userservice Userservice) Find(user *(entity.User)) (*entity.User, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	doc := mogo.NewDoc(entity.User{}).(*(entity.User))
	err := doc.FindOne(bson.M{"email": user.Email}, doc)

	if err != nil {
		return nil, err
	}
	return doc, nil
}

//Find user from email
func (userservice Userservice) FindByEmail(email string) (*entity.User, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	user := new(entity.User)
	user.Email = email
	return userservice.Find(user)
}
