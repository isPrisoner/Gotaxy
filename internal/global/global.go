package global

import (
	"context"
	"database/sql"

	"github/JustGopher/Gotaxy/internal/config"
	"github/JustGopher/Gotaxy/internal/pool"

	"go.uber.org/zap"
)

var (
	Ctx      context.Context
	Cancel   context.CancelFunc
	ConnPool *pool.Pool
	Log      *zap.SugaredLogger
	DB       *sql.DB
	Config   config.Config
)
