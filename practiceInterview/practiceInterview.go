package practiceinterview

import (
	"sync"
)

func Test(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	// 有向有权图最短路问题
	// Dijkstra_Alg()

	// // dijkstra()

	// B+查找树
	// BPlusTreeFunc()

	// 希尔排序
	// start := time.Now() // 获取当前时间
	// shellSort(tempSlice)
	// log.Println(tempSlice, "希尔排序\r\n", time.Since(start))

	// 堆排序
	// start := time.Now() // 获取当前时间
	// heapSort(tempSlice)
	// log.Println("\r\n", time.Since(start))

	// 归并排序
	// start_1 := time.Now() // 获取当前时间
	// segmentation(tempSlice)
	// log.Println("未优化/n", time.Since(start_1))

	// start_2 := time.Now() // 获取当前时间
	// pointerSplitMergeSort(tempSlice1, 0, len(tempSlice1))
	// log.Println("堆排序优化后\r\n", time.Since(start_2))
}

func Dijkstra_Alg(graph [][]int, start int) {

}

/**
 * @description: 有向有权图最短路问题
 * @param {[][]int} graph 图抽象的邻接表(二维数组)
 * @param {int} start 开始的点
 * @param {int} end 结束的点
 * @return {[]int} 最短路线
 * @return {int} 最短路线距离
 */
// func dijkstra(graph [][]int, start, end int) ([]int, int) {
// 	lenGraph := len(graph)       //节点数量
// 	pre := make([]int, lenGraph) //记录前驱
// 	// vis := make([]int, lenGraph) //记录节点遍历状态
// 	dis := []int{} //保存最短距离
// 	for i := 0; i < lenGraph; i++ {
// 		dis = append(dis, i)
// 	}
// 	// road := make([]int, lenGraph) //保存最短路径
// 	// roads := []int{}

// 	//初始化起点到其他点的距离
// 	for i := 0; i < lenGraph; i++ {
// 		if i == start {
// 			dis[i] = 0
// 		} else {
// 			dis[i] = graph[start][i]
// 		}
// 		if graph[start][i] != 999 {
// 			pre[i] = start
// 		} else {
// 			pre[i] = -1
// 		}
// 	}

// 	return []int{}, -1
// }

// 节点数据结构***************************************************************
type BPlusItem struct {
	Key int
	Val interface{}
}

type BPlusNode struct {
	MaxKey int          //用于存储子树的最大关键字
	Nodes  []*BPlusNode //结点的子树（叶子结点的Nodes=nil）
	Items  []BPlusItem  //叶子结点的数据记录（索引结点的Items=nil）
	Next   *BPlusNode   //叶子结点中指向下一个叶子结点，用于实现叶子结点链表
}

// B+树的定义
type BPlusTree struct {
	mutex sync.RWMutex //多线程读写锁
	root  *BPlusNode   //指向B+树的根结点
	width int          //用于表示B+树的阶
	halfw int          //B树至少有费根非叶 m(阶数)/2 棵子树
}

// B+树的初始化****************************************************************
func NewLeafNode(width int) (BPlusNode *BPlusNode) {
	BPlusNode.Items = make([]BPlusItem, width+1)
	BPlusNode.Items = BPlusNode.Items[0:0]
	return
}

func NewBPlusTree(width int) (BPlusTree *BPlusTree) {
	if width < 3 {
		width = 3
	}

	BPlusTree.root = NewLeafNode(width)
	BPlusTree.width = width
	BPlusTree.halfw = (BPlusTree.width + 1) / 2
	return
}

/**
 * @description:      B+树的查询****************************************************************
 * @param {int} key   查询的值
 * @return {*}        结果
 */
// func (_BPlusTree *BPlusTree) get(key int) interface{} {

// 	_BPlusTree.mutex.Lock()
// 	defer _BPlusTree.mutex.Unlock()

