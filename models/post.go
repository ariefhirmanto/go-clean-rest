package models

// base model
type Post struct {
	ID       int64
	Title    string
	Slug     string
	Content  string
	ImageURL string
	Category string
}

// response
// type PostResponse struct {
// 	ID       int64  `json:"product_id,omitempty" db:"id"`
// 	Title    string `json:"title,omitempty" db:"title"`
// 	Content  string `json:"content,omitempty" db:"content"`
// 	ImageURL string `json:"image_url,omitempty" db:"image_url"`
// 	Category string `json:"category,omitempty" db:"price"`
// }
