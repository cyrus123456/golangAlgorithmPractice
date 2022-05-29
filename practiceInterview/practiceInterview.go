package practiceinterview

import (
	"fmt"
	"time"
)

func Test() {

	// 希尔排序
	start := time.Now() // 获取当前时间
	shellSort(tempSlice)
	fmt.Println(tempSlice, "希尔排序\r\n", time.Since(start))

	// 堆排序
	// start := time.Now() // 获取当前时间
	// heapSort(tempSlice)
	// fmt.Println("\r\n", time.Since(start))

	// 归并排序
	// start_1 := time.Now() // 获取当前时间
	// segmentation(tempSlice)
	// fmt.Println("未优化/n", time.Since(start_1))

	start_2 := time.Now() // 获取当前时间
	pointerSplitMergeSort(tempSlice1, 0, len(tempSlice1))
	fmt.Println("堆排序优化后\r\n", time.Since(start_2))
}

func shellSort(slice_arg []int) {
	StepLength := len(slice_arg) / 2
	for gap := StepLength; gap >= 1; gap = gap / 2 {

		for rightIndex := gap; rightIndex < len(slice_arg); rightIndex++ {
			for leftIndex := rightIndex - gap; leftIndex >= 0; leftIndex = leftIndex - gap {
				if slice_arg[leftIndex] > slice_arg[leftIndex+gap] {
					slice_arg[leftIndex], slice_arg[leftIndex+gap] = slice_arg[leftIndex+gap], slice_arg[leftIndex]
				}
			}
		}
	}
}

/**
 * @name:
 * @msg:
 * @param {[]int} slice_arg  数组切片
 * @return {*}
 */
func heapSort(slice_arg []int) {

	for nonLeafNode := len(slice_arg)/2 - 1; nonLeafNode >= 0; nonLeafNode-- {
		adjustHeap(slice_arg, nonLeafNode, len(slice_arg))
	}

	for endIndex := len(slice_arg) - 1; endIndex > 0; endIndex-- {
		slice_arg[endIndex], slice_arg[0] = slice_arg[0], slice_arg[endIndex]
		adjustHeap(slice_arg, 0, endIndex)
	}

}

/**
 * @name:
 * @msg:												三角堆节点顺序调整
 * @param {[]int} slice_arg 		数组切片
 * @param {*} index_arg 				三角树顶
 * @param {int} sliceLength_arg 数组切片长度
 * @return {*}
 */
func adjustHeap(slice_arg []int, index_arg, sliceLength_arg int) {

	temp := slice_arg[index_arg]

	for childIndex := 2*index_arg + 1; childIndex < sliceLength_arg; childIndex = 2*childIndex + 1 {

		if childIndex+1 < sliceLength_arg && slice_arg[childIndex] < slice_arg[childIndex+1] {
			childIndex++ //指向右子节点
		}

		if slice_arg[childIndex] > slice_arg[index_arg] {
			slice_arg[index_arg] = slice_arg[childIndex] //子节点大于父节点、交换位置
			index_arg = childIndex
		} else {
			break
		}

		slice_arg[index_arg] = temp

	}

}

// 优化方法1---------------------------------------------------------------------
func pointerSplitMergeSort(slice_arg []int, startIndex_arg, endInde_arg int) {

	if endInde_arg-startIndex_arg < 2 {
		return
	}

	mideleIndex := (startIndex_arg + endInde_arg) >> 1

	pointerSplitMergeSort(slice_arg, startIndex_arg, mideleIndex)
	pointerSplitMergeSort(slice_arg, mideleIndex, endInde_arg)

	pointerSplitMerge(slice_arg, startIndex_arg, mideleIndex, endInde_arg)

}

func pointerSplitMerge(slice_arg []int, startIndex_arg, mideleIndex_arg, endInde_arg int) {

	leftSliceLen := mideleIndex_arg - startIndex_arg
	leftSlice := make([]int, leftSliceLen)

	copy(leftSlice, slice_arg[startIndex_arg:mideleIndex_arg])

	leftIndex, startIndex, rightIndex := 0, startIndex_arg, mideleIndex_arg
	for leftIndex < leftSliceLen {
		if rightIndex < endInde_arg && slice_arg[rightIndex] < leftSlice[leftIndex] {
			slice_arg[startIndex] = slice_arg[rightIndex]
			startIndex++
			rightIndex++
		} else {
			slice_arg[startIndex] = leftSlice[leftIndex]
			startIndex++
			leftIndex++
		}
	}

}

// 普通方法2-------------------------------------------------------------------------
// 分割切片
func segmentation(_tempSlice []int) []int {
	if len(_tempSlice) < 2 {
		return _tempSlice
	}
	// 找到中间下标
	middleIndeex := len(_tempSlice) / 2
	// 左半部分切片只剩两个
	leftSlice := segmentation(_tempSlice[:middleIndeex])
	// 右半部分切片只剩两个
	rightSlice := segmentation(_tempSlice[middleIndeex:])
	return mergeSlices(leftSlice, rightSlice)
}

// 合并切片
func mergeSlices(_leftSlice, _rightSlice []int) []int {

	resultSlice := make([]int, 0)
	leftIndex, rightIndex := 0, 0
	leftLen, rightLen := len(_leftSlice), len(_rightSlice)

	for leftIndex < leftLen && rightIndex < rightLen {

		if _leftSlice[leftIndex] > _rightSlice[rightIndex] {
			resultSlice = append(resultSlice, _rightSlice[rightIndex])
			rightIndex++

		} else {
			resultSlice = append(resultSlice, _leftSlice[leftIndex])
			leftIndex++

		}
	}

	resultSlice = append(resultSlice, _rightSlice[rightIndex:]...)
	resultSlice = append(resultSlice, _leftSlice[leftIndex:]...)

	return resultSlice
}
