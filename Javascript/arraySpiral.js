const arraySpiral = (arr) => {
    const result = []

    function helper(arr) {
        if (arr.length < 1 ) {
            return
        } else if (arr.length === 1) {
            console.log("spread: ", ...arr)
            result.push(...arr[0])
        } else {
            for (let i = 0; i < arr.length; i++) {
                if (i === 0) {
                    result.push(...arr[i])
                } else if (i === arr.length - 1) {
                    const reversedArr = arr[arr.length - 1].reverse()
                    result.push(...reversedArr)
                } else {
                    const lastElement = arr[i][arr[i].length-1]
                    if (lastElement) {
                        result.push(lastElement)
                    }
                }
            }

            if (arr.length > 2) {
                for (let k = arr.length - 2; k > 0; k--) {
                    result.push(arr[k][0])
                }
            }
        }

        const modifiedArr = []

        for (let j = 1; j < arr.length - 1; j++) {
            const slicedArr = arr[j].slice(1, arr[j].length - 1)
            console.log("sliced arr: ", slicedArr)
            modifiedArr.push(slicedArr)
        }

        return helper(modifiedArr)
        
    }

    helper(arr)

    return result
}


const test1 = [
    [1, 2, 3],
    [8, 9, 4],
    [7, 6, 5]
]

console.log(arraySpiral(test1))

