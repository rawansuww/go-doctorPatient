package controllers //packages always in smol letters

import (
	"github.com/gin-gonic/gin"
	"github.com/rawansuww/go-doctorpatient/services"
) //uppercase is public and small is private

type AuthController struct {
	AuthService *services.AuthService
}

//route handler to pass it to rg.get
func (ac *AuthController) RegisterUser(ctx *gin.Context) {

	ctx.JSON(200, gin.H{"status": "registered"}) //gin.H creates a jason object
}

//in params, variable name comes first
func (ac *AuthController) SetupRoutes(rg *gin.RouterGroup) {
	rg.GET("/register", ac.RegisterUser) //assign route to handler func
}

//fake constructor of AuthController, which takes AuthService
func NewAuthController(as *services.AuthService) *AuthController {
	return &AuthController{
		AuthService: as,
	}
}
