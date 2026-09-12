package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/BradleyZhang/ecommerce/user/usecase"

	"github.com/BradleyZhang/ecommerce/user/repository"

	"github.com/BradleyZhang/ecommerce/user/validater"

	"github.com/BradleyZhang/ecommerce/user/delivery/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	repo := repository.NewMemRepo()
	validater := validater.NewTodoValidater()
	usecase := usecase.NewUserUsecase(repo, validater)
	handler := rest.NewUserHandler(usecase)
	handler.SetRouter(server)

	srv := &http.Server{
		Addr:    ":1717",
		Handler: server,
	}

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Print("failed to start HTTP server: " + err.Error())
	}
	fmt.Print("server exited")
}
