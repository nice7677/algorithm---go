package main

import "fmt"

/**
Depth First Search(깊이 우선 탐색)
깊은 부분을 우선적으로 탐색하는 알고리즘

DFS는 스택 자료구조(혹은 재귀 함수)를 이용한다
1. 탐색 시작 노드를 스택에 삽입하고 방문 처리
2. 스택의 최상단 노드에 방문하지 않은 인접 노드가 하나라도 있으면 그 노드를 스택에 넣고 방문 처리한다.
   방문하지 않은 인접 노드가 없으면 스택에서 최상단 노드를 꺼낸다.
3. 더 이상 2번의 과정을 수행할 수 없을 때까지 반복한다.
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
	dfs(graph, 1, visited)
}

/**
  DFS가 어떻게 진행되는지 보는 연습
*/
func dfs(graph [][]int, v int, visited []bool) {
	// 처음에 들어오면 방문 처리
	visited[v] = true
	// 탐색 순서 확인 하기
	fmt.Println("Current visited node number is", v)
	// 노드 내부 연결된 노드들 확인
	for _, item := range graph[v] {
		// 방문 한적이없는 노드는 방문
		if !visited[item] {
			dfs(graph, item, visited)
		}
	}
}
