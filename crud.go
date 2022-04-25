package main

import (
	"context"
	"fmt"
	"test/ent"
	"test/ent/user"
)

func CRUD01(ctx context.Context, client *ent.Client) error {
	aaa, err := client.User.
		Create().
		SetName("aaa").
		SetAge(20).Save(ctx)
	if err != nil {
		return fmt.Errorf("create user : %w", err)
	}
	fmt.Println("user1: ", aaa)

	bbb, err := client.User.
		Create().
		SetName("bbb").
		SetAge(20).Save(ctx)

	if err != nil {
		return fmt.Errorf("create user : %w", err)
	}

	fmt.Println("user2: ", bbb)

	dog := client.Pet.
		Create().
		SetName("dog").
		SetOwner(aaa).
		SaveX(ctx)
	fmt.Println("pet: ", dog)

	group := client.Group.
		Create().
		SetName("Group1").
		AddUsers(aaa, bbb).
		SaveX(ctx)

	fmt.Println("group: ", group)

	return nil
}

func CRUD02(ctx context.Context, client *ent.Client) error {
	ccc := client.User.
		Create().
		SetName("ccc").
		SetAge(20).SaveX(ctx)

	fmt.Println("ccc: ", ccc)

	n := client.User.
		UpdateOneID(ccc.ID).
		SetAge(25).
		SaveX(ctx)

	fmt.Println("n: ", n)

	nn := client.User.
		Update().
		Where(
			user.Name("ccc"),
		).
		SetAge(30).SaveX(ctx)

	fmt.Println("nn: ", nn)

	return nil
}

func CRUD03(ctx context.Context, client *ent.Client) error {
	users := client.User.
		Query(). //Query All User
		AllX(ctx)

	fmt.Println("users: ")
	for _, user := range users {
		fmt.Println(user)
	}

	aaa := client.User.
		Query().
		Where(
			user.Name("aaa"), //Query User Name aaa
		).OnlyX(ctx) //panic if more than 1 user returned

	fmt.Println("aaa: ", aaa)

	pet := client.User.
		QueryPets(aaa). //Query User aaa's Pet
		OnlyX(ctx)

	fmt.Println("pet: ", pet)
	fmt.Println("owner: ", pet.Edges.Owner)

	pet2 := client.User.
		QueryPets(aaa).
		WithOwner(). //Query User aaa's Pet with her's Owner
		OnlyX(ctx)

	fmt.Println("pet: ", pet2)
	fmt.Println("owner: ", pet2.Edges.Owner)

	return nil
}

func CRUD04(ctx context.Context, client *ent.Client) error {

	ddd, err := client.User.
		Create().
		SetName("ddd").
		SetAge(40).Save(ctx)
	if err != nil {
		return fmt.Errorf("create user : %w", err)
	}

	users := client.User.
		Query(). //Query All User
		AllX(ctx)

	fmt.Println("users: ")
	for _, user := range users {
		fmt.Println(user)
	}

	client.User.
		DeleteOne(ddd).
		ExecX(ctx)

	users = client.User.
		Query(). //Query All User
		AllX(ctx)

	fmt.Println("users: ")
	for _, user := range users {
		fmt.Println(user)
	}
	return nil
}

func CRUD05(ctx context.Context, client *ent.Client) error {

	users := client.User.
		Query(). //Query All User
		AllX(ctx)

	fmt.Println("users: ")
	for _, user := range users {
		fmt.Println(user)
	}

	client.User.
		Delete().
		Where(
			user.AgeGTE(30), //Delete user who has age greater or equa 30
		).
		ExecX(ctx)

	users = client.User.
		Query(). //Query All User
		AllX(ctx)

	fmt.Println("users: ")
	for _, user := range users {
		fmt.Println(user)
	}
	return nil
}
