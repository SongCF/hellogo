package main

import (
	"fmt"
)

func main() {
	//保存需要排序的Slice
	arr := []int{9, 3, 4, 7, 2, 1, 0, 11, 12, 11, 13, 4, 7, 2, 1, 0, 11, 12, 11}
	//实际用于排序的Slice
	list := make([]int, len(arr))

	copy(list, arr)
	BubbleSortX(list)
	fmt.Println("冒泡排序：\t", list)

	copy(list, arr)
	QuickSort(list, 0, len(arr)-1)
	fmt.Println("快速排序：\t", list)

	copy(list, arr) //将arr的数据覆盖到list，重置list
	InsertSort(list)
	fmt.Println("直接插入排序：\t", list)

	copy(list, arr)
	ShellSort(list)
	fmt.Println("希尔排序：\t", list)

	copy(list, arr)
	MergeSort(list)
	fmt.Println("二路归并排序：\t", list)

	copy(list, arr)
	SelectSort(list)
	fmt.Println("简单选择排序：\t", list)

	copy(list, arr)
	HeapSort(list)
	fmt.Println("堆排序：     \t", list)

}

//region 冒泡排序
//1，正宗的冒泡排序
/*
每趟排序过程中通过两两比较，找到第 i 个小（大）的元素，将其往上排。
*/
func BubbleSort(list []int) {
	var temp int // 用来交换的临时数
	var i int
	var j int
	// 要遍历的次数
	for i = 0; i < len(list)-1; i++ {
		// 从后向前依次的比较相邻两个数的大小，遍历一次后，把数组中第i小的数放在第i个位置上
		for j = len(list) - 1; j > i; j-- {
			// 比较相邻的元素，如果前面的数大于后面的数，则交换
			if list[j-1] > list[j] {
				temp = list[j-1]
				list[j-1] = list[j]
				list[j] = temp
			}
		}
	}
}

//2，冒泡排序优化
/*
对冒泡排序常见的改进方法是加入标志性变量exchange，用于标志某一趟排序过程中是否有数据交换。
如果进行某一趟排序时并没有进行数据交换，则说明所有数据已经有序，可立即结束排序，避免不必要的比较过程。
*/
func BubbleSortX(list []int) {
	var exchange bool = false
	var temp int // 用来交换的临时数
	var i int
	var j int
	// 要遍历的次数
	for i = 0; i < len(list)-1; i++ {
		// 从后向前依次的比较相邻两个数的大小，遍历一次后，把数组中第i小的数放在第i个位置上
		for j = len(list) - 1; j > i; j-- {
			// 比较相邻的元素，如果前面的数大于后面的数，则交换
			if list[j-1] > list[j] {
				temp = list[j-1]
				list[j-1] = list[j]
				list[j] = temp
				exchange = true
			}
		}
		if !exchange {
			break
		}
		exchange = false
	}
}

//endregion

//region 快速排序
func division(list []int, left int, right int) int {

	// 以最左边的数(left)为基准
	var base int = list[left]
	for left < right {
		// 从序列右端开始，向左遍历，直到找到小于base的数
		for left < right && list[right] >= base {
			right--
		}
		// 找到了比base小的元素，将这个元素放到最左边的位置
		list[left] = list[right]
		// 从序列左端开始，向右遍历，直到找到大于base的数
		for left < right && list[left] <= base {
			left++
		}
		// 找到了比base大的元素，将这个元素放到最右边的位置
		list[right] = list[left]

	}
	// 最后将base放到left位置。此时，left位置的左侧数值应该都比left小
	// 而left位置的右侧数值应该都比left大。
	list[left] = base //此时left == right
	//fmt.Println("DONE: base:", base, "\tleft:", left, "\tright:", right)
	return left
}

func QuickSort(list []int, left int, right int) {
	// 左下标一定小于右下标，否则就越界了
	if left < right {
		//对数组进行分割，取出下次分割的基准标号
		var base int = division(list, left, right)
		//对“基准标号“左侧的一组数值进行递归的切割，以至于将这些数值完整的排序
		QuickSort(list, left, base-1)
		//对“基准标号“右侧的一组数值进行递归的切割，以至于将这些数值完整的排序
		QuickSort(list, base+1, right)
	}

}

//endregion

