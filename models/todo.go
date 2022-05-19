package models

import "bubble/dao"

//Todo 增删改查
//创建todo
func CreateATodo(todo *dao.TodoList) (err error) {
	err = dao.DB.Create(&todo).Error

	return err
}
func GetAllTodo() (todo []*dao.TodoList, err error) {
	if err = dao.DB.Find(&todo).Error; err != nil {
		return nil, err
	}
	return
}
func GetTodoById(id string) (todo *dao.TodoList, err error) {
	todo = new(dao.TodoList)
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}
func UpdataATodo(todo *dao.TodoList) (err error) {
	err = dao.DB.Save(todo).Error
	return
}
func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&dao.TodoList{}).Error
	return
}
