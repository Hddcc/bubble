package models

import "bubble/dao"

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
Todo这个Model的增删改查操作都在这里
*/

// CreateATodo 创建Todo
func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetTodoList() (todolist []*Todo, err error) {

	if err := dao.DB.Find(&todolist).Error; err != nil {
		return nil, err
	}
	return
}

func GetTodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleTodo(id string)(err error){
	err = dao.DB.Where("id=?",id).Delete(&Todo{}).Error
	return
}