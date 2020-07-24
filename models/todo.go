package models

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

// Todo 的结构定义
type Todo struct {
	ID         bson.ObjectId `bson:"_id"`
	Content    string        `bson:"content"`
	Owner      string        `bson:"owner"`
	Completion string        `bson:"completion"`
}

// GetTodoByUserName 获取用户的 todos
func GetTodoByUserName(userName string) (todos []Todo) {
	err := CollectionTodo.Find(bson.M{"owner": userName}).All(&todos)
	if err != nil {
		fmt.Println(err)
		return
	}
	return todos
}

// GetTodoByTodoStatus 根据完成情况获取 todos
func GetTodoByTodoStatus(userName, completion string) (todos []Todo) {
	err := CollectionTodo.Find(bson.M{"owner": userName, "completion": completion}).All(&todos)
	if err != nil {
		fmt.Println(err)
		return
	}
	return todos
}

// CreateTodo 新建 todo
func CreateTodo(todo Todo) bool {
	err := CollectionTodo.Insert(&todo)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// DeleteTodo 删除 todo
func DeleteTodo(objectID string) bool {
	err := CollectionTodo.Remove(bson.M{"_id": bson.ObjectIdHex(objectID)})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// DeleteTodosByUser 清空用户所有 todo
func DeleteTodosByUser(userName string) bool {
	_, err := CollectionTodo.RemoveAll(bson.M{"owner": userName, "completion": true})
	if err != nil {
		println(err)
		return false
	}
	return true
}

// CompleteAllTodos 将当前所有 todo 设为已完成
func CompleteAllTodos(userName string) bool {
	selector := bson.M{"owner": userName, "completion": "false"}
	data := bson.M{"$set": bson.M{"completion": "true"}}
	changeInfo, err := CollectionTodo.UpdateAll(selector, data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Printf("%+v\n", changeInfo)
	return true
}
