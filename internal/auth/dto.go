package auth

type (
	FullCredentials struct {
		Name string `json:"name" validate:"required"`
		Credentials
	}

	Credentials struct {
		Login    string `json:"login" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)
