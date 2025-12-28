package main

import (
	"fmt"
	"github.com/google/uuid"
	"context"
	"time"
	"os"
	"github.com/jtyler139/blogaggregator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	ctx := context.Background()

	_, err := s.db.GetUser(ctx, name)
	if err != nil {
		fmt.Printf("User <%s> does not exist\n", name)
		os.Exit(1)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	ctx := context.Background()

	u, err := s.db.GetUser(ctx, name)
	if err == nil {
		fmt.Printf("user <%s> already exists\n", u.Name)
		os.Exit(1)
	}


	user, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID: 		uuid.New(),
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
		Name: 		name,
	})
	if err != nil {
		return fmt.Errorf("error registering new user: %v", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}


	fmt.Printf("New user %s was created\n", user.Name)

	return nil
}