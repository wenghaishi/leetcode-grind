// 49. Group Anagrams

const groupAnagrams = (strs) => {
  if (!strs) return;
  const m = new Map();
  for (const str of strs) {
    const sorted = str.split("").sort().join();
    const current = m.get(sorted) || [];
    m.set(sorted, [str, ...current]);
  }

  let result = [];

  for (const key of m.keys()) {
    result.push(m.get(key));
  }

  return result;
};

console.log(groupAnagrams(["ate", "eat", "bat"]));
