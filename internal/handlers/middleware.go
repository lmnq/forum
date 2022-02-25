package handlers

import (
	"fmt"
	"forum/internal/router"
	"net/http"
	"strconv"
)

const sessionHeaderKey = "session"

// AuthMiddleware ..
func (f *Forum) AuthMiddleware(next router.Handler, block bool) router.Handler {
	return func(ctx *router.Context) {
		session, err := ctx.Request.Cookie(sessionHeaderKey)
		if err != nil {
			if block {
				ctx.WriteError(http.StatusUnauthorized)
				return
			}
			next(ctx)
			return
		}
		fmt.Println("session:", session)
		userID, err := f.Service.GetUserSession(session.Value)
		if err != nil && block {
			ctx.WriteError(http.StatusUnauthorized)
			return
		}
		ctx.SetParam("userID", strconv.Itoa(userID))
		next(ctx)
	}
}
