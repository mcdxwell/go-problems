package adder


func Adder[V int | string](a, b V) V {

	return a + b

}

