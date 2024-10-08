// Given a signed 32-bit integer x, return x with its digits reversed.
// If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

var reverse = function (x) {
  const MIN_INT32 = -2147483648;
  const MAX_INT32 = 2147483647;

  const numAsString = x.toString();
  const stringArray = numAsString.split("");
  let left = stringArray[0] === "-" ? 1 : 0;
  let right = stringArray.length;

  while (left < right) {
    let temp = stringArray[left];
    stringArray[left] = stringArray[right];
    stringArray[right] = temp;
    left++;
    right--;
  }

  const reversedNum = parseInt(stringArray.join(""));

  if (reversedNum < MIN_INT32 || reversedNum > MAX_INT32) return 0;
  return reversedNum;
};