// 	// 循环寻找所在节点
// 	node := _BPlusTree.root
// 	for i := 0; i < len(node.Nodes); i++ {
// 		if key <= node.Nodes[i].MaxKey {
// 			node = node.Nodes[i]
// 			i = 0
// 		}
// 	}
// 	// 没有到达叶子结点
// 	if len(node.Nodes) > 0 {
// 		return nil
// 	}

// 	// 循环节点的阶(数据)
// 	for i := 0; i < len(node.Items); i++ {
// 		if node.Items[i].Key == key {
// 			return node.Items[i].Val
// 		}
// 	}

// 	return nil
// }

/**
 * @description: 节点添加数据处理函数
 * @param {int} key
 * @param {interface{}} value
 * @return {*}
 */
// func (_BPlusNode *BPlusNode) setValue(key int, value interface{}) (ok bool, err error) {

// 	item := BPlusItem{key, value} //要插入的真实数据
// 	num := len(_BPlusNode.Items)  //当前节点的有几条真实数据

// 	if num < 1 { //当前节点没有存储数据
// 		_BPlusNode.Items = append(_BPlusNode.Items, item) //将数据加入当前节点存储
// 		_BPlusNode.MaxKey = item.Key                      //当前节点最大值设为插入节点的key
// 		return true, nil
// 	} else if key < _BPlusNode.Items[0].Key { //设置的key值小于当前节点存储数据第一条,
// 		_BPlusNode.Items = append([]BPlusItem{item}, _BPlusNode.Items...) //将数据插入到第一位
// 		return true, nil
// 	} else if key > _BPlusNode.Items[num-1].Key { //设置的key值大于当前节点存储数据第最后一条
// 		_BPlusNode.Items = append(_BPlusNode.Items, item) //将数据插入到第最后一位
// 		_BPlusNode.MaxKey = item.Key                      //当前节点最大值设置为插入的数据
// 		return true, nil
// 	}

// 	for i := 0; i < num; i++ { //如果数据大小在中间，循环插入
// 		if _BPlusNode.Items[i].Key > key { //如果该节点key大于将要设置的值key
// 			_BPlusNode.Items = append(_BPlusNode.Items, BPlusItem{}) //追加一个空数据
// 			copy(_BPlusNode.Items[i+1:], _BPlusNode.Items[i:])       //整体数据后移一位
// 			_BPlusNode.Items[i] = item                               // 插入数据
// 			return true, nil
// 		} else if _BPlusNode.Items[i].Key == key { //如果发现相同主键key
// 			_BPlusNode.Items[i] = item //直接赋值
// 			return true, nil
// 		}
// 	}

// 	return false, errors.New("BPlusNode节点添加数据失败") //走到这里
// }

/**
 * @description:分割节点
 * @param {*BPlusNode} node
 * @return {*}
 */
// func (_BPlusTree *BPlusTree) splitNode(node *BPlusNode) *BPlusNode {

// 	if len(node.Nodes) > _BPlusTree.width { //如果节点子节点数量大于B+树的阶值

// 		halfw := _BPlusTree.width/2 + 1 //中间阶数
// 		node2 := &BPlusNode{
// 			Items:  make([]BPlusItem, _BPlusTree.width),
// 			Nodes:  node.Nodes[halfw:len(node.Nodes)],
// 			MaxKey: node.Nodes[len(node.Nodes)-1].MaxKey,
// 		}

// 		node.Nodes = node.Nodes[0:halfw]
// 		node.MaxKey = node.Nodes[len(node.Nodes)-1].MaxKey

// 		return node2

// 	} else if len(node.Items) > _BPlusTree.width {

// 		return nil
// 	}

// 	return nil
// }

/**
 * @description: 删除节点
 * @param {int} key
 * @return {*}
 */
func (_BPlusTree *BPlusTree) Remove(key int) {
	_BPlusTree.mutex.Lock()
	defer _BPlusTree.mutex.Unlock()

}

/**
 * @description: 删除数据
 * @param {*BPlusNode} parent
 * @param {*BPlusNode} node
 * @param {int} key
 * @return {*}
 */
