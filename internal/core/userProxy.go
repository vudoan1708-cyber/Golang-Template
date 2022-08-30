package core

import (
	"example/Go/models"
	"log"
)

type UserProxy interface {
	Get() (models.User, error)
	Update(args models.User) (models.User, error)
}

// This is a private struct, its information can only be retrieved via interacting with the above interface
type userProxyBasic struct {
	_id     string
	name    string
	age     uint8
	address string
}

func NewUserProxy(_id string, name string, age uint8, address string) (UserProxy, error) {
	up := userProxyBasic{_id, name, age, address}

	return &up, nil
}

// Implementation
func (up *userProxyBasic) Get() (models.User, error) {
	log.Println("userProxyBasic.Get() called")

	mUser := models.User{
		ID:      up._id,
		Name:    up.name,
		Age:     up.age,
		Address: up.address,
	}

	return mUser, nil
}

func (up *userProxyBasic) Update(args models.User) (models.User, error) {
	var mUser models.User = models.User{
		ID:      up._id,
		Name:    up.name,
		Age:     up.age,
		Address: up.address,
	}
	if args.ID != "" {
		mUser.ID = args.ID
	}
	if args.Name != "" {
		mUser.Name = args.Name
	}
	if args.Age > 0 {
		mUser.Age = args.Age
	}
	if args.Address != "" {
		mUser.Address = args.Address
	}

	return mUser, nil
}
