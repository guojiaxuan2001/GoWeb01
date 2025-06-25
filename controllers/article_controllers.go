package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"goweb01/global"
	"goweb01/models"
	"net/http"
)

func CreateArticle(ctx *gin.Context) {
	var article models.Article

	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := global.Db.AutoMigrate(&article); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err := global.Db.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusCreated, article)
}

func GetAllArticles(ctx *gin.Context) {
	var articles []models.Article

	if err := global.Db.Find(&articles).Error; err != nil { //相当于SELECT * FROM Article;
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, articles)
}

func GetArticleByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var article models.Article

	if err := global.Db.Where("id = ?", id).First(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		return
	}

	ctx.JSON(http.StatusOK, article)
}
