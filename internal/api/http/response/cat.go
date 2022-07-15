package response

type CatResp struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Age   uint   `json:"age"`
	Breed string `json:"breed"`
	Sex   string `json:"sex"`
	Color string `json:"color"`
}

type CatRespList struct {
	Items []CatResp `json:"items"`
	Meta  MetaData  `json:"meta"`
}

type MetaData struct {
	Total uint64 `json:"total"`
}
