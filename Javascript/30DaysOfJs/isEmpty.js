// 2727. Is Object Empty

// Given an object or an array, return if it is empty.

// An empty object contains no key-value pairs.
// An empty array contains no elements.
// You may assume the object or array is the output of JSON.parse.

const isEmpty = (obj) => {
  const keys = Object.keys(obj);

  if (keys.length === 0) {
    return true;
  } else {
    return false;
  }
};

