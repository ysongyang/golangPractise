package main

//递归的使用案例 迷宫
import "fmt"

//开始找路
func setWay(maps *[8][7]int, i, j int) bool {
	if maps[6][5] == 2 { //找到了终点
		return true
	} else {
		//继续找
		if maps[i][j] == 0 {
			//假设这个点是通的，需要探测
			maps[i][j] = 2
			if setWay(maps, i+1, j+1) { //往右边斜着走
				return true
			} else if setWay(maps, i+1, j) { //i+1,j 往下走
				return true
			} else if setWay(maps, i, j+1) { //i+1,j 往右走
				return true
			} else if setWay(maps, i-1, j) { //i-1,j 往上走
				return true
			} else if setWay(maps, i, j-1) { //i+1,j 往做走
				return true
			} else {
				maps[i][j] = 3
				return false
			}

		} else {
			//1 表示一面墙
			return false
		}
	}
}

//初始化地图
func initMap(maps *[8][7]int) {
	//先把地图最上边最下边设置为1
	for i := 0; i < 7; i++ {
		maps[0][i] = 1
		maps[7][i] = 1
	}

	for i := 0; i < 8; i++ {
		maps[i][0] = 1
		maps[i][6] = 1
	}

	maps[3][1] = 1
	maps[3][2] = 1

	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(maps[i][j], "  ")
		}
		fmt.Println()
	}
}

func main() {
	//创建一个二维数组 模拟迷宫

	//规则
	//1,如果元素的值为1 则为墙
	//2,如果元素的值为0 没有走过的路
	//3,如果元素的值为2 是一个通路
	//4,如果元素的值为3 是走过的点,且路不通
	var maps [8][7]int
	initMap(&maps)

	setWay(&maps, 1, 1)

	fmt.Println("走完路的效果：")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(maps[i][j], "  ")
		}
		fmt.Println()
	}
}
