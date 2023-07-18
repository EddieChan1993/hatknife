package hatknife

// SliceIsMember slice中是否包含该元素
func SliceIsMember(s []int32, e int32) bool {
	for _, i := range s {
		if e == i {
			return true
		}
	}
	return false
}

// SliceIsMember64 slice中是否包含该元素
func SliceIsMember64(s []int64, e int64) bool {
	for _, i := range s {
		if e == i {
			return true
		}
	}
	return false
}

// SliceCopyAppend slice拷贝追加
func SliceCopyAppend(base, add []int32) []int32 {
	res := make([]int32, 0, len(base)+len(add))
	res = append(res, base...)
	res = append(res, add...)
	return res
}

// SliceGet 获取切片元素
func SliceGet(index int32, base []int32) (int32, bool) {
	if index < 0 {
		return 0, false
	}
	if int32(len(base)) < index+1 {
		return 0, false
	}
	return base[index], true
}

// SliceGetSafe 安全获取切片元素,如果没找到就获取最后一个
func SliceGetSafe(index int32, base []int32) int32 {
	if index < 0 {
		return base[0]
	}
	if int32(len(base)) < index+1 {
		return base[len(base)-1]
	}
	return base[index]
}

// SliceIndexRangeScore 目标位于哪个索引，[0,20)
func SliceIndexRangeScore(target int32, slice []int32) int32 {
	if target <= slice[0] {
		//比最小的小
		return 0
	}
	if target >= slice[len(slice)-1] {
		//比最大的大
		return int32(len(slice) - 1)
	}
	for index, nums := range slice {
		if target <= nums {
			return int32(index)
		}
	}
	return int32(len(slice) - 1)
}

// SliceRem 删除切片元素
func SliceRem(arr []int32, elem int32) []int32 {
	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			arr = append(arr[:i], arr[i+1:]...)
			return arr
		}
	}
	return arr
}
