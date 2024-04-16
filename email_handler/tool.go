package emailhandler

import "regexp"

func isValidEmail(email string) bool {
	// 使用正则表达式检查电子邮件地址的格式
	regexPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(regexPattern, email)
	if err != nil {
		return false
	}
	return match
}
