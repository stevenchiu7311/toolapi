package apiserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"toolapi/dao"
)

const (
	APIV1MAP = "shared_map"
)

var sharedMap = make(map[string]string)

func InitShareMapApi(e *gin.Engine, bundle *dao.OptBundle) {
	e.POST("/" + APIV1MAP + "/put", PutValue)
	e.POST("/" + APIV1MAP + "/get", GetValue)
}

type SharedMapRequest struct {
	Key     string `json:"key"`
	Value	string `json:"value"`
}

type SharedMapResponse struct {
	Key     string `json:"key"`
	Value	string `json:"value"`
}

func PutValue(c *gin.Context) {
	var request SharedMapRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{keyError: InputDataIncorrect})
		return
	}

	sharedMap[request.Key] = request.Value
	fmt.Println("store key", request.Key)
	c.JSON(http.StatusOK, &request)
}

func GetValue(c *gin.Context) {
	var request SharedMapRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{keyError: InputDataIncorrect})
		return
	}

	fmt.Println("get key", request.Key)

	c.JSON(http.StatusOK, &SharedMapResponse{Key : request.Key, Value : sharedMap[request.Key]})
}