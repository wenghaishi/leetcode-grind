// There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].

// You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to
// its next (i + 1)th station. You begin the journey with an empty tank at one of the gas stations.

// Given two integer arrays gas and cost, return the starting gas station's index if you can travel around
// the circuit once in the clockwise direction, otherwise return -1. If there exists a solution, it is guaranteed to be unique

// difficulty: medium

const gas = [1, 2, 3, 4, 5];
const cost = [3, 4, 5, 1, 2];

const canCompleteCircuit = (gas, cost) => {
  let curTank = 0,
    totalTank = 0,
    pos = 0;
  for (let i = 0; i < gas.length; i++) {
    curTank += gas[i] - cost[i];
    totalTank += gas[i] - cost[i];
    if (curTank < 0) {
      curTank = 0;
      pos = i + 1;
    }
  }
  return totalTank < 0 ? -1 : pos;
};
