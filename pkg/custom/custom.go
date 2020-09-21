package custom

import (
	"context"
	"fmt"
	"net/http"
)

type customService struct{}

func NewService() Service { return &customService{} }

func (c *customService) TestStatus(_ context.Context) (int, error) {
	fmt.Println("Checking the custom service status...")
	return http.StatusOK, nil
}
