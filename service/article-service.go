package service

import (
	"bensi-api/dto"
	"bensi-api/entity"
	"bensi-api/repository"
)

type ArticleService interface {
	GetGuestArticle() []entity.Article
	InsertArticle(art dto.ArticleCreateDTO, img string, slug string) entity.Article
	UpdateArticle(art dto.ArticleUpdateDTO, img string, slug string) entity.Article
	DeleteArticle(art entity.Article, slug string)
	GetArticleBySlug(slug string) entity.Article
	IsAllowedToEdit(articleID uint64) bool
	GetImageData(slug string) string
	GetTitle(title string) string
}

type articleService struct {
	articleRepository repository.ArticleRepository
}

func NewArticleService(artRepo repository.ArticleRepository) ArticleService {
	return &articleService{
		articleRepository: artRepo,
	}
}

func (serv *articleService) GetGuestArticle() []entity.Article {
	return serv.articleRepository.GetGuestArticle()
}

func (serv *articleService) InsertArticle(art dto.ArticleCreateDTO, img string, slug string) entity.Article {
	articles := entity.Article{
		Title:    art.Title,
		Slug:     slug,
		MetaDesc: art.MetaDesc,
		Content:  art.Content,
		Cover:    img,
	}

	res := serv.articleRepository.InsertArticle(articles)
	return res
}

func (serv *articleService) UpdateArticle(art dto.ArticleUpdateDTO, img string, slug string) entity.Article {
	articles := entity.Article{
		Title:    art.Title,
		Slug:     slug,
		MetaDesc: art.MetaDesc,
		Content:  art.Content,
		Cover:    img,
	}

	res := serv.articleRepository.UpdateArticle(articles, slug)
	return res
}

func (serv *articleService) DeleteArticle(art entity.Article, slug string) {
	serv.articleRepository.DeleteArticle(art, slug)
}

func (serv *articleService) GetArticleBySlug(slug string) entity.Article {
	var articles entity.Article
	return articles
}

func (serv *articleService) IsAllowedToEdit(articleID uint64) bool {
	return false
}

func (serv *articleService) GetImageData(slug string) string {
	return serv.articleRepository.GetImageData(slug)
}

func (serv *articleService) GetTitle(title string) string {
	return serv.articleRepository.GetTitle(title)
}
