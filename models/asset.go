package models

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Asset struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
	FamilyID uint   `json:"family_id"`
}

func (a *Asset) CreateAsset(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	// check if family exists
	var family Family
	err := db.Where("id = ?", a.FamilyID).First(&family).Error
	if err != nil {
		return errors.New("family not found")
	}
	err = db.Create(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Asset) UpdateAsset(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	var asset Asset
	err := db.Where("id = ?", a.ID).First(&asset).Error
	if err != nil {
		return err
	}
	err = db.Model(&asset).Updates(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Asset) DeleteAsset(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	var asset Asset
	err := db.Where("id = ?", a.ID).First(&asset).Error
	if err != nil {
		return err
	}
	err = db.Delete(&asset).Error
	if err != nil {
		return errors.New("asset not found")
	}
	return nil
}

func (a *Asset) GetAssets(ctx *gin.Context) ([]Asset, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var assets []Asset
	assetsA := db.Find(&assets)
	fmt.Println(assetsA)
	err := db.Find(&assets).Error
	fmt.Println(assets)
	if err != nil {
		return nil, err
	}
	return assets, nil
}

func (a *Asset) GetAssetByFamilyID(ctx *gin.Context) ([]Asset, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var assets []Asset
	err := db.Where("family_id = ?", a.FamilyID).Find(&assets).Error
	if err != nil {
		return nil, err
	}
	return assets, nil
}
