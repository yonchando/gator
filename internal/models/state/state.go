package state

import (
	"github.com/yonchando/gator/internal/config"
	"github.com/yonchando/gator/internal/database"
)

type State struct {
	Cfg *config.Config
	Db  *database.Queries
}
