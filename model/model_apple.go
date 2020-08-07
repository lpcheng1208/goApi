package model

type AppleIdToken struct {
	Iss            string `json:"iss,omitempty"`
	Aud            string `json:"aud,omitempty"`
	Exp            int    `json:"exp,omitempty"`
	Iat            int    `json:"iat,omitempty"`
	Sub            string `json:"sub,omitempty"`
	AtHash         string `json:"at_hash,omitempty"`
	Email          string `json:"email,omitempty"`
	EmailVerified  string `json:"email_verified,omitempty"`
	IsPrivateEmail string `json:"is_private_email,omitempty"`
	AuthTime       int    `json:"auth_time,omitempty"`
	NonceSupported bool   `json:"nonce_supported,omitempty"`
}

type AppleAuth struct {
	AccessToken  string `json:"access_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	IdToken      string `json:"id_token,omitempty"`
}
