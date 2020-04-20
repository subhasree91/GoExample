package database

import (
	conf "GoWeb1/config"
	"encoding/json"
	"fmt"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetSession() *mgo.Session {
	info := &mgo.DialInfo{
		Addrs:    []string{conf.Config.DB.Address},
		Timeout:  60 * time.Second,
		Database: conf.Config.DB.DBname,
		Username: conf.Config.DB.User,
		Password: conf.Config.DB.Password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		fmt.Println("Session error:", err)
		os.Exit(1)

	}

	session.SetMode(mgo.Monotonic, true)

	return session
}

func Create(collection string, data interface{}) error {

	s := GetSession()
	defer s.Close()
	c := s.DB("goapp").C(collection)
	err := c.Insert(data)
	return err
}

func GetDoc(collection string, findQuery map[string]interface{}) ([]byte, error) {

	var data interface{}

	s := GetSession()
	defer s.Close()
	c := s.DB("goapp").C(collection)
	err := c.Find(findQuery).One(&data)
	if err != nil {
		return nil, err
	}
	return json.Marshal(data)

}

func Get(collection string, findQuery map[string]interface{}) ([]byte, error) {

	var data []interface{}

	s := GetSession()
	defer s.Close()
	c := s.DB("goapp").C(collection)
	err := c.Find(findQuery).All(&data)
	if err != nil {
		return nil, err
	}
	return json.Marshal(data)

}

func Update(collection string, updateQuery map[string]interface{}, findQuery map[string]interface{}) error {

	s := GetSession()
	defer s.Close()
	c := s.DB("goapp").C(collection)
	err := c.Update(findQuery, bson.M{"$set": updateQuery})

	return err
}

// db.createUser(
// 	{ user: "admin",
// 	pwd: "password",

// 	roles:[{role: "userAdminAnyDatabase" , db:"admin"}]})
