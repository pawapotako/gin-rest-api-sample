package handler

import (
	"go-project/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type defaultHandler struct {
}

func InitDefaultHandler(e *gin.Engine) {
	handler := defaultHandler{}

	e.GET("/health-check", handler.healthCheck)
	e.POST("/covid/summary", handler.covidSummary)
}

func (h defaultHandler) healthCheck(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h defaultHandler) covidSummary(c *gin.Context) {

	request := model.DefaultPayload[[]model.CovidSummaryRequest]{}
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	var provinceCount = make(map[string]int)
	var ageGroupCount = make(map[string]int)
	for _, item := range request.Data {

		if item.Province == nil {
			provinceCount["N/A"]++
		} else {
			provinceCount[*item.Province]++
		}

		if item.Age == nil {
			ageGroupCount["N/A"]++
		} else if *item.Age >= 0 && *item.Age <= 30 {
			ageGroupCount["0-30"]++
		} else if *item.Age > 30 && *item.Age <= 60 {
			ageGroupCount["31-60"]++
		} else if *item.Age > 60 {
			ageGroupCount["61+"]++
		}
	}

	response := model.CovidSummaryResponse{
		Province: provinceCount,
		AgeGroup: ageGroupCount,
	}

	c.JSON(http.StatusOK, response)
}
