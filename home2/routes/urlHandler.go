package routes

import "github.com/gin-gonic/gin"

type urlHandler interface {
	Index(c *gin.Context)
	Login(c *gin.Context)
	Signup(c *gin.Context)
	GetById(c *gin.Context)
}
