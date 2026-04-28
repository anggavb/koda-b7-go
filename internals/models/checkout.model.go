package models

import "errors"

type Bank struct {
	Name string
}
func (b Bank) Method(pays []int) (total int, paymentMethod string, err error) {
	total = 0
	for _, p := range pays {
		total += p
	}

	if total <= 0 {
		return 0, "", errors.New("Total payment must be at least 0")
	}

	return total, "bank", nil
}

type Online struct {
	Name string
}
func (o Online) Method(pays []int) (total int, paymentMethod string, err error) {
	total = 0
	for _, p := range pays {
		total += p
	}

	if total <= 0 {
		return 0, "", errors.New("Total payment must be at least 0")
	}

	return total, "online", nil
}