package logic

import (
	"context"
	"fmt"
	"gold/model/request"
	"gold/public/common"
	"gold/public/tools"
	"gold/service/isql"

	"github.com/gin-gonic/gin"
)

type TradeUser struct{}

// 列出用户
func (t TradeUser) ListTradeUser(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	ctx := context.Background()
	usList, err := isql.User.ListTradeUser(ctx)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取用户数据失败: %v", err.Error()))
	}
	result := usList
	return result, nil
}

// 删除用户
func (t TradeUser) DeleteTradeUser(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	ctx := context.Background()
	r, ok := req.(*request.DeleteTradeUserReq)
	if !ok {
		common.Log.Errorf("转换数据结构失败")
		return nil, ReqAssertErr
	}

	err := isql.User.DeleteTradeUser(ctx, r.UserID)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("删除用户失败: %v", err.Error()))
	}
	return nil, nil
}

// 创建用户
func (t TradeUser) CreateTradeUser(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	ctx := context.Background()
	r, ok := req.(*request.CreateTradeUserReq)
	if !ok {
		common.Log.Errorf("转换数据结构失败")
		return nil, ReqAssertErr
	}
	fmt.Println(r)
	err := isql.User.CreateTradeUser(ctx, r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("创建用户失败: %v", err.Error()))
	}
	return nil, nil
}

// 修改用户
func (t TradeUser) UpdateTradeUser(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	ctx := context.Background()
	r, ok := req.(*request.UpdateTradeUserReq)
	if !ok {
		common.Log.Errorf("转换数据结构失败")
		return nil, ReqAssertErr
	}

	err := isql.User.UpdateTradeUser(ctx, r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("更新用户失败: %v", err.Error()))
	}
	return nil, nil
}
