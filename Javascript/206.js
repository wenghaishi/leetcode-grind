function ListNode(val, next) {
    this.val = (val===undefined ? 0 : val)
    this.next = (next===undefined ? null : next)
}

function solution(head) {
    let current = head
    let prev = null
    let temp

    while (current != null) {
        temp = current.next
        current.next = prev
        prev = current
        current = temp
    }

    return prev;
}

let l1 = ListNode(1)
let l2 = ListNode(2, l1)
let l3 = ListNode(3, l2)

console.log(solution(l3))