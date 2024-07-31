package request

type ListTradeUserReq struct {
}

type DeleteTradeUserReq struct {
	UserID int `json:"id"`
}

type CreateTradeUserReq struct {
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Phone     string `json:"phone"`
	UserRole  string `json:"user_role"`
	UserGroup string `json:"user_group"`
}

type UpdateTradeUserReq struct {
	UserID    int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	UserRole  string `json:"user_role"`
	UserGroup string `json:"user_group"`
	Status    bool   `json:"status"`
}
