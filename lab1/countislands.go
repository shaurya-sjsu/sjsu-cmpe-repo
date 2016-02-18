//Shaurya Sharma Lab 1
package main

import (
	"container/list"
)

func CountIslands(grid [][]int) int {
	rows := len(grid)
	columns := len(grid[0])
	var visited = make([][]int, rows)
	for i := range visited {
		visited[i] = make([]int, columns)
	}
	//fmt.Println(visited)
	var i, j, ti, tj, is, temp int
	is = 0
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			if grid[i][j] == 1 && visited[i][j] != 1 {
				visited[i][j] = 1
				is = is + 1
				l := list.New()

				l.PushBack(10*i + j)

				for e := l.Front(); e != nil; e = e.Next() {
					temp = (e.Value).(int)
					ti = temp / 10
					tj = temp % 10
					visited[ti][tj] = 1

					if ti != 0 && grid[ti-1][tj] == 1 && visited[ti-1][tj] != 1 {
						l.PushBack(10*(ti-1) + tj)
					}
					if tj != 0 && grid[ti][tj-1] == 1 && visited[ti][tj-1] != 1 {
						l.PushBack(10*ti + (tj - 1))
					}
					if ti != (rows-1) && grid[ti+1][tj] == 1 && visited[ti+1][tj] != 1 {
						l.PushBack(10*(ti+1) + tj)
					}
					if tj != (columns-1) && grid[ti][tj+1] == 1 && visited[ti][tj+1] != 1 {
						l.PushBack(10*ti + (tj + 1))
					}

				}
			}
		}
	}
	return is
}
