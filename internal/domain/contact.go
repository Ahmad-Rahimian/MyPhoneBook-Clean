package domain

// Contact is our core entity (exported fields for JSON/DB scanning)
type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}
