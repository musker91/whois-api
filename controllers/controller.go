package controllers

import (
	"net/http"
	"whois-api/libs/logger"
	"whois-api/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		"state": models.WhoisStateQueryFail,
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ApiRequestFail, "data": data, "msg": "request params error"})
		logger.Echo.WithFields(logrus.Fields{
			"routers": c.Request.URL.Path,
			"err":     err,
			"query":   requestForm,
		}).Error("get whois request form params fail")
		return
	}
	whoisInfo := &models.WhoisInfo{
		RequestForm: requestForm,
	}
	err = whoisInfo.Whois()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ApiRequestFail, "data": data, "msg": "query fail"})
		logger.Echo.WithFields(logrus.Fields{
			"routers": c.Request.URL.Path,
			"err":     err,
			"query":   requestForm,
		}).Error("query whois info fail")
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
