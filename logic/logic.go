package logic

import (
	"fmt"

	"gold/public/tools"

	jsoniter "github.com/json-iterator/go"
)

var (
	ReqAssertErr = tools.NewRspError(tools.SystemErr, fmt.Errorf("请求异常"))

	User = &TradeUser{}
	//DomainCloud   DomainCloudProvider = &DomainAwsLogic{}

	json = jsoniter.ConfigCompatibleWithStandardLibrary
)
