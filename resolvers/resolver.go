package resolvers

import (
	"example/Go/internal/core"
	"log"
)

type Resolver struct {
	// user factory
	UserFactory core.UserFactory
	// Array of Users
	Users []core.UserProxy
}

func ResolverInstantiation() *Resolver {
	// Instantiate a User Factory
	wf, err := core.NewUserFactory()

	if err != nil {
		log.Fatalf("Error when instantiating a user factory: %s", err)
	}

	resolver := &Resolver{
		UserFactory: *wf,
	}

	return resolver
}
