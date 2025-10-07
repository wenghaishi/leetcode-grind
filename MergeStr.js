// You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, 
//starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

// Return the merged string.

const mergeStr = (w1, w2) => {
  const arr1 = w1.split('')
  const arr2 = w2.split('');
  let result = []

  for (let i = 0; i < Math.max(arr1.length, arr2.length); i++) {
    const l1 = arr1[i]
    const l2 = arr2[i]

    if (l1) {
      result.push(l1);
    }
    if (l2) {
      result.push(l2);
    }
  }

  return result.join();
}

console.log(mergeStr('02468', '13579'))
console.log(mergeStr("aa", 'bbbbbb'))
console.log(mergeStr('', 'aaa'))

