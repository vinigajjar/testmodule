package calc

func Add(i ... int)int{
	sum := 0
	for _,v := range i {
		sum = sum + v
	}
	return sum
}
