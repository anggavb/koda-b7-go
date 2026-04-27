package models

type Biodata struct {
	Name string
	Photo string
	Email string
	Age uint
	Telp string
	IsMarriage bool
	Education []Education
}