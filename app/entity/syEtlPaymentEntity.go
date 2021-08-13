package entity

import (
	"time"
)

type SyEtlPayment struct {
	ID                int           `json:"id" ;gorm:"AUTO_INCREMENT"`
	ProposalNumber    string        `json:"proposal_number"`
	PolicyNumber      string        `gorm:"column:policy_number;unique;" json:"policy_number"`
	PaidDate          time.Time     `json:"paid_date"`
	PaymentMethodName string        `json:"payment_method_name"`
	PolicyStatus      string        `json:"policy_status"`
	TotalPremium      float64       `json:"total_premium"`
	UpdatedAt         time.Time     `json:"updated_at"`
	OdsEtlPayment     OdsEtlPayment `gorm:"foreignkey:PolicyNumber;AssociationForeignKey:PolicyNumber"`
	// LumpSumPayment    LumpSumPayment `gorm:"foreignkey:PolicyNumber;AssociationForeignKey:PolicyNumber"`
	// UserPolicy        UserPolicy `gorm:"foreignkey:PolicyNumber;AssociationForeignKey:PolicyNumber"`
}

func (syEtlPayment *SyEtlPayment) TableName() string {
	return "etl_sy_payments"
}
