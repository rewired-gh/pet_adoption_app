package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rewired-gh/pet_adoption_app/db/sqlc"
	"github.com/rewired-gh/pet_adoption_app/util"
)

func getCategoryIDAnyhow(server *Server, ctx *gin.Context, param sqlc.GetCategoryIDParams) (id int32, err error) {
	id, err = server.store.GetCategoryID(ctx, param)
	if err != nil {
		if err == sql.ErrNoRows {
			categoryParam := sqlc.CreateCategoryParams{
				Species: param.Species,
				Color:   param.Color,
				Gender:  param.Gender,
			}

			createErr := server.store.CreateCategory(ctx, categoryParam)

			if createErr != nil {
				return
			}

			id, err = server.store.GetCategoryID(ctx, param)
		}
	}
	return
}

type createPetRequest struct {
	Nickname    string
	Uid         int32
	Description string
	Birthday    string
	Species     string
	Color       string
	Gender      string
}

func (server *Server) createPet(ctx *gin.Context) {
	var req createPetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	param := sqlc.GetCategoryIDParams{
		Species: req.Species,
		Color:   util.ToNullString(req.Color),
		Gender:  util.ToNullString(req.Gender),
	}

	categoryID, err := getCategoryIDAnyhow(server, ctx, param)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	birthday, err := time.Parse(time.RFC3339, req.Birthday)

	petParam := sqlc.CreatePetParams{
		Nickname:    req.Nickname,
		Uid:         req.Uid,
		CategoryID:  categoryID,
		Description: req.Description,
		Birthday: sql.NullTime{
			Valid: err == nil,
			Time:  birthday,
		},
	}

	err = server.store.CreatePet(ctx, petParam)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
}
