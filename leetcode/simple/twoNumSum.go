package main

/**
## 总结
* 感觉自己用了大量的判断才完成，没用用到该有的算法。。。

 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 *ListNode = &ListNode{
		0,
		nil,
	}
	//l1.Next = &ListNode{
	//	4,
	//	nil,
	//}
	//l1.Next.Next = &ListNode{
	//	8,
	//	nil,
	//}

	var l2 *ListNode = &ListNode{
		7,
		nil,
	}
	l2.Next = &ListNode{
		3,
		nil,
	}
	//l2.Next.Next = &ListNode{
	//	4,
	//	nil,
	//}

	addTwoNumbers(l1, l2)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		index1, index2                    int
		carryNum, val1, val2              int
		newNode, originNode, tmpL1, tmpL2 *ListNode
	)
	tmpL1 = l1;
	tmpL2 = l2
	carryNum = 0
	for index1 = 0; ; index1++ {
		if tmpL1 != nil {
			val1 = tmpL1.Val
		} else {
			val1 = 0
		}
		tmpL2 = l2;
		for index2 = 0; index2 <= index1; index2++ {
			if index1 == index2 {
				break;
			} else {
				if tmpL2 != nil {
					tmpL2 = tmpL2.Next
				}
			}
		}
		if tmpL2 != nil {
			val2 = tmpL2.Val
		} else {
			val2 = 0
		}
		valTmp := val1 + val2 + carryNum
		node := &ListNode{
			valTmp,
			nil,
		}
		carryFlag := valTmp % 10
		carryNum = valTmp / 10
		if carryNum > 0 {
			node.Val = carryFlag
		}
		if newNode != nil {
			newNode.Next = node
			newNode = newNode.Next
		} else {
			newNode = node
			originNode = newNode
		}
		if tmpL1 != nil {
			tmpL1 = tmpL1.Next
		}
		if tmpL1 == nil && carryNum == 0 && (tmpL2==nil || tmpL2.Next==nil) {
			break
		}
	}

	return originNode
}
