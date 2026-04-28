package checkout

type Payment interface {
	Method(pays []int) (total int, paymentMethod string, err error)
}

func Checkout(payment Payment, pays []int) (total int, method string, err error) {
	total, method, err = payment.Method(pays)

	return total, method, err
}