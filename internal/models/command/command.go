package command

import (
	"github.com/yonchando/gator/internal/models/state"
	"log"
)

type Commands struct {
	Handlers map[string]func(*state.State, Command) error
}

type Command struct {
	Name string
	Args []string
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) {
	c.Handlers[name] = f
}

func (c *Commands) Run(s *state.State, cmd Command) error {

	f, ok := c.Handlers[cmd.Name]
	if !ok {
		log.Fatalln("command not found")
	}

	err := f(s, cmd)

	if err != nil {
		return err
	}

	return nil
}
