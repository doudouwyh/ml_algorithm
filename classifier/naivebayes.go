package classifier

import (
	"fmt"
	"wh_ml/lib"
	"sort"
)

const (
	S  int = 0
	M  int = 1
	L  int = 2
)

/**
*   朴素贝叶斯方法属于生成模型，之所以称为朴素，是由于
*   特征条件独立性假设,即：
*   P(X=x|Y=CK) = P(X1=x1,X2=x2,...,Xn=xn|Y=CK) =  PAI P(Xj=xj|Y=CK) (j=1,2,3...,m)
*   m为X的特征数，也即X特征空间的维度
*   y = arg max P(Y=CK) PAI P(Xj = xj|Y=CK)  (j=1,2,3...,m)
*   朴素贝叶斯是多分类方法，本例只测试两分类
*/
func NaiveBayes(){
	fmt.Println("\nNaiveBayes:")
	
	n := 15
	//coordinate特征维，value分类
	var dots []lib.Dot = make([]lib.Dot,n)
	dots[0].Corordinate = []int{1,S}
	dots[0].Value = -1
	dots[1].Corordinate = []int{1,M}
	dots[1].Value = -1
	dots[2].Corordinate = []int{1,M}
	dots[2].Value = 1
	dots[3].Corordinate = []int{1,S}
	dots[3].Value = 1
	dots[4].Corordinate = []int{1,S}
	dots[4].Value = -1
	dots[5].Corordinate = []int{2,S}
	dots[5].Value = -1
	dots[6].Corordinate = []int{2,M}
	dots[6].Value = -1
	dots[7].Corordinate = []int{2,M}
	dots[7].Value = 1
	dots[8].Corordinate = []int{2,L}
	dots[8].Value = 1
	dots[9].Corordinate = []int{2,L}
	dots[9].Value = 1
	dots[10].Corordinate = []int{3,L}
	dots[10].Value = 1
	dots[11].Corordinate = []int{3,M}
	dots[11].Value = 1
	dots[12].Corordinate = []int{3,M}
	dots[12].Value = 1
	dots[13].Corordinate = []int{3,L}
	dots[13].Value = 1
	dots[14].Corordinate = []int{3,L}
	dots[14].Value = -1

	lib.PrintDots(dots)

	var x1 []int
	var x2 []int
	var  y  []int
	for i:=0; i<15; i++{
		x1 = append(x1,dots[i].Corordinate[0])
		x2 = append(x2,dots[i].Corordinate[1])
		y   = append(y,int(dots[i].Value))
	}
	
	diffy   := getdiff(y)
	propy   := getpy(diffy,dots)
	
	//测试: input(2,S)
	var c int
	var max float32
	for i:=0; i<len(propy);i++{
		p := propy[i]* getcondprop(2,diffy[i],dots,1)*getcondprop(S,diffy[i],dots,2)
		fmt.Println("i=",i,"y=",diffy[i],"p=",p)
		if p > max {  //arg max f(x)
			max = p
			c = diffy[i]
		}
	}
	fmt.Println("input, x1:",2,"x2:",S)
	fmt.Print("class:",c,",propability:",max)
}

//找一组数中不同的数
func getdiff(nums lib.IntSlice) (out []int) {
	if len(nums) < 1{
		return
	}
	sort.Sort(nums)
	out = append(out,nums[0])
	for i:=1; i<len(nums);i++{
		if nums[i] != out[len(out)-1]{
			out = append(out,nums[i])
		}
	}
	return
}

// p(C=CK)
func getpy(diffy []int,dots []lib.Dot) (prop []float32){
	if len(diffy) ==0 || len(dots)==0{
		panic("param error!")
	}
	
	prop = make([]float32,len(diffy))
	for i:=0; i<len(diffy); i++{
		count :=0
		for j:=0; j<len(dots); j++{
			if int(dots[j].Value) == diffy[i]{
				count++
			}
		}
		prop[i] = float32(count)/float32(len(dots))
	}
	return
}

/*
*   P(X|Y)
*   dim: 1 -- x1, 2 -- x2
*/
func getcondprop(x,y int, dots []lib.Dot, dim int) float32{
	if len(dots) < 1{
		return 0
		return 0
	}
	county := 0
	countxy := 0
	for i:=0; i<len(dots); i++{
		if int(dots[i].Value) == y{
			county++
			if dots[i].Corordinate[dim-1] == x{
				countxy++
			}
		}
	}
	
	return float32(countxy)/float32(county)
}

