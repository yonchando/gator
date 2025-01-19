package middlewarepAuth

import (
	"context"
	"errors"

	"github.com/yonchando/gator/internal/database"
	"github.com/yonchando/gator/internal/models/command"
	"github.com/yonchando/gator/internal/models/state"
)

func MiddlewareLoggedIn(handler func(s *state.State, cmd command.Command, user database.User) error) func(*state.State, command.Command) error {

	return func(s *state.State, cmd command.Command) error {

		var err error

		var user database.User

		user, err = s.Db.GetUserByName(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return errors.New("user not logged in")
		}

		err = handler(s, cmd, user)

		if err != nil {
			return err
		}

		return nil
	}

}
