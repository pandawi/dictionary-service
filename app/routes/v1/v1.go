package v1

import (
	"dictionary-service/app"
	"github.com/gin-gonic/gin"
	"sync"
)

func SetupV1Router(routes *gin.RouterGroup, app *app.App) *gin.RouterGroup {
	v1Routes := routes.Group("v1")
	v1Routes.GET("similar", func(c *gin.Context) {
		ch := make(chan *gin.Context, 1)
		ch <- c
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			app.HandleSearch(ch)
		}()
		wg.Wait()
	})
	v1Routes.GET("stats", app.GetStats)
	return routes
}
