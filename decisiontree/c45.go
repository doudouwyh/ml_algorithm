package decisiontree

import (
	"wh_ml/lib"
	"fmt"
)

func C45(){
	fmt.Println("C4.5:")
	n := 15
	var dots []lib.Dot = make([]lib.Dot,n)
	//前面四列是特征维，最后一列是目标维
	dots[0].Corordinate = []int{YOUTH,NO,NO,NORMAL,NO}
	dots[1].Corordinate = []int{YOUTH,NO,NO,GOOD,NO}
	dots[2].Corordinate = []int{YOUTH,YES,NO,GOOD,YES}
	dots[3].Corordinate = []int{YOUTH,YES,YES,NORMAL,YES}
	dots[4].Corordinate = []int{YOUTH,NO,NO,NORMAL,NO}
	dots[5].Corordinate = []int{MID,NO,NO,NORMAL,NO}
	dots[6].Corordinate = []int{MID,NO,NO,GOOD,NO}
	dots[7].Corordinate = []int{MID,YES,YES,GOOD,YES}
	dots[8].Corordinate = []int{MID,NO,YES,VERY,YES}
	dots[9].Corordinate = []int{MID,NO,YES,VERY,YES}
	dots[10].Corordinate = []int{OLD,NO,YES,VERY,YES}
	dots[11].Corordinate = []int{OLD,NO,YES,GOOD,YES}
	dots[12].Corordinate = []int{OLD,YES,NO,GOOD,YES}
	dots[13].Corordinate = []int{OLD,YES,NO,VERY,YES}
	dots[14].Corordinate = []int{OLD,NO,NO,NORMAL,NO}

	lib.PrintDots(dots)

	indexs := sort_by_info_gain_ratio(dots,4)
	fmt.Println("gratioindexs:",indexs)

	var Tree DTree
	DTree_Create(dots,indexs,0,nil,&Tree)
	Print_dtree(Tree)
}

func sort_by_info_gain_ratio(dots [] lib.Dot, dim int)[]int {
	var infogainratios []float64
	for j:=0; j<dim; j++{
		infogainratios = append(infogainratios,Get_info_gain_ratio(dots,j))
	}

	fmt.Println(infogainratios)
	var count []int = make([]int,dim)
	for i:=0; i<dim; i++{
		for j:=i+1; j<dim; j++{
			if infogainratios[i] >= infogainratios[j]{
				count[i] += 1
			}else{
				count[j] += 1
			}
		}
	}
	var index []int = make([]int,dim)
	for i:=0; i<dim; i++{
		index[dim-count[i]-1] = i //降序
	}
	return index
	return nil
}

