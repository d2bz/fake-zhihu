// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type RegisterRequest struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}
