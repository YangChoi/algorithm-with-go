/**
flood fill: 색칠 알고리즘
연결되어 있는 픽셀을 시작 픽셀(sr, sc)과 같은 색(값)으로 새로 칠한다.
그렇지 않을 경우엔 칠하지 않는다.

재귀 이용
*/
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	// 시작픽셀의 색
	if color != newColor {
		// 아직 색을 칠하지 않았다면?
		changeColor(image, sr, sc, &color, &newColor)
		// 색을 칠한다.
	}
	return image
}

func changeColor(image [][]int, x, y int, color, newColor *int) {
	if image[x][y] == *color {
		image[x][y] = *newColor
		if x >= 1 {
			changeColor(image, x-1, y, color, newColor)
		}
		if y >= 1 {
			changeColor(image, x, y-1, color, newColor)
		}
		if x+1 < len(image) {
			changeColor(image, x+1, y, color, newColor)
		}
		if y+1 < len(image[0]) {
			changeColor(image, x, y+1, color, newColor)
		}
	}
}