package comm

import "github.com/yangsen996/ExamplesWebSite/model"

type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

type LoginResponse struct {
	User      model.User `json:"user"`
	Token     string     `json:"token"`
	ExpiresAt int64      `json:"expires_at"`
}

type RegisterReq struct {
	Username string `json:"user_name"` // 用户名
	Password string `json:"password"`  // 密码
}
