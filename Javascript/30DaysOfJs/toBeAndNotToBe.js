// Write a function expect that helps developers test their code.
// It should take in any value val and return an object with the following two functions.

// toBe(val) accepts another value and returns true if the two values === each other.
// If they are not equal, it should throw an error "Not Equal".
// notToBe(val) accepts another value and returns true if the two values !== each other. If they are equal,
// it should throw an error "Equal".

const expect = (val) => {
  const toBe = (val2) => {
    if (val === val2) {
      return true;
    } else {
      return false;
    }
  };

  const notToBe = (val2) => {
    if (val === val2) {
      return true;
    } else {
      return false;
    }
  };

  return { toBe, notToBe };
};

console.log(expect(5).toBe(5));
console.log(expect(5).toBe(4));
