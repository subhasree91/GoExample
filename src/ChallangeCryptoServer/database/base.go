package database

import (
	conf "ChallangeCryptoServer/config"
	"fmt"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
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

func Get(collection string, findQuery map[string]interface{}) ([]byte, error) {

}
