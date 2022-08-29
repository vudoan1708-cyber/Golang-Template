package core

import (
	"fmt"
	"log"
)

/*
	The Factory does not deal directly with the models, but deals with them via the proxy struct
*/

type UserFactory struct {
	userCount int
}

// NewWristbandFactory creates a new instance of a factory to produce user instances
func NewUserFactory() (*UserFactory, error) {
	log.Println("UserFactory.NewUserFactory called")

	uf := UserFactory{
		userCount: 0,
	}

	return &uf, nil
}

func (uf *UserFactory) NewUser(name string, age uint8, address string) (UserProxy, error) {
	uf.userCount += 1

	up, err := NewUserProxy(
		fmt.Sprint(uf.userCount),
		name,
		age,
		address,
	)

	if err != nil {
		return nil, err
	}

	return up, nil
}
