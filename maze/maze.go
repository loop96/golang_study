package main

import (
	"fmt"
	"os"
)

func readMaze(filepath string) [][]int {
	file, err := os.Open(filepath)
	if nil != err {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	fmt.Printf("maze row=%d,col=%d\n", row, col)
	maze := make([][]int, row)
	for i := 0; i < row; i++ {
		maze[i] = make([]int, col)
		for j := 0; j < col; j++ {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

func checkPrint(maze [][]int) {
	for _, rowsData := range maze {
		for _, cel := range rowsData {
			fmt.Printf("%2d  ", cel)
		}
		fmt.Println()
	}
}

// 点位构造体
type point struct {
	i, j int
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

// at 返回迷宫中的点位数据和是否越界
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

// 点位移动的四个方向 上 左 下 右
var directions = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

// walk
// maze 迷宫数据
// start,end 出入口
func walk(maze [][]int, start, end point) [][]int {
	//定义走过的路线
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	//定义待探索的队列
	Q := []point{start}

	//如果待探索的队列为空就结束
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		//如果当前位置在出口，停止探索
		if cur == end {
			break
		}
		//对当前点位四个方向进行探索
		for _, d := range directions {
			next := cur.add(d)
			// next 只能在maze不能碰到墙壁(不为1的点),且不能为开始的点位
			if val, ok := next.at(maze); !ok || next == start || val == 1 {
				continue
			}
			// next 只能在steps走过没走过的路(为0的点),且不能为开始的点位
			if val, ok := next.at(steps); !ok || next == start || val != 0 {
				continue
			}
			//给steps地图中新增路径
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	//读取maze.in成二维切片
	maze := readMaze("maze/maze.in")
	//打印校验一下
	checkPrint(maze)
	fmt.Println("===walk===")
	steps := walk(maze, point{0, 0}, point{5, 6})
	checkPrint(steps)
}
