package controllers

import (
	"net/http"
	"whois-api/models"

	"github.com/gin-gonic/gin"
)

const (
	ApiRequestSuccess = iota
	ApiRequestFail
)

func SiteHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func WhoisQuery(c *gin.Context) {
	var requestForm models.WhoisRequestForm
	err := c.Bind(&requestForm)
	data := gin.H{
		"data":  gin.H{},
		"state": models.WhoisStateRequestParamsError,
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": ApiRequestFail, "data": data, "msg": "request params error"})
		return
	}
	whoisInfo := &models.WhoisInfo{
		RequestForm: requestForm,
	}
	err = whoisInfo.Whois()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": ApiRequestFail, "data": data, "msg": "query fail"})
		return
	}
	data["state"] = whoisInfo.State
	if whoisInfo.State == models.WhoisStateQuerySuccess {
		if whoisInfo.RequestForm.OutType == models.WhoisOutJsonType {
			data["data"] = whoisInfo.JsonInfo
		} else {
			data["data"] = whoisInfo.TextInfo
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": ApiRequestSuccess, "data": data, "msg": ""})
}
