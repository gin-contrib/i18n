package utils

// IsIntSliceContains -- check slice contain string
func IsIntSliceContains(searchInt int, intSlice []int) bool {
	for _, value := range intSlice {
		if value == searchInt {
			return true
		}
	}
	return false
}

// IndexInSliceInt --
func IndexInSliceInt(num int, list []int) int {
	for index, v := range list {
		if v == num {
			return index
		}
	}
	return -1
}

// RemoveElementInSliceInt --
func RemoveElementInSliceInt(list []int, index int) []int {
	list = append(list[:index], list[index+1:]...)
	return list
}

// ToIntMap --
func ToIntMap(list []int) map[int]struct{} {
	m := make(map[int]struct{})
	for i := range list {
		m[list[i]] = struct{}{}
	}
	return m
}

// IntKeys --
func IntKeys(mmap map[int]int) []int {
	keys := make([]int, 0, len(mmap))
	for k := range mmap {
		keys = append(keys, k)
	}
	return keys
}
