package models

import (
	"fmt"

	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User 的结构定义
type User struct {
	ID       bson.ObjectId `bson:"_id"`
	UserName string        `bson:"userName"`
	Password string        `bson:"password"`
}

func init() {
	session, _ = mgo.Dial(URL)
	db := session.DB("todo")
	CollectionUser = db.C("user")
}

// CheckUser 登录校验
func CheckUser(username, password string) bool {
	var user User
	err := CollectionUser.Find(bson.M{"userName": username, "password": password}).One(&user)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// CreateUser 注册新用户
func CreateUser(userName, password string) bool {
	var user User
	user.ID = bson.NewObjectId()
	user.UserName = userName
	user.Password = password
	fmt.Println(user)
	err := CollectionUser.Insert(&user)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// IsExistUserByName 校验该用户名是否存在
func IsExistUserByName(username string) bool {
	var user User
	err := CollectionUser.Find(bson.M{"userName": username}).One(&user)
	if err != nil {
		return false
	}
	return true
}

// //管理员获取所有用户
// func GetUserIdByName(user_name string) int {
// 	userss := User{}
// 	db.Select("user_id").Where("user_name=?", user_name).First(&userss)
// 	println("id", userss.User_ID)
// 	return userss.User_ID
// }
