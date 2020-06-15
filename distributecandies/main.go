package distributecandies

// 第一版
func distributeCandies(candies []int) int {
	sister := map[int]bool{}
	for _, candy := range candies {
		if _, exist := sister[candy]; !exist {
			sister[candy] = true
			if len(sister) == len(candies)/2 {
				break
			}
		}
	}
	return len(sister)
}

// 第二版
func distributeCandies2(candies []int) int {
	sister := map[int]bool{}
	for _, candy := range candies {
		if !sister[candy] {
			sister[candy] = true
			if len(sister) == len(candies)/2 {
				break
			}
		}
	}
	return len(sister)
}
