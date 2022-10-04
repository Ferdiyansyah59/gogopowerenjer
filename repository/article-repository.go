package repository

import (
	"bensi-api/entity"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	GetGuestArticle() []entity.Article
	InsertArticle(article entity.Article) entity.Article
	UpdateArticle(article entity.Article, slug string) entity.Article
	DeleteArticle(article entity.Article, slug string)
	GetArticleBySlug(slug string) entity.Article
	GetImageData(slug string) string
	GetTitle(title string) string
}

type articleConnection struct {
	connection *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleConnection{
		connection: db,
	}
}

func (db *articleConnection) GetGuestArticle() []entity.Article {
	var articles []entity.Article
	db.connection.Find(&articles)
	return articles
}

func (db *articleConnection) InsertArticle(article entity.Article) entity.Article {
	db.connection.Save(&article)
	return article
}

func (db *articleConnection) UpdateArticle(article entity.Article, slug string) entity.Article {
	var articles entity.Article
	db.connection.Model(&articles).Where("slug = ?", slug).Updates(&article)
	return articles
}

func (db *articleConnection) DeleteArticle(article entity.Article, slug string) {
	db.connection.Where("slug = ?", slug).Delete(&article)
}

func (db *articleConnection) GetArticleBySlug(slug string) entity.Article {
	var articles entity.Article
	db.connection.Find(&articles, slug)
	return articles
}

func (db *articleConnection) GetImageData(slug string) string {
	var article entity.Article
	db.connection.Select("cover").Where("slug = ?", slug).First(&article)
	return article.Cover
}

func (db *articleConnection) GetTitle(title string) string {
	var article entity.Article
	db.connection.Select("title").Where("title = ?", title).First(&article)
	return article.Title
}
