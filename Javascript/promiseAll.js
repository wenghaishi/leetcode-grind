Promise.all = function promiseAll(values) {
  return new Promise((resolve, reject) => {
    const promises = Array.from(values);
    const results = new Array(promises.length);
    let pending = promises.length;

    promises.forEach((promise) => {
      Promise.resolve(promise).then((res) => results.push(res));
    });
  });
};


