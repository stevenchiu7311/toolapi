package apiserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toolapi/dao"
)

const (
	APIV1REDIRECT = "redirect"
)

var count = 0
var redirectList = []string{
	"172.30.100.24:30342",
	"172.30.100.33:30342",
	"172.30.100.23:30342",
	"172.30.100.25:30342",
	"172.30.100.34:30342",
}
var redirectCache = make(map[string]*string)

func InitRedirectApi(e *gin.Engine, bundle *dao.OptBundle) {
	e.GET("/" + APIV1REDIRECT + "/list", GetList)
	e.POST("/" + APIV1REDIRECT + "/candidate", GetCandidate)
}

type RedirectCandidateRequest struct {
	CallID string `json:"call_id"`
}

type RedirectResponse struct {
	List []string `json:"list"`
}

func GetCandidate(c *gin.Context) {
	var request RedirectCandidateRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{keyError: InputDataIncorrect})
		return
	}
	cid := request.CallID
	cacheCandidate := redirectCache[cid]
	if cacheCandidate != nil {
		c.JSON(http.StatusOK, &RedirectResponse{List:[]string{*cacheCandidate}})
	} else {
		count = count + 1
		count = count % len(redirectList)
		candidate := redirectList[count]
		redirectCache[cid] = &candidate
		c.JSON(http.StatusOK, &RedirectResponse{List:[]string{candidate}})
	}
}

func GetList(c *gin.Context) {
	c.JSON(http.StatusOK, &RedirectResponse{List:redirectList})
}