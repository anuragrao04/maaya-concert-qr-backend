package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PRN       string `gorm:"index"`
	SRN       string
	Email     string
	Name      string
	Semester  string
	Branch    string
	IsPresent bool

	// only for outside pes people
	IsPesticide           bool   `json:"isPesticide"`
	PesticideReferralSRN  string `json:"pesticideReferralSRN"`
	PesticideReferralName string `json:"pesticideReferralName"`
}
