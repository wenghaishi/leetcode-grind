




const solution = (ip_address) => {
  const arr = ip_address.split('.')
  const numArr = arr.map((element) => {
    return parseInt(element)
  })

  for (let i = 0; i < numArr.length; i++) {
    if (numArr[i] < 0 || numArr[i] > 255) {
      return -1
    }
  }
  
  if (numArr[0] < 100) {
    return 1
  } else if ( numArr[0] > 100 && numArr[0] < 200) {
    return 2
  } else {
    return 3
  }

}

console.log(solution("11.222.222.256"))
console.log(solution("155.222.222.222"))
console.log(solution("220.222.222.222"))