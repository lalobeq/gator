package main

import (
	"context"
	"fmt"
	"github.com/lalobec/gator/internal/database"
	"github.com/google/uuid"
	"time"
	"github.com/lib/pq"
)

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	feed_to_fetch, err := s.db.GetNextFeedToFetch(ctx)		
	if err != nil {
		return fmt.Errorf("Error getting the feed: %v", err)
	}

	err = s.db.MarkFeedFetched(ctx, feed_to_fetch.ID)
	if err != nil {
		return fmt.Errorf("Can not mark feed as fetched: %v", err)
	}
	
	rssFeed, err := fetchFeed(ctx, feed_to_fetch.Url)
	if err != nil {
		return fmt.Errorf("An error ocurred while fetching: %v", err)
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Println(item.Title)
		fmt.Println("Saving to database...")

		// Parsing time
		parsedTime, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			fmt.Printf("An error ocurred in parsing: %v", err)
			continue
		}

		argsPost := database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(), 
			UpdatedAt: time.Now().UTC(),
			Title: item.Title,
			Url: item.Link,
			Description: item.Description,
			PublishedAt: parsedTime,
			FeedID: feed_to_fetch.ID,
		}

		_, err = s.db.CreatePost(ctx, argsPost)
		if err != nil {
			if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
        // duplicate url, ignore
        continue
    	}
    	return fmt.Errorf("create post: %w", err)
		}
		fmt.Println("Post created successfully")
		//fmt.Println(post)

	}

	return nil
}
