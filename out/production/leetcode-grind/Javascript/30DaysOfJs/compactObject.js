const compactObj = (obj) => {

  const helper = (item) => {
    if (Array.isArray(item)) {
      // Filter out falsy values from arrays and process inner objects/arrays
      return item
        .filter(Boolean) // Remove falsy values
        .map(helper);    // Recursively process inner objects/arrays
    } else if (typeof item === 'object' && item !== null) {
      const result = {};
      for (const key in item) {
        if (item[key]) { // Only keep truthy values
          result[key] = helper(item[key]); // Recursively process inner objects/arrays
        }
      }
      return result;
    } else {
      return item; // For non-object, non-array types, return the value as is
    }
  };

  return helper(obj);

};
