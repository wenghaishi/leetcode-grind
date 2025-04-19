Array.prototype.flattenArr = function flatten() {
  let result = [];

  function flattenElement(e) {
    console.log(typeof e);
    if (typeof e != "object") {
      result.push(e);
    } else {
      for (const ele of e) {
        flattenElement(ele);
      }
    }
  }

  for (const element of this) {
    flattenElement(element);
  }

  return result;
};

const arr = [1, 2, [1, 2, 3, [1, 2, 3]]];

const flattenedArr = arr.flattenArr();

console.log(flattenedArr);
