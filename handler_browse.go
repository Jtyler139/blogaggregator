package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jtyler139/blogaggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, u database.User) error {
	limit := 2
	if len(cmd.Args) > 1 {
		fmt.Errorf("usage %s <limit>", cmd.Name)
	}
	if len(cmd.Args) == 1 {
		limitStr := cmd.Args[0]
		l, err := strconv.Atoi(limitStr)
		if err != nil {
			return fmt.Errorf("invalid limit: %w", err)
		}
		limit = l
	}
	

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID:	 	u.ID,
		Limit:		int32(limit),
	})
	if err != nil {
		return fmt.Errorf("Couldn't get posts: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("Found post: %s\n", post.Title)
	}

	return nil
}