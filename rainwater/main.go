package rainwater

import (
	"math"
)

func compute(arg []int) int {
	// その数の左右を見て、その数より大きい数が左右ともにあれば水はたまる
	// たまる量はその少ない方
	var totalRain int
	for val, i := range arg {
		// 左側の最大値取得
		// 右側の最大値取得
		// 小さい方とvalの差分が水の量
	}
	return 0
}

func getLeftMax(arr []int, index int) int {
	var max float64
	for i := 0; i < index; i++ {
		max = math.Max(float64(max), float64(arr[i]))
	}
	return int(max)
}
