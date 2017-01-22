package test

import (
	"wh_ml/lib"
	"math/rand"
)

func Dot_Test(){
	test_dim2()
	test_dim3()
}

func test_dim2(){
	//2-dim
	var dot2 []lib.Dot = make([]lib.Dot,50)
	for i:=0; i<50; i++{
		var cor []int = make([]int,2)
		dot2[i] = lib.NewDot(2)
		cor[0] = i
		cor[1] = i+ rand.Int() % 10
		dot2[i].Set(cor,rand.Float64()*15)
	}

	for i:=0; i<50;i++{
		dot2[i].Print()
	}
}

func test_dim3(){
	//3-dim
	var dot3 []lib.Dot = make([]lib.Dot,50)
	for i:=0; i<50; i++{
		var cor []int = make([]int,3)
		dot3[i] = lib.NewDot(3)
		cor[0] = i
		cor[1] = i+ rand.Int() % 10
		cor[2] = i+ rand.Int() % 10
		dot3[i].Set(cor,rand.Float64()*15)
	}

	for i:=0; i<50;i++{
		dot3[i].Print()
	}
}







