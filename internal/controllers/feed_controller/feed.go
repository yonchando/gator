package feed_controller

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
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

func HandleAgg(s *state.State, cmd command.Command) error {

	d, err := time.ParseDuration(cmd.Args[2])
	if err != nil {
		return err
	}

	fmt.Println("Collecting feeds every 1m0s")

	ticker := time.NewTicker(d)

	for ; ; <-ticker.C {
		fmt.Println("Fetching...")
		scapeFeeds(s)
	}
}

func scapeFeeds(s *state.State) error {

	ctx := context.Background()

	feed, err := s.Db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	markParams := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: time.Now(),
		ID:        feed.ID,
	}

	err = s.Db.MarkFeedFetched(ctx, markParams)
	if err != nil {
		return err
	}

	var rss *models.RSSFeed

	rss, err = models.FetchFeed(ctx, feed.Url)

	for _, v := range rss.Channel.Item {

		var published_at time.Time
		published_at, err = time.Parse("RFC3339", v.PubDate)

		postParams := database.CreatePostParams{
			ID:    uuid.New(),
			Title: v.Title,
			Url:   v.Link,
			Description: sql.NullString{
				String: v.Description,
				Valid:  true,
			},
			PublishedAt: sql.NullTime{
				Time:  published_at,
				Valid: err == nil,
			},
			FeedID:    feed.ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		_, err = s.Db.CreatePost(ctx, postParams)

		if e, ok := err.(*pq.Error); ok {
			if e.Code != "23505" {
				log.Println(e.Code.Name())
			}
		} else {
			log.Println(err)
		}
	}

	return nil
}
