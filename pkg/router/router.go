package router

import (
	"io/fs"
	"net/http"
	"time"

	"github.com/alkaidchen/invitation-card/assets"
	"github.com/alkaidchen/invitation-card/pkg/api"
	"github.com/alkaidchen/invitation-card/pkg/middleware/cros"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var DefaultTheme = "default"

func NewRouter(logger *zap.Logger, embedServer *assets.EmbedServer, debug bool) (*gin.Engine, error) {
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	g.Use(cros.Cors())
	g.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	g.Use(ginzap.RecoveryWithZap(logger, true))

	g.POST("api/submit", api.Submit)

	g.StaticFS("/ui/dist", http.FS(assets.FrontendAssets))
	g.NoRoute(func(c *gin.Context) {
		//设置响应状态
		c.Writer.WriteHeader(http.StatusOK)
		indexHTML, err := fs.ReadFile(assets.FrontendAssets, "ui/dist/index.html")
		if err != nil {
			c.Writer.WriteString("index.html not found, err: " + err.Error())
			return
		}
		_, err = c.Writer.Write(indexHTML)
		if err != nil {
			c.Writer.WriteString("write index.html failed, err: " + err.Error())
			c.Redirect(http.StatusInternalServerError, "/error")
			return
		}
		c.Writer.Header().Add("Accept", "text/html")
		c.Writer.Flush()
		return
	})

	return g, nil
}
