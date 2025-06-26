package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"goweb01/global"
	"net/http"
)

func LikeArticle(c *gin.Context) {
	articleID := c.Param("id")

	likeKey := "article:" + articleID + ":likes"

	if err := global.RedisDB.Incr(likeKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully liked article"})
}

func GetArticleLikes(c *gin.Context) {
	articleID := c.Param("id")

	likeKey := "article:" + articleID + ":likes"

	likes, err := global.RedisDB.Get(likeKey).Result()

	if errors.Is(err, redis.Nil) {
		likes = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
}
