package forms

type (
	InsertCategory struct {
		Name string `json:"name" binding:"required"`
	}

	UpdateCategory struct {
		Name string `json:"name" binding:"required"`
	}
)