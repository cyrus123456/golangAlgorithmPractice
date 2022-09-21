package practiceinterview

import (
	"math"
)

// Ordered 限制所有支持比较运算符的类型。
// 也就是说符合条件的类型都支持 <, <=, >, 和 >= 运算符。
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

// 节点结构体
type node_avlTree[K Ordered, V any] struct {
	key    K
	value  V
	left   *node_avlTree[K, V]
	right  *node_avlTree[K, V]
	height int
}

// 主体类结构体
type avlTree[K Ordered, V any] struct {
	root *node_avlTree[K, V]
	size int
}

/**
 * @description: 获取平衡因子
 * @param {*node_avlTree[K} node 节点
 * @return {int} 左减去右，正数左高，相等就是0,
 */
func (_avlTree *avlTree[K, V]) getBalanceFactor(node *node_avlTree[K, V]) int {
	return node.left.height - node.right.height
}

/**
 * @description: 右旋操作
 * @param {*node_avlTree[K} node 节点
 * @return {*node_avlTree[K, V]}
 */
func (avlTree *avlTree[K, V]) rightRotate(node *node_avlTree[K, V]) *node_avlTree[K, V] {
	tempLeft := node.left
	tempLeftRight := tempLeft.right
	// 右旋转
	tempLeft.right = node
	node.left = tempLeftRight
	// 更换新高度
	node.height = int(math.Max(float64(node.left.height), float64(node.right.height))) + 1
	tempLeft.height = int(math.Max(float64(tempLeft.left.height), float64(tempLeft.right.height))) + 1
	return tempLeft
}

/**
 * @description: 左旋操作
 * @param {*node_avlTree[K} node 节点
 * @return {*node_avlTree[K, V]}
 */
func (_avlTree *avlTree[K, V]) leftRoutate(node *node_avlTree[K, V]) *node_avlTree[K, V] {
	tempRight := node.right
	tempRightLeft := tempRight.left
	//左旋
	tempRight.left = node
	node.right = tempRightLeft
	// 更新高度
	node.height = int(math.Max(float64(node.left.height), float64(node.right.height))) + 1
	tempRight.height = int(math.Max(float64(tempRight.left.height), float64(tempRight.right.height))) + 1
	return node
}

/**
 * @description: avl树根节点赋值
 * @param {K} key 键-可比较类型
 * @param {V} value 值-任意类型
 * @return {*}
 */
func (_avlTree *avlTree[K, V]) add_avlTree(key K, value V) {
	_avlTree.root = _avlTree.add_node_avlTree(_avlTree.root, key, value)
}

/**
 * @description: 递归添加节点
 * @param {node_avlTree[K, V]} root 根节点
 * @param {K} key 键-可比较类型
 * @param {V} value 值-任意类型
 * @return {node_avlTree[K, V]} 节点
 */
func (_avlTree *avlTree[K, V]) add_node_avlTree(node *node_avlTree[K, V], key K, value V) *node_avlTree[K, V] {

	// 寻找合适的位置
	if key < node.key {
		node.left = _avlTree.add_node_avlTree(node.left, key, value)
	} else if node.key > key {
		node.right = _avlTree.add_node_avlTree(node.right, key, value)
	} else {
		node.value = value
	}

	// 节点高度
	node.height = 1 + int(math.Max(float64(node.left.height), float64(node.right.height)))

	// 当前节点平衡因子
	balanceFactor := _avlTree.getBalanceFactor(node)

	// 平衡修复
	if int(math.Abs(float64(balanceFactor))) > 1 {

		// 当前和左子节点都向左倾斜-又旋转
		if balanceFactor > 1 && _avlTree.getBalanceFactor(node.left) >= 0 {
			return _avlTree.rightRotate(node)
		}

		// 当前和右子节点都向右勤快些-向左转
		if balanceFactor < -1 && _avlTree.getBalanceFactor(node.right) <= 0 {
			return _avlTree.leftRoutate(node)
		}

		// 当前节点左偏，子节点右偏
		if balanceFactor > 1 && _avlTree.getBalanceFactor(node.left) < 0 {
			node.right = _avlTree.leftRoutate(node.right)
			return _avlTree.rightRotate(node)
		}

		// 当前节点右偏，子节点左偏
		if balanceFactor < -1 && _avlTree.getBalanceFactor(node.right) > 0 {
			node.left = _avlTree.rightRotate(node.left)
			return _avlTree.leftRoutate(node)
		}

	}

	return node
}

/**
 * @description: 获取值
 * @param {*node_avlTree[K} node 节点
 * @param {K} key 主键
 * @return {*node_avlTree[K, V]} 节点
 */
func (_avlTree *avlTree[K, V]) getNode(node *node_avlTree[K, V], key K) *node_avlTree[K, V] {

	if key == node.key { //找到了
		return node
	} else if key < node.key { //值比当前节点小
		return _avlTree.getNode(node.left, key)
	} else { //值比当前节点大
		return _avlTree.getNode(node.right, key)
	}
}

/**
 * @description: 删除
 * @param {K} key 要删除的主键
 * @return {interface{}} 返回任意数据类型
 */
func (_avlTree *avlTree[K, V]) remove_avlTree(key K) interface{} {
	_avlTree.root = _avlTree.remove_node_avlTree(_avlTree.root, key)
	return nil
}

/**
 * @description: 删除递归
 * @param {*node_avlTree[K} node 节点
 * @param {K} key 要删除的
 * @return {*node_avlTree[K, V]} 节点
 */
func (_avlTree *avlTree[K, V]) remove_node_avlTree(node *node_avlTree[K, V], key K) *node_avlTree[K, V] {

	return node
}
