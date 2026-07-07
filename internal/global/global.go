package global

import (
	"ent-scaffold/internal/conf"

	"go.uber.org/zap"
)

var (
	Config *conf.Config
	Log    *zap.Logger
)
