package models

type Data struct {
	Id int64 `json:"id,omitempty"`

	Age int8 `json:"age,omitempty"`

	Name string `json:"name,omitempty"`

	Total float32 `json:"total,omitempty"`

	Verified bool `json:"verified,omitempty"`
}
