package cli

import (
	"context"
	"fmt"
	"net/url"

	"github.com/bzelaznicki/gator/internal/rss"
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
