package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func makeRoutes() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			c.SetCookie("gin_cookie", time.Now().Format("2006-01-02T15:04:05Z07:00"), 3600, "/", "localhost", false, true)
		}

		if cookie == "" {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title":      "Main website",
				"today":      time.Now().Format("2006-01-02T15:04:05Z07:00"),
				"first_time": time.Now().Format("2006-01-02T15:04:05Z07:00"),
			})
			return
		}

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":      "Main website",
			"today":      time.Now().Format("2006-01-02T15:04:05Z07:00"),
			"first_time": cookie,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.AsciiJSON(http.StatusOK, gin.H{
			"test":  "ok",
			"test2": 3,
			"test3": []int{1, 2, 3},
		})
	})

	r.GET("test", func(c *gin.Context) {
		t := c.DefaultQuery("t", "")
		if t == "" {
			c.AsciiJSON(http.StatusOK, gin.H{
				"nope": "1",
			})
			return
		}

		c.AsciiJSON(http.StatusOK, gin.H{
			"test":  "ok",
			"test2": t,
		})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.GET("user", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		c.JSON(http.StatusOK, gin.H{"user": user, "value": 1})
	})

	return r
}
