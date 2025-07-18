package models

type RentalReport struct {
	Total_created_rent   int64  `json:"total_created_rent"`
	Total_completed_rent int64  `json:"total_completed_rent"`
	Rents                []User `json:"rents"`
}

type SuccessRentalResponse = SuccessResponseWithMeta[RentalReport]
