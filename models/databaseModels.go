package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PRN       string `gorm:"index" json:"prn"`
	SRN       string `json:"srn"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Semester  string `json:"semester"`
	Branch    string `json:"branch"`
	IsPresent bool   `json:"isPresent"`

	// only for outside pes people
	IsPesticide           bool   `json:"isPesticide"`
	AadharNo              string `json:"aadharNo"`
	PesticideReferralSRN  string `json:"pesticideReferralSRN"`
	PesticideReferralName string `json:"pesticideReferralName"`
}
