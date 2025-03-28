package utils

func AppendIfNotPresent(slice []string, item string) []string {
	for _, existingItem := range slice {
		if existingItem == item {
			// Item already exists, return the original slice
			return slice
		}
	}

	// Append if not found
	return append(slice, item)
}