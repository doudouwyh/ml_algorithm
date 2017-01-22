package decisiontree

import (
	"wh_ml/lib"
	"fmt"
)


/**
* CART (classify and regression tree)
*
*
*
*
*
*
*
********************************************/
func CART(){
	fmt.Println("CART:")
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

	indexs := sort_by_entropy_gini(dots,4)
	fmt.Println("giniindexs:",indexs)
}


func sort_by_entropy_gini(dots [] lib.Dot, dim int)[]int {
	var ginis []float64
	for j:=0; j<dim; j++{
		ginis = append(ginis,get_gini(dots,j))
	}

	fmt.Println(ginis)
	var count []int = make([]int,dim)
	for i:=0; i<dim; i++{
		for j:=i+1; j<dim; j++{
			if ginis[i] >= ginis[j]{
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
}

/**基尼值**/
func get_gini(dots[] lib.Dot, dim int)float64{
	values := lib.GetDimValue(dots,dim)
	return  get_gini_by_values(values)
}

func get_gini_by_values(values []int) float64{
	pairs   := Getdiff_count(values)
	var p2 float64 = 0
	for _, v := range pairs{
		tmp := float64(v.Count)/float64(len(values))
		p2 += tmp*tmp
	}
	return 1-p2
}

/**获取特征维不同属性值的基尼指数
*  continuous 标记连续还是离散
*
***********************************/
func get_gini_index(dots[] lib.Dot, dim int, continuous bool) map[int]float64{
	var mpgini map[int]float64 = make(map[int]float64)
	dlen := len(dots)
	if dlen <= 0 {
		return mpgini
	}
	values := lib.GetDimValue(dots,dim)
	targetdim := len(dots[0].Corordinate) - 1
	if !continuous {
		pairs := Getdiff_count(values)
		for i := 0; i < len(pairs); i++ {
			var di []int
			var dni []int
			v := pairs[i].Value
			for j := 0; j < len(dots); j++ {
				if dots[j].Corordinate[dim] == v {
					di = append(di, dots[j].Corordinate[targetdim])
				} else {
					dni = append(dni, dots[j].Corordinate[targetdim])
				}
			}
			mpgini[v] = float64(len(di))/float64(dlen) * get_gini_by_values(di) + float64(len(dni))/float64(dlen) * get_gini_by_values(dni)
		}
	}else{
		splits := get_continuous_split(values)
		for i := 0; i < len(splits); i++ {
			var le []int  //less and equal
			var gt []int
			for j := 0; j < len(dots); j++ {
				if dots[j].Corordinate[dim] <= i {
					le = append(le, dots[j].Corordinate[targetdim])
				} else {
					gt = append(gt, dots[j].Corordinate[targetdim])
				}
			}
			mpgini[i] = float64(len(le))/float64(dlen) * get_gini_by_values(le) + float64(len(gt))/float64(dlen) * get_gini_by_values(gt)
		}
	}
	return mpgini
}

/***连续特征维分裂点***/
func get_continuous_split(values []int)[]int{
	var out []int
	if len(values) <= 0 {
		return out
	}

	if len(values) <= 1 {
		out = append(out,values[0])
		return out
	}

	for i:=0; i<len(values)-1; i++{
		split := int(float64(values[i]+values[i+1]) / 2)  //直接取整，不考虑小数
		out = append(out,split)
	}
	return out
}

