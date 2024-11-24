package utils

func SliceDiff(slice1, slice2 []string) (diff []string) {
	for _, val1 := range slice1 {
		found := false

		for _, elem2 := range slice2 {
			if val1 == elem2 {
				found = true
				break
			}
		}
		
		if !found {
			diff = append(diff, val1)
		}
	}
	return diff
}
