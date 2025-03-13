Promise.resAll = function promiseAll(promises) {
  return new Promise((resolve, reject) => {
    let completed = 0;
    let results = [];

    promises.forEach((p) => {
      Promise.resolve(p)
        .then((res) => {
          results.push(res);
          completed++;
          if (completed === promises.length) {
            resolve(results);
          }
        })
        .catch((err) => reject(err));
    });
  });
};

const apis = [
  "https://jsonplaceholder.typicode.com/posts/1",
  "https://jsonplaceholder.typicode.com/posts/2",
  "https://jsonplaceholder.typicode.com/posts/3",
];

const promises = apis.map((api) => fetch(api));

const res = await Promise.resAll(promises);

const data = await res.json();

console.log(data);
