function pathSum(root, target) {
  if (!root) return [];

  const result = [];

  function dfs(root, arr, currentRemainder) {
    const currentPath = [...arr, root.val];

    if (!root.left && !root.right) {
      // check if remainder - root.val === 0
      if (root.val === currentRemainder) {
        result.push(currentPath);
      }
    }

    if (root.left) {
      dfs(root.left, currentPath, currentRemainder - root.val);
    }

    if (root.right) {
      dfs(root.right, currentPath, currentRemainder - root.val);
    }
  }

  dfs(root, [], target);

  return result;
}

// Test case
class TreeNode {
  constructor(val = 0, left = null, right = null) {
    this.val = val;
    this.left = left;
    this.right = right;
  }
}

const root = new TreeNode(5);
root.left = new TreeNode(4);
root.right = new TreeNode(8);
root.left.left = new TreeNode(11);
root.left.left.left = new TreeNode(7);
root.left.left.right = new TreeNode(2);
root.right.left = new TreeNode(13);
root.right.right = new TreeNode(4);
root.right.right.right = new TreeNode(1);

console.log(pathSum(root, 22)); // Expected: [[5,4,11,2], [5,8,4,1]]
