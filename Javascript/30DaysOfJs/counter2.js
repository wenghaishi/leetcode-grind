const createCounter = (init) => {
  let counter = init;

  const increment = () => {
    return ++init;
  };

  const decrement = () => {
    return --init;
  };

  const reset = () => {
    init = counter;
    return init;
  };

  return {
    increment,
    decrement,
    reset,
  };
};

const counter = createCounter(5);
console.log(counter.increment());
console.log(counter.increment());
console.log(counter.decrement());
console.log(counter.reset());
console.log(counter.increment());
