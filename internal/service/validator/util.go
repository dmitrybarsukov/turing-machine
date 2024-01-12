package validator

func compare(value1, value2 int) Compare {
	if value1 < value2 {
		return Less
	} else if value1 > value2 {
		return More
	} else {
		return Equal
	}
}

func getOrderStrict(arr []int) Order {
	var isAsc, isDesc = true, true
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] <= arr[i+1] {
			isDesc = false
		}

		if arr[i] >= arr[i+1] {
			isAsc = false
		}
	}

	if isAsc {
		return Ascending
	}

	if isDesc {
		return Descending
	}

	return None
}

func getParity(value int) Parity {
	return []Parity{Even, Odd}[value%2]
}
