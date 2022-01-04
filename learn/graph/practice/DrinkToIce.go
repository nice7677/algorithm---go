package main

import (
	"fmt"
	"github.com/oleiade/lane"
)

var (
	n         = 4
	m         = 5
	dfsResult = 0 // 얼음의 처음 갯수는 0
	bfsResult = 0 // 얼음의 처음 갯수는 0
	dfsGraph  = [][]int{
		{0, 0, 1, 1, 0}, // 0
		{0, 0, 0, 1, 1}, // 1
		{1, 1, 1, 1, 1}, // 2
		{0, 0, 0, 0, 0}, // 3
	}
	bfsGraph = [][]int{
		{0, 0, 1, 1, 0}, // 0
		{0, 0, 0, 1, 1}, // 1
		{1, 1, 1, 1, 1}, // 2
		{0, 0, 0, 0, 0}, // 3
	}
)

/**
음료수 얼려 먹기

연결 요소 찾기
*/
func main() {
	for i := 0; i < n; i++ { // 노드만큼 for loop (1)
		for j := 0; j < m; j++ { // (1)의 연결 노드 만큼 for loop (2)
			if dfs(i, j) {
				dfsResult++ // 덩어리를 찾았다는것
			}
		}
	}
	fmt.Println("dfs :", dfsResult)
	for i := 0; i < n; i++ { // 노드만큼 for loop (1)
		for j := 0; j < m; j++ { // (1)의 연결 노드 만큼 for loop (2)
			if bfs(i, j) {
				bfsResult++ // 덩어리를 찾았다는것
			}
		}
	}
	fmt.Println("bfs :", bfsResult)
}

func dfs(x int, y int) bool {
	// 조건내 범위를 벗어나면 false
	if x <= -1 || x >= n || y <= -1 || y >= m {
		return false
	}
	// 아직 방문하지 않으면 0 값이다.
	if dfsGraph[x][y] == 0 {
		// 방문했으니 1로 꾼다
		dfsGraph[x][y] = 1
		/**
		방문을 했으니 +1은 확정,
		그리고 이제 덩어리를 전부 찾아 바꿔준다.
		*/
		// 아래로 찾아본다
		dfs(x+1, y)
		// 위로 찾아본다
		dfs(x-1, y)
		// 우로 찾아본다
		dfs(x, y+1)
		// 좌로 찾아본다
		dfs(x, y-1)

		/**
		graph[x][y]가 0일때 true는 확정
		확정이고 위아래로 자기와 파트너가 있는지 찾는것.
		찾으면 파트너들을 전부 방문처리 후 main function의 for loop(line 32:)으로 이동
		그리고 자기와 파트너들을 전부 찾았으니 result를 +1 해줌
		*/
		return true
	}
	// 이미 값이 1, 방문이 됬다고 된넘들은 false
	return false
}

var (
	dx   = []int{1, -1, 0, 0}
	dy   = []int{0, 0, 1, -1}
	dLen = len(dx)
)

// 시작값 x, y
func bfs(x int, y int) bool {
	// if 1이면 false
	if bfsGraph[x][y] == 1 {
		return false
	}
	// 큐 생성
	queue := lane.NewQueue()
	// 큐 삽입
	queue.Enqueue([]int{x, y})
	// 큐에 값이 없을때까지 for loop
	for !queue.Empty() {
		// ------------ (1) ------------
		// 큐 Dequeue
		v := queue.Dequeue().([]int)
		x = v[0]           // loop 용 x
		y = v[1]           // loop 용 y
		bfsGraph[x][y] = 1 // x,y 방문 처리
		// 상하좌우 만큼 for loop, because bfs do not use recursive functions?
		for i := 0; i < dLen; i++ { // dLen is 4 (line :82)
			/** 검색해본다.
			i = 0 밑
			i = 1 위
			i = 2 우
			i = 3 좌
			*/
			nx := x + dx[i]
			ny := y + dy[i]
			// 범위에서 벗어나는 값, 그리고 방문하지 않은 값 0 확인
			if nx >= 0 && nx < n && ny >= 0 && ny < m && bfsGraph[nx][ny] == 0 {
				// 큐 삽입 후 for loop (1)로 복귀
				queue.Enqueue([]int{nx, ny})
			}
		}
	}
	/**
	for loop이 시작 된 순간 방문돼지않은 노드임.
	return true 확정 후 queue로 돌아가면서 노드들 검색 후 방문처리로 덩어리로 만듬
	그리고 return 되면 main function에서 value 0을 찾아서 다시 돌림. 덩어리 +1
	*/
	return true

}
