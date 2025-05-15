
const mergeLists = (l1, l2) => {
    const dummy  = new ListNode(null)

    let tail = dummy

    while(l1 != null && l2 != null) {
        if (l1.val > l2.val) {
            tail.next = l2
            l2 = l2.next
        } else {
            tail.next = l1
            l1 = l1.next
        }

        tail = tail.next
    }

    if (l1 != null) {
        tail.next = l1
    } else {
        tail.next = l1
    }

    return dummy;
}