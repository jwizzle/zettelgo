package util

func StringInSlice(substr string, slice []string) (bool) {
	for _, item := range slice {
		if item == substr {
			return true
		}
	}
	return false
}
