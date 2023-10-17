package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rewired-gh/pet_adoption_app/db/sqlc"
)

type createAdoptionRequest struct {
	PetID int32
	Uid   int32
}

func (server *Server) createAdoption(ctx *gin.Context) {
	var req createAdoptionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	param := sqlc.CreateAdoptionParams{
		Uid:   req.Uid,
		PetID: req.PetID,
	}

	err := server.store.CreateAdoption(ctx, param)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.String(http.StatusOK, "")
}

type listUserAdoptionRequest struct {
	Uid int32
}

type listUserAdoptionResponseEntry struct {
	PetID      int32
	Nickname   string
	Species    string
	IsReviewed bool
	IsApproved bool
}

func (server *Server) listUserAdoption(ctx *gin.Context) {
	var req listUserAdoptionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	adoptions, err := server.store.ListUserAdoption(ctx, req.Uid)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := make([]listUserAdoptionResponseEntry, len(adoptions))
	for i, adoption := range adoptions {
		res[i] = listUserAdoptionResponseEntry{
			PetID:      adoption.PetID,
			Nickname:   adoption.Nickname,
			Species:    adoption.Species,
			IsReviewed: adoption.IsReviewed,
			IsApproved: adoption.IsApproved,
		}
	}

	ctx.JSON(http.StatusOK, res)
}
