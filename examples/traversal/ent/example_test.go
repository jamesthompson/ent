// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"

	"github.com/facebookincubator/ent/dialect/sql"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleGroup() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the group's edges.
	u0 := client.User.
		Create().
		SetAge(1).
		SetName("string").
		SaveX(ctx)
	log.Println("user created:", u0)
	u1 := client.User.
		Create().
		SetAge(1).
		SetName("string").
		SaveX(ctx)
	log.Println("user created:", u1)

	// create group vertex with its edges.
	gr := client.Group.
		Create().
		SetName("string").
		AddUsers(u0).
		SetAdmin(u1).
		SaveX(ctx)
	log.Println("group created:", gr)

	// query edges.
	u0, err = gr.QueryUsers().First(ctx)
	if err != nil {
		log.Fatalf("failed querying users: %v", err)
	}
	log.Println("users found:", u0)

	u1, err = gr.QueryAdmin().First(ctx)
	if err != nil {
		log.Fatalf("failed querying admin: %v", err)
	}
	log.Println("admin found:", u1)

	// Output:
}
func ExamplePet() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the pet's edges.
	pe0 := client.Pet.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("pet created:", pe0)

	// create pet vertex with its edges.
	pe := client.Pet.
		Create().
		SetName("string").
		AddFriends(pe0).
		SaveX(ctx)
	log.Println("pet created:", pe)

	// query edges.
	pe0, err = pe.QueryFriends().First(ctx)
	if err != nil {
		log.Fatalf("failed querying friends: %v", err)
	}
	log.Println("friends found:", pe0)

	// Output:
}
func ExampleUser() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the user's edges.
	pe0 := client.Pet.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("pet created:", pe0)
	u1 := client.User.
		Create().
		SetAge(1).
		SetName("string").
		SaveX(ctx)
	log.Println("user created:", u1)

	// create user vertex with its edges.
	u := client.User.
		Create().
		SetAge(1).
		SetName("string").
		AddPets(pe0).
		AddFriends(u1).
		SaveX(ctx)
	log.Println("user created:", u)

	// query edges.
	pe0, err = u.QueryPets().First(ctx)
	if err != nil {
		log.Fatalf("failed querying pets: %v", err)
	}
	log.Println("pets found:", pe0)

	u1, err = u.QueryFriends().First(ctx)
	if err != nil {
		log.Fatalf("failed querying friends: %v", err)
	}
	log.Println("friends found:", u1)

	// Output:
}