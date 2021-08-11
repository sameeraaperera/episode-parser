package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sap4001/episode-parser/internal/episode"
)

type Output struct {
	Response []Show
}

type Show struct {
	Image string
	Slug  string
	Title string
}

func JSONParser(c *gin.Context) {
	var ed episode.Data
	err := c.ShouldBindJSON(&ed)
	if err != nil {
		log.Printf("Error parsing input JSON %q", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not decode request: JSON parsing failed",
		})
		return
	}

	output := Output{}
	for _, val := range ed.Payload {
		if episode.IsDRMEnabled(val) {
			show := Show{
				val.Image.ShowImage,
				val.Slug,
				val.Title,
			}
			output.Response = append(output.Response, show)
		}

	}

	c.JSON(http.StatusOK, output)

}
