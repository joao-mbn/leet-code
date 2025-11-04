/**
 * Definition for isBadVersion()
 *
 * @param {integer} version number
 * @return {boolean} whether the version is bad
 * isBadVersion = function(version) {
 *     ...
 * };
 */

/**
 * @param {function} isBadVersion()
 * @return {function}
 */
var solution = function(isBadVersion) {
  /**
   * @param {integer} n Total versions
   * @return {integer} The first bad version
   */
  return function(n) {
    let start = 0
    let end = n

    while (true) {
      const halfway = Math.ceil((start+end)/2)
      const isCurrentBad = isBadVersion(halfway)
      const isPrevBad = isBadVersion(halfway-1)

      if (isCurrentBad && !isPrevBad) {
        return halfway
      }

      if (isCurrentBad) {
        end = halfway
      } else {
        start = halfway
      }
    }
  };
};