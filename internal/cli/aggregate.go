package cli

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bzelaznicki/gator/internal/database"
	"github.com/bzelaznicki/gator/internal/rss"
)

func HandlerAgg(s *state, cmd Command) error {

	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("agg requires a time duration. Usage: agg <time_between_reqs>")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("cannot parse time: %v", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)

	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			fmt.Printf("Error scraping feeds: %v\n", err)
		}
	}

	return nil
}

func scrapeFeeds(s *state) error {
	feedToScrape, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get next feed: %v", err)
	}
	var timestamp sql.NullTime
	timestamp.Time = time.Now()
	timestamp.Valid = true

	fetchParams := database.MarkFeedFetchedParams{
		LastFetchedAt: timestamp,
		UpdatedAt:     time.Now(),
		ID:            feedToScrape.ID,
	}
	_, err = s.db.MarkFeedFetched(context.Background(), fetchParams)

	if err != nil {
		return fmt.Errorf("failed to mark feed as fetched: %v", err)
	}

	feed, err := rss.FetchFeed(context.Background(), feedToScrape.Url)

	if err != nil {
		return err
	}
	fmt.Printf("On %s at %v found:\n\n", feedToScrape.Name, time.Now())

	for i := range feed.Channel.Item {
		fmt.Printf("Title: %s\n", feed.Channel.Item[i].Title)
	}
	fmt.Printf("Taken from: %s", feedToScrape.Url)
	fmt.Printf("\n\n")

	return nil
}
