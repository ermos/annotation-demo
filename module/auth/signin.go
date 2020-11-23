package auth

/*
	@Route("POST", "/auth/signin")
	@Desc("sign in to the application")
	@Auth(["admin", "user"])
	@Response([200, 500])
	@Payload("username", string)
	@Payload("password", string)
	@?Payload("2fa", string)
*/
func (Handler) SignIn() {
	// Logic here
}