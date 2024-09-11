const flattenArr = (arr, n) => {
  const res = [];

  const helper = (arr, depth) => {
    for (const element of arr) {
      if (typeof element === "object" && depth < n) {
        helper(element, depth + 1);
      } else {
        res.push(element);
      }
    }
  };

  helper(arr, 0);

  return res;
};
