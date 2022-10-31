from ch2_add_two_numbers import addTwoNumbers, ListNode, make_list

if __name__ == '__main__':

    l1 = ListNode(val=6, next=ListNode(val=8))
    l2 = ListNode(val=5, next=ListNode(val=8))

    res = addTwoNumbers(make_list([9, 9, 9, 9, 9]), make_list([9, 9, 9]))
    while res is not None:
        print(res.val)
        res = res.next
