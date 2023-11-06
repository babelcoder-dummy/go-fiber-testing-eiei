package model

import (
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Slug  string  `gorm:"uniqueIndex;not null"`
	Name  string  `gorm:"uniqueIndex;not null"`
	Desc  string  `gorm:"not null"`
	Price float64 `gorm:"not null"`
	Image string  `gorm:"not null"`
}

func (f *Product) BeforeSave(tx *gorm.DB) (err error) {
	f.Slug = regexp.MustCompile(" +").
		ReplaceAllString(strings.ToLower(f.Name), "-")

	return
}
