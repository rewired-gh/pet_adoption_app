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
			categoryParam := sqlc.CreateCategoryParams(param)

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

	ctx.String(http.StatusOK, "")
}

func petRowsToResList(rows []sqlc.ListPetRow) (res []listPetResponseEntry) {
	res = make([]listPetResponseEntry, len(rows))
	for i, pet := range rows {
		birthdayStr := pet.Birthday.Time.Format(time.RFC3339)
		if !pet.Birthday.Valid {
			birthdayStr = ""
		}
		colorStr := pet.Color.String
		if !pet.Color.Valid {
			colorStr = ""
		}
		genderStr := pet.Gender.String
		if !pet.Gender.Valid {
			genderStr = ""
		}
		res[i] = listPetResponseEntry{
			PetID:       pet.PetID,
			Uid:         pet.Uid,
			Username:    pet.Username,
			CategoryID:  pet.CategoryID,
			Nickname:    pet.Nickname,
			Birthday:    birthdayStr,
			IsAdopted:   pet.IsAdopted,
			PublishDate: pet.PublishDate,
			Species:     pet.Species,
			Color:       colorStr,
			Gender:      genderStr,
			Description: pet.Description,
		}
	}
	return
}

type listPetRequest struct {
	Uid int32
}

type listPetResponseEntry struct {
	PetID       int32
	Uid         int32
	Username    string
	CategoryID  int32
	Nickname    string
	Birthday    string
	IsAdopted   bool
	PublishDate time.Time
	Species     string
	Color       string
	Gender      string
	Description string
}

func (server *Server) listPet(ctx *gin.Context) {
	var req listPetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	pets, err := server.store.ListPet(ctx, req.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := petRowsToResList(pets)

	ctx.JSON(http.StatusOK, res)
}

type listAvailablePetRequest struct {
	Uid int32
}

func (server *Server) listAvailablePet(ctx *gin.Context) {
	var req listAvailablePetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	pets, err := server.store.ListAvailablePet(ctx, req.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	petRows := make([]sqlc.ListPetRow, len(pets))
	for i, pet := range pets {
		petRows[i] = sqlc.ListPetRow(pet)
	}
	res := petRowsToResList(petRows)

	ctx.JSON(http.StatusOK, res)
}

type deletePetRequest struct {
	PetID int32
}

func (server *Server) deletePet(ctx *gin.Context) {
	var req deletePetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeletePet(ctx, req.PetID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.String(http.StatusOK, "")
}

type updatePetRequest struct {
	Nickname    string
	Description string
	Birthday    string
	Species     string
	Color       string
	Gender      string
	PetID       int32
}

func (server *Server) updatePet(ctx *gin.Context) {
	var req updatePetRequest
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

	petParam := sqlc.UpdatePetParams{
		PetID:       req.PetID,
		Nickname:    req.Nickname,
		CategoryID:  categoryID,
		Description: req.Description,
		Birthday: sql.NullTime{
			Valid: err == nil,
			Time:  birthday,
		},
	}

	err = server.store.UpdatePet(ctx, petParam)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.String(http.StatusOK, "")
}
