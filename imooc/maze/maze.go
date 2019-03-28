package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d\n", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
		fmt.Fscanf(file, "\n")
	}

	return maze
}

type point struct {
	i, j int
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	// 若超出范围
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

// 往上左下右四个方向移动时每一步的步长
var dirs = [4]point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func walk(maze [][]int, start, end point) [][]int {
	// 记录走过的步子
	// 行列和maze相同, 全部填充0
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	queue := []point{start}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)
			val, ok := next.at(maze)
			// 越界或者撞墙
			if !ok || val == 1 {
				continue
			}
			val, ok = next.at(steps)
			// 越界或者走过
			if !ok || val != 0 {
				continue
			}
			// 回到起点
			if next == start {
				continue
			}
			// 当前格子的值, 起点的话就是0
			curStep, _ := cur.at(steps)
			// 下一步格子的值加1
			steps[next.i][next.j] = curStep + 1

			// 加入队列
			queue = append(queue, next)
		}
	}

	return steps
}

func main() {
	// 广度优先搜索算法实现走迷宫
	// 加载迷宫
	maze := readMaze("imooc/maze/maze.in")
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%2d ", val)
		}
		fmt.Println()
	}
	fmt.Println()
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%2d ", val)
		}
		fmt.Println()
	}

}
