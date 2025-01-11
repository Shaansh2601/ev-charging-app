package handler

import (
	"backend/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createPost(c *gin.Context) {
	var input model.Post
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	username, exists := c.Get("username")
	if !exists {
		newErrorResponse(c, http.StatusUnauthorized, "username not found")
		return
	}

	input.Username = username.(string)

	id, err := h.services.Chat.CreatePost(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":       id,
		"username": input.Username,
	})
}

func (h *Handler) getAllPosts(c *gin.Context) {

	posts, err := h.services.Chat.GetAllPosts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, posts)

}

func (h *Handler) createComment(c *gin.Context) {
	var input model.Comment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, 400, err.Error())
		return
	}

	input.PostID = id

	username, exists := c.Get("username")
	if !exists {
		newErrorResponse(c, http.StatusUnauthorized, "username not found")
		return
	}

	input.Username = username.(string)

	id, err = h.services.Chat.CreateComment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":       id,
		"username": input.Username,
	})

}

func (h *Handler) GetAllComments(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, 400, err.Error())
		return
	}

	comments, err := h.services.Chat.GetAllComments(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, comments)

}
