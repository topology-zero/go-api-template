package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"go-api-template/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

var writeList = []string{
	"/login",
	"/metrics",
}

func RequestLog(c *gin.Context) {
	if c.Request.Method == http.MethodOptions {
		c.Next()
		return
	}

	var requestBody string

	if c.Request.Method == http.MethodGet {
		requestBody = c.Request.URL.RawQuery
	} else {
		switch c.ContentType() {
		case binding.MIMEJSON:
			body, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			requestBody = string(body)
		case binding.MIMEPOSTForm:
			c.Request.ParseForm()
			requestBody = c.Request.PostForm.Encode()
		case binding.MIMEMultipartPOSTForm:
			c.Request.ParseMultipartForm(10e6)
			requestBody = c.Request.PostForm.Encode()
		}
	}

	// 开始时间
	startTime := time.Now()

	c.Next()

	for _, url := range writeList {
		if c.Request.URL.Path == url {
			return
		}
	}

	requestId, _ := c.Get(util.TrafficKey)

	data := logrus.Fields{
		"useTime":    time.Since(startTime),
		"uri":        "[" + c.Request.Method + "]" + c.Request.URL.Path,
		"status":     c.Writer.Status(),
		"ip":         c.ClientIP(),
		util.TraceId: requestId,
	}

	if len(requestBody) > 0 {
		data["request"] = requestBody
	}

	if c.Err() != nil {
		data["err"] = c.Err().Error()
	}

	logrus.WithFields(data).Info()
}
