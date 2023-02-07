package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/nga1hte/booking/db/sqlc"
)

type createBookingRequest struct {
	BookedBy      int64     `json:"bookedBy" binding:"required,min=1"`
	BookingStarts time.Time `json:"bookStarts" binding:"required"`
	BookingEnds   time.Time `json:"bookEnds" binding:"required"`
}

func (server *Server) createBooking(ctx *gin.Context) {
	var req createBookingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.BookingTxParams{
		BookedBy:      req.BookedBy,
		BookingStarts: req.BookingStarts,
		BookingEnds:   req.BookingEnds,
	}

	booking, err := server.store.BookingTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, booking)
}
