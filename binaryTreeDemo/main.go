package main

import "fmt"

//二叉树的使用案例

type Hero struct {
	No        int
	Name      string
	LeftHero  *Hero
	RightHero *Hero
}

//前序遍历
//先输出根节点
//再输出左支树，然后再输出右支数

func preOrder(node *Hero) {
	if node != nil {
		fmt.Printf("编号：%d 名字：%s\n", node.No, node.Name)
		preOrder(node.LeftHero)
		preOrder(node.RightHero)
	}
}

//中序遍历
//先输出根节点左支树
//再输出跟节点，然后再输出右支数
func indexOrder(node *Hero) {
	if node != nil {
		preOrder(node.LeftHero)
		fmt.Printf("编号：%d 名字：%s\n", node.No, node.Name)
		preOrder(node.RightHero)
	}
}

//后序遍历
//先输出根节点左支树
//再输出右支数，再输出跟节点
func postOrder(node *Hero) {
	if node != nil {
		preOrder(node.LeftHero)
		preOrder(node.RightHero)
		fmt.Printf("编号：%d 名字：%s\n", node.No, node.Name)
	}
}

func main() {
	root := &Hero{
		No:   1,
		Name: "宋江",
	}

	left1 := &Hero{
		No:   2,
		Name: "吴用",
	}

	left11 := &Hero{
		No:   10,
		Name: "吴用10",
	}

	right12 := &Hero{
		No:   11,
		Name: "吴用11",
	}

	left1.LeftHero = left11
	left1.RightHero = right12

	right1 := &Hero{
		No:   3,
		Name: "卢俊义",
	}

	right2 := &Hero{
		No:   4,
		Name: "林冲",
	}

	right1.RightHero = right2

	root.LeftHero = left1
	root.RightHero = right1

	preOrder(root)

	fmt.Println("中序输出：")

	indexOrder(root)

	fmt.Println("后序输出：")

	postOrder(root)
}