//region 直接插入排序
func InsertSort(list []int) {
	var temp int
	var i int
	var j int
	// 第1个数肯定是有序的，从第2个数开始遍历，依次插入有序序列
	for i = 1; i < len(list); i++ {
		temp = list[i] // 取出第i个数，和前i-1个数比较后，插入合适位置
		// 因为前i-1个数都是从小到大的有序序列，所以只要当前比较的数(list[j])比temp大，就把这个数后移一位
		for j = i - 1; j >= 0 && temp < list[j]; j-- {
			list[j+1] = list[j]
		}
		list[j+1] = temp
	}
}

//endregion

//region 希尔排序
func ShellSort(list []int) {
	for gap := (len(list) + 1) / 2; gap >= 1; gap = gap / 2 {
		for i := 0; i < len(list)-gap; i++ {
			InsertSort(list[i:(gap + i + 1)]) //list[i:(gap + i + 1)]表示list索引i到gap+i的元素组成的slice
		}
	}
}

//region

//region 简单选择排序
/*
简单排序处理流程：
（1）从待排序序列中，找到关键字最小的元素；
（2）如果最小元素不是待排序序列的第一个元素，将其和第一个元素互换；
（3）从余下的 N - 1 个元素中，找出关键字最小的元素，重复（1）、（2）步，直到排序结束。
*/
func SelectSort(list []int) {
	var temp int
	var index int
	var i int
	var j int

	// 需要遍历获得最小值的次数
	// 要注意一点，当要排序 N 个数，已经经过 N-1 次遍历后，已经是有序数列
	for i = 0; i < len(list)-1; i++ {
		temp = 0
		index = i // 用来保存最小值得索引
		// 寻找第i个小的数值
		for j = i + 1; j < len(list); j++ {
			if list[index] > list[j] {
				index = j
			}
		}
		// 将找到的第i个小的数值放在第i个位置上
		temp = list[index]
		list[index] = list[i]
		list[i] = temp
	}
}

//endregion

//region 堆排序
func heapAdjust(list []int, parent int, length int) {
	temp := list[parent]  // temp保存当前父节点
	child := 2*parent + 1 // 先获得左孩子

	for child < length {
		// 如果有右孩子结点，并且右孩子结点的值大于左孩子结点，则选取右孩子结点
		if child+1 < length && list[child] < list[child+1] {
			child++
		}

		// 如果父结点的值已经大于孩子结点的值，则直接结束
		if temp >= list[child] {
			break
		}

		// 把孩子结点的值赋给父结点
		list[parent] = list[child]

		// 选取孩子结点的左孩子结点,继续向下筛选
		parent = child
		child = 2*child + 1
	}

	list[parent] = temp
}

func HeapSort(list []int) {
	// 循环建立初始堆
	for i := len(list) / 2; i >= 0; i-- {
		heapAdjust(list, i, len(list)-1)
	}

	// 进行n-1次循环，完成排序
	for i := len(list) - 1; i > 0; i-- {
		// 最后一个元素和第一元素进行交换
		temp := list[i]
		list[i] = list[0]
		list[0] = temp

		// 筛选 R[0] 结点，得到i-1个结点的堆
		heapAdjust(list, 0, i)
	}
}

//endregion

//region 归并排序(二路归并)
func merge(list []int, low int, mid int, high int) {
	var i int = low                  // i是第一段序列的下标
	var j int = mid + 1              // j是第二段序列的下标
	var k int = 0                    // k是临时存放合并序列的下标
	list2 := make([]int, high-low+1) // list2是临时合并序列
	// 扫描第一段和第二段序列，直到有一个扫描结束
	for i <= mid && j <= high {
		// 判断第一段和第二段取出的数哪个更小，将其存入合并序列，并继续向下扫描
		if list[i] <= list[j] {
			list2[k] = list[i]
			i++
			k++
		} else {
			list2[k] = list[j]
			j++
			k++
		}
	}
	// 若第一段序列还没扫描完，将其全部复制到合并序列
	for i <= mid {
		list2[k] = list[i]
		i++
		k++
	}

	// 若第二段序列还没扫描完，将其全部复制到合并序列
	for j <= high {
		list2[k] = list[j]
		j++
		k++
	}
	// 将合并序列复制到原始序列中
	k = 0
	for i = low; i <= high; i++ {
		list[i] = list2[k]
		k++
	}
}

func MergeSort(list []int) {
	for gap := 1; gap < len(list); gap = 2 * gap {
		var i int
		// 归并gap长度的两个相邻子表
		for i = 0; i+2*gap-1 < len(list); i = i + 2*gap {
			merge(list, i, i+gap-1, i+2*gap-1)
		}
		// 余下两个子表，后者长度小于gap
		if i+gap-1 < len(list) {
			merge(list, i, i+gap-1, len(list)-1)
		}
	}

}

//endregion