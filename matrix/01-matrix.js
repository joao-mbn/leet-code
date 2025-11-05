/**
 * @param {number[][]} mat
 * @return {number[][]}
 */
var updateMatrix = function(mat) {
  const output = []
  const queue = []

  for (let i = 0; i < mat.length; i++) {
    for (let j = 0; j < mat[0].length; j++) {
      output[i] ??= []

      if (mat[i][j] === 0) {
        output[i][j] = 0
        queue.push([i,j])
      } else {
        output[i][j] = -1
      }
    }
  }

  while (queue.length) {
    const [i, j] = queue.shift()

    const north = [i-1, j]
    const south = [i+1, j]
    const west = [i, j-1]
    const east = [i, j+1]

    for (const [di, dj] of [north, south, west, east]) {
      if (output[di]?.[dj] === -1) {
        output[di][dj] = output[i][j] + 1
        queue.push([di, dj])
      }
    }
  }

  return output
};