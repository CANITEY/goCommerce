package models

type Product struct {
	ID uint
	Name string
	Description string
	Price float32
	Image []byte
	Reviews []string
}
