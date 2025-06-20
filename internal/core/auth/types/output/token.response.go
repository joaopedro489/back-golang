package output

type TokenResponse struct {
	Token string `json:"token"`
}

func NewTokenResponse(token string) *TokenResponse {
	return &TokenResponse{
		Token: token,
	}
}
