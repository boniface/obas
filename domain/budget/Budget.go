package budget

import (
	"math/big"
	"time"
)

type Awards struct {
	Id     string    `json:"id"`
	Email  string    `json:"email"`
	Amount big.Float `json:"amount"`
	Date   time.Time `json:"date"`
}
type Budget struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
type Disbusement struct {
	Id              string    `json:"id"`
	AwardId         string    `json:"awardId"`
	DisbusementType string    `json:"disbusementType"`
	Date            time.Time `json:"date"`
}
type DisbursementType struct {
	Id               string `json:"id"`
	DisbursementType string `json:"disbursement_type"`
	Description      string `json:"description"`
}
