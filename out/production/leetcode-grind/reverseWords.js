// Given an input string s, reverse the order of the words.

// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

// Return a string of the words in reverse order concatenated by a single space.

const s1 = "   hi jack";

const s2 = "hi jack  james ";

const solution = (s) => {
  const newArr = s.split(" ");
  let removeEmpty = [];
  newArr.forEach((element) => {
    if (element.length > 0) {
      removeEmpty.push(element);
    }
  });
  let upper = removeEmpty.length - 1;
  let lower = 0;
  while (upper > lower) {
    let tem = removeEmpty[lower];
    removeEmpty[lower] = removeEmpty[upper];
    removeEmpty[upper] = tem;
    lower++;
    upper--;
  }

  return removeEmpty.join().replaceAll(",", " ");
};


console.log(solution(s1));

console.log(solution(s2));



