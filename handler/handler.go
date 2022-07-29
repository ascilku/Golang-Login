package handler

import (
	"go-4/errors"
	"go-4/member"
	"go-4/respons"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handler struct {
	sevices member.Sevices
}

func NewHandler(service member.Sevices) *handler {
	return &handler{service}
}

func (h *handler) SaveHandler(g *gin.Context) {
	inputMember := member.InputMember{}
	err := g.ShouldBindJSON(&inputMember)
	if err != nil {
		// var errors []string
		// for _, e := range err.(validator.ValidationErrors) {
		// 	errors = append(errors, e.Error())
		// }

		errorValidation := errors.ErrorValidation(err)
		errorMessage := gin.H{"errors": errorValidation}

		APIRespons := respons.APIRespons("Failed Input", http.StatusOK, "Failed", errorMessage)
		g.JSON(http.StatusBadRequest, APIRespons)
	} else {
		keyHandler, err := h.sevices.SaveSevices(inputMember)
		if err != nil {
			g.JSON(http.StatusBadRequest, nil)
		} else {
			formatter := member.Formatter(keyHandler, "token token")
			APIRespons := respons.APIRespons("Success Input", http.StatusOK, "Success", formatter)
			g.JSON(http.StatusOK, APIRespons)
		}
	}
}

func (h *handler) SaveLogin(g *gin.Context) {
	var memberLogin member.LoginInput
	err := g.ShouldBindJSON(&memberLogin)

	if err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}

		errorMessage := gin.H{"errors": errors}
		newAPIRespons := respons.APIRespons("Login Gagal", http.StatusUnprocessableEntity, "Gagal", errorMessage)
		g.JSON(http.StatusUnprocessableEntity, newAPIRespons)
	} else {
		newLoginServis, err := h.sevices.LoginServis(memberLogin)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			newAPIRespons := respons.APIRespons("Login Gagal", http.StatusBadRequest, "Gagal", errorMessage)
			g.JSON(http.StatusBadRequest, newAPIRespons)
		} else {
			newFormatter := member.Formatter(newLoginServis, "token token ini")
			newAPIRespons := respons.APIRespons("Login Sukses", http.StatusOK, "Sukses", newFormatter)
			g.JSON(http.StatusOK, newAPIRespons)
		}
	}
}
