package chat

import "sort"

func uniqueSliceInt64(s []int64) []int64 {
	uniqueUserIDs := map[int64]struct{}{}

	for _, userID := range s {
		uniqueUserIDs[userID] = struct{}{}
	}

	result := make([]int64, 0, len(uniqueUserIDs))

	for userID := range uniqueUserIDs {
		result = append(result, userID)
	}

	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result
}
