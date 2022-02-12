package util

func String_in_slice(substr string, slice []string) (bool) {
	for _, item := range slice {
		if item == substr {
			return true
		}
	}
	return false
}
