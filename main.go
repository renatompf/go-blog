package main

import (
	"github.com/gin-gonic/gin"
	"github.com/renatompf/controllers"
	"github.com/renatompf/initializers"
	"github.com/renatompf/migrate"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToPostgres()
	migrate.Migrate()
}

func main() {
	r := gin.Default()

	// Register Posts route
	controllers.RegisterPostRoutes(r)

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
