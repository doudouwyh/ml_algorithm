package decisiontree

import (
	"wh_ml/lib"
	"sort"
	"math"
	"fmt"
	"container/list"
)

const (
	YOUTH int  = 1
	MID      int = 2
	OLD      int  = 3
	YES       int = 1
	NO        int =  2
	NORMAL int = 1
	GOOD   int =2
	VERY     int =3
)


/**
*  熵 entropy, H(X) = - sigma(pi*log(pi))  , pi 为X中每个不同值出现的概率
*  条件熵: condtion entropy, H(Y|X)  = sigma(pi*H(Y|xi)), pi 为X重xi出现的概率,xi 为X中的不同值
*  信息增益: Gain(D,A) = H(D) - H(D|A)，信息(特征)A提供后，D的不确定性下降的程度
*  信息增益比：Gain(R)(D,A) = Gain(D,A)/H(A)(D)
*  ID3算法用信息增益来挑选最佳特征，这会导致不同值多的特征占优势
*  C4.5算法用信息增益比来挑选最佳特征，弥补了ID3的缺陷
*  CART通过基尼系数来挑选最佳特征
****************************************************/
func ID3(){
	fmt.Println("ID3:")
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

	attributesidx := sort_by_entropy(dots,4)
	fmt.Println("eindexs:",attributesidx)

	attributesidx = sort_by_info_gain(dots,4)
	fmt.Println("gindexs:",attributesidx)

	attributesidx = []int{0,1,2,3}
	var Tree DTree
	DTree_Create(dots,attributesidx,0,nil,&Tree)
	Print_dtree(Tree)
}

