package api

import (
	"net/http"
	db "tesfayprep/simplebank/db/sqlc"
	"tesfayprep/simplebank/util"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedpass, err := util.HashedPassword(req.Password)
	if err != nil {

	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedpass,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	account, err := server.store.Createaccount(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
