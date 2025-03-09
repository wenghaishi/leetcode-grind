const hasPathSum = function (root, targetSum) {
  if (!root) return false;
  if (!root.left && !root.right) {
    if (targetSum !== root.val) {
      return false;
    }
    return true;
  }

  if (root.left && root.right) {
    return (
      hasPathSum(root.left, targetSum - root.val) ||
      hasPathSum(root.right, targetSum - root.val)
    );
  }

  if (root.left && !root.right) {
    return hasPathSum(root.left, targetSum - root.val);
  }

  if (root.right && !root.left) {
    return hasPathSum(root.right, targetSum - root.val);
  }
};
