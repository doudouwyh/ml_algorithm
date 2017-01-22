package em

import (
	"wh_ml/lib"
	"math/rand"
	"wh_ml/classifier"
	"fmt"
)

func K_Means(){
	fmt.Println("\nK_Means:")
	n := 11
	var dots []lib.Dot = make([]lib.Dot,n)
	dots[0].Corordinate = []int{6,2}
	dots[0].Dim = 2
	dots[1].Corordinate = []int{2,3}
	dots[1].Dim = 2
	dots[2].Corordinate = []int{4,1}
	dots[2].Dim = 2
	dots[3].Corordinate = []int{2,4}
	dots[3].Dim = 2
	dots[4].Corordinate = []int{1,4}
	dots[4].Dim = 2
	dots[5].Corordinate = []int{5,1}
	dots[5].Dim = 2
	dots[6].Corordinate = []int{6,1}
	dots[6].Dim = 2
	dots[7].Corordinate = []int{0,5}
	dots[7].Dim = 2
	dots[8].Corordinate = []int{5,2}
	dots[8].Dim = 2
	dots[9].Corordinate = []int{4,0}
	dots[9].Dim = 2
	dots[10].Corordinate = []int{1,3}
	dots[10].Dim = 2


	k := 3//k=2

	//随机选k个中心点
	var centers []lib.Dot
	random := rand.Perm(len(dots))
	for i:=0; i<k; i++{
		centers = append(centers,dots[random[i]])   //直接取前k个
	}

	fmt.Println("centers:",centers)

	var clusters map[int][]lib.Dot
	rn := 0
	for rn < 50 {  //迭代50次
		centers,clusters = get_cluster(dots, centers)
		rn++
	}

	for k,v := range clusters{
		fmt.Println("class:",k)
		lib.PrintDots(v)
	}
}

func get_cluster(dots []lib.Dot, centers []lib.Dot) ([]lib.Dot,map[int][]lib.Dot){
	lib.PrintDots(centers)
	var clusters map[int][]lib.Dot = make(map[int][]lib.Dot )
	for j:=0;  j<len(dots); j++{
		var key int = 0
		var value lib.Dot
		var mindistance float64 = 1000000000000
		for i:=0; i<len(centers); i++{
			d := classifier.Get_distance(dots[j],centers[i])
			if d < mindistance{
				key = i
				value = dots[j]
				mindistance = d
			}
		}
		clusters[key] = append(clusters[key],value)
	}

	var newcenters []lib.Dot
	for _,v :=range clusters{
		newcenters = append(newcenters,get_center(v))
	}
	return newcenters,clusters
}

func get_center(dots []lib.Dot) lib.Dot{
	if len(dots) <= 0 {
		panic("parameter error!")
	}
	var cord []int
	dim := len(dots[0].Corordinate)
	for i:=0; i<dim; i++{
		values := lib.GetDimValue(dots,i)
		cord = append(cord,int(float64(classifier.Get_sum(values))/float64(len(values)))) //取整
	}

	var dot lib.Dot
	dot.Corordinate = cord
	dot.Dim = dim
	return dot
}














