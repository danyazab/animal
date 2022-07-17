package request

type (
	CreateCatReq struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Breed       string  `json:"breed"`
		Birthday    string  `json:"birthday"`
		Sex         string  `json:"sex"`
		TailLength  uint    `json:"tailLength"`
		Color       string  `json:"color"`
		WoolType    string  `json:"woolType"`
		IsChipped   bool    `json:"isChipped"`
		Weight      float32 `json:"weight"`
	}
	CatListReq struct {
		PrePage uint `query:"perPage"`
	}
)
