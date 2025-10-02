// WARN: *LISTNODE* is implemented by LeetCode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	// Get the head of node
	dummy, carry := &ListNode{}, 0

	// Stop when have nothing more to add
	for current := dummy; l1 != nil || l2 != nil || carry != 0; current = current.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: carry % 10} // The carry mod is the value of next node
		carry /= 10                               // Example: 10 + 7 = 17 -> 7 will be the node value and the carry will get 1
	}

	return dummy.Next
}
