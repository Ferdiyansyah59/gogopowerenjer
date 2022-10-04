package controller

import (
	"bensi-api/dto"
	"bensi-api/helper"
	"bensi-api/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
	GetAllUser(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["email"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}

	userUpdateDTO.ID = id
	u := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

// GET USER
func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["email"])
	user := c.userService.Profile(id)
	res := helper.BuildResponse(true, "OK!", user)
	context.JSON(http.StatusOK, res)
}

// GET ALL USER
func (c *userController) GetAllUser(context *gin.Context) {
	user := c.userService.GetAllUser()
	res := helper.BuildResponse(true, "OK!", user)
	context.JSON(http.StatusOK, res)
}
