package server

import (
	"gopkg.in/mgo.v2"
)

func GetCollection(DbName string) *mgo.Collection  {
	session := GetSession()
	collection := session.DB("payroll_api").C(DbName)
	return collection
}