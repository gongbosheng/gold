package isql

import (
	"context"
	"gold/model"
	"gold/model/request"
	"gold/public/common"
)

type TradeUserService struct {
}

func (t *TradeUserService) ListTradeUser(ctx context.Context) ([]model.TradeUser, error) {
	var usList []model.TradeUser
	err := common.DB.WithContext(ctx).Model(&model.TradeUser{}).Find(&usList).Error
	if err != nil {
		return usList, err
	}
	return usList, nil
}

func (t *TradeUserService) DeleteTradeUser(ctx context.Context, id int) error {
	err := common.DB.WithContext(ctx).Where("user_id = ?", id).Delete(&model.TradeUser{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TradeUserService) CreateTradeUser(ctx context.Context, r *request.CreateTradeUserReq) error {
	data := model.TradeUser{
		Username:  r.Username,
		Email:     r.Email,
		Phone:     r.Phone,
		UserRole:  r.UserRole,
		UserGroup: r.UserGroup,
	}
	err := common.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TradeUserService) UpdateTradeUser(ctx context.Context, r *request.UpdateTradeUserReq) error {
	data := model.TradeUser{
		Username:  r.Username,
		Email:     r.Email,
		Phone:     r.Phone,
		UserRole:  r.UserRole,
		UserGroup: r.UserGroup,
	}
	err := common.DB.WithContext(ctx).Model(&model.TradeUser{}).Where("user_id = ?", r.UserID).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}
