package tools

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/levigross/grequests"
	"k8s.io/klog/v2"
)

const (
	SystemErr    = 500
	MySqlErr     = 501
	LdapErr      = 505
	OperationErr = 506
	ValidatorErr = 412
)

type RspError struct {
	code int
	err  error
}

func (re *RspError) Error() string {
	return re.err.Error()
}

func (re *RspError) Code() int {
	return re.code
}

// NewRspError New
func NewRspError(code int, err error) *RspError {
	return &RspError{
		code: code,
		err:  err,
	}
}

// NewMySqlError mysql错误
func NewMySqlError(err error) *RspError {
	return NewRspError(MySqlErr, err)
}

// NewValidatorError 验证错误
func NewValidatorError(err error) *RspError {
	return NewRspError(ValidatorErr, err)
}

// NewLdapError ldap错误
func NewLdapError(err error) *RspError {
	return NewRspError(LdapErr, err)
}

// NewOperationError 操作错误
func NewOperationError(err error) *RspError {
	return NewRspError(OperationErr, err)
}

// ReloadErr 重新加载错误
func ReloadErr(err interface{}) *RspError {
	rspErr, ok := err.(*RspError)
	if !ok {
		rspError, ok := err.(error)
		if !ok {
			return &RspError{
				code: SystemErr,
				err:  fmt.Errorf("unknow error"),
			}
		}
		return &RspError{
			code: SystemErr,
			err:  rspError,
		}
	}
	return rspErr
}

// Success http 成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

// Err http 错误
func Err(c *gin.Context, err *RspError, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": err.Code(),
		"msg":  err.Error(),
		"data": data,
	})
}

// 返回前端
func Response(c *gin.Context, httpStatus int, code int, data gin.H, message string) {
	c.JSON(httpStatus, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	})
}

type HttpThirdClient struct {
	host      string
	headers   map[string]string
	postBody  map[string]string
	getParams map[string]string
	json      interface{}
	xml       string
}

// NewHttpClient new http client
func NewHttpClient(host string, headers, postBody, getParams map[string]string, json interface{}, xml string) *HttpThirdClient {
	return &HttpThirdClient{
		host:      host,
		headers:   headers,
		postBody:  postBody,
		getParams: getParams,
		json:      json,
		xml:       xml,
	}
}

func (third *HttpThirdClient) buildRequestOptions() *grequests.RequestOptions {
	if !strings.HasPrefix(third.host, "http://") && !strings.HasPrefix(third.host, "https://") {
		third.host = fmt.Sprintf("%s%s", "http://", third.host)
	}
	return &grequests.RequestOptions{
		Headers: third.headers,
		Data:    third.postBody,
		Params:  third.getParams,
		Host:    third.host,
		JSON:    third.json,
		XML:     third.xml,
	}
}

// GET send http GET request
func (third *HttpThirdClient) GET() ([]byte, error) {
	ro := third.buildRequestOptions()
	rsp, err := grequests.Get(ro.Host, ro)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code is %d", rsp.StatusCode)
	}
	if !rsp.Ok {
		return nil, rsp.Error
	}
	bodyByte, err := ioutil.ReadAll(rsp.RawResponse.Body)
	if err != nil {
		return nil, err
	}
	return bodyByte, nil
}

// POST send http POST request
func (third *HttpThirdClient) POST() ([]byte, error) {
	ro := third.buildRequestOptions()
	host := ro.Host
	if ro.JSON != nil {
		ro.Host = ""
	}
	ro.InsecureSkipVerify = true
	rsp, err := grequests.Post(host, ro)

	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code is %d", rsp.StatusCode)
	}
	if !rsp.Ok {
		return nil, rsp.Error
	}
	bodyByte, err := ioutil.ReadAll(rsp.RawResponse.Body)
	if err != nil {
		return nil, err
	}
	return bodyByte, nil
}

func (third *HttpThirdClient) HttpRequestFunc(requestType string) ([]byte, error) {
	targetUrl := third.host

	payload := bytes.NewReader(third.json.([]byte))
	req, _ := http.NewRequest(requestType, targetUrl, payload)

	for k, v := range third.headers {
		req.Header.Add(k, v)
	}
	q := req.URL.Query()
	for k, v := range third.getParams {
		q.Add(k, v)
	}

	response, err := http.DefaultClient.Do(req)
	defer func() {
		Err := response.Body.Close()
		if Err != nil {
			klog.Errorf("httpclient: %s body close failed: %v", targetUrl, err)
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return body, err
	}

	return body, err
}

// HttpPostFormWithHeader post发送form表单
func (third *HttpThirdClient) HttpPostFormWithHeader() ([]byte, error) {
	targetUrl := third.host

	formData := third.json.([]byte)
	header := third.headers

	payload := bytes.NewReader(formData)
	req, err := http.NewRequest("POST", targetUrl, payload)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}

	response, err := http.DefaultClient.Do(req)
	defer func() {
		Err := response.Body.Close()
		if Err != nil {
			klog.Errorf("httpclient: %s body close failed: %v", targetUrl, err)
		}
	}()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return body, err
	}
	return body, err
}
