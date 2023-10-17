package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rewired-gh/pet_adoption_app/db/sqlc"
)

type applyRoleRequest struct {
	Uid         int32
	IsReviewer  bool
	IsUserAdmin bool
}

func (server *Server) applyRole(ctx *gin.Context) {
	var req applyRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var err error
	if req.IsReviewer || req.IsUserAdmin {
		err = server.store.ApplyAllRoles(ctx, sqlc.ApplyAllRolesParams{Uid: req.Uid})
	} else {
		err = server.store.ApplyUserRole(ctx, req.Uid)
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.String(http.StatusOK, "")
}
