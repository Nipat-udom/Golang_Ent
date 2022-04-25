package main

import (
	"context"
	"fmt"
	"test/ent"
)

func Do3(ctx context.Context, client *ent.Client) error {
	hub := client.Group.
		Create().
		SetName("GitHub").
		SaveX(ctx)
	lab := client.Group.
		Create().
		SetName("GitLab").
		SaveX(ctx)
	prayad := client.User.
		Create().
		SetAge(30).
		SetName("Prayad").
		AddGroups(hub, lab).
		SaveX(ctx)
	prawet := client.User.
		Create().
		SetAge(28).
		SetName("Prawet").
		AddGroups(hub).
		SaveX(ctx)

	// Query the edges.
	groups, err := prayad.
		QueryGroups().
		All(ctx)
	if err != nil {
		return fmt.Errorf("querying a8m groups: %w", err)
	}
	fmt.Println(groups)

	groups, err = prawet.
		QueryGroups().
		All(ctx)
	if err != nil {
		return fmt.Errorf("querying nati groups: %w", err)
	}
	fmt.Println(groups)

	return nil
}
