package command

import (
	"errors"
	"fmt"

	"github.com/yonchando/gator/internal/config"
)

type state struct {
	*config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func (c *commands) run() error {

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("Expect 1 arg")
	}

	s.SetUser(cmd.args[0])
	fmt.Printf("User set to %v \n", cmd.args[0])

	return nil
}
