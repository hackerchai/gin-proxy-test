package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "https://vue-proxy.zeabur.app")
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
	// config, err := rest.InClusterConfig()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// // creates the clientset
	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	panic(err.Error())
	// }
	e := gin.Default()
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
	// e.Any("/k8s", func(ctx *gin.Context) {
	// 	dl, err := clientset.AppsV1().Deployments("zeabur-system").List(ctx, v1.ListOptions{})
	// 	if err != nil {
	// 		ctx.JSON(200, err)
	// 		return
	// 	}
	// 	ctx.JSON(200, dl.Items)
	// })
	log.Fatalln(e.Run())
}
