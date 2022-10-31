#
# You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse
# order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
#
# You may assume the two numbers do not contain any leading zero, except the number 0 itself.

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def make_list(arr: [int]) -> ListNode:
    head = node = ListNode()
    for v in arr:
        node.next = ListNode(val=v)
        node = node.next
    return head.next


def addTwoNumbers(l1: ListNode | None, l2: ListNode | None) -> ListNode | None:
    sentinel = ListNode()

    overflow = 0
    node = sentinel

    while l1 is not None and l2 is not None:
        val = l1.val + l2.val + overflow
        overflow = int(val / 10)

        node.next = ListNode(val=val % 10)
        node = node.next
        l1 = l1.next
        l2 = l2.next
    else:
        # One of the lists ended. At most one of the loops executes.
        while l1 is not None:
            val = l1.val + overflow
            overflow = int(val / 10)
            node.next = ListNode(val=val % 10)
            node = node.next
            l1 = l1.next
        while l2 is not None:
            val = l2.val + overflow
            overflow = int(val / 10)
            node.next = ListNode(val=val % 10)
            node = node.next
            l2 = l2.next

    if overflow == 1:
        node.next = ListNode(val=1)

    return sentinel.next

