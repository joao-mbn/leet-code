/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function (prices) {
  let profit = 0;
  let buy = prices[0];

  for (let i = 1; i < prices.length; i++) {
    const price = prices[i];

    profit = Math.max(price - buy, profit);
    buy = Math.min(price, buy);
  }

  return profit;
};
