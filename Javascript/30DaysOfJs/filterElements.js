// Given an integer array arr and a filtering function fn, 
// return a filtered array filteredArr.

// The fn function takes one or two arguments:

// arr[i] - number from the arr
// i - index of arr[i]
// filteredArr should only contain the elements from the arr for which 
// the expression fn(arr[i], i) evaluates to a truthy value. 
// A truthy value is a value where Boolean(value) returns true.

// Please solve it without the built-in Array.filter method.



const filter = function (arr, fn) {
  let resultArr = [];
  for (let i = 0; i < arr.length; i++) {
    const resultE = fn(arr[i], i);
    if (resultE) {
      resultArr.push(e);
    }
  }
  return resultArr;
};
