package domain

type LoginSucessResponse struct {
	Message    string `json:"message"`
	Status     int    `json:"status"`
	AcessToken string `json:"access_token"`
	UserData   User   `json:"user_data"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
