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
	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feedInfo.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), newFeedFollow)
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %v", err)
	}
	fmt.Printf("Feed created successfully:\nName: %s\n URL: %s\n", feedInfo.Name, feedInfo.Url)

	return nil
}

func HandlerFeeds(s *state, cmd Command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds: %v", err)
	}

	if len(feeds) == 0 {
		return fmt.Errorf("no feeds found. add one by using addfeed <name> <url>")
	}

	for _, feed := range feeds {
		fmt.Printf("Name: %s\nURL: %s\nCreated by: %s\n\n", feed.Name, feed.Url, feed.CreatedBy)
	}

	return nil
}

func HandlerFollow(s *state, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("follow requires a feed URL. Usage: follow <url>")
	}

	feedUrl := cmd.Arguments[0]
	_, err := url.Parse(feedUrl)
	if err != nil {
		return err
	}

	feedInfo, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("failed to get feed information: %v", err)
	}

	userInfo, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get user information: %v", err)
	}

	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userInfo.ID,
		FeedID:    feedInfo.ID,
	}

	createdFeedFollow, err := s.db.CreateFeedFollow(context.Background(), newFeedFollow)
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %v", err)
	}

	fmt.Printf("Feed created successfully!\nUser: %s\nFeed name: %s\nCreated at: %v\n Updated at: %v\n", createdFeedFollow.UserName, createdFeedFollow.FeedName, createdFeedFollow.CreatedAt, createdFeedFollow.UpdatedAt)

	return nil
}

func HandlerFollowing(s *state, cmd Command) error {
	userInfo, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get user information: %v", err)
	}
	followedFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), userInfo.ID)
	if err != nil {
		return fmt.Errorf("failed to get user feeds: %v", err)
	}
	if len(followedFeeds) == 0 {
		return fmt.Errorf("%s is not following any feeds", userInfo.Name)
	}

	fmt.Printf("%s is following the following feeds:\n", userInfo.Name)
	for _, feed := range followedFeeds {
		fmt.Printf("* %s\n", feed.FeedName)
	}
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
