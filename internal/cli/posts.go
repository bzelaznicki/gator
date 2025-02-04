package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bzelaznicki/gator/internal/database"
)

func HandlerBrowse(s *state, cmd Command, user database.User) error {
	var limit int32
	if len(cmd.Arguments) == 0 {
		limit = 2
	} else {
		setLimit, err := strconv.ParseInt(cmd.Arguments[0], 10, 32)
		if err != nil {
			return fmt.Errorf("invalid value in limit: %s", err)
		}
		limit = int32(setLimit)
	}

	args := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}
	posts, err := s.db.GetPostsForUser(context.Background(), args)
	if err != nil {
		return fmt.Errorf("failed to get posts: %s", err)
	}
	if len(posts) == 0 {
		fmt.Printf("you don't have any feeds yet. subscribe to some! to view available feeds use feeds or add a new one by using addfeed <name> <url>\n")
		return nil
	}
	fmt.Printf("Recent %d posts from your feed:\n\n", limit)
	for _, post := range posts {

		fmt.Printf("Post from %s:\n", post.FeedName)
		fmt.Printf("Title: %s\n", post.Title.String)
		fmt.Printf("Published at: %s\n", post.PublishedAt.Time)
		fmt.Printf("Description: %s\n", post.Description.String)
		fmt.Printf("URL: %s\n", post.Url)
		fmt.Printf("========================\n\n\n\n")
	}

	return nil
}
