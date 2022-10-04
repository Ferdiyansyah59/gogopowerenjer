package controller

import (
	"bensi-api/dto"
	"bensi-api/entity"
	"bensi-api/helper"
	"bensi-api/service"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type ArticleContorller interface {
	GetAllArticle(ctx *gin.Context)
	InsertArticle(ctx *gin.Context)
	DeleteArticle(ctx *gin.Context)
	UpdateArticle(ctx *gin.Context)
}

type articleController struct {
	articleService service.ArticleService
	jwtService     service.JWTService
}

func NewArticleController(art service.ArticleService, jwt service.JWTService) ArticleContorller {
	return &articleController{
		articleService: art,
		jwtService:     jwt,
	}
}

func (c *articleController) GetAllArticle(ctx *gin.Context) {
	articles := c.articleService.GetGuestArticle()
	res := helper.BuildResponse(true, "Success", articles)
	ctx.JSON(http.StatusOK, res)
}

func (c *articleController) InsertArticle(ctx *gin.Context) {
	var artDTO dto.ArticleCreateDTO
	errDTO := ctx.ShouldBind(&artDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	file, err := ctx.FormFile("cover")
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension

	if err := ctx.SaveUploadedFile(file, "C:/xampp/htdocs/Cordova/bensi/controller/asset/"+newFileName); err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	// getTitle := c.articleService.GetTitle(artDTO.Title)
	slugTxt := slug.Make(artDTO.Title) + strconv.FormatInt(int64(rand.Intn(1000)), 10)
	art := c.articleService.InsertArticle(artDTO, newFileName, slugTxt)
	res := helper.BuildResponse(true, "Success", art)
	ctx.JSON(http.StatusOK, res)
	// if artDTO.Title == getTitle {
	// 	if errDTO != nil {
	// 		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
	// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	// 		return
	// 	}

	// } else {
	// 	if errDTO != nil {
	// 		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
	// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	// 		return
	// 	}
	// 	slugTxt := slug.Make(artDTO.Title)
	// 	art := c.articleService.InsertArticle(artDTO, newFileName, slugTxt)
	// 	res := helper.BuildResponse(true, "Success", art)
	// 	ctx.JSON(http.StatusOK, res)
	// }

}

func (c *articleController) DeleteArticle(ctx *gin.Context) {
	var articles entity.Article
	slugUrl := ctx.Param("slug")
	articles.Slug = slugUrl

	imgData := c.articleService.GetImageData(articles.Slug)
	e := os.Remove("C:/xampp/htdocs/Cordova/bensi/controller/asset/" + imgData)
	if e != nil {
		log.Fatal(e)
	}

	c.articleService.DeleteArticle(articles, slugUrl)
	res := helper.BuildResponse(true, "Deleted Success", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (c *articleController) UpdateArticle(ctx *gin.Context) {
	var articles entity.Article
	var artDTO dto.ArticleUpdateDTO
	errDTO := ctx.ShouldBind(&artDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	slugUrl := ctx.Param("slug")
	articles.Slug = slugUrl

	imgData := c.articleService.GetImageData(articles.Slug)
	e := os.Remove("C:/xampp/htdocs/Cordova/bensi/controller/asset/" + imgData)
	if e != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	file, err := ctx.FormFile("cover")
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension

	if err := ctx.SaveUploadedFile(file, "C:/xampp/htdocs/Cordova/bensi/controller/asset/"+newFileName); err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	slugTxt := slug.Make(artDTO.Title) + strconv.FormatInt(int64(rand.Intn(1000)), 10)
	art := c.articleService.UpdateArticle(artDTO, newFileName, slugTxt)
	res := helper.BuildResponse(true, "Success", art)
	ctx.JSON(http.StatusOK, res)
	// getTitle := c.articleService.GetTitle(artDTO.Title)

	// if artDTO.Title == getTitle {

	// } else {
	// 	if errDTO != nil {
	// 		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
	// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	// 		return
	// 	}
	// 	slugTxt := slug.Make(artDTO.Title)
	// 	art := c.articleService.UpdateArticle(artDTO, newFileName, slugTxt)
	// 	res := helper.BuildResponse(true, "Success", art)
	// 	ctx.JSON(http.StatusOK, res)
	// }
}
