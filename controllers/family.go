package controllers

import (
	"github.com/MicBun/TestJavan/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CreateFamilyInput struct {
	Name         string `json:"name" binding:"required"`
	Sex          string `json:"sex" binding:"required"`
	DescendantID uint   `json:"descendant_id"`
}

type UpdateFamilyInput struct {
	Name         string `json:"name"`
	Sex          string `json:"sex" binding:"required"`
	DescendantID uint   `json:"descendant_id"`
}

// CreateFamily godoc
// @Summary Create a family
// @Description Create a family
// @Tags Family
// @Accept  json
// @Produce  json
// @Param  family body CreateFamilyInput  true "Family"
// @Success 200 {object} models.Family
// @Failure 400 {object} string
// @Router /families [post]
func CreateFamily(ctx *gin.Context) {
	var input CreateFamilyInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	family := models.Family{Name: input.Name, Sex: input.Sex, DescendantID: input.DescendantID}
	err := family.CreateFamily(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": family})
}

// UpdateFamily godoc
// @Summary Update a family
// @Description Update a family
// @Tags Family
// @Accept  json
// @Produce  json
// @Param  family body UpdateFamilyInput true "Family"
// @Success 200 {object} models.Family
// @Failure 400 {object} string
// @Router /families/{id} [put]
func UpdateFamily(ctx *gin.Context) {
	var input UpdateFamilyInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	family := models.Family{ID: uint(id), Name: input.Name, Sex: input.Sex, DescendantID: input.DescendantID}
	err := family.UpdateFamily(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": family})
}

// DeleteFamily godoc
// @Summary Delete a family
// @Description Delete a family
// @Tags Family
// @Accept  json
// @Produce  json
// @Param id path int true "Family ID"
// @Success 200 {object} models.Family
// @Failure 400 {object} string
// @Router /families/{id} [delete]
func DeleteFamily(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	family := models.Family{ID: uint(id)}
	err := family.DeleteFamily(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

// TotalPriceAssetByApi godoc
// @Summary Get total price asset by api
// @Description Get total price asset by api
// @Tags Family
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Family
// @Router /families/totalPriceAssetByApi [get]
func TotalPriceAssetByApi(ctx *gin.Context) {
	family := models.Family{}
	totalPrice := family.TotalAssetPriceEachFamily(ctx)
	ctx.JSON(http.StatusOK, gin.H{"data": totalPrice})
}

// GetFamilies godoc
// @Summary Get all families
// @Description Get all families
// @Tags Family
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Family
// @Router /families [get]
func GetFamilies(ctx *gin.Context) {
	families := models.Family{}
	familiesRes, _ := families.GetFamilies(ctx)
	ctx.JSON(http.StatusOK, gin.H{"data": familiesRes})
}
