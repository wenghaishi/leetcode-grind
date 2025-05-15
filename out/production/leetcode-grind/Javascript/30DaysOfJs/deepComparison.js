const deepComparison = (object1, object2) => {
  if (object1.length !== object2.length) return false;
  if (typeof object1 !== typeof object2) return false;

  // both are objects
  if (typeof object1 === "object") {
    const object1Keys = Object.keys(object1);
    const object2Keys = Object.keys(object2);

    for (let i = 0; i < object1Keys.length; i++) {
      if (object1Keys[i] !== object2Keys[i]) {
        return false;
      }

      if (object1[object1Keys[i]] !== object2[object2Keys[i]]) {
        return false;
      }
    }
  } else if (typeof object1 === "array") {
    for (let j = 0; j < object1.length; j++) {
      if (object1[j] !== object2[j]) {
        return false;
      }
    }
  } else {
    if (object1 !== object2) {
      return false;
    }
  }

  //

  return true;
};

const test1 = { key1: "value1", key2: "value2" };
const test2 = { key1: "value1", key3: "value2" };
const test3 = { key1: "value1", key3: "value2" };

const test4 = [1, 2, 3];
const test5 = [1, 2, 4];
const test8 = [1, 2, 4];


const test6 = 1
const test7 = 5

// equal
console.log(deepComparison(test1, test1));

// different value
console.log(deepComparison(test1, test3));

// different key
console.log(deepComparison(test1, test2));

// equal arr
console.log(deepComparison(test4, test4));

// different value array
console.log(deepComparison(test4, test5));

// different num
console.log(deepComparison(test6, test7));

console.log(test5 === test8)
