package redBlackTree

import "fmt"

/**
## 红黑树定义:
	红黑树（Red Black Tree） 是一种自平衡二叉查找树，是在计算机科学中用到的一种数据结构，典型的用途是实现关联数组。
	红黑树是一种特定类型的二叉树，它是在计算机科学中用来组织数据比如数字的块的一种结构。所有数据块都存储在节点中。这些节点中的某一个节点总是担当起始位置的功能，它不是任何节点的儿子，我们称之为根节点或根。它有最多两个"儿子"，都是它连接到的其他节点。所有这些儿子都可以有自己的儿子，以此类推。这样根节点就有了把它连接到在树中任何其他节点的路径。
## 红黑树历史：
	它是在1972年由Rudolf Bayer发明的，当时被称为平衡二叉B树（symmetric binary B-trees）。后来，在1978年被 Leo J. Guibas 和 Robert Sedgewick 修改为如今的“红黑树”。
## 红黑树性质：
* 红黑树是每个节点都带有颜色属性的二叉查找树，颜色或红色或黑色。在二叉查找树强制一般要求以外，对于任何有效的红黑树我们增加了如下的额外要求:
- 性质1. 节点是红色或黑色。
- 性质2. 根节点是黑色。
- 性质3 每个叶节点（NIL节点，空节点）是黑色的。
- 性质4 每个红色节点的两个子节点都是黑色。(从每个叶子到根的所有路径上不能有两个连续的红色节点)
- 性质5. 从任一节点到其每个叶子的所有路径都包含相同数目的黑色节点。(黑高度)
* 这些约束强制了红黑树的关键性质: 从根到叶子的最长的可能路径不多于最短的可能路径的两倍长。结果是这个树大致上是平衡的。因为操作比如插入、删除和查找某个值的最坏情况时间都要求与树的高度成比例，这个在高度上的理论上限允许红黑树在最坏情况下都是高效的，而不同于普通的二叉查找树。

## 应用
* 主要是用它来存储有序的数据，它的时间复杂度是O(lgn)，效率非常之高。
* 典型的用途是实现关联数组

## 其他
* 红黑树比AVL树优秀的地方之一在于黑父的情况比较常见，从而使红黑树需要旋转的几率相对AVL树来说会少一些。
* 红黑树节点的数据结构和avl有所差异，红黑树需要存储父节点


## 参考资料
* https://baike.baidu.com/item/%E7%BA%A2%E9%BB%91%E6%A0%91/2413209?fr=aladdin
* http://www.cnblogs.com/skywang12345/p/3245399.html
* 思路清晰的描述 http://www.cnblogs.com/skywang12345/p/3624177.html

*/

/**
红黑树的颜色
0红色，1黑色
*/
const (
	Red = iota
	Black
)

type DataType int;

type RedBlackNode struct {
	Data                DataType      //节点值
	Color               int           //节点颜色
	Parent, Left, Right *RedBlackNode //节点的父节点、左孩子、右孩子
}

// 树的根节点
type RBRoot RedBlackNode;

var TreeRoot *RedBlackNode;

//创建一个新树，返回该树的根节点
func CreateTree() *RedBlackNode {
	TreeRoot = new(RedBlackNode)
	return TreeRoot
}

//新增节点
func (_this *RedBlackNode) AddNode(val DataType) *RedBlackNode {
	if _this == nil {
		return &RedBlackNode{
			val,
			Black,
			nil,
			nil,
			nil,
		}
	}
	if val < _this.Data {
		_this.Left = _this.Left.AddNode(val)
	} else if val > _this.Data {
		_this.Right = _this.Right.AddNode(val)
	} else {
		fmt.Println("已经存在相同值的节点！")
		return nil
	}

	return _this
}

// todo 左侧旋转 LL型旋转 this.p,thisL.p,lr.p
func (_this *RedBlackNode) RRRotation() *RedBlackNode {
	var lNode *RedBlackNode
	lNode = _this.Left
	//右节点可以为空
	_this.Left = lNode.Right
	_this.Left.Parent = _this.Parent
	_this.Parent = lNode
	lNode.Right = _this
	lNode.Right.Parent = _this

	return lNode;
}

// todo LR型旋转
func (_this *RedBlackNode) LRRotation() *RedBlackNode {
	_this.Right = _this.Right.LLRotation()
	_this.RRRotation()
	return _this
}

// todo RL型旋转
func (_this *RedBlackNode) RLRotation() *RedBlackNode {
	_this.Left = _this.Left.LLRotation()
	_this.RRRotation()
	return _this
}

// todo 向左侧旋转 LL型旋转 this.p,thisR.p,rl.p这3个parent会发生变化
func (_this *RedBlackNode) LLRotation() *RedBlackNode {
	var rNode *RedBlackNode
	rNode = _this.Right
	//根节点的左节点可以为空
	_this.Right = rNode.Left
	//如果this.Parent为空，则将this置为根节点
	if _this.Parent == nil {
		TreeRoot = _this
	} else {
		_this.Right.Parent = _this.Parent
		_this.Parent = rNode
	}
	rNode.Left = _this
	rNode.Left.Parent = _this

	return rNode
}

// todo 前序遍历
func (_this *RedBlackNode) PrevTraverse() {
	if _this == nil {
		return
	}
	fmt.Printf("%d\t", _this.Data)
	_this.Left.PrevTraverse()
	_this.Right.PrevTraverse()
}

func (_this *RedBlackNode) Print() {
	_this.PrevTraverse()
}

//双红修正
func SolveDoubleRed() {

}

//失黑修正
func SolveLostBlack() {

}

func IsEmpty() {

}

//删除
func removeTree() {

}

//清楚整棵树
func clear() {

}
