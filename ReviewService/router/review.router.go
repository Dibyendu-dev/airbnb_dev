package router

import (
	"ReviewService/controllers"
	"ReviewService/middlewares"

	"github.com/go-chi/chi/v5"
)

type ReviewRouter struct {
	reviewController *controllers.ReviewController
}

func NewReviewRouter(_reviewController *controllers.ReviewController) Router{
	return &ReviewRouter{
		reviewController: _reviewController,
	}
}

func (rr *ReviewRouter) Register(r chi.Router){

	//CRUD -fn
	r.With(middlewares.ReviewCreateRequestValidator).Post("/review",rr.reviewController.CreateReview)
	r.Get("/reviews",rr.reviewController.GetAllReview)
	r.Get("/review/{id}",rr.reviewController.GetReviewById)
	r.With(middlewares.ReviewUpdateRequestValidator).Put("/review/{id}",rr.reviewController.UpdateReview)
	r.Delete("/review/{id}",rr.reviewController.DeleteReview)

	// filter-op  -- send data in query params
	r.Get("/review/user",rr.reviewController.GetReviewsByUserId)
	r.Get("/review/hotel",rr.reviewController.GetReviewsByHotelId)
	r.Get("/review/booking",rr.reviewController.GetReviewsByBookingId)

}

