package dto

type LoginResponse struct {
	Token                  string `json:"token"`
	TokenExpiryTime        int64  `json:"tokenExpiryTime"`
	RefreshToken           string `json:"refreshToken"`
	RefreshTokenExpiryTime int64  `json:"refreshTokenExpiryTime"`
}
