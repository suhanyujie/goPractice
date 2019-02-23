package bubble

/**
冒泡排序算法
 */
func BubbleSort(values []int) error {
	length := len(values)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if values[j] < values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
	return nil
}
