package user

import (
	"context"
	"github.com/yonchando/gator/internal/database"
	"github.com/yonchando/gator/internal/models/state"
)

func GetUser(s *state.State, name string) (database.User, error) {

	ctx := context.Background()

	user, err := s.Db.GetUserByName(ctx, name)
	if err != nil {
		return database.User{}, err
	}

	return user, nil
}
