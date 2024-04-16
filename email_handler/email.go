package emailhandler

import (
	"fmt"

	"gopkg.in/gomail.v2"
	"lxtend.com/friend_link/config"
	"lxtend.com/friend_link/logger"
)

var e EmailHandler

type EmailHandler struct {
	server string
	port   int
	user   string
	passwd string
}

func Init(server string, port int, user string, passwd string) {
	e.Init(server, port, user, passwd)
}

func (e *EmailHandler) Init(server string, port int, user string, passwd string) {
	e.server = server
	e.port = port
	e.user = user
	e.passwd = passwd
}

func GetMailHandler() *EmailHandler {
	return &e
}

func (e *EmailHandler) SendFriendApplicationByEmail(name string, url string, description string, avatar string, approve_token string) error {
	approveUrl := fmt.Sprintf("https://%s/friend/approve.php?token=%s", config.GetConfig().Domain.Api, approve_token)
	rejectUrl := fmt.Sprintf("https://%s/friend/reject.php?token=%s", config.GetConfig().Domain.Api, approve_token)
	return e.MailTo(config.GetConfig().Email.AdminMail, "友链申请", fmt.Sprintf("来自用户%s发起的友链申请<br>网站地址为<a href=%s>%s</a><br>网站描述为 %s<br>头像地址为 %s<br><a href=%s>批准请求</a><br><a href=%s>拒绝请求</a>", name, url, url, description, avatar, approveUrl, rejectUrl))
}

func (e *EmailHandler) SendApplicationUploaded(peer_mail string) error {
	return e.MailTo(peer_mail, "友链申请", "你的友链申请已经上传成功，等待博主审核中<br>请注意查收邮件，审核结果将通过邮件通知你<br>如果你未请求添加友链，请忽略此邮件")
}

func (e *EmailHandler) SendApplicationApproved(peer_mail string) error {
	msg := fmt.Sprintf("你的友链申请已经通过审核，快访问<a href=https://%s/friend/>https://%s/friend/</a>查看吧", config.GetConfig().Domain.Host, config.GetConfig().Domain.Host)
	return e.MailTo(peer_mail, "友链申请", msg)
}

func (e *EmailHandler) SendApplicationRejected(peer_mail string) error {
	return e.MailTo(peer_mail, "友链申请", "你的友链申请未通过审核，请联系博主以获取更多信息")
}

func (e *EmailHandler) MailTo(to string, subject string, content string) error {
	if !isValidEmail(to) {
		logger.Error("invalid email address:", to)
		return fmt.Errorf("invalid email address: %s", to)
	}
	// Set up the email message
	m := gomail.NewMessage()
	m.SetHeader("From", e.user)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	// Set up the SMTP server configuration
	d := gomail.NewDialer(e.server, e.port, e.user, e.passwd)

	// Send the email
	return d.DialAndSend(m)
}
