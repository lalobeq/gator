package main

import (
	"context"
	"fmt"
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
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Println(item.Title)
	}

	return nil
}
