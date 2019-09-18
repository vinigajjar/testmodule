package calc

import "errors"

func Add(i ... int) (int, error){
	sum := 0
	if len(i) < 2 {
		errors.New("Less number of arguments passed")
	}
	for _,v := range i {
		sum = sum + v
	}
	return sum, nil
}
