package models

type Unique struct {
	Id int64 `json:"id,omitempty"`

	Age int8 `json:"age,omitempty"`

	Unique string `json:"unique,omitempty"`

	Verified bool `json:"verified,omitempty"`
}
