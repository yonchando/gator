package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/yonchando/gator/internal/models/command"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/yonchando/gator/internal/config"
	"github.com/yonchando/gator/internal/controllers/auth"
	"github.com/yonchando/gator/internal/controllers/feed"
	"github.com/yonchando/gator/internal/database"
	"github.com/yonchando/gator/internal/models/state"
)

func handlerReset(s *state.State, cmd command.Command) error {

	_ = cmd

	ctx := context.Background()

	err := s.Db.DeleteAllUser(ctx)

	if err != nil {
		return err
	}

	err = s.Cfg.SetUser("")

	if err != nil {
		return err
	}

	return nil
}

func main() {

	cfg, err := config.Read()

	if err != nil {
		log.Fatalln(err)
	}

	var db *sql.DB

	db, err = sql.Open("postgres", cfg.DbUrl)

	dbQueries := database.New(db)

	s := state.State{
		Cfg: &cfg,
		Db:  dbQueries,
	}

	cmds := command.Commands{
		Handlers: map[string]func(*state.State, command.Command) error{},
	}

	cmds.Register("reset", handlerReset)
	cmds.Register("agg", feed.HandleAgg)

	cmds.Register("login", auth.HandlerLogin)
	cmds.Register("register", auth.HandlerRegister)
	cmds.Register("users", auth.HandlerGetUsers)

	cmds.Register("addfeed", feed.HandleAddFeed)
	cmds.Register("feeds", feed.HandleFeed)
	cmds.Register("follow", feed.HandleFeedFollow)

	args := os.Args

	if len(args) == 1 {
		fmt.Println("not enough arguments were provided")
		os.Exit(1)
	}

	cmd := command.Command{
		Name: args[1],
		Args: args,
	}

	err = cmds.Run(&s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
