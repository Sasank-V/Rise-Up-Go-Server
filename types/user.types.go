package types

type MessageResponse struct {
	Message string `json:"message"`
}

type SigninRequest struct {
	GoogleID     string `json:"google_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Picture      string `json:"picture"`
	Role         string `json:"role"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    string `json:"expires_at"`
}
