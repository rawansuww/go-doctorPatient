package main

//importing package
import (
	"github.com/gin-gonic/gin"
	"github.com/rawansuww/go-doctorpatient/controllers"
	"github.com/rawansuww/go-doctorpatient/services"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	var as *services.AuthService
	as = services.NewAuthService()
	//declare and init.
	ac := controllers.AuthController{
		AuthService: as,
	}
	var rg gin.RouterGroup
	rg = *r.Group("/API") //instantiating, creating a routing group
	ac.SetupRoutes(&rg)   //calling the function

	//anything u do after run is blocked
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
