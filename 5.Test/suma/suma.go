package suma

func Add(a, b int32) int32 {
	return a + b
}

func AddList(nums ...int) int {
	resp := 0

	for _, v := range nums {
		resp += v
	}

	return resp
}
