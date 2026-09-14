package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/BradleyZhang/ecommerce/user/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase *usecase.UserUsecase
}

func NewUserHandler(u *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: u,
	}
}

func (h *UserHandler) SetRouter(router *gin.Engine) {
	apiRouter := router.Group("api")
	{
		userRoute := apiRouter.Group("/user")
		{
			userRoute.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "user Server is live",
				})
			})

			userRoute.POST("/register", h.Register)
		}
	}
}

type RegistRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var request RegistRequest
	err := decodeJson(c.Request.Body, &request)
	if err != nil {
		apiError(c, "invalid params")
		return
	}
	if err := h.UserUsecase.Registration(request.Username, request.Password); err != nil {
		apiError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
}

func decodeJson(reader io.Reader, v any) error {
	return json.NewDecoder(reader).Decode(v)
}

func apiError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"message": msg,
	})
}
