package algo

/*
There is an mxn
m×n character matrix called grid, which contains different types of cells:

'0' represents an empty land that you can pass through
'X' represents an obstacle that you cannot pass through
'#' represents a bubble tea cup, and there may be multiple cups of bubble tea in the matrix
'*' represents your position, and there is exactly one cell with '*' in the matrix
You need to start from your position and find the shortest path to reach a cell with a cup of bubble tea. Finally, return the length of the shortest path to obtain any cup of bubble tea. If such a path does not exist, return -1.

sample input
grid = [
  "XXXXXX",
  "X*OOOX",
  "XOO#OX",
  "XXXXXX"
]
*/

type node struct {
	row int
	col int
}

func GetBubbleTea(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	queue := make([]node, 0)
	walk := 0

LOOP:
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// find my location
			if grid[i][j] == '*' {
				queue = append(queue, node{row: i, col: j})
				break LOOP
			}
		}
	}

	for len(queue) > 0 {
		walk++
		size := len(queue)
		for i := 0; i < size; i++ {
			n := queue[0]
			queue = queue[1:]
			// check 4 direction
			for j := 0; j < 4; j++ {
				var x int
				var y int
				//left
				if j == 0 {
					x = n.row
					y = n.col - 1
					// up
				} else if j == 1 {
					x = n.row - 1
					y = n.col
					// right
				} else if j == 2 {
					x = n.row
					y = n.col + 1
					// down
				} else {
					x = n.row + 1
					y = n.col
				}
				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
					if grid[x][y] == '#' {
						return walk
					} else if grid[x][y] == 'O' {
						//visited
						grid[x][y] = 'X'
						queue = append(queue, node{row: x, col: y})
					}
				}
			}
		}
	}
	return -1
}
