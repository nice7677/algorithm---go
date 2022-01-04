package main

import (
	"fmt"
	"github.com/oleiade/lane"
)

/**
Breadth-First Search(너비 우선 탐색)

그래프에서 가까운 노드부터 우선적으로 탐색하는 알고리즘

큐를 이용하며 동작 과정은 다음과 같음.

1. 탐색 시작 노드를 큐에 삽입하고 방문 처리를 함.
2. 큐에서 노드를 꺼낸 뒤에 해당 노드의 인접 노드 중에서 방문하지 않은 노드를 모두 큐에 삽입하고 방문처리함.
3. 더 이상 2번의 과정을 수행할 수 없을 때까지 반복.

최단거리,
재귀 사용x
*/
func main() {
	/**
	각 노드의 연결 정보
	graph[i] <- i가 노드의 번호
	ex) graph[1]은 1번 노드이고 연결된 노드는 2,3,8
	*/
	graph := [][]int{
		{},
		{2, 3, 8},
		{1, 7},
		{1, 4, 5},
		{3, 5},
		{3, 4},
		{7},
		{2, 6, 8},
		{1, 7},
	}
	visited := make([]bool, len(graph))
	bfs(graph, 1, visited)
}

func bfs(graph [][]int, start int, visited []bool) {
	queue := lane.NewQueue()
	// 큐에 노드 넣기
	queue.Enqueue(start)
	// 해당 노드 방문 처리
	visited[start] = true
	for queue.Size() > 0 { // queue가 빌때까지
		// --------- (1) --------- //
		// 큐에서 노드 꺼내기
		v := queue.Dequeue().(int)
		fmt.Println("Current dequeue element is ", v)
		// 노드랑 연결된 노드들 하나씩 방문 확인하면서 처리
		for _, item := range graph[v] {
			// 방문하지 않았을때 방문
			if !visited[item] {
				// 큐에 노드 추가
				queue.Enqueue(item)
				// 해당 노드 방문 처리
				visited[item] = true
			}
			// 연결된 큐가 끝나면 다시 (1)로 이동 후 큐 비우면서 방문 처리
		}
	}
}
