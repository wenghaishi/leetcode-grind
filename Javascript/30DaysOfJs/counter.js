// Given an integer n, return a counter function.
// This counter function initially returns n and then returns
// 1 more than the previous value every subsequent time it is called (n, n + 1, n + 2, etc).

var createCounter = function (n) {
  return function () {
    return n++;
  };
};

const counter = createCounter(10);
console.log(counter());
console.log(counter());
console.log(counter());

// explanation

// In this case, n is like a private variable in OOP the inner function still has access to n
