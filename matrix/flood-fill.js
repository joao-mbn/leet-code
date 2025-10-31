/**
 * @param {number[][]} image
 * @param {number} sr
 * @param {number} sc
 * @param {number} color
 * @return {number[][]}
 */
var floodFill = function(image, sr, sc, color) {
  const stack = [[sr, sc]]
  const originalValue = image[sr][sc]

  while (stack.length) {
    const cell = stack.pop()
    const unvisited = getNeighbors(image, ...cell).filter(([nsr, nsc]) => image[nsr][nsc] === originalValue)

    stack.push(...unvisited)
    image[cell[0]][cell[1]] = -1

    for (const [usr, usc] of unvisited) {
      image[usr][usc] = -1
    }
  }

  for (let i = 0; i < image.length; i++) {
    for (let j = 0; j < image[0].length; j++) {
      if (image[i][j] === -1) {
        image[i][j] = color
      }
    }
  }

  return image
};

function getNeighbors(image, sr, sc) {
  const neighbors = []
  if (sr > 0) {
    neighbors.push([sr-1,sc])
  }
  if (sr < image.length-1) {
    neighbors.push([sr+1,sc])
  }
  if (sc > 0) {
    neighbors.push([sr,sc-1])
  }
  if (sc < image[0].length-1) {
    neighbors.push([sr,sc+1])
  }

  return neighbors
}

floodFill([[1,1,1],[1,1,0],[1,0,1]], 1, 1, 2)