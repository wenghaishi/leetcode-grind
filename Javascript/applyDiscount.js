

const applyDiscount = (items, discount) => {
    currentTotal = 0
    appliedDiscount = false

    const maxPrice = Math.max(...items)

    for (const item of items ) {
        if (item === maxPrice && !appliedDiscount) {
            discountAmount = item * discount / 100
            discountedPrice = Math.floor(item - discountAmount)
            currentTotal += discountedPrice
            appliedDiscount = true
        } else {
            currentTotal += item
        }
    }

    return currentTotal

}

console.log(applyDiscount([40,2,1], 50))