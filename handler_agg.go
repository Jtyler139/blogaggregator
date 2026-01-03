package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"database/sql"

	"github.com/jtyler139/blogaggregator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s...", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)
		return
	}
	log.Println("Found a feed to fetch!")
	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}
	for _, item := range feedData.Channel.Item {
		t, err := convertTime(item.PubDate)
		if err != nil {
			log.Printf("Couldn't convert time: %w", err)
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:				uuid.New(),
			CreatedAt:		time.Now().UTC(),
			UpdatedAt:		time.Now().UTC(),
			Title:			item.Title,
			Url:			item.Link,
			Description:	sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt:	t,
			FeedID:			feed.ID,
		})
		if err != nil {
			if isPQUniqueViolation(err) {
				continue
			}
			log.Printf("Couldn't create post: %w", err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}


func convertTime(timeString string) (time.Time, error) {
	layouts := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC3339,
	}

	for _, layout :=  range layouts {
		t, err := time.Parse(layout, timeString)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("couldn't convert time")
}

func isPQUniqueViolation(err error) bool {
	if err == nil {
		return false
	}
	if pqErr, ok := err.(*pq.Error); ok {
		return pqErr.Code == "23505"
	}
	return false
}