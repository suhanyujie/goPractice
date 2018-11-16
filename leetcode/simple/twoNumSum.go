package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 *ListNode = &ListNode{
		2,
		nil,
	}
	l1.Next = &ListNode{
		4,
		nil,
	}
	l1.Next = &ListNode{
		3,
		nil,
	}


	var l2 *ListNode = &ListNode{
		5,
		nil,
	}
	l2.Next = &ListNode{
		6,
		nil,
	}
	l2.Next = &ListNode{
		4,
		nil,
	}


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
		index1, index2 int
		val1, val2     int
		newNode        *ListNode
		tmpL1, tmpL2   *ListNode
		carryNum       int
	)
	tmpL1 = l1;
	tmpL2 = l2
	carryNum = 0
	for index1 = 0; ; index1++ {
		val1 = tmpL1.Val
		tmpL2 = l2;
		for index2 = 0; index2 <= index1; index2++ {
			tmpL2 = tmpL2.Next
			if index1==index2 {
				break;
			}
		}
		val2 = tmpL2.Val
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
		} else {
			newNode = node
		}

		tmpL1 = tmpL1.Next
		if tmpL1 == nil {
			break
		}
	}

	return newNode
}
