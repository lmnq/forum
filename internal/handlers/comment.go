package handlers

import (
	"forum/internal/app"
	"forum/internal/router"
	"log"
	"net/http"
	"strconv"
)

// CommentPostHandler ..
func (f *Forum) CommentPostHandler(ctx *router.Context) {
	postID, _ := strconv.Atoi(ctx.Params["postID"])
	userID, _ := strconv.Atoi(ctx.Params["userID"])
	comment := app.Comment {
		PostID:   postID,
		AuthorID: userID,
		Content:  ctx.Request.FormValue("comment"),
	}
	if err := f.Service.ValidateComment(comment); err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest)
		return
	}
	if err := f.Service.AddNewCommentToPost(comment); err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest)
		return
	}
	ctx.ResponseWriter.WriteHeader(http.StatusCreated)
}
