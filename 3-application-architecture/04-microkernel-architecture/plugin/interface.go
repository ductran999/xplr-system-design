package plugin

type PaymentPlugin interface {
	Name() string
	Process(amount float64) bool
}
