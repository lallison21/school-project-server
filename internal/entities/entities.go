package entity

type Role struct {
	Id          int    `json:"id"`
	FullName    string `json:"full_name"`
	AccessLevel int    `json:"access_level"`
}
