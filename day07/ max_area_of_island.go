/**
섬은 수평, 수직 상관없이 4방향으로 연결된 1의 모임
따라서 물(0)으로 둘러쌓인 grid의 4개의 꼭지점을 추정할 수 있다.
섬 중 가장 면적이 넓은 grid를 가진 것을(1의 모임 중 가장 큰) 찾아 리턴
*/
func maxAreaOfIsland(grid [][]int) int {
	max := 0
	height, width := len(grid), len(grid[0])
	visited := make([][]int, height)
	for i := range visited {
		visited[i] = make([]int, width)
	}

	for i, row := range grid {
		for j := range row {
			area := getArea(i, j, grid, visited, height, width)
			if area > max {
				max = area
			}
		}
	}
	return max
}

func getArea(i int, j int, grid [][]int, visited [][]int, height int, width int) int {
	if i < 0 || i >= height || j < 0 || j >= width || grid[i][j] != 1 || visited[i][j] == 1 {
		// 바다의 조건들
		return 0
	}
	area := 1
	visited[i][j] = 1
	// 그 외에는 섬, 섬은 방문한 흔적을 남긴다
	area += getArea(i-1, j, grid, visited, height, width)
	area += getArea(i+1, j, grid, visited, height, width)
	area += getArea(i, j-1, grid, visited, height, width)
	area += getArea(i, j+1, grid, visited, height, width)

	return area
}