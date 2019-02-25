package service

import (
	"fmt"

	"assignment/user/app/model"
	"assignment/user/app/repository"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

)

func GetSession() (*mgo.Session, error) {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")
	// Check if connection error, is mongo running?
	if err != nil {
		return nil, err
	}
	// Deliver session
	return s, nil
}

type DbSession struct {
	mgo *mgo.Session
}

func NewUserDbClient(conn *mgo.Session) repository.UserDB {
	return &DbSession{mgo: conn}
}

// AddUser Add New User In DataBase
func (s *DbSession) AddUser(user *model.UserSignUp) error {
	con := s.mgo.Copy()
	c := con.DB("assignment").C("user")
	userCount, err := c.Find(bson.M{"email": user.Email}).Count()
	if err != nil {
		return err
	}
	if userCount > 0 {
		return fmt.Errorf("resource %s already exists", user.Email)
	}
	// Add an Id
	user.Id = bson.NewObjectId()
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)
	defer con.Close()
	if err := c.Insert(user); err != nil {
		return err
	}
	return nil
}

// Get User login user after verify credential
func (s *DbSession) GetUser(email, password string, user *model.UserSignUp) error {
	// Fetch user
	con := s.mgo.Copy()
	if err := con.DB("assignment").C("user").Find(bson.M{"email": email}).One(user); err != nil {
		return fmt.Errorf("user not exist")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return err
	}

	defer con.Close()
	if err := con.DB("assignment").C("user").Find(bson.M{"email": email, "password": user.Password}).One(user); err != nil {
		return fmt.Errorf("incorrect password")
	}

	return nil
}

func (s *DbSession) UpdateUser(userId string, data *model.UserUpdate) error {
	con := s.mgo.Copy()
	c := con.DB("assignment").C("user")
	change := bson.M{"$set": bson.M{
		"name":        data.Name,
		"phonenumber": data.PhoneNumber,
	}}
	defer con.Close()
	if err := c.Update(bson.M{"_id": bson.ObjectIdHex(userId)}, change); err != nil {
		return err
	}
	if err := con.DB("assignment").C("user").Find(bson.M{"_id": bson.ObjectIdHex(userId)}).One(data); err != nil {
		return err
	}
	return nil
}
// Remove user
func (s *DbSession) DeleteUser(id string) error {
	con := s.mgo.Copy()
	defer con.Close()
	if err := con.DB("assignment").C("user").RemoveId(bson.ObjectIdHex(id)); err != nil {
		return err
	}
	return nil
}
