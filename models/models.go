package models

import "gopkg.in/mgo.v2"

const URL = "127.0.0.1:27017"

var CollectionUser *mgo.Collection
var CollectionTodo *mgo.Collection
var session *mgo.Session

func init() {
	session, _ = mgo.Dial(URL)
	db := session.DB("todo")
	CollectionTodo = db.C("todo")
	CollectionUser = db.C("user")
}
