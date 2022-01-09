package main

import (
	"github.com/wyy-go/wsession"
	"github.com/wyy-go/wsession/memstore"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := memstore.NewStore([]byte("secret"))
	r.Use(wsession.New("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		session := wsession.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8080")
}
