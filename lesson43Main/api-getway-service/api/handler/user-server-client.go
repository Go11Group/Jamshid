package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func (h *Handler) UserClient(gn *gin.Context) {
	method := gn.Request.Method
	url := gn.Request.URL.Path
	body := gn.Request.Body
	client := &http.Client{}
	req, err := http.NewRequest(method, "http://localhost:8090"+url, body)
	if err != nil {
		BadRequest(gn, err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			BadRequest(gn, err)
		}
	}(res.Body)

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		BadRequest(gn, err)
		fmt.Println("error:", err.Error())
		return
	}

	var data1 map[string]interface{}
	err = json.Unmarshal(resp, &data1)
	if err != nil {
		gn.JSON(res.StatusCode, gin.H{"data": string(resp)})
	} else {
		gn.JSON(res.StatusCode, data1)
	}
}
