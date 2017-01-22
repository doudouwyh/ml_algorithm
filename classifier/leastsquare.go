package classifier

/**
* 最小二乘法通过最小化误差的平方和寻找数据的最佳函数匹配。它本身不是分类方法，也不是回归方法，
* 但这些算法都可以用最小二乘来来作为其损失函数。
* 最小二乘的基本原来是，对于训练样本，使预测的输出和实际输出之间的差的平方和最小，也即是最小化
* 误差的平方和。推导过程见picture.
*
* 示例是用最小二乘法来学习一元线性回归
*****************************************/


import (
	"wh_ml/lib"
	"math/rand"
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func LeastSquare(){
	fmt.Println("\nLeastSquare:")
	n := 50
	var dot2 []lib.Dot = make([]lib.Dot,n)
	var x []int
	var y []int
	for i:=0; i<n; i++{
		var cor []int = make([]int,2)
		dot2[i] = lib.NewDot(2)
		cor[0] = rand.Int() % 15   //x
		cor[1] = i+ rand.Int() % 10//y
		x = append(x,cor[0])
		y = append(y,cor[1])
		dot2[i].Set(cor,1)
	}

	for i:=0; i<n;i++{
		dot2[i].Print()
	}

	fmt.Println("x:",x)
	fmt.Println("y:",y)

	var sigmaxi2 int = 0
	var sigmaxi   int = 0
	var sigmaxiyi int = 0
	var sigmayi    int =0
	for i:=0; i<n; i++{
		sigmaxi2 += dot2[i].Corordinate[0] * dot2[i].Corordinate[0]
		sigmaxi   += dot2[i].Corordinate[0]
		sigmayi   += dot2[i].Corordinate[1]
		sigmaxiyi += dot2[i].Corordinate[0]*dot2[i].Corordinate[1]
	}
	fmt.Println("sigmaxi2:",sigmaxi2)
	fmt.Println("sigmaxi:",sigmaxi)
	fmt.Println("sigmaxiyi:",sigmaxiyi)
	fmt.Println("sigmayi:",sigmayi)

	var a float32
	var b float32
	deno := float32(n*sigmaxi2-sigmaxi*sigmaxi)
	a =  (float32(n*sigmaxiyi -sigmaxi*sigmayi))/deno
	b = float32(sigmayi)/float32(n) - float32(sigmaxi)/float32(n) *a

	fmt.Println("a:",a)
	fmt.Println("b:",b)

	fmt.Println("\nForcast:")
	running := true
	reader := bufio.NewReader(os.Stdin)
	for running {
		fmt.Print("please input a integer,input stop to exit:")
		data, _, _ := reader.ReadLine()
		inputstring := string(data)
		if inputstring == "stop" {
			running = false
		}else{
			input, err := strconv.Atoi(inputstring)
			if err != nil{
				fmt.Println("input error! please input a integer!")
			}else{
				fmt.Println("your input is:",inputstring,"forcast is:",a*float32(input)+b)
				return
			}
		}
	}
}

