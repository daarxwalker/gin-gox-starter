package gox_service

import (
	"github.com/daarxwalker/gox"
	"github.com/gin-gonic/gin"
)

type GoxService struct {
	c *gin.Context
}

func New(c *gin.Context) *GoxService {
	return &GoxService{
		c: c,
	}
}

func (s *GoxService) Render(status int, children ...gox.Node) {
	s.c.Header("Content-Type", "text/html; charset=utf-8")
	s.c.Status(status)
	s.c.String(status, gox.Render(children...))
}
