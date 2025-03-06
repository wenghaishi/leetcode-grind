const twoSum = (num, target) => {
  const myMap = new Map()

  for (let i = 0; i < num.length; i++) {
    if (myMap.has(target - num[i])) {
        return [myMap.get(target - num[i]), i]
    } else {
        myMap.set(num[i], i)
    }
  }
}