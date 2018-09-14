package utils

// Response JSON from KKHC Server
type Response struct {
	Msg string `json:"msg"`
}

// User obj...
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Collection obj...
type Collection struct {
	Name string `json:"Name"`
}

// Avatar obj...
type Avatar struct {
	NameOnDisc string `json:"nameOnDisc"`
	Extenison  string `json:"extension"`
}

// HTTPHeader is a key-value pair of strings
// passed to the HTTPRequest as header.
type HTTPHeader struct {
	key, value string
}
