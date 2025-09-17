/**
 * @param {number} num
 * @return {string}
 */
var intToRoman = function (num) {
  const decimals = {
    0: ["I", "V"],
    1: ["X", "L"],
    2: ["C", "D"],
    3: ["M"],
  };

  let roman = "";

  const size = num.toString().length;

  for (let i = 0; i < size; i++) {
    const digit = Number(num.toString()[i]);
    const nDecimals = size - i - 1;

    if (digit < 4) {
      for (let j = 0; j < digit; j++) {
        roman += decimals[nDecimals][0];
      }
    } else if (digit === 4) {
      roman += decimals[nDecimals][0];
      roman += decimals[nDecimals][1];
    } else if (digit === 5) {
      roman += decimals[nDecimals][1];
    } else if (digit < 9) {
      roman += decimals[nDecimals][1];
      for (let j = 0; j < digit - 5; j++) {
        roman += decimals[nDecimals][0];
      }
    } else {
      roman += decimals[nDecimals][0];
      roman += decimals[nDecimals + 1][0];
    }
  }

  return roman;
};
