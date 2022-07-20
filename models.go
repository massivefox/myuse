package github.com/massivefox/myuse

type CreateInput struct {
	Id    string `json:"id"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type UpdateInput struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
}
