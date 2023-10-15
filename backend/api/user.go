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

type uidResponse struct {
	Uid int32
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

	authParam := sqlc.AuthUserParams{
		Username: req.Username,
		Pword:    req.Pword,
	}

	uid, err := server.store.AuthUser(ctx, authParam)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, uidResponse{
		Uid: uid,
	})
}

type authUserRequest struct {
	Username string
	Pword    string
}

func (server *Server) authUser(ctx *gin.Context) {
	var req authUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	param := sqlc.AuthUserParams{
		Username: req.Username,
		Pword:    req.Pword,
	}

	uid, err := server.store.AuthUser(ctx, param)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, uidResponse{
		Uid: uid,
	})
}

type getUserRequest struct {
	Uid int32
}

type getUserResponse struct {
	Username string
	Gender   string
	Birthday time.Time
	Province string
	City     string
	District string
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Uid)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	city := user.City.String
	if !user.City.Valid {
		city = ""
	}

	ctx.JSON(http.StatusOK, getUserResponse{
		Username: user.Username,
		Gender:   user.Gender,
		Birthday: user.Birthday,
		Province: user.Province,
		City:     city,
		District: user.District,
	})
}

type addContactsRequest struct {
	Uid      int32
	Contacts []struct {
		Kind    string
		Content string
	}
}

func (server *Server) addContacts(ctx *gin.Context) {
	var req addContactsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	for _, contact := range req.Contacts {
		param := sqlc.AddContactParams{
			Uid:     req.Uid,
			Kind:    contact.Kind,
			Content: contact.Content,
		}

		err := server.store.AddContact(ctx, param)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	ctx.String(http.StatusOK, "")
}

type getRolesRequest struct {
	Uid int32
}

type getRolesResponse struct {
	Roles []string
}

func (server *Server) getRoles(ctx *gin.Context) {
	var req getRolesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	roles, err := server.store.GetRoles(ctx, req.Uid)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, getRolesResponse{
		Roles: roles,
	})
}
