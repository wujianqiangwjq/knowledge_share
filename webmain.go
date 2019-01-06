package main

import (
	"elastic"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	var item elastic.Item
	res := c.BindJSON(&item)
	if res == nil {
		if item.AddData() {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": res, "status": "failed",
		})
	}

}

func Search(c *gin.Context) {
	status := c.DefaultQuery("q", "")
	log.Println(status)
	if status != "" {
		status = strings.TrimSpace(status)
	}
	data := elastic.SearchData(status)
	c.JSON(http.StatusOK, gin.H{"data": data})

}
func main() {
	r := gin.Default()
	group := r.Group("/api")
	group.POST("/issues/add/", Add)
	group.GET("/issues/search/", Search)
	r.Run(":80")

}
