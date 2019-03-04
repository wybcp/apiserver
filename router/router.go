package router

import (
	"apiserver/handler/sd"
	"apiserver/handler/user"
	"apiserver/router/middleware"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// Load loads the middleware, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// pprof router
	pprof.Register(g)

	// Middleware
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	g.POST("/login", user.Login)
	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:id", user.Index)
		//u.GET("/:username", user.GetUser)
	}

	// The health check handlers
	serverStatus := g.Group("/sd")
	{
		serverStatus.GET("/health", sd.HealthCheck)
		serverStatus.GET("/disk", sd.DiskCheck)
		serverStatus.GET("/cpu", sd.CPUCheck)
		serverStatus.GET("/ram", sd.RAMCheck)
	}

	return g
}
