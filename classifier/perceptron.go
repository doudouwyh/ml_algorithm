package classifier
import (
	"fmt"
	"wh_ml/lib"
)

/**
* 感知器算法是二分类算法
* 模型：分离超平面
* 策略：最小化误分类点到分类超平面的距离之和
* 算法：梯度下降
*  损失函数：误分类点到分离超平面的距离
* 推导：
*	original: picture/perceptron1/perceptron1.jpg
*	 dual:      picture/perceptron1/perceptron2.jpg
****************************************/
func Perceptron(){
	PerceptronOrigin()//原始方法
	PerceptronDual()//对偶方法
	PerceptronDualGram()//对偶查表
}

/**原始方法
*   f(x)=sign(w.x+b)
*   梯度下降的第i轮迭代：
*   w <- w + yita*xi*yi
*   b <- b + yita*yi
*/
func PerceptronOrigin(){
	fmt.Println("\nPerceptronOrigin:")

	n := 3
	//coordinate特征维，value分类
	var dots []lib.Dot = make([]lib.Dot,n)
	dots[0].Corordinate = []int{3,3}
	dots[0].Value = 1
	dots[1].Corordinate = []int{4,3}
	dots[1].Value = 1
	dots[2].Corordinate = []int{1,1}
	dots[2].Value = -1

	//train  (w0=1,w1=1,b=-3)
	yita := 1
	count := 0
	w0 := 0
	w1 := 0
	b  := 0
	for i:=0; i<n; i++{
	    wt0 := 0
	    wt1 := 0
	    bt := 0
 	    fmt.Println("wt0:",wt0,"wt1:",wt1,"bt:",bt)
	    y := int(dots[i].Value) * (wt0*dots[i].Corordinate[0] + wt1*dots[i].Corordinate[1] + bt)
	    if y > 0 {  //分类正确
			count++
			if count == n{  //没有误分类点，终止
				w0 = wt0
				w1 = wt1
				b  = bt
				break
			}
	    }else{//分类错误
		    count = 0
		    wt0 = wt0 + yita *dots[i].Corordinate[0]*int(dots[i].Value)
		    wt1 = wt1 + yita *dots[i].Corordinate[1]*int(dots[i].Value)
		    bt  = bt + yita * int(dots[i].Value)
	    }
	}

	//predict
	x1 := 2
	x2 := 3
	fmt.Println("input:",x1,x2,"\npredict:",sign(w0*x1+w1*x2+b))
}

func sign(n int) int{
	if n >0 {
		return 1
	}else{
		return -1
	}
}

/** f(x)=sign(sigma(alphaj*yj*xj).xi+b)
*   梯度下降的第i轮迭代：
*   alphai <- alphai + yita
*   b      <- b + yita * yi
*/
func PerceptronDual(){
	fmt.Println("\nPerceptronDual:")
	
	n := 3
	var dots []lib.Dot = make([]lib.Dot,n)
	dots[0].Corordinate = []int{3,3}
	dots[0].Value = 1
	dots[1].Corordinate = []int{4,3}
	dots[1].Value = 1
	dots[2].Corordinate = []int{1,1}
	dots[2].Value = -1
	
	//train  w0=1,w1=1,b=-3
	yita := 1
	count := 0
	alpha := []int{0,0,0}
	b  := 0
	for i:=0; i<n; i++{
		alphat := []int{0,0,0}
		
		bt := 0
		fmt.Println("alphat0:",alphat[0],"alphat1:",alphat[1],"alphat2:",alphat[2],"bt:",bt)
		
		sum := []int{0,0}
		for j:=0; j<n; j++{
			sum[0] += alphat[j] * int(dots[j].Value)*dots[j].Corordinate[0]
			sum[1] += alphat[j] * int(dots[j].Value)*dots[j].Corordinate[1]
		}
				
		y := int(dots[i].Value) * (sum[0]*dots[i].Corordinate[0] + sum[1]*dots[i].Corordinate[1] + bt)
		if y > 0 {  //分类正确
			count++
			if count == n{
				alpha = alphat
				b  = bt
				break
			}
		}else{//分类错误
			count = 0
			alphat[i] += yita
			bt  = bt + yita * int(dots[i].Value)
		}
	}
	
	x1 := 2
	x2 := 3
	sum := []int{0,0}
	for j:=0; j<n; j++{
		sum[0] += alpha[j] * int(dots[j].Value)*dots[j].Corordinate[0]
		sum[1] += alpha[j] * int(dots[j].Value)*dots[j].Corordinate[1]
	}
	fmt.Println(sign(sum[0]*x1+sum[0]*x2+b))
}

/** 对对偶形式的另一种计算，引入Gram矩阵，减少运算量
*/
func PerceptronDualGram(){
	fmt.Println("\nPerceptronDualGram:")
	
	n := 3
	var dots []lib.Dot = make([]lib.Dot,n)
	dots[0].Corordinate = []int{3,3}
	dots[0].Value = 1
	dots[1].Corordinate = []int{4,3}
	dots[1].Value = 1
	dots[2].Corordinate = []int{1,1}
	dots[2].Value = -1
	
	/* gram matrix:
	**       18 21 6
	**       21 25 7
	**        6  7 2
	*/
	gram := []int{18,21,6,21,25,7,6,7,2}
	
	//train  (w0=1,w1=1,b=-3)
	yita := 1
	count := 0
	alpha := []int{0,0,0}
	b  := 0
	for i:=0; i<n; i++{
		alphat := []int{0,0,0}
		
		bt := 0
		fmt.Println("alphat0:",alphat[0],"alphat1:",alphat[1],"alphat2:",alphat[2],"bt:",bt)
		
		sum := 0
		for j:=0; j<n; j++{
			sum += alphat[j] * int(dots[j].Value)*gram[j*3 + i]
		}
		
		y := int(dots[i].Value) * (sum + bt)
		if y > 0 {  //分类正确
			count++
			if count == n{
				alpha = alphat
				b  = bt
				break
			}
		}else{//分类错误
			count = 0
			alphat[i] += yita
			bt  = bt + yita * int(dots[i].Value)
		}
	}

	//predict
	x1 := 2
	x2 := 3
	sum := []int{0,0}
	for j:=0; j<n; j++{
		sum[0] += alpha[j] * int(dots[j].Value)*dots[j].Corordinate[0]
		sum[1] += alpha[j] * int(dots[j].Value)*dots[j].Corordinate[1]
	}
	fmt.Println("input:",x1,x2,"\npredict:",sign(sum[0]*x1+sum[0]*x2+b))
}
