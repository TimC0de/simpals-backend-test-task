package data_source

type DocumentTemp struct {
	Id         string     `json:"id,omitempty"`
	Categories Categories `json:"categories,omitempty"`
	Title      Title      `json:"title,omitempty"`
	Type       string     `json:"type,omitempty"`
	Posted     float32    `json:"posted,omitempty"`
}

type Categories struct {
	Subcategory string `json:"subcategory,omitempty"`
}

type Title struct {
	Ro string `json:"ro,omitempty"`
	Ru string `json:"ru,omitempty"`
}
