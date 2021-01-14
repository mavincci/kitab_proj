package model

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model

	AuthorID		uint
	User		User `gorm:"foreignkey:AuthorID"`

	Title		string
	Price		float32

	Download	int32

	Rate		int32	// No of rates
	AvgRating	float32

	Thumbnail	string
	Description	string
	Location	string
}
