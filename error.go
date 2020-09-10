package selfvindication

import (
	"bytes"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ErrorRec(req *http.Request, logPayload bool, headers ...string) {
	bs := bytes.Buffer{}
	bs.WriteString("method:")
	bs.WriteString(req.Method)
	bs.WriteString("\n")
	bs.WriteString("URL:")
	bs.WriteString(req.URL.String())
	bs.WriteString("\n")
	for _, header := range headers {
		bs.WriteString(header + ":")
		bs.WriteString(req.Header.Get(header))
		bs.WriteString("\n")
	}
	if logPayload {
		respByte, _ := ioutil.ReadAll(req.Body)
		req.ContentLength = int64(len(respByte))
		result := new(bytes.Buffer)
		result.Write(respByte)
		req.Body = ioutil.NopCloser(result)
		bs.WriteString("payload:")
		bs.Write(respByte)
		bs.WriteString("\n")
	}
	glog.Error(bs.String())
}
func ErrorRes(err error, resp *http.Response, decodingFunc func(res *http.Response) string, headers ...string) {
	if err != nil {
		glog.Errorf("%+v \n", err)
		return
	}
	bs := bytes.Buffer{}
	bs.WriteString("StatusCode:")
	bs.WriteString(strconv.Itoa(resp.StatusCode))
	bs.WriteString("\n")
	for _, header := range headers {
		bs.WriteString(header + ":")
		bs.WriteString(resp.Header.Get(header))
		bs.WriteString("\n")
	}
	if decodingFunc != nil {
		bs.WriteString("payload:")
		bs.WriteString(decodingFunc(resp))
		bs.WriteString("\n")
	}
	glog.Error(bs.String())
}
