package calc

func Add(i ... int)int{
	summ := 0
	for _,v := range i {
		summ = summ + v
	}
	return summ
}
