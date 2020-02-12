package types

import (
	"fmt"
	"math"

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

//CalculateASchedule calculates the amortization schedule
func (data *Loan) CalculateASchedule(schedChan chan AmortizationSched) {

	aSched := AmortizationSched{OriginalBalance: data.Balance}

	iRMonthly := (data.APR / 12) / 100
	TotalPaid := 0.00
	TotalPaidInterest := 0.00

	DF := (math.Pow((1+iRMonthly), data.LengthInMonths) - 1) / (iRMonthly * math.Pow((1+iRMonthly), data.LengthInMonths))
	fmt.Println(DF)
	MonthlyPayment := math.Round(data.Balance / DF)
	aSched.MonthlyPayment = MonthlyPayment

	//total interest paid if no extra principle is paid
	ti := (data.LengthInMonths * MonthlyPayment) - data.Balance
	aSched.InterestPaidOrig = ti

	balance := data.Balance
	Aint := make([]float64, int(data.LengthInMonths)+1)
	for i := 1; true; i++ {
		MonthlyInterest := iRMonthly * balance
		PriPaid := MonthlyPayment - MonthlyInterest
		TotalPriPaid := PriPaid + data.ExtraPaidPri
		balance = balance - TotalPriPaid
		Aint[i-1] = MonthlyInterest
		if balance < 1 {
			break
		}
	}

	for i := 1; true; i++ {
		//MonthlyInterest := iRMonthly * data.Balance
		PriPaid := MonthlyPayment - Aint[i-1] //MonthlyInterest
		TotalPriPaid := PriPaid + data.ExtraPaidPri
		data.Balance = data.Balance - TotalPriPaid

		monthB := MonthBreakdown{
			RemainingBalance: data.Balance,
			InterestPaid:     Aint[i-1],
			PrinciplePaid:    TotalPriPaid,
			TotalPayment:     TotalPriPaid + Aint[i-1],
		}

		aSched.MonthBreakdown = append(aSched.MonthBreakdown, monthB)

		TotalPaid = TotalPaid + TotalPriPaid + Aint[i-1]  //MonthlyInterest
		TotalPaidInterest = TotalPaidInterest + Aint[i-1] //MonthlyInterest
		if data.Balance < 1 {
			aSched.InterestPaidAdj = TotalPaidInterest
			aSched.InterestSaved = ti - TotalPaidInterest
			aSched.TotalPayment = TotalPaid
			break
		}
	}
	schedChan <- aSched
}

//AmortizationSched represents the Amortization Schedule data structure for the customer's loan
type AmortizationSched struct {
	OriginalBalance  float64          `json:"originalBalance"`
	MonthlyPayment   float64          `json:"monthlyPayment"`
	InterestPaidOrig float64          `json:"interestPaidOrig"`
	InterestPaidAdj  float64          `json:"interestPaidAdj"`
	InterestSaved    float64          `json:"interestSaved"`
	TotalPayment     float64          `json:"totalPayment"`
	MonthBreakdown   []MonthBreakdown `json:"aSchedTable"`
}

//MonthBreakdown is the breakdown of the loan details for the entire month
type MonthBreakdown struct {
	RemainingBalance float64 `json:"remainingBalance"`
	InterestPaid     float64 `json:"interestPaid"`
	PrinciplePaid    float64 `json:"principlePaid"`
	TotalPayment     float64 `json:"totalPayment"`
}
