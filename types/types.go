package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User struct
type User struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Loans map[string]Loan    `json:"loan,omitempty" bson:"loan,omitempty"`
}

//Loan is a single loan the user has
type Loan struct {
	APR            float64 `json:"apr" bson:"apr"`
	Balance        float64 `json:"balance" bson:"balance"`
	LengthInMonths float64 `json:"lengthInMonths" bson:"lengthInMonths"`
	ExtraPaidPri   float64 `json:"extraPaidPri" bson:"extraPaidPri"`
}
