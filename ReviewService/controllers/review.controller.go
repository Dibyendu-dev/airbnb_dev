package controllers

import (
	"ReviewService/dto"
	"ReviewService/services"
	"ReviewService/utils"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ReviewController struct {
	ReviewService services.ReviewService
}

func NewReviewController(_reviewService services.ReviewService) *ReviewController {
	return &ReviewController{
		ReviewService: _reviewService,
	}
}

func (rc *ReviewController) GetReviewById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fetching review by id in reviewController")

	reviewId:= chi.URLParam(r, "id")
	if reviewId== ""{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest,"review id is required",fmt.Errorf("missing review id"))
		return
	}
	review,err := rc.ReviewService.GetReviewById(reviewId)
	if err != nil{
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError,"failed to fetch review",err)
		return
	}

	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Review fetched successfully",review)
	fmt.Println("Review fetched successfully")
}

func (rc *ReviewController) CreateReview(w http.ResponseWriter, r *http.Request) {

	payload := r.Context().Value("payload").(dto.CreateReviewRequestDTO)

	fmt.Println("Payload recived:",payload)

	review,err := rc.ReviewService.CreateReview(&payload)
	if err != nil{
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError,"failed to create review",err)
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Review created successfully",review)
	fmt.Println("Review created successfully")
}

func (rc *ReviewController) UpdateReview(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating review in ReviewController")

	reviewId:=chi.URLParam(r,"id")
	if reviewId== ""{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest,"review id is required",fmt.Errorf("missing review id"))
		return
	}

	payload := r.Context().Value("payload").(dto.UpdateReviewRequestDTO)

	fmt.Println("Payload recived:",payload)

	review,err := rc.ReviewService.UpdateReview(reviewId,&payload)
	if err != nil{
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError,"failed to update review",err)
		return
	}

	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Review update successfully",review)
	fmt.Println("Review update successfully")
}


func (rc *ReviewController) DeleteReview (w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleting review in ReviewController")

	reviewId:=chi.URLParam(r,"id")
	if reviewId== ""{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest,"review id is required",fmt.Errorf("missing review id"))
		return
	}

	err:= rc.ReviewService.DeleteReview(reviewId)
	if err != nil{
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError,"failed to delete review",err)
		return
	}

	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Review deleted successfully",nil)
	fmt.Println("Review deleted successfully")

}

func (rc *ReviewController) GetAllReview(w http.ResponseWriter, r *http.Request){
	fmt.Println("Fetching all reviews in ReviewController")

	reviews,err := rc.ReviewService.GetAllReviews()
	if err != nil{
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError,"failed to fetched review",err)
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Review fetched successfully",reviews)
	fmt.Println("Review fetched successfully, count:",len(reviews))
}


func (rc *ReviewController) GetReviewsByUserId(w http.ResponseWriter, r *http.Request){
	fmt.Println("Fetching reviews by user ID in ReviewController")

	userId:= r.URL.Query().Get("user_id")
	if userId== ""{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest,"user id is required",fmt.Errorf("missing user id"))
		return
	}

	reviews,err := rc.ReviewService.GetReviewsByUserId(userId)
	if err != nil{
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError,"failed to fetched review by user id",err)
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Review fetched successfully",reviews)
	fmt.Println("Review fetched successfully for userid:",userId, "count:",len(reviews))
}


func (rc *ReviewController) GetReviewsByHotelId(w http.ResponseWriter, r *http.Request){
	fmt.Println("Fetching reviews by hotel ID in ReviewController")

	hotelId:= r.URL.Query().Get("hotel_id")
	if hotelId== ""{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest,"hotel id is required",fmt.Errorf("missing hotel id"))
		return
	}

	reviews,err := rc.ReviewService.GetReviewsByHotelId(hotelId)
	if err != nil{
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError,"failed to fetched review by hotel id",err)
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Review fetched successfully",reviews)
	fmt.Println("Review fetched successfully for hotelid:",hotelId, "count:",len(reviews))
}


func (rc *ReviewController) GetReviewsByBookingId(w http.ResponseWriter, r *http.Request){
	fmt.Println("Fetching reviews by booking ID in ReviewController")

	bookingId:= r.URL.Query().Get("booking_id")
	if bookingId== ""{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest,"booking id is required",fmt.Errorf("missing booking id"))
		return
	}

	reviews,err := rc.ReviewService.GetReviewsByBookingId(bookingId)
	if err != nil{
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError,"failed to fetched review by booking id",err)
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Review fetched successfully",reviews)
	fmt.Println("Review fetched successfully for bookingid:",bookingId, "count:",len(reviews))
}
