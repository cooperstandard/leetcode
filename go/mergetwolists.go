package main

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	head := &ListNode{}
// 	cursor := head
// 	if list1 == nil && list2 == nil {
// 		return nil
// 	}
//
// 	for list1 != nil || list2 != nil {
// 		if list1 == nil {
// 			cursor.Next = list2.Next
// 			cursor.Val = list2.Val
// 			break
// 		}
// 		if list2 == nil {
// 			cursor.Next = list1.Next
// 			cursor.Val = list1.Val
// 			break
// 		}
// 		if list1.Val <= list2.Val {
// 			cursor.Val = list1.Val
// 			list1 = list1.Next
// 			cursor.Next = &ListNode{}
// 			cursor = cursor.Next
// 		} else {
// 			cursor.Val = list2.Val
// 			list2 = list2.Next
// 			cursor.Next = &ListNode{}
// 			cursor = cursor.Next
// 		}
// 	}
//
// 	return head
// }
//
