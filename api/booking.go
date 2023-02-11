package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/nga1hte/booking/db/sqlc"
	"github.com/nga1hte/booking/token"
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

	if !server.validAccount(ctx, req.BookedBy) {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.BookedBy != authPayload.Uid {
		err := errors.New("user not authorised")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
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

func (server *Server) validAccount(ctx *gin.Context, userID int64) bool {
	_, err := server.store.GetUser(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	return true
}

type getUserBookingsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getUserBookings(ctx *gin.Context) {
	var req getUserBookingsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.GetUserBookingsParams{
		BookedBy: authPayload.Uid,
		Limit:    req.PageSize,
		Offset:   (req.PageID - 1) * req.PageSize,
	}

	bookings, err := server.store.GetUserBookings(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, bookings)

}
