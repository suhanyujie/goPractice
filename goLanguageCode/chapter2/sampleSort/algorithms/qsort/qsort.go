package qsort

func QuickSort(values []int) (err error) {
	left := 0
	right := len(values) - 1
	quikSort(values, left, right)
	return nil
}

func quikSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right
	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}
		if values[i] <= temp && i <= p {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quikSort(values, left, p-1)
	}
	if right-p > 1 {
		quikSort(values, p+1, right)
	}
}

/**
## 快速排序


## 参考
* 百度百科-快速排序算法 https://baike.baidu.com/item/%E5%BF%AB%E9%80%9F%E6%8E%92%E5%BA%8F%E7%AE%97%E6%B3%95/369842?fromtitle=%E5%BF%AB%E9%80%9F%E6%8E%92%E5%BA%8F&fromid=2084344&fr=aladdin#2


 */
