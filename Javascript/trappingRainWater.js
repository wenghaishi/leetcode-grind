// Slow solution

var trap = function (height) {
  let mH = height;
  let total = 0;
  let max = Math.max(...height);

  while (max > 0) {
    let leftBound = 0;
    let rightBound = 0;
    for (let i = 0; i < mH.length; i++) {
      if (mH[i] > 0) {
        leftBound = i;
        break;
      }
    }
    for (let j = mH.length - 1; j > 0; j--) {
      if (mH[j] > 0) {
        rightBound = j;
        break;
      }
    }

    console.log("max:", max, "LB: ", leftBound, "RB: ", rightBound);

    for (let k = 1; k < mH.length - 1; k++) {
      if (mH[k] === 0 && k > leftBound && k < rightBound) {
        total++;
      }
    }

    for (let j = 0; j < mH.length; j++) {
      if (mH[j] > 0) {
        mH[j] = mH[j] - 1;
      }
    }
    max--;
  }

  return total;
};

// Optimum solution

const trapWater = (height) => {
  let leftMax = [];
  let rightMax = [];
  const arrLength = height.length;
  let result = 0;

  let max = height[0];

  for (let i = 0; i < arrLength; i++) {
    if (i === 0 || i === arrLength) {
      leftMax.push(0);
    } else {
      leftMax.push(max);
      if (height[i] > max) {
        max = height[i];
      }
    }
  }

  max = height[arrLength - 1];

  for (let j = arrLength - 1; j >= 0; j--) {
    console.log("max:", max);
    console.log(j);
    if (j === 0 || j === arrLength - 1) {
      rightMax = [max, ...rightMax];
    } else {
      rightMax = [max, ...rightMax];
      if (height[j] > max) {
        max = height[j];
      }
    }
  }

  console.log("left max: ", leftMax);
  console.log("right max: ", rightMax);

  for (let k = 0; k < arrLength; k++) {
    const trappedWater = Math.min(leftMax[k], rightMax[k]) - height[k];
    console.log("trapped water: ", trappedWater);

    if (trappedWater > 0) {
      result += trappedWater;
    }
  }

  return result;
};

console.log("res:", trapWater([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]));
