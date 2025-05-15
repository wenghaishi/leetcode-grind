const arrangeInt = (input) => {
  // convert integer arr to no string
  const strArr = input.map((e) => {
    return e.toString();
  });

  const sortedArr = strArr.sort((a, b) => {
    const stringA = parseInt(`${a}${b}`);
    const stringB = parseInt(`${b}${a}`);

    if (stringA > stringB) return 1;

    return -1;
  });

  return sortedArr;
};

console.log(arrangeInt([20, 10]));

console.log(arrangeInt[(50, 1, 2, 8)]);
