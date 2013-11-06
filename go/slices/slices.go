package slices

func expandSlice(slice []int, data ...int) []int {
	currentLength := len(slice)
	neededLength := currentLength + len(data)

	if neededLength > cap(slice) {
		newSlice := make([]int, len(slice), (neededLength+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}

	view := slice[len(slice):cap(slice)]

	for i := range data {
		view[i] = data[i]
	}

	return slice[:neededLength]
}
