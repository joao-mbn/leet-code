/**
 * @param {number[][]} points
 * @param {number} k
 * @return {number[][]}
 */
var kClosest = function(points, k) {
  const swapTarget = k - 1
  let swapIndex = -1
  let pivotIndex = points.length - 1

  while (true) {
    const pivotCoordinates = points.at(pivotIndex)
    const pivotDistance = Math.sqrt(
      Math.pow(pivotCoordinates[0], 2) +
      Math.pow(pivotCoordinates[1], 2)
    )
    const swapIndexBefore = swapIndex

    for (let i = swapIndex + 1; i <= pivotIndex; i++) {
      const [x, y] = points[i]
      const distance = Math.sqrt(
        Math.pow(x, 2) +
        Math.pow(y, 2)
      )

      if (distance > pivotDistance) {
        continue
      }

      swapIndex++
      if (swapIndex === i) {
        continue
      }

      points[i] = points[swapIndex]
      points[swapIndex] = [x, y]
    }

    if (swapIndex === swapTarget) {
      return points.slice(0, k)
    } else if (swapIndex > swapTarget) {
      pivotIndex = swapIndex - 1
      swapIndex = swapIndexBefore
    }
  }
};

console.log(kClosest([[1,3],[-2,2]], 1))
console.log(kClosest([[3,3],[5,-1],[-2,4]], 2))
console.log(kClosest([[-95,76],[17,7],[-55,-58],[53,20],[-69,-8],[-57,87],[-2,-42],[-10,-87],[-36,-57],[97,-39],[97,49]], 5))