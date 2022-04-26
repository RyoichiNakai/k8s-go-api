package main

import (
	"github.com/gin-gonic/gin"

	"fmt"
)

func main() {
	g := gin.Default()
	
	g.GET("/", helloWorld)
	g.GET("/:user", helloUser)
	g.GET("/:user/:region", greedingUser)

	g.Run(":8080")
}

func helloWorld(ctx *gin.Context) {
	ctx.Writer.WriteString("Hello World!")
}

func helloUser(ctx *gin.Context) {
	user := ctx.Param("user")
	ctx.Writer.WriteString("Hello " + user + "!")
}

func greedingUser(ctx *gin.Context) {
	region := ctx.Param("region")
	user := ctx.Param("user")
	switch region {
  case "ja":
		ctx.Writer.WriteString(fmt.Sprintf("こんにちは、%s!", user))
	case "fr":
		ctx.Writer.WriteString(fmt.Sprintf("Bonjour %s!", user))
  case "de":
		ctx.Writer.WriteString(fmt.Sprintf("Guten Tag %s!", user))
	default:
  	ctx.Writer.WriteString(fmt.Sprintf("Hello %s!", user))
	}
}
