package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func main() {
	logger := log.Default()
	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(func(ctx *gin.Context) {
		logger.Printf(
			"request",
			"method", ctx.Request.Method,
			"path", ctx.Request.URL.Path,
			"ip", ctx.ClientIP(),
			"ua", ctx.Request.UserAgent(),
			"query", ctx.Request.URL.RawQuery,
			"form", ctx.Request.PostForm.Encode(),
			"header", ctx.Request.Header,
		)
		ctx.Next()
	})
	e.Use(Cors())
	e.Any("/api/setcookie", func(ctx *gin.Context) {
		ctx.SetCookie("gin_cookie", "test", 3600, "/", ".zeabur.app", false, false)
	})
	e.Any("/api/getcookie", func(ctx *gin.Context) {
		s, err := ctx.Cookie("gin_cookie")
		if err != nil {
			ctx.JSON(200, err)
			return
		}
		ctx.JSON(200, s)
	})
	e.Any("/api/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "pong")
	})
	e.Any("/kanaries", func(ctx *gin.Context) {
		b, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(200, err)
		}
		fmt.Println(string(b))
		ctx.JSON(200, "OK")
	})
	log.Fatalln(e.Run())
}
