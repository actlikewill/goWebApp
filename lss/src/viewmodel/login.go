package viewmodel

// Login implements the login struct.
type Login struct {
	Title string
	Active string
	Email string
	Password string
}

// NewLogin initiates a new Login.
func NewLogin() Login {
	result := Login{
		Active: "home",
		Title: "Lemonade Stand Supply",
	}
	return result
}