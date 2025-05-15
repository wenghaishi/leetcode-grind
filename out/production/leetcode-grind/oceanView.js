// Buildings that are ocean view

// There are n buildings in a line. You are given an integer array heights of size n that represents the
// heights of the buildings in the line. The ocean is to the right of the buildings. A building has an ocean
// view if the building can see the ocean without obstructions. Formally, a building has an ocean view if all
// the buildings to its right have a smaller height. Return a list of indices (0-indexed) of buildings that have
// an ocean view, sorted in increasing order.

// difficulty: medium

const heights = [2, 3, 5, 1, 2, 3];

const isBuildingOceanView = (heights) => {
  let max = 0;
  let output = [];
  for (let i = heights.length - 1; i >= 0; i--) {
    if (heights[i] > max) {
      output.push(heights[i]);
      max = heights[i];
    }
  }
};

// explanation: iterate through the array from the index closest to the ocean. If the current number is more than the
// max so far, push it into the output and assign it as the max

