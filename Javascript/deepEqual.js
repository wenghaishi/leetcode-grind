
const deepEqual = (item1, item2) => {
    const item1Type = typeof item1
    const item2Type = typeof item2

    if (item1Type !== item2Type) {
        console.log("diff types")
        return false
    }
    if (item1.length != item2.length) {
        console.log(item1.length)
        console.log(item2.length)

        console.log("diff length")
        return false
    }

    // Handle primitived
    if (item1Type !== 'object' && item2Type !== 'object') {
        return item1 === item2
    }

    // Handle arrays
    if (Array.isArray(item1) && Array.isArray(item2)) {
        if (item1.length !== item2.length) return false

        for (let i = 0; i < item1.length;i++) {
            if(deepEqual(item1[i], item2[i]) === false) {
                return false
            }
        }

        return true
    }

    // Handle objects
    if (!Array.isArray(item1) && !Array.isArray(item2)) {
        const item1Keys = Object.keys(item1)
        const item2Keys = Object.keys(item2)

        const item1Value = Object.values(item1)
        const item2Value = Object.values(item2)

        return deepEqual(item1Keys, item2Keys) && deepEqual(item1Value, item2Value)
    }

}



const res1 = deepEqual(null, undefined)
console.log(res1)

