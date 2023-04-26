package auth

type (
	User struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
		Name  string `json:"name"`
		Role  string `json:"role"`
	}

	Creds struct {
		Login    string `json:"login"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Role     string `json:"role"`
	}
)
