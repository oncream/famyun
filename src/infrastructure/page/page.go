package page

type Page struct {
	TotalElements int64 `json:"totalElements"`
	TotalPages    int   `json:"totalPages"`
	PageNumber    int   `json:"number" form:"page"`
	PageSize      int   `json:"size" form:"size"`
	Content       any   `json:"content"`
}
