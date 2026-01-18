package services

func ComparePasswords(passwordInMessage string, passwordInProfile string) bool {
	passwordInMessagemd5 := Makemd5(passwordInMessage)
	return passwordInMessagemd5 == passwordInProfile
}
