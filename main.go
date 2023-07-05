package main

import (
	"dictionary-service/app"
	v1 "dictionary-service/app/routes/v1"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	appInstance, appErr := app.InitializeApp()
	if appErr != nil {
		fmt.Printf("error instantiating app: %v", appErr)
	}
	r := setupRouter(appInstance)
	err := r.Run(":8000")
	if err != nil {
		fmt.Printf("error running web server: %v", err)
	}
}

func setupRouter(app *app.App) *gin.Engine {
	r := gin.Default()
	apiRoutes := r.Group("api")
	v1.SetupV1Router(apiRoutes, app)

	return r
}
