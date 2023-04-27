package tools

var (
	// 验证码长度
	CodeSize = 6
	// 邮箱密钥
	EmailPassword="xxxxx"
)

// SendCodeToEmail
// 发送验证码到邮箱
func SendCodeToEmail(mail, code string) error {
	e := email.NewEmail()
	e.From = "golang-developer <$email>"
	e.To = []string{mail}
	e.Subject = "验证码请求"
	e.HTML = []byte("<h3>你的验证码为: " + code + "</h3>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "$email", EmailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	return err
}

// RandCode
// 生成随机验证码
func RandCode() string {
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < CodeSize; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}
	return code
}
