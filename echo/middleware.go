package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// UserID ヘッダーのuserIDをcontextに埋め込むミドルウェア
func UserID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		h := c.Request().Header
		userIDstr := h.Get("X-User-Id")
		if userIDstr == "" {
			log.Println("there is no user id in header")
			return echo.NewHTTPError(http.StatusUnauthorized, "user id not found")
		}

		userID, err := strconv.ParseInt(userIDstr, 10, 64)
		if err != nil {
			log.Println("failed to parse int user id")
			return echo.NewHTTPError(http.StatusUnauthorized, "user id not found")
		}

		ctx := c.Request().Context()
		newCtx := WithUserID(ctx, userID)
		req := c.Request().WithContext(newCtx)
		c.SetRequest(req)
		return next(c)
	}
}

type ctxKey int

const CtxKeyUserID ctxKey = iota + 1

func WithUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, CtxKeyUserID, userID)
}
