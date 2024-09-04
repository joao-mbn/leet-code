/**
 * @param {string} s
 * @return {number}
 */
var romanToInt = function (s) {
  const translationTable = new Map([
    ["I", 1],
    ["V", 5],
    ["X", 10],
    ["L", 50],
    ["C", 100],
    ["D", 500],
    ["M", 1000],
  ]);

  let total = 0;

  for (let i = 0; i < s.length; i++) {
    const value = translationTable.get(s[i]);
    const nextValue = s[i + 1] && translationTable.get(s[i + 1]);

    if (value < nextValue) {
      total -= value;
    } else {
      total += value;
    }
  }

  return total;
};
