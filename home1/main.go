package main

import "errors"
import "fmt"

/*
*
实现删除切片特定下标元素的方法。
要求一：能够实现删除操作就可以。
要求二：考虑使用比较高性能的实现。
要求三：改造为泛型方法
要求四：支持缩容，并旦设计缩容机制。
*/
func main() {
	slice, err := DeleteSlice(2, []string{"1", "w4e3", "4355", "abcad1"})
	fmt.Println(slice, err)

	lv1Slice := []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff"}
	slice1, err1 := DeleteSliceLv1(3, 4, lv1Slice)
	fmt.Println(slice1, err1)

	//lv2Slice := []int{1, 2, 3, 4, 5, 6}
	lv2Slice1 := []string{"aa", "bb", "cc", "dd"}
	res, err := DeleteSliceLv2(2, 3, lv2Slice1)
	fmt.Println(res, err)
}

// DeleteSlice 要求一：能够实现删除操作就可以。
func DeleteSlice(num int, str []string) ([]string, error) {
	//获取切片长度
	len := len(str)
	//如果切片是0，直接返回空切片
	if len == 0 {
		return []string{}, errors.New("长度不能为空")
	}
	//如果下标长度超过切片长度，返回错误
	if num < 1 || num > len {
		return nil, errors.New("超出范围")
	}
	//拼接
	start := num - 1
	newStr := append(str[:start], str[start+1:]...)
	return newStr, nil
}

// DeleteSliceLv1 优化版本
func DeleteSliceLv1(num int, end int, str []string) ([]string, error) {
	// 获取切片长度
	len := len(str)
	// 如果切片是空的，直接返回空切片
	if len == 0 {
		return []string{}, errors.New("长度不能为空")
	}
	//如果超出范围就返回错误
	if num < 1 || num > len || end < num || end > len {
		return nil, errors.New("超出范围")
	}
	//重置初始化下标
	start := num - 1
	//计算要删除的下标
	switch {
	case start == 0 && end == len: //整个切片就直接返回空
		return []string{}, nil
	case start == 0:
		return str[end:], nil
	case end == len:
		return str[start:], nil
	default: //使用copy元素，不会改变切片的容量。用copy 替代 append
		return str[:start+copy(str[start:], str[end:])], nil
	}
}

// DeleteSliceLv2 改成泛式
func DeleteSliceLv2[T any](num int, end int, slice []T) ([]T, error) {
	// 获取切片长度
	len := len(slice)
	// 如果切片是空的，直接返回空切片
	if len == 0 {
		return nil, errors.New("长度不能为空")
	}
	//如果超出范围就返回错误
	if num < 1 || num > len || end < num || end > len {
		return nil, errors.New("超出范围")
	}
	//重置初始化下标
	start := num - 1
	//计算要删除的下标
	switch {
	case start == 0 && end == len: //整个切片就直接返回空
		return nil, nil
	case start == 0:
		return slice[end:], nil
	case end == len:
		return slice[start:], nil
	default: //使用copy元素，不会改变切片的容量。用copy 替代 append
		return slice[:start+copy(slice[start:], slice[end:])], nil
	}
}
