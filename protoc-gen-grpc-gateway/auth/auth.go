package auth

import (
	"net/http"

	"golang.org/x/net/context"
)

type AuthHandler interface {
	CheckAuth(ctx context.Context, req *http.Request) (context.Context, error)
	SetHTTPError(ctx context.Context, w http.ResponseWriter, req *http.Request, err error)
}
