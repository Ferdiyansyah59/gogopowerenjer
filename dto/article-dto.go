package dto

import (
	"mime/multipart"
)

// Data Transfer Object Artikel
type ArticleUpdateDTO struct {
	// ID       uint64                `json:"id" form:"id"`
	Title    string                `json:"title" form:"title" binding:"required"`
	Slug     string                `json:"slug" form:"slug"`
	MetaDesc string                `json:"meta_desc" form:"meta_desc" binding:"required"`
	Content  string                `json:"content" form:"content" binding:"required"`
	Cover    *multipart.FileHeader `json:"cover" form:"cover" binding:"required"`
}

type ArticleCreateDTO struct {
	Title    string                `json:"title" form:"title" binding:"required"`
	Slug     string                `json:"slug" form:"slug"`
	MetaDesc string                `json:"meta_desc" form:"meta_desc" binding:"required"`
	Content  string                `json:"content" form:"content" binding:"required"`
	Cover    *multipart.FileHeader `json:"cover" form:"cover" binding:"required"`
}
