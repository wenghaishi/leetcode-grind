// A chunked array contains the original elements in arr, but consists of subarrays each of length size.
// The length of the last subarray may be less than size if arr.length is not evenly divisible by size.

const chunkArr = (arr, size) => {
  let count = 0;
  const result = [];
  let currentArr = [];

  for (let i = 0; i < arr.length; i++) {
    count++;
    currentArr.push(arr[i]);

    if (count === size) {
      count = 0;
      result.push(currentArr);
      currentArr = [];
    }
  }

  if (currentArr.length != 0) result.push(currentArr);

  return result;
};
