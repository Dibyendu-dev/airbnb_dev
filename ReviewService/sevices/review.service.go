package sevices

import (
	db "ReviewService/db/repository"
	"ReviewService/dto"
	"ReviewService/models"
	"fmt"
	"strconv"
)

type ReviewService interface {
	GetReviewById(id string) (*models.Review, error)
	CreateReview(payload *dto.CreateReviewRequestDTO) (*models.Review, error)
	UpdateReview(id string, payload *dto.UpdateReviewRequestDTO) (*models.Review, error)
	DeleteReview(id string) error
	GetAllReviews() ([]*models.Review, error)
	GetReviewsByUserId(userId string) ([]*models.Review, error)
	GetReviewsByHotelId(hotelId string) ([]*models.Review, error)
	GetReviewsByBookingId(bookingId string) ([]*models.Review, error)
}

type ReviewServiceImpl struct {
	reviewRepository db.ReviewRepository
}

func NewReviewService(_reviewRepository db.ReviewRepository) ReviewService {
	return &ReviewServiceImpl{
		reviewRepository: _reviewRepository,
	}
}

func (r *ReviewServiceImpl) GetReviewById(id string) (*models.Review, error) {

	fmt.Println("fetching review in review service")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error parsing review ID:", err)
		return nil, fmt.Errorf("invalid review ID")
	}
	review, err := r.reviewRepository.GetById(idInt)
	if err != nil {
		fmt.Println("Error fetching review:", err)
		return nil, err
	}
	return review, nil
}

func (r *ReviewServiceImpl) CreateReview(payload *dto.CreateReviewRequestDTO) (*models.Review, error) {
	fmt.Println("Creating review in ReviewService")

	//validate rating
	if payload.Rating < 1 || payload.Rating > 5 {
		return nil, fmt.Errorf("rating must be between 1 to 5")
	}

	review, err := r.reviewRepository.Create(payload.UserId, payload.BookingId, payload.HotelId, payload.Comment, payload.Rating)
	if err != nil {
		fmt.Println("Error creating review:", err)
		return nil, err
	}

	fmt.Println("Review created successfully:", review)
	return review, nil
}

func (r *ReviewServiceImpl) UpdateReview(id string, payload *dto.UpdateReviewRequestDTO) (*models.Review, error) {

	fmt.Println("Updating review in ReviewService")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error parsing review ID:", err)
		return nil, fmt.Errorf("invalid review ID")
	}

	//validate rating
	if payload.Rating < 1 || payload.Rating > 5 {
		return nil, fmt.Errorf("rating must be between 1 to 5")
	}
	review, err := r.reviewRepository.Update(idInt, payload.Comment, payload.Rating)
	if err != nil {
		fmt.Println("Error updating review:", err)
		return nil, err
	}

	fmt.Println("Review updating successfully:", review)
	return review, nil
}

func (r *ReviewServiceImpl) DeleteReview(id string) error {

	fmt.Println("Deleting review in ReviewService")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error parsing review ID:", err)
		return fmt.Errorf("invalid review ID")
	}

	delErr := r.reviewRepository.Delete(idInt)
	if delErr != nil {
		fmt.Println("Error deleting review:", delErr)
		return delErr
	}

	fmt.Println("Review deleting successfully")
	return nil

}

func (r *ReviewServiceImpl) GetAllReviews() ([]*models.Review, error) {

	fmt.Println("Fetching all reviews in ReviewService")

	reviews, err := r.reviewRepository.GetAll()
	if err != nil {
		fmt.Println("Error fetching review:", err)
		return nil, err
	}
	return reviews, nil
}

func (r *ReviewServiceImpl) GetReviewsByUserId(userId string) ([]*models.Review, error) {

	fmt.Println("Fetching reviews by user ID in ReviewService")

	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		fmt.Println("Error parsing user ID:", err)
		return nil, fmt.Errorf("invalid user ID")
	}
	reviews, err := r.reviewRepository.GetByUserId(userIdInt)
	if err != nil {
		fmt.Println("Error fetching reviews by user ID:", err)
		return nil, err
	}
	return reviews, nil
}

func (r *ReviewServiceImpl) GetReviewsByHotelId(hotelId string) ([]*models.Review, error) {

	fmt.Println("Fetching reviews by hotel ID in ReviewService")

	hotelIdInt, err := strconv.ParseInt(hotelId, 10, 64)
	if err != nil {
		fmt.Println("Error parsing hotel ID:", err)
		return nil, fmt.Errorf("invalid hotel ID")
	}
	reviews, err := r.reviewRepository.GetByUserId(hotelIdInt)
	if err != nil {
		fmt.Println("Error fetching reviews by hotel ID:", err)
		return nil, err
	}
	return reviews, nil
}

func (r *ReviewServiceImpl) GetReviewsByBookingId(bookingId string) ([]*models.Review, error) {

	fmt.Println("Fetching reviews by booking ID in ReviewService")

	bookingIdInt, err := strconv.ParseInt(bookingId, 10, 64)
	if err != nil {
		fmt.Println("Error parsing booking ID:", err)
		return nil, fmt.Errorf("invalid booking ID")
	}
	reviews, err := r.reviewRepository.GetByUserId(bookingIdInt)
	if err != nil {
		fmt.Println("Error fetching reviews by booking ID:", err)
		return nil, err
	}
	return reviews, nil
}
