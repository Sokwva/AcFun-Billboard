package common

//[1 2 3 4 5 6 7] 原
//[1 3 4 5 6 7 7] 处理后原来的
//[1 3 4 5 6 7] 返回的

//删除元素
func DelFromSliceDestory[T comparable](slc []T, item T) []T {
	i := 0
	for _, v := range slc {
		if v != item {
			slc[i] = v
			i++
		}
	}
	return slc[:i]
}

func DelFromSlice[T comparable](slc []T, item T) []T {
	i := -1
	for index, v := range slc {
		if v == item {
			i = index
			break
		}
	}
	if i == -1 {
		return slc
	}
	return append(slc[:i], slc[i+1:]...)
}
