package calc

func Add(i ... int)int, err{
	sum := 0
	for _,v := range i {
		sum = sum + v
	}
	return sum, nil
}
