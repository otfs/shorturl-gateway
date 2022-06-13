package main

import (
	"github.com/gin-gonic/gin"
)

// ShortUrl 短链实体
type ShortUrl struct {
	Id       int64  `db:"id"`        // 短链Id
	Url      string `db:"url"`       // 原始url地址
	ExpireAt int64  `db:"expire_at"` // 过期时间轴（毫秒）
}

// 短链跳转
func shortUrlJumpHandle(c *gin.Context) {
	slug := c.Param("slug")
	ids, err := hashIds.DecodeInt64WithError(slug)
	if err != nil || len(ids) == 0 {
		c.Status(404)
		return
	}
	short, err := getShortUrl(ids[0])
	if err != nil {
		c.Status(500)
		return
	}
	if short == nil {
		c.Status(404)
		return
	}

	c.Status(301)
	c.Header("Location", short.Url)
}

// 获取短链信息
func getShortUrl(id int64) (*ShortUrl, error) {
	short := new(ShortUrl)
	err := db.Get(short, "select id, url, expire_at from short_url where id = ?", id)
	if err != nil {
		return nil, err
	}
	return short, nil
}
