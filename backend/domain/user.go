package domain

type RentalRepository interface {
	CreateNewRental(rentalPayload *User) error
	UpdateRental(rental *User) error
	DeleteRental(rentalId uint) error
	CompleteRental(rentalId uint) error
	GetRentalReport(page int, pageSize int, queryTerm string) (RentalReport, MetaModel, error)
	GetRentals(page int, pageSize int, queryTerm string) ([]User, MetaModel, error)
}

type (
	User struct {
		Base
		Full_name   string      `json:"full_name" validate:"required,min=3"`
		Address     string      `json:"address"`
		Pre_payment uint        `json:"pre_payment"`
		Active      bool        `gorm:"default:true" json:"active"`
		Phones      []string    `gorm:"serializer:json" json:"phones"`
		Date        string      `json:"date" validate:"required"`
		RentTools   []RentTools `json:"rent_tools" validate:"required"`
	}
)

type UpdateRentalResponse = SuccessResponseWithData[User]
type RentalListResponse = SuccessResponseWithMeta[[]User]
