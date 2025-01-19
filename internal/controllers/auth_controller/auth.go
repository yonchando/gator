package auth_controller

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/yonchando/gator/internal/database"
	"github.com/yonchando/gator/internal/models/command"
	"github.com/yonchando/gator/internal/models/state"
)

func HandlerLogin(s *state.State, cmd command.Command) error {

	if len(cmd.Args) != 3 {
		fmt.Println("username is required")
		os.Exit(1)
	}

	var user database.User
	var err error

	user, err = s.Db.GetUserByName(context.Background(), cmd.Args[2])
	if err != nil {
		return err
	}

	fmt.Println(user.Name)

	if user.Name == "" {
		return errors.New("user not found")
	}

	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("User set to %v \n", cmd.Args[2])

	return nil
}

func HandlerRegister(s *state.State, cmd command.Command) error {
	if len(cmd.Args) != 3 {
		fmt.Println("username is required")
		os.Exit(1)
	}

	ctx := context.Background()

	exits, _ := s.Db.GetUserByName(ctx, cmd.Args[2])

	if exits.Name != "" {
		return errors.New("User is already exits.")
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      cmd.Args[2],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	var user database.User
	var err error

	user, err = s.Db.CreateUser(ctx, userParams)

	if err != nil {
		return err
	}

	err = s.Cfg.SetUser(user.Name)

	if err != nil {
		return err
	}

	return nil
}

func HandlerGetUsers(s *state.State, _ command.Command) error {

	var ctx = context.Background()
	var users []database.User
	var err error

	users, err = s.Db.GetUsers(ctx)

	if err != nil {
		return err
	}

	for _, user := range users {
		fmt.Printf("* %s ", user.Name)
		if s.Cfg.CurrentUserName == user.Name {
			fmt.Println("(current)")
		} else {
			fmt.Println("")
		}
	}

	return nil
}
