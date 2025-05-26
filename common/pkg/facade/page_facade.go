package facade

import (
	"context"
	
	"common/pkg/service/page_service"
)

func Page(c context.Context) *page_service.PageService {
	service, ok := c.Value(page_service.Token).(*page_service.PageService)
	if !ok {
		panic(page_service.Token + " not found in context")
	}
	return service
}
