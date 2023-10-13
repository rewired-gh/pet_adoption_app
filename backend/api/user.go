package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rewired-gh/pet_adoption_app/db/sqlc"
)

type createUserRequest struct {
	LocationID int32
	Username   string
	Pword      string
	Gender     string
	Birthday   time.Time
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	param := sqlc.CreateUserParams{
		LocationID: req.LocationID,
		Username:   req.Username,
		Pword:      req.Pword,
		Gender:     req.Gender,
		Birthday:   req.Birthday,
	}

	err := server.store.CreateUser(ctx, param)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.String(http.StatusOK, "")
}

type getUserRequest struct {
	Username string
	Pword    string
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	param := sqlc.GetUserParams{
		Username: req.Username,
		Pword:    req.Pword,
	}

	user, err := server.store.GetUser(ctx, param)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
