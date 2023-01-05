package public

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"pt-auto/model"

	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
)

var Configs model.Config

var Redis *redis.Client

var Log *zap.Logger

func PostHttp(url string, data map[string]string, header ...http.Header) (body []byte, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
	}
	jsonByte, err := json.Marshal(data)
	if err != nil {
		return
	}
	res, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonByte))
	if err != nil {
		return
	}
	if len(header) != 0 {
		res.Header = header[0]
	}

	res.Header.Add("content-type", "application/json; charset=utf-8")
	resp, err := client.Do(res)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
