package models

type Post struct {
	Id         uint64 `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	AuthorId   uint64 `json:"authorid,omitempty"`
	AuthorName string `json:"authorname,omitempty"`
	Likes      uint64 `json:"likes"`
	RegDate    string `json:"regdate,omitempty"`
}
