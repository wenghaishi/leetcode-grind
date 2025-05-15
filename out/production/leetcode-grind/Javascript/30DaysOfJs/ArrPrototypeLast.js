// Write code that enhances all arrays such that you can call the array.last() method on any array
// and it will return the last element.
// If there are no elements in the array, it should return -1.

Array.prototype.last = function () {
  const length = this.length;
  if (length === 0) return -1;
  return this[length - 1];
};

const arr = [1, 2, 3];

console.log(arr.last());
