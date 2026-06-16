/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    var numb, numb1, numb2, currentVal int
    over := 0
	var l3 *ListNode
	var current *ListNode

	for i := 0; l1 != nil || l2 != nil || over != 0; i++ {

        if l1 != nil {
            numb1 = l1.Val
        } else {
            numb1 = 0
        }

        if l2 != nil {
            numb2 = l2.Val
        } else {
            numb2 = 0
        }

        numb = numb1 + numb2 + over

        currentVal = numb % 10
        over = numb / 10

        if l1 != nil {
            l1 = l1.Next
        }
		if l2 != nil {
            l2 = l2.Next
        }

		node := &ListNode{Val: currentVal}

		if l3 == nil {
			l3 = node
			current = node
		} else {
			current.Next = node
			current = node
		}
	}

	return l3
}
