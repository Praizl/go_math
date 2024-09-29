package geometry

import "errors"

func Cube(side int) (int, error) {
	if side != 0 {
		return side * side * side, nil
	} else {
		return 0, errors.New("cube side dimension invalid")

	}
}
