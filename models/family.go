package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Family struct {
	ID           uint    `json:"id" gorm:"primary_key"`
	Name         string  `json:"name" gorm:"type:varchar(255)"`
	Sex          string  `json:"sex" gorm:"type:varchar(255)"`
	DescendantID uint    `json:"descendant_id" gorm:"foreignKey:ID;association_foreignKey:DescendantID"`
	Assets       []Asset `json:"-" gorm:"foreignKey:FamilyID"`
}

func (f *Family) CreateFamily(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Create(&f).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *Family) UpdateFamily(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	var family Family
	err := db.Where("id = ?", f.ID).First(&family).Error
	if err != nil {
		return err
	}
	err = db.Model(&family).Updates(&f).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *Family) DeleteFamily(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	var family Family
	err := db.Where("id = ?", f.ID).First(&family).Error
	if err != nil {
		return err
	}
	err = db.Delete(&family).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *Family) GetFamilies(ctx *gin.Context) ([]Family, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var families []Family
	err := db.Find(&families).Error
	if err != nil {
		return nil, err
	}
	return families, nil
}

func FindFamilyByID(ctx *gin.Context, id uint) (Family, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var family Family
	err := db.Where("id = ?", id).First(&family).Error
	if err != nil {
		return family, err
	}
	return family, nil
}

type TotalPriceAssetByApiOutput struct {
	FamilyName string `json:"family_name"`
	TotalPrice int    `json:"total_price"`
}

type ResponseAPI struct {
	Products []Product `json:"products"`
}

type Product struct {
	Title string `json:"title"`
	Price string `json:"price"`
}

func (f *Family) TotalAssetPriceEachFamily(ctx *gin.Context) []TotalPriceAssetByApiOutput {
	var totalPriceAssetByApiOutput []TotalPriceAssetByApiOutput
	families, err := f.GetFamilies(ctx)
	if err != nil {
		return nil
	}
	for _, family := range families {
		totalPrice := 0
		assetTemp := Asset{FamilyID: family.ID}
		assetByFamily, _ := assetTemp.GetAssetByFamilyID(ctx)
		for _, asset := range assetByFamily {
			url := "https://dummyjson.com/products/search?q=" + asset.Name
			response, _ := http.Get(url)
			data, _ := ioutil.ReadAll(response.Body)
			var responseAPI ResponseAPI
			json.Unmarshal(data, &responseAPI)
			price, _ := strconv.Atoi(responseAPI.Products[0].Price)
			totalPrice += price
		}
		totalPriceAssetByApiOutput = append(totalPriceAssetByApiOutput, TotalPriceAssetByApiOutput{
			FamilyName: family.Name,
			TotalPrice: totalPrice,
		})
	}
	return totalPriceAssetByApiOutput
}
