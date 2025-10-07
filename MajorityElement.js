// 169. Majority Element

// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. 
// You may assume that the majority element always exists in the array.

const majorityE = (arr) => {
  const min = arr.length / 2;
  let map = {};

  for (let i = 0; i < arr.length; i++) {
    const obj = arr[i];
    if (map[obj]) {
      map[obj] = map[obj] + 1;
      if (map[obj] >= min) return obj;
    } else {
      map[obj] = 1;
    }
  }
};

console.log(majorityE([1, 1, 1, 2]));
console.log(majorityE([1, 1, 4]));
console.log(majorityE([1, 1, 2, 2]));
console.log(majorityE([1, 1, 1]));
