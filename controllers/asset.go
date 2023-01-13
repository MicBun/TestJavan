package controllers

import (
	"github.com/MicBun/TestJavan/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CreateAssetInput struct {
	Name     string `json:"name" binding:"required"`
	FamilyID uint   `json:"family_id" binding:"required"`
}

type UpdateAssetInput struct {
	Name     string `json:"name"`
	FamilyID uint   `json:"family_id" binding:"required"`
}

// CreateAsset godoc
// @Summary Create a asset
// @Description Create a asset
// @Tags Asset
// @Accept  json
// @Produce  json
// @Param  asset body CreateAssetInput  true "Asset"
// @Success 200 {object} models.Asset
// @Failure 400 {object} string
// @Router /assets [post]
func CreateAsset(ctx *gin.Context) {
	var input CreateAssetInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	asset := models.Asset{Name: input.Name, FamilyID: input.FamilyID}
	err := asset.CreateAsset(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": asset})
}

// UpdateAsset godoc
// @Summary Update a asset
// @Description Update a asset
// @Tags Asset
// @Accept  json
// @Produce  json
// @Param id path int true "Asset ID"
// @Param asset body UpdateAssetInput true "Asset"
// @Success 200 {object} models.Asset
// @Failure 400 {object} string
// @Router /assets/{id} [put]
func UpdateAsset(ctx *gin.Context) {
	var input UpdateAssetInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	asset := models.Asset{ID: uint(id), Name: input.Name, FamilyID: input.FamilyID}
	err := asset.UpdateAsset(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": asset})
}

// DeleteAsset godoc
// @Summary Delete a asset
// @Description Delete a asset
// @Tags Asset
// @Accept  json
// @Produce  json
// @Param id path int true "Asset ID"
// @Success 200 {object} models.Asset
// @Failure 400 {object} string
// @Router /assets/{id} [delete]
func DeleteAsset(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	asset := models.Asset{ID: uint(id)}
	err := asset.DeleteAsset(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": asset})
}

type TotalPriceAssetByApiOutput struct {
	FamilyName string `json:"family_name"`
	TotalPrice int    `json:"total_price"`
}

// GetAssets godoc
// @Summary Get all assets
// @Description Get all assets
// @Tags Asset
// @Accept  json
// @Produce  json
// @Router /assets [get]
func GetAssets(ctx *gin.Context) {
	var asset models.Asset
	assets, _ := asset.GetAssets(ctx)
	ctx.JSON(http.StatusOK, gin.H{"data": assets})
}
