package vo

import "fmt"

type Amount struct {
	Value float64 `json:"value"`
}

func NewAmount(value float64) Amount {
	return Amount{Value: value}
}
func (a *Amount) Add(amount Amount) {
	a.Value += amount.Value
}
func (a *Amount) Compare(amount Amount) bool {
	return a.Value == amount.Value
}
func (a *Amount) Subtract(amount Amount) {
	a.Value -= amount.Value
}
func (a *Amount) IsNegative() bool {
	return a.Value < 0
}
func (a *Amount) String() string {
	return fmt.Sprintf("%.2f", a.Value)
}