// func (_BPlusTree *BPlusTree) deleteItem(parent *BPlusNode, node *BPlusNode, key int) {
// 	lenNode := len(node.Nodes)
// 	for i := 0; i < lenNode; i++ { //递归寻找节点
// 		if key <= node.Nodes[i].MaxKey {
// 			_BPlusTree.deleteItem(node, node.Nodes[i], key)
// 			break
// 		}
// 	}

// 	if lenNode < 1 {

// 	}

// }

// func (_BPlusNode *BPlusNode) deleteItem(key int) bool {
// 	lenItem := len(_BPlusNode.Items)
// 	for i := 0; i < lenItem; i++ {

// 		if _BPlusNode.Items[i].Key > key {
// 			return false
// 		} else if _BPlusNode.Items[i].Key == key {
// 			copy(_BPlusNode.Items[i:], _BPlusNode.Items[i+1:]) //前移删除数据
// 			_BPlusNode.Items = _BPlusNode.Items[0 : len(_BPlusNode.Items)-1]
// 			_BPlusNode.MaxKey = _BPlusNode.Items[len(_BPlusNode.Items)-1].Key
// 			return true
// 		}

// 	}

// 	return false

// }

/**
 * @description: 树设置值
 * @param {*BPlusNode} parent 父节点
 * @param {*BPlusNode} node 要操作的节点
 * @param {int} key
 * @param {interface{}} value
 * @return {*}
 */
// func (_BPlusTree *BPlusTree) setValueTree(parent *BPlusNode, node *BPlusNode, key int, value interface{}) {

// 	for i := 0; i < len(node.Nodes); i++ { //循环所有子节点
// 		if key <= node.Nodes[i].MaxKey || i == len(node.Nodes)-1 { //如果设置的值小于某个子节点的最大值，或者循环完了所有子节点
// 			_BPlusTree.setValueTree(node, node.Nodes[i], key, value) //进行递归
// 			break
// 		}
// 	}

// 	if len(node.Nodes) < 1 { //如果节点没有子节点
// 		node.setValue(key, value) //直接设置值
// 	}

// }

/**
 * @description: 树结构体设置值
 * @param {int} key
 * @param {interface{}} value
 * @return {*}
 */
func (_BPlusTree *BPlusTree) SetTree(key int, value interface{}) (bool, error) {
	_BPlusTree.mutex.Lock()
	defer _BPlusTree.mutex.Unlock()

	return false, nil
}

// B+树
func BPlusTreeFunc() {

}

// func shellSort(slice_arg []int) {
// 	StepLength := len(slice_arg) / 2
// 	for gap := StepLength; gap >= 1; gap = gap / 2 {

// 		for rightIndex := gap; rightIndex < len(slice_arg); rightIndex++ {
// 			for leftIndex := rightIndex - gap; leftIndex >= 0; leftIndex = leftIndex - gap {
// 				if slice_arg[leftIndex] > slice_arg[leftIndex+gap] {
// 					slice_arg[leftIndex], slice_arg[leftIndex+gap] = slice_arg[leftIndex+gap], slice_arg[leftIndex]
// 				}
// 			}
// 		}
// 	}
// }

/**
 * @name:
 * @msg:
 * @param {[]int} slice_arg  数组切片
 * @return {*}
 */
// func heapSort(slice_arg []int) {

// 	for nonLeafNode := len(slice_arg)/2 - 1; nonLeafNode >= 0; nonLeafNode-- {
// 		adjustHeap(slice_arg, nonLeafNode, len(slice_arg))
// 	}

// 	for endIndex := len(slice_arg) - 1; endIndex > 0; endIndex-- {
// 		slice_arg[endIndex], slice_arg[0] = slice_arg[0], slice_arg[endIndex]
// 		adjustHeap(slice_arg, 0, endIndex)
// 	}

// }

