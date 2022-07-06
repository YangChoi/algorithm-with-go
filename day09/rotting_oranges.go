/**
0은 빈 셀,
1은 신선한 오랜지,
2는 썩은 오랜지이다.
매 분 마다 신선한 오랜지는 4방향으로 인접한 썩은 오랜지를 가지게 된다.
신선한 오랜지가 남지 않을 때까지 최소 몇분이 걸릴 지 알아내라.
만약 불가능하다면 -1 반환
*/
/**
0은 빈 셀,
1은 신선한 오랜지,
2는 썩은 오랜지이다.
매 분 마다 신선한 오랜지는 4방향으로 인접한 썩은 오랜지를 가지게 된다.
신선한 오랜지가 남지 않을 때까지 최소 몇분이 걸릴 지 알아내라.
만약 불가능하다면 -1 반환
*/
type Coordinate struct {
	X int
	Y int
}

type Grid struct {
	C        [][]Coordinate
	Quantity int
	Count    int
	Flag     bool
}

func orangesRotting(grid [][]int) int {

	g := new(Grid)
	g.Quantity = 0
	g.Count = 0
	g.Flag = true

	tmpC := []Coordinate{}
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {

			//How many fresh oranges?
			if grid[j][i] == 1 {
				g.Quantity++

				//How many rotten oranges?
			} else if grid[j][i] == 2 {
				tmpC = append(tmpC, Coordinate{i, j})
			}

		}
	}

	if g.Quantity == 0 || tmpC == nil {
		return 0
	}

	g.C = append(g.C, tmpC)

	//checking function
	g.check(grid)

	//If fresh oranges remain.
	if g.Quantity > 0 {
		return -1
	}

	return g.Count - 1
}

//check() look for rotten oranges.
func (g *Grid) check(grid [][]int) {

	if g.Flag {

		g.Flag = false
		arr := []Coordinate{}

		for i := 0; i < len(g.C[g.Count]); i++ {

			x := g.C[g.Count][i].X
			y := g.C[g.Count][i].Y

			//left
			if x >= 1 {
				g.setData(grid, x-1, y, &arr)
			}
			//bottom
			if y >= 1 {
				g.setData(grid, x, y-1, &arr)
			}
			//right
			if x+1 < len(grid[0]) {
				g.setData(grid, x+1, y, &arr)
			}
			//top
			if y+1 < len(grid) {
				g.setData(grid, x, y+1, &arr)
			}

		}

		g.C = append(g.C, arr)
		g.Count++

		//Check again.
		g.check(grid)
	}
}

func (g *Grid) setData(grid [][]int, x, y int, arr *[]Coordinate) {
	if grid[y][x] == 1 {
		grid[y][x] = 2
		g.Quantity--
		*arr = append(*arr, Coordinate{x, y})
		g.Flag = true
	}
}
