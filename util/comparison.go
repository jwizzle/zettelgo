package util

func StringInSlice(substr string, slice []string) (bool) {
	for _, item := range slice {
		if item == substr {
			return true
		}
	}
	return false
}

func CompareSlicesLoose(x []string, y []string) (bool) {
	for _, i := range x {
		if ! StringInSlice(i, y) {
			return false
		}
	}
	return true
}


func MapToSlice(inputmap map[string]string) ([]string) {
	var out []string

	for _, v := range inputmap {
		out = append(out, v)
	}

	return out
}
