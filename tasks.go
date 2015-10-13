package main

import (
	"strconv"
)

type User struct{
	Name string
	Tasks [] *Task
}

type Task struct{
	NameTask string
	Description string
}

func (u *User) addTaskToUser(nameTask, descriptionTask string) bool{
	task := &Task{NameTask: nameTask, Description: descriptionTask}

	u.Tasks = append(u.Tasks, task)

	return true
}

func (u *User) getAllUserTasks() [] *Task{
	return u.Tasks
}

func (u *User) removeAllTasks() bool {
	u.Tasks = u.Tasks[:0]

	return true
}

var users [] *User

func getUserByName(name string) *User {
	for _, user := range users {
		if user.Name == name {
			return user
		}
	}

	return nil
}

func createNewUser(name string) *User {
	user := &User{Name: name}
	users = append(users, user)

	return user
}

func CreateNewTask(nameUser, nameTask, descriptionTask string) bool {
	user := getUserByName(nameUser)

	if user == nil {
		user = createNewUser(nameUser)
	}

	user.addTaskToUser(nameTask, descriptionTask)

	return true
}

func GetAllUserTasks(nameUser string) string {
	user := getUserByName(nameUser)

	if user == nil {
		return "User with this name not found"
	}

	tasks := user.getAllUserTasks()

	var message string

	for key, task := range tasks {
		message = message +  strconv.Itoa(key) + ". " + task.NameTask + ": " + task.Description
	}

	return message
}

func RemoveTaskByTaskName(nameUser, nameTask string) bool {
	user := getUserByName(nameUser)

	if user == nil {
		return false
	}

	user.removeAllTasks()

	return true
}