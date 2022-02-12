package handlers

import (
	"forum/internal/router"
	"net/http"
	"strconv"
)

const sessionHeaderKey = "session"

// AuthMiddleware ..
func (f *Forum) AuthMiddleware(next router.Handler) router.Handler {
	return func(ctx *router.Context) {
		session := ctx.Request.Header.Get(sessionHeaderKey)
		userID, err := f.Service.GetUserSession(session)
		if err != nil {
			ctx.WriteError(http.StatusUnauthorized)
			return
		}
		ctx.SetParam("userID", strconv.Itoa(userID))
		next(ctx)
	}
}