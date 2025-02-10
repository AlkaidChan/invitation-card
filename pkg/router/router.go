package router

import (
	"time"

	"github.com/alkaidchen/invitation-card/pkg/middleware/cros"
	"github.com/alkaidchen/invitation-card/pkg/api"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var DefaultTheme = "default"

func NewRouter(logger *zap.Logger, debug bool) (*gin.Engine, error) {
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	g.Use(cros.Cors())
	g.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	g.Use(ginzap.RecoveryWithZap(logger, true))

	g.POST("/api/submit", api.Submit)

	return g, nil
}
