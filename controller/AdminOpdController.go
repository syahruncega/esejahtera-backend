package controller

import (
	"fmt"
	"kemiskinan/helper"
	"kemiskinan/request"
	"kemiskinan/responses"
	"kemiskinan/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type adminOpdController struct {
	adminOpdService service.AdminOpdService
}

func NewAdminOpdController(adminOpdService service.AdminOpdService) *adminOpdController {
	return &adminOpdController{adminOpdService}
}

func (c *adminOpdController) GetAdminOpds(cntx *gin.Context) {
	var adminOpds, err = c.adminOpdService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var adminOpdsResponse []responses.AdminOpdResponse

	for _, adminOpd := range adminOpds {
		var adminOpdResponse = helper.ConvertToAdminOpdResponse(adminOpd)
		adminOpdsResponse = append(adminOpdsResponse, adminOpdResponse)
	}

	cntx.JSON(http.StatusOK, adminOpdsResponse)
}

func (c *adminOpdController) GetAdminOpd(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var adminOpd, err = c.adminOpdService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var adminOpdResponse = helper.ConvertToAdminOpdResponse(adminOpd)

	cntx.JSON(http.StatusOK, adminOpdResponse)
}

func (c *adminOpdController) GetAdminOpdByUserId(cntx *gin.Context) {
	var userIdString = cntx.Param("userid")
	var userId, _ = strconv.Atoi(userIdString)

	var adminOpd, err = c.adminOpdService.FindByUserId(userId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var adminOpdResponse = helper.ConvertToAdminOpdResponse(adminOpd)

	cntx.JSON(http.StatusOK, adminOpdResponse)
}

func (c *adminOpdController) CreateAdminOpd(cntx *gin.Context) {
	var adminOpdRequest request.CreateAdminOpdRequest

	var err = cntx.ShouldBindJSON(&adminOpdRequest)
	if err != nil {
		var errorMessages = []string{}

		for _, e := range err.(validator.ValidationErrors) {
			var errorMessage = fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	adminOpd, err := c.adminOpdService.Create(adminOpdRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var adminOpdResponse = helper.ConvertToAdminOpdResponse(adminOpd)

	cntx.JSON(http.StatusOK, adminOpdResponse)
}

func (c *adminOpdController) UpdateAdminOpd(cntx *gin.Context) {
	var adminOpdRequest request.UpdateAdminOpdRequest

	var err = cntx.ShouldBindJSON(&adminOpdRequest)
	if err != nil {
		var errorMessages = []string{}

		for _, e := range err.(validator.ValidationErrors) {
			var errorMessage = fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	adminOpd, err := c.adminOpdService.Update(id, adminOpdRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var adminOpdResponse = helper.ConvertToAdminOpdResponse(adminOpd)

	cntx.JSON(http.StatusOK, adminOpdResponse)
}

func (c *adminOpdController) DeleteAdminOpd(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.adminOpdService.Delete(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data gagal dihapus",
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"status": "Data berhasil dihapus",
	})
}
