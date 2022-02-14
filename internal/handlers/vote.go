package handlers

import (
	"fmt"
	"forum/internal/router"
	"net/http"
	"strconv"
)

// VotePostHandler ..
func (f *Forum) VotePostHandler(ctx *router.Context) {
	postID, _ := strconv.Atoi(ctx.Params["postID"])
	userID, _ := strconv.Atoi(ctx.Params["userID"])
	rate, err := strconv.Atoi(ctx.Request.FormValue("rate"))

	if err != nil || (rate != 1 && rate != -1) {
		ctx.WriteError(http.StatusBadRequest)
		return
	}

	if err := f.Service.VotePost(postID, userID, rate); err != nil {
		ctx.WriteError(http.StatusBadRequest)
		return
	}

	http.Redirect(ctx.ResponseWriter, ctx.Request, fmt.Sprintf("/post/%d", postID), 302)
}

// VoteCommentHandler ..
func (f *Forum) VoteCommentHandler(ctx *router.Context) {
	commentID, _ := strconv.Atoi(ctx.Params["commentID"])
	userID, _ := strconv.Atoi(ctx.Params["userID"])
	postID, _ := strconv.Atoi(ctx.Params["postID"])
	rate, err := strconv.Atoi(ctx.Request.FormValue("rate"))

	if err != nil || (rate != 1 && rate != -1) {
		ctx.WriteError(http.StatusBadRequest)
		return
	}

	if err := f.Service.VoteComment(commentID, userID, rate); err != nil {
		ctx.WriteError(http.StatusBadRequest)
		return
	}

	http.Redirect(ctx.ResponseWriter, ctx.Request, fmt.Sprintf("/post/%d", postID), 302)
}
