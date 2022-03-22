package utils

import "errors"

// import (
// 	  "cashmachine/Bank"
// )

func Calculate(sum uint) (out Payment, err error) {
	for sum > Banknotes[0] {
		out[Banknotes[0]] += 1
		sum -= Banknotes[0]
	}
	for sum > Banknotes[1] {
		out[Banknotes[1]] += 1
		sum -= Banknotes[1]
	}
	for sum > Banknotes[2] {
		out[Banknotes[2]] += 1
		sum -= Banknotes[2]
	}
	for sum > Banknotes[3] {
		out[Banknotes[3]] += 1
		sum -= Banknotes[3]
	}
	for sum > Banknotes[4] {
		out[Banknotes[4]] += 1
		sum -= Banknotes[4]
	}
	for sum > Banknotes[5] {
		out[Banknotes[5]] += 1
		sum -= Banknotes[5]
	}
	if sum > 0 {
		return nil, errors.New("invalid sum")
	}
	return
}
