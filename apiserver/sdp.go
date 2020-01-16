package apiserver

import (
	"fmt"
	"regexp"
	"github.com/gin-gonic/gin"
	"net/http"
	"toolapi/dao"
)

const (
	APIV1SDP = "/sdp"
)

var ReserveP2pCandidate = false
var sdpMap = make(map[string]string)
func InitSdpApi(e *gin.Engine, bundle *dao.OptBundle) {
	e.POST("/" + APIV1SDP, QuerySdp)
	e.POST("/store_sdp", StoreSdp)
	e.POST("/get_sdp", GetSdp)
	ReserveP2pCandidate = *bundle.ReserveP2pCandidate == "true"
	fmt.Println("ReserveP2pCandidate:",ReserveP2pCandidate)
}

type QuerySdpResponse struct {
	Sdp string `json:"sdp"`
}

type QuerySdpRequest struct {
	OriginSdp string `json:"sdp"`
	Offer     string `json:"offer"`
	Tag   	  string `json:"tag"`
}

func QuerySdp(c *gin.Context) {
	var request QuerySdpRequest
	c.BindJSON(&request)

	retSdp := request.OriginSdp
	fmt.Println("Before process:",retSdp)
	if (!ReserveP2pCandidate) {
		re := regexp.MustCompile("(?m)[\r\n]+^.*172.31.*$")
		retSdp = re.ReplaceAllString(retSdp, "")
	}

	fmt.Println("After process:",retSdp)
	c.JSON(http.StatusOK, &QuerySdpResponse{Sdp : retSdp})
}

func StoreSdp(c *gin.Context) {
	var request QuerySdpRequest
	c.BindJSON(&request)

	key := "sdp_" + request.Tag
	sdpMap[key] = request.OriginSdp
	fmt.Println("store key", key)
	c.JSON(http.StatusOK, &QuerySdpResponse{Sdp : ""})
}

func GetSdp(c *gin.Context) {
	var request QuerySdpRequest
	c.BindJSON(&request)
	key := "sdp_" + request.Tag
	retSdp := sdpMap[key]
	fmt.Println("get key", key,)
	//re := regexp.MustCompile("(?m)[\r\n]+^.*extmap.*$")
	//retSdp = re.ReplaceAllString(retSdp, "")
	//re1 := regexp.MustCompile("(?m)[\r\n]+^.*rtcp.*$")
	//retSdp = re1.ReplaceAllString(retSdp, "")
	//
	//retSdp = strings.ReplaceAll(retSdp, "127.0.0.1", "192.168.82.70")
	//retSdp = strings.ReplaceAll(retSdp, "150.116.170.95", "192.168.82.70")
	c.JSON(http.StatusOK, &QuerySdpResponse{Sdp : retSdp})
}