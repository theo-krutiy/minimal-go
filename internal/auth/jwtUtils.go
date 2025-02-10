package auth

func issueNewJWT() (string, error) {
	return "DummyJWT", nil
}

func validateJWT(token string) error {
	return nil
}