/**
 * @name:
 * @msg:												三角堆节点顺序调整
 * @param {[]int} slice_arg 		数组切片
 * @param {*} index_arg 				三角树顶
 * @param {int} sliceLength_arg 数组切片长度
 * @return {*}
 */
// func adjustHeap(slice_arg []int, index_arg, sliceLength_arg int) {

// 	temp := slice_arg[index_arg]

// 	for childIndex := 2*index_arg + 1; childIndex < sliceLength_arg; childIndex = 2*childIndex + 1 {

// 		if childIndex+1 < sliceLength_arg && slice_arg[childIndex] < slice_arg[childIndex+1] {
// 			childIndex++ //指向右子节点
// 		}

// 		if slice_arg[childIndex] > slice_arg[index_arg] {
// 			slice_arg[index_arg] = slice_arg[childIndex] //子节点大于父节点、交换位置
// 			index_arg = childIndex
// 		} else {
// 			break
// 		}

// 		slice_arg[index_arg] = temp

// 	}

// }

// 优化方法1---------------------------------------------------------------------
// func pointerSplitMergeSort(slice_arg []int, startIndex_arg, endInde_arg int) {

// 	if endInde_arg-startIndex_arg < 2 {
// 		return
// 	}

// 	mideleIndex := (startIndex_arg + endInde_arg) >> 1

// 	pointerSplitMergeSort(slice_arg, startIndex_arg, mideleIndex)
// 	pointerSplitMergeSort(slice_arg, mideleIndex, endInde_arg)

// 	pointerSplitMerge(slice_arg, startIndex_arg, mideleIndex, endInde_arg)

// }

// func pointerSplitMerge(slice_arg []int, startIndex_arg, mideleIndex_arg, endInde_arg int) {

// 	leftSliceLen := mideleIndex_arg - startIndex_arg
// 	leftSlice := make([]int, leftSliceLen)

// 	copy(leftSlice, slice_arg[startIndex_arg:mideleIndex_arg])

// 	leftIndex, startIndex, rightIndex := 0, startIndex_arg, mideleIndex_arg
// 	for leftIndex < leftSliceLen {
// 		if rightIndex < endInde_arg && slice_arg[rightIndex] < leftSlice[leftIndex] {
// 			slice_arg[startIndex] = slice_arg[rightIndex]
// 			startIndex++
// 			rightIndex++
// 		} else {
// 			slice_arg[startIndex] = leftSlice[leftIndex]
// 			startIndex++
// 			leftIndex++
// 		}
// 	}

// }

// 普通方法2-------------------------------------------------------------------------
// 分割切片

// func segmentation(_tempSlice []int) []int {
// 	if len(_tempSlice) < 2 {
// 		return _tempSlice
// 	}
// 	// 找到中间下标
// 	middleIndeex := len(_tempSlice) / 2
// 	// 左半部分切片只剩两个
// 	leftSlice := segmentation(_tempSlice[:middleIndeex])
// 	// 右半部分切片只剩两个
// 	rightSlice := segmentation(_tempSlice[middleIndeex:])
// 	return mergeSlices(leftSlice, rightSlice)
// }

// 合并切片
// func mergeSlices(_leftSlice, _rightSlice []int) []int {

// 	resultSlice := make([]int, 0)
// 	leftIndex, rightIndex := 0, 0
// 	leftLen, rightLen := len(_leftSlice), len(_rightSlice)

// 	for leftIndex < leftLen && rightIndex < rightLen {

// 		if _leftSlice[leftIndex] > _rightSlice[rightIndex] {
// 			resultSlice = append(resultSlice, _rightSlice[rightIndex])
// 			rightIndex++

// 		} else {
// 			resultSlice = append(resultSlice, _leftSlice[leftIndex])
// 			leftIndex++

// 		}
// 	}

// 	resultSlice = append(resultSlice, _rightSlice[rightIndex:]...)
// 	resultSlice = append(resultSlice, _leftSlice[leftIndex:]...)

// 	return resultSlice
// }
