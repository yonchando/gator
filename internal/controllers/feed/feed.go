package feed

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/yonchando/gator/internal/models/command"
	"github.com/yonchando/gator/internal/models/state"
	"github.com/yonchando/gator/internal/models/user"

	"github.com/google/uuid"
	"github.com/yonchando/gator/internal/database"
	"github.com/yonchando/gator/internal/models"
	modelUser "github.com/yonchando/gator/internal/models/user"
)

func HandleAddFeed(s *state.State, cmd command.Command) error {

	if len(cmd.Args) != 4 {
		return errors.New("not enough args")
	}

	var user database.User
	var err error
	ctx := context.Background()

	user, err = modelUser.GetUser(s, s.Cfg.CurrentUserName)
	if err != nil {
		return err
	}

	if user.Name == "" {
		return errors.New("user not found")
	}

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

	return nil
}

func HandleFeed(s *state.State, cmd command.Command) error {

	ctx := context.Background()

	feeds, err := s.Db.GetFeedWithUser(ctx)

	if err != nil {
		return err
	}

	fmt.Println(feeds)

	return nil
}

func HandleFeedFollow(s *state.State, cmd command.Command) error {

	if len(cmd.Args) == 2 {
		return errors.New("not enough args")
	}

	var dbUser database.User

	dbUser, err := user.GetUser(s, s.Cfg.CurrentUserName)
	if err != nil {
		return err
	}

	if dbUser.Name == "" {
		return errors.New("user not exists")
	}

	var feed database.Feed

	ctx := context.Background()
	feed, err = s.Db.GetFeedByUrl(ctx, cmd.Args[2])
	if err != nil {
		return err
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    dbUser.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = s.Db.CreateFeedFollow(ctx, feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("%s has been follow %s\n", dbUser.Name, feed.Name)

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
