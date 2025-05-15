// Given a function fn, an array of arguments args, and a timeout t in milliseconds, return a cancel function cancelFn.

// After a delay of cancelTimeMs, the returned cancel function cancelFn will be invoked.


const cancellable = async (fn, args, t) => {
  const timer = setTimeout(() => {
    fn(...args);
  }, t);

  const cancelTimer = () => clearTimeout(timer);

  return cancelTimer;
};
