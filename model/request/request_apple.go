package request

type AppleLoginCodeAndId struct {
	AuthorizationCode string `json:"authorization_code" form:"authorizationCode"`
	UserIdentifier    string `json:"user_identifier" form:"user_identifier"`
}
