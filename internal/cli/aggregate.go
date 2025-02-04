package cli

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bzelaznicki/gator/internal/database"
	"github.com/bzelaznicki/gator/internal/rss"
	"github.com/google/uuid"
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
	fmt.Printf("Querying %s at %v\n\n", feedToScrape.Name, time.Now())

	for i := range feed.Channel.Item {

		layout := "Mon, 02 Jan 2006 15:04:05 -0700"
		pubDate, err := time.Parse(layout, feed.Channel.Item[i].PubDate)
		if err != nil {
			return fmt.Errorf("failed to parse pubDate: %s", err)
		}

		postParams := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: sql.NullString{
				String: feed.Channel.Item[i].Title,
				Valid:  true,
			},
			Url: feed.Channel.Item[i].Link,
			Description: sql.NullString{
				String: feed.Channel.Item[i].Description,
				Valid:  true,
			},
			PublishedAt: sql.NullTime{
				Time:  pubDate,
				Valid: true,
			},
			FeedID: uuid.NullUUID{
				UUID:  feedToScrape.ID,
				Valid: true,
			},
		}
		err = s.db.CreatePost(context.Background(), postParams)
		if err != nil {
			return fmt.Errorf("failed to add post: %v", err)
		}

	}
	fmt.Printf("\n")

	return nil
}
