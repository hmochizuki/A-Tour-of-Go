package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	//	fmt.Println(t)
	if t == nil {
		return
	}
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func WalkTest() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		value := <-ch
		fmt.Print(value, ",")
	}
	fmt.Print("\n")
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	v1, v2, isSame := 0, 0, false

	for i := 0; i < 10; i++ {
		v1, v2 = <-ch1, <-ch2
		isSame = v1 == v2
		if isSame == false {
			break
		}
	}
	return isSame
}
func SameTest() {
	isSame1 := Same(tree.New(1), tree.New(1))
	isSame2 := Same(tree.New(1), tree.New(2))

	if isSame1 == true && isSame2 == false {
		fmt.Println("Success!!", "isSame1:", isSame1, ", isSame2:", isSame2)
	} else {
		fmt.Println("Fail...")
	}
}

func main() {
	fmt.Print("Test Walk func. \nResult: ")
	WalkTest()
	fmt.Print("Test Same func. \nResult: ")
	SameTest()
}
