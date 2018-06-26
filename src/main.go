package main

import (
  "strconv"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/*a", func(c *gin.Context) {
    s := c.DefaultQuery("s", "200")
    hs, err := strconv.Atoi(s)
    if (err == nil) {
      hs = 200
    }

    c.String(hs, "")
  })
  r.POST("/*a", func(c *gin.Context) {
    s := c.DefaultQuery("s", "201")
    hs, err := strconv.Atoi(s)
    if (err == nil) {
      hs = 201
    }
    c.String(hs, "")
  })

  r.PUT("/*a", func(c *gin.Context) {
    s := c.DefaultQuery("s", "201")
    hs, err := strconv.Atoi(s)
    if (err == nil) {
      hs = 201
    }
    c.String(hs, "")
  })

  r.DELETE("/*a", func(c *gin.Context) {
    s := c.DefaultQuery("s", "204")
    hs, err := strconv.Atoi(s)
    if (err == nil) {
      hs = 209
    }
    c.String(hs, "")
  })

  r.Run(":80")
}
