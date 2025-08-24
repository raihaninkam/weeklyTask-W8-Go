package usermanager

import (
	"errors"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

type UserManager struct {
	users map[int]*User
}

func NewUserManager() *UserManager {
	return &UserManager{
		users: make(map[int]*User),
	}
}

func (um *UserManager) AddUser(id int, name string) error {
	if _, exist := um.users[id]; exist {
		return errors.New(fmt.Sprintf("Error: User dengan ID %d sudah terdaftar", id))
	}
	um.users[id] = &User{
		ID: id,
		Name: name,
	}
	fmt.Printf("User %s (ID: %d) berhasil ditambahkan\n",name, id)
	return nil
}

func (um *UserManager) GetUser(id int) (*User, error){
	user, exists := um.users[id]
	if !exists{
		return nil, errors.New(fmt.Sprintf("Error: User dengan ID %d tidak ditemukan",id))
	}
	return user, nil
}

func (um *UserManager) DisplayUsers(){
	fmt.Println("\nDaftar User:")
	for id, user := range um.users{
		fmt.Printf("ID: %d, Name: %s\n",id, user.Name)
	}
}
