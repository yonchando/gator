package feed

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/yonchando/gator/internal/models/command"
	"github.com/yonchando/gator/internal/models/state"

	"github.com/google/uuid"
	"github.com/yonchando/gator/internal/database"
	"github.com/yonchando/gator/internal/models"
)

func HandleAddFeed(s *state.State, cmd command.Command, user database.User) error {

	if len(cmd.Args) != 4 {
		return errors.New("not enough args")
	}

	var err error
	ctx := context.Background()

	var feed database.Feed

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      cmd.Args[2],
		Url:       cmd.Args[3],
		UserID:    user.ID,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	feed, err = s.Db.CreateFeed(ctx, feedParams)
	if err != nil {
		return err
	}

	fmt.Println(feed)

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = s.Db.CreateFeedFollow(ctx, feedFollowParams)

	if err != nil {
		return err
	}

	return nil
}

func HandleFeed(s *state.State, cmd command.Command, user database.User) error {

	ctx := context.Background()

	feeds, err := s.Db.GetFeedWithUser(ctx)

	if err != nil {
		return err
	}

	fmt.Println(feeds)

	return nil
}

func HandleFeedFollow(s *state.State, cmd command.Command, user database.User) error {

	if len(cmd.Args) == 2 {
		return errors.New("not enough args")
	}

	var feed database.Feed
	var err error

	ctx := context.Background()
	feed, err = s.Db.GetFeedByUrl(ctx, cmd.Args[2])
	if err != nil {
		return err
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = s.Db.CreateFeedFollow(ctx, feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("%s has been follow %s\n", user.Name, feed.Name)

	return nil
}

func HandleFeedUnfollow(s *state.State, cmd command.Command, user database.User) error {

	var err error
	var feed database.Feed
	ctx := context.Background()

	feed, err = s.Db.GetFeedByUrl(ctx, cmd.Args[2])
	if err != nil {
		return err
	}

	err = s.Db.DeleteFeedFollowByUrlAndUserId(ctx, database.DeleteFeedFollowByUrlAndUserIdParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return err
	}

	fmt.Printf("%s unfollow %s\n", user.Name, feed.Name)

	return nil

}

func HandleFeedFollowing(s *state.State, _ command.Command, user database.User) error {

	var feedFollowsForUser []database.GetFeedFollowsForUserRow
	var err error

	feedFollowsForUser, err = s.Db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return err
	}

	for _, feedFollow := range feedFollowsForUser {
		fmt.Println(feedFollow.FeedName)
	}

	return nil
}

func HandleAgg(_ *state.State, _ command.Command) error {

	ctx := context.Background()

	rss, err := models.FetchFeed(ctx, "https://www.wagslane.dev/index.xml")

	if err != nil {
		return err
	}

	fmt.Println(rss)

	return nil
}
