package server

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type Person struct {
	Name string
	Phone string
}

func GetSession() *mgo.Session {
	fmt.Println("Connecting MongoDB")
	session, err := mgo.Dial("mongodb://habib:habib786@ds135624.mlab.com:35624/payroll_api")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}