/**	按特征维熵降序排
*	dim : 特征维总维数
*/
func sort_by_entropy(dots [] lib.Dot, dim int)[]int {
	var entropys []float64
	for j:=0; j<dim; j++{
		entropys = append(entropys,Get_entropy(dots,j))
	}

	fmt.Println(entropys)
	var count []int = make([]int,dim)
	for i:=0; i<dim; i++{
		for j:=i+1; j<dim; j++{
			if entropys[i] >= entropys[j]{
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

/**   按特征维信息增益降序排
*	dim : 特征维总维数
*/
func sort_by_info_gain(dots [] lib.Dot, dim int)[]int{
	var infogains []float64
	for j:=0; j<dim; j++{
		infogains = append(infogains,Get_info_gain(dots,j))
	}

	fmt.Println(infogains)
	var count []int = make([]int,dim)
	for i:=0; i<dim; i++{
		for j:=i+1; j<dim; j++{
			if infogains[i] >= infogains[j]{
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

type DTree_Node struct{
	SplitDim int
	IsLeaf  bool
	Parent *DTree_Node
	Childs map[int]*DTree_Node
	Class   int
}

func (node *DTree_Node) print(){
	fmt.Println("splitdim:",node.SplitDim)
	fmt.Println("isleaf:",node.IsLeaf)
	fmt.Println("class:",node.Class)
}

type DTree struct{
	Root *DTree_Node
}



func DTree_Create(dots []lib.Dot,attributes []int,curdimindex int, parent *DTree_Node,Tree *DTree){
	if len(dots) <= 0{
		return
	}

	//特征维为空，或已遍历所有特征，取多数为最终分类
	if curdimindex == len(attributes){
		var classes []int = make([]int,len(dots))
		classes = lib.GetDimValue(dots,curdimindex)
		fmt.Println("maxdim,classes:",classes)
		pairs := Getdiff_count(classes)
		cls :=0
		count := 0
		for _, v:=range pairs{
			if v.Count > count {
				cls = v.Value
			}
		}
		var newnode DTree_Node
		newnode.Class = cls
		newnode.IsLeaf = true
		newnode.Parent = parent
		if parent == nil{
			Tree.Root = &newnode
		}else{
			if curdimindex > 0 {
				parent.Childs[dots[0].Corordinate[attributes[curdimindex-1]]] = &newnode
			}
		}
		return
	}

	//fmt.Println("DTree_Create:")
	//lib.PrintDots(dots)
	//fmt.Println("curdimindex:",curdimindex)

	var values []int = make([]int,len(dots))
	values = lib.GetDimValue(dots,attributes[curdimindex])
	fmt.Println("values:",values)
	pairs := Getdiff_count(values)

	//该特征维上的数据都是同一类
	if len(pairs) <= 1{
		fmt.Println("one class,values:",values)
		var newnode DTree_Node
		newnode.Class = pairs[0].Value
		newnode.IsLeaf = true
		newnode.Parent = parent
		if parent == nil{
			Tree.Root = &newnode
		}else{
			if curdimindex > 0 {
				parent.Childs[dots[0].Corordinate[attributes[curdimindex-1]]] = &newnode
			}
		}
		return
	}else{
		var newnode DTree_Node
		newnode.Parent = parent
		newnode.SplitDim = curdimindex+1
		if parent == nil{
			Tree.Root = &newnode
		}else{
			if curdimindex > 0 {
				parent.Childs[dots[0].Corordinate[curdimindex-1]] = &newnode//上一个分裂维的数据相同
			}
		}
		newnode.Childs = make(map[int]*DTree_Node,len(pairs))
		dotlists := make(map[int][]lib.Dot)
		for i:=0; i<len(dots); i++{
			t := dots[i].Corordinate[attributes[curdimindex]]
			dotlists[t] = append(dotlists[t] ,dots[i])
		}
		fmt.Println("dotlists:",dotlists)

		for _,v := range dotlists{
			DTree_Create(v,attributes,curdimindex+1,&newnode,Tree)
		}
	}
}

func DTree_Search(t DTree, target lib.Dot){

}


type CountPair struct{
	Value int
	Count int
}

/**
*	获取不同值及每个值的个数
*
****************************/
func Getdiff_count(nums lib.IntSlice) (out []CountPair) {
	if len(nums) < 1{
		return
	}
	sort.Sort(nums)
	var pair CountPair
	pair.Value = nums[0]
	pair.Count = 1
	out = append(out,pair)
	for i:=1; i<len(nums);i++{
		if nums[i] != out[len(out)-1].Value{
			var p CountPair
			p.Value = nums[i]
			p.Count = 1
			out = append(out,p)
		}else{
			out[len(out)-1].Count++
		}
	}
	return
}

/**
*	计算某个特征维的熵
*	dim = 0,1,...
**/
func Get_entropy(dots[] lib.Dot, dim int) float64{
	values := lib.GetDimValue(dots,dim)
	return Get_entropy_by_values(values)
}

/**
*	计算某个特征维对目标维的信息增益
*	dim = 0,1,...
**/
func Get_info_gain(dots[] lib.Dot, dim int) float64{
	var gain float64 = 0
	if len(dots) <= 0 {
		return gain
	}

	//目标维熵
	targetdim := len(dots[0].Corordinate)-1
	entropyd := Get_entropy(dots,targetdim)

	//条件熵
	values := lib.GetDimValue(dots,dim)
	pairs   := Getdiff_count(values)
	var entropydk float64 = 0
	for _,v := range pairs{
		var d []int
		for i:=0; i<len(dots); i++{
			if dots[i].Corordinate[dim] == v.Value{
				d = append(d,dots[i].Corordinate[targetdim])
			}
		}
		entropydk += float64(len(d))/float64(len(values))*Get_entropy_by_values(d)
	}

	//信息增益
	gain = entropyd - entropydk
	return gain
}

/**
*	计算某个特征维对目标维的信息增益比
*	dim = 0,1,...
**/
func Get_info_gain_ratio(dots[] lib.Dot, dim int) float64{
	gain := Get_info_gain(dots, dim)
	return gain / Get_entropy(dots,dim)
}

func Get_entropy_by_values(values []int) float64{
	pairs   := Getdiff_count(values)
	var entropy float64 = 0
	for _, v := range pairs{
		tmp := float64(v.Count)/float64(len(values))
		entropy += -( tmp * math.Log2(tmp))
	}
	return entropy
}

func print_tree_bylayer(root *DTree_Node){
	if root == nil{
		fmt.Println("the tree is empty!")
		return
	}

	l := list.New()
	l.PushBack(root)
	cur_level := 1
	cur_level_num := 1
	next_level_num := 0
	for  l.Len() > 0 {
		n := l.Front()
		l.Remove(n)
		n.Value.(*DTree_Node).print()
		cur_level_num--

		for _,v := range n.Value.(*DTree_Node).Childs{
			l.PushBack(v)
			next_level_num++
		}

		if  cur_level_num == 0 {
			fmt.Println("")
			cur_level++
			cur_level_num = next_level_num
			next_level_num = 0
		}
	}
}

func Print_dtree(t DTree){
	fmt.Println("print_tree:")
	print_tree_bylayer(t.Root)
}