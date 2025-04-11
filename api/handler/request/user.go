package request

type UserRegisterReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}
