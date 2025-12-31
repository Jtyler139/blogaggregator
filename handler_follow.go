package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jtyler139/blogaggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	newFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt:	time.Now().UTC(),
		UserID:		user.ID,
		FeedID:		feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	fmt.Println("Feed followed successfully:")
	printFollow(newFollow.ID, newFollow.CreatedAt, newFollow.UpdatedAt, user, feed)
	fmt.Println("=====================================")

	return nil
}

func handlerListFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get list of follows: %w", err)
	}

	if len(follows) == 0 {
		fmt.Println("No follows found.")
		return nil
	}

	fmt.Printf("Found %d follows:\n", len(follows))
	for _, follow := range follows {
		feed, err := s.db.GetFeedByID(context.Background(), follow.FeedID)
		if err != nil {
			return fmt.Errorf("couldn't get feed: %w", err)
		}
		printFollow(follow.ID, follow.CreatedAt, follow.UpdatedAt, user, feed)
		fmt.Println("=====================================")
	}

	return nil
}

func printFollow(id uuid.UUID, created_at, updated_at time.Time, user database.User, feed database.Feed) {
	fmt.Printf("* ID:            %s\n", id)
	fmt.Printf("* Created:       %s\n", created_at)
	fmt.Printf("* Updated:       %s\n", updated_at)
	fmt.Printf("* Feed:          %s\n", feed.Name)
	fmt.Printf("* Feed ID:       %s\n", feed.ID)
	fmt.Printf("* User:          %s\n", user.Name)
	fmt.Printf("* User ID:       %s\n", user.ID)
}