package model

type GitlabProject struct {
	ID   int    `json:"id" gorm:"not null;unique;primary_key"`
	Name string `json:"name" gorm:"not null;unique"`
	Url  string `gorm:"type:longText" json:"url"`
}
