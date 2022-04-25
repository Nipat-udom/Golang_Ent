package main

import (
	"context"
	"fmt"
	"log"
	"test/ent"
)

func Do(ctx context.Context, client *ent.Client) error {
	user, err := client.User.
		Create().
		SetAge(30).
		SetName("Somchai").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	log.Println("user:", user)
	card1, err := client.Card.
		Create().
		SetOwner(user).
		SetNumber("1020").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating card: %w", err)
	}
	log.Println("card:", card1)

	// Only returns the card of the user,
	card2, err := user.QueryCard().Only(ctx)
	if err != nil {
		return fmt.Errorf("querying card: %w", err)
	}
	log.Println("card:", card2)
	// its back-reference.
	owner, err := card2.QueryOwner().Only(ctx)
	if err != nil {
		return fmt.Errorf("querying owner: %w", err)
	}
	log.Println("owner:", owner)
	return nil
}
