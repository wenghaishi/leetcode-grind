// Given two promises promise1 and promise2, return a new promise. 
// promise1 and promise2 will both resolve with a number. 
// The returned promise should resolve with the sum of the two numbers.

const addTwoPromises = async (p1, p2) => {
  const [num1, num2] = await Promise.all([p1, p2]);
  return num1 + num2;
};
