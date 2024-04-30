package utils

func Rsshell2(arr []int) {
	size := len(arr)

	for interval := size / 2; interval > 0; interval /= 2 {
		for i := interval; i < size; i++ {
			temp := arr[i]
			j := i

			for ; j >= interval && arr[j-interval] < temp; j -= interval {
				arr[j] = arr[j-interval]
			}
			arr[j] = temp
		}
	}
}
