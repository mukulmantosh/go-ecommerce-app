package common_errors

type CartEmptyError struct{}

func (e *CartEmptyError) Error() string {
	return "No items available in the cart"
}
