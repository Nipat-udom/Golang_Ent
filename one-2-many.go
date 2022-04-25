package main

import (
	"context"
	"fmt"
	"test/ent"
)

func Do2(ctx context.Context, client *ent.Client) error {
	// Create the 2 pets.
	pedro, err := client.Pet.
		Create().
		SetName("pedro").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating pet: %w", err)
	}
	lola, err := client.Pet.
		Create().
		SetName("lola").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating pet: %w", err)
	}
	// Create the user, and add its pets on the creation.
	user, err := client.User.
		Create().
		SetAge(30).
		SetName("Shinichi").
		AddPets(pedro, lola).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	fmt.Println("User created:", user)
	// Output: User(id=1, age=30, name=a8m)

	// Query the owner. Unlike `Only`, `OnlyX` panics if an error occurs.
	owner := pedro.QueryOwner().OnlyX(ctx)
	fmt.Println(owner.Name)
	// Output: a8m

	pets := pedro.
		QueryOwner(). // Shinichi
		QueryPets().  // pedro, lola
		AllX(ctx)

	fmt.Println(pets)
	return nil
}
