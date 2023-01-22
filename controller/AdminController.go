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

type adminController struct {
	adminService service.AdminService
}

func NewAdminController(adminService service.AdminService) *adminController {
	return &adminController{adminService}
}

func (c *adminController) GetAdmins(cntx *gin.Context) {
	var admins, err = c.adminService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var adminsResponse []responses.AdminResponse

	for _, admin := range admins {
		var adminResponse = helper.ConvertToAdminResponse(admin)
		adminsResponse = append(adminsResponse, adminResponse)
	}

	cntx.JSON(http.StatusOK, adminsResponse)
}

func (c *adminController) GetAdmin(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var admin, err = c.adminService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var adminResponse = helper.ConvertToAdminResponse(admin)

	cntx.JSON(http.StatusOK, adminResponse)
}

func (c *adminController) GetAdminWithRelation(cntx *gin.Context) {
	var admins, err = c.adminService.FindAllRelation()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var adminsResponse []responses.AdminWithRelationResponse

	for _, admin := range admins {
		var adminResponse = helper.ConvertToAdminWithRelationResponse(admin)
		adminsResponse = append(adminsResponse, adminResponse)
	}

	cntx.JSON(http.StatusOK, adminsResponse)
}

func (c *adminController) CreateAdmin(cntx *gin.Context) {
	var adminRequest request.CreateAdminRequest

	var err = cntx.ShouldBindJSON(&adminRequest)
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

	admin, err := c.adminService.Create(adminRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var adminResponse = helper.ConvertToAdminResponse(admin)

	cntx.JSON(http.StatusCreated, gin.H{
		"data": adminResponse,
	})
}

func (c *adminController) UpdateAdmin(cntx *gin.Context) {
	var adminRequest request.UpdateAdminRequest

	var err = cntx.ShouldBindJSON(&adminRequest)
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

	admin, err := c.adminService.Update(id, adminRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var adminResponse = helper.ConvertToAdminResponse(admin)

	cntx.JSON(http.StatusOK, gin.H{
		"data": adminResponse,
	})
}

func (c *adminController) DeleteAdmin(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.adminService.Delete(id)
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
