package api

import (
	"errors"
	"net/http"

	db "github.com/Golang/bank_management_system/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateTransferReq struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"account" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (Server *Server) CreateTransfer(ctx *gin.Context) {
	var req CreateTransferReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !Server.validateAccount(ctx, req.FromAccountID, req.Currency) {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("invalid fromAccID")))
		return
	}

	if !Server.validateAccount(ctx, req.ToAccountID, req.Currency) {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("invalid toAccID")))
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	transfer, err := Server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transfer)
}

func (Server *Server) validateAccount(ctx *gin.Context, accoountId int64, currency string) bool {
	account, err := Server.store.GetAccount(ctx, accoountId)
	if err != nil {
		return false
	}
	if account.Currency != currency {
		return false
	}
	return true
}
