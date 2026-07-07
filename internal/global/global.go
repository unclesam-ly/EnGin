package global

import (
	"EnGin/internal/conf"

	"go.uber.org/zap"
)

var (
	Config *conf.Config
	Log    *zap.Logger
)
