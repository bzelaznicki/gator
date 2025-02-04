package cli

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/bzelaznicki/gator/internal/database"
	"github.com/bzelaznicki/gator/internal/rss"
	"github.com/google/uuid"
)

func HandlerAgg(s *state, cmd Command) error {

	feedUrl := "https://www.wagslane.dev/index.xml"
	_, err := url.Parse(feedUrl)
	if err != nil {
		return err
	}

	feed, err := rss.FetchFeed(context.Background(), feedUrl)

	if err != nil {
		return err
	}

	fmt.Printf("%v\n", feed)

	return nil
}

func HandlerAddFeed(s *state, cmd Command) error {

	if len(cmd.Arguments) < 2 {
		return fmt.Errorf("addfeed command requires a title and an URL. Usage: rss <title> <url>")
	}
	title := cmd.Arguments[0]
	feedUrl := cmd.Arguments[1]
	_, err := url.Parse(feedUrl)
	if err != nil {
		return err
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("user not found: %v", err)
	}
	params := database.InsertFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      title,
		Url:       feedUrl,
		UserID:    user.ID,
	}

	feedInfo, err := s.db.InsertFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	fmt.Printf("Feed created successfully:\nName: %s\n URL: %s\n", feedInfo.Name, feedInfo.Url)

	return nil
}

/*
func HandlerRss(s *state, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("rss command requires an URL. Usage: rss <url>")
	}

	feedUrl := cmd.Arguments[0]
	_, err := url.Parse(feedUrl)
	if err != nil {
		return err
	}

	feed, err := rss.FetchFeed(context.Background(), feedUrl)

	if err != nil {
		return err
	}

	fmt.Printf("%v", feed)

	return nil
}
*/
