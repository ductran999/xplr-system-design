package model

type Order struct {
	ID          int
	RawAmount   float64
	FinalAmount float64
	IsValid     bool
}
