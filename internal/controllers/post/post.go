package post_controller

import (
	"context"
	"fmt"
	"strconv"

	"github.com/yonchando/gator/internal/models/command"
	"github.com/yonchando/gator/internal/models/state"
)

func HandleBrowse(s *state.State, cmd command.Command) error {

	var limit int32

	if len(cmd.Args) != 3 {
		limit = 2
	} else {
		arg, err := strconv.Atoi(cmd.Args[2])

		if err != nil {
			limit = 2
		} else {
			limit = int32(arg)
		}
	}

	ctx := context.Background()
	posts, err := s.Db.GetPosts(ctx, int32(limit))

	if err != nil {
		return err
	}

	fmt.Println(posts)

	return nil
}
