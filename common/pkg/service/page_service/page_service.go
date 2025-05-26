package page_service

import "github.com/gin-gonic/gin"

type PageService struct {
	c     *gin.Context
	title string
}

const (
	Token = "page_service"
)

func New(c *gin.Context) *PageService {
	return &PageService{
		c: c,
	}
}

func (s *PageService) GetTitle() string {
	return s.title
}

func (s *PageService) SetTitle(title string) *PageService {
	s.title = title
	return s
}
