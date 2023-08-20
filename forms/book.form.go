package forms

type (
	InsertBook struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		ImageURL    string `json:"image_url"`
		ReleaseYear int    `json:"release_year"`
		Price       string `json:"price"`
		TotalPage   int    `json:"total_page"`
		Thickness   string `json:"thickness"`
		CategoryID  int    `json:"category_id"`
	}

	UpdateBook struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		ImageURL    string `json:"image_url"`
		ReleaseYear int    `json:"release_year"`
		Price       string `json:"price"`
		TotalPage   int    `json:"total_page"`
		Thickness   string `json:"thickness"`
		CategoryID  int    `json:"category_id"`
	}
)