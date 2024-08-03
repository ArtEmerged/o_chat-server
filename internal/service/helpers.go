package service

func uniqueSliceInt64(s []int64) []int64 {
	uniqueUserIDs := map[int64]struct{}{}

	for _, userID := range s {
		uniqueUserIDs[userID] = struct{}{}
	}

	result := make([]int64, 0, len(uniqueUserIDs))

	for userID := range uniqueUserIDs {
		result = append(result, userID)
	}
	return result
}
