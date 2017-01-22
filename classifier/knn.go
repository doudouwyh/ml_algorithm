package classifier

import (
	"fmt"
	"wh_ml/lib"
	"sort"
	"math"
	"container/heap"
	"container/list"
	"wh_ml/decisiontree"
)

/**K近邻算法是多分类算法，回归方法
*	KNN 是无参数分类法，不需要学习
*	KNN算法思想，对输入的点x，找与其距离最近的k个点，这k个点中属于哪一类的点最多，x就属于该类
*	这里的距离可以是欧氏距离，马氏距离(协方差)，曼哈顿距离，P距离等。一般选马氏距离
*	本例通过构建kdtree并搜索来进行分类
*****************************************/
type KdTree_Node struct{
	Layer int
	Split int
	Value []lib.Dot
	Parent *KdTree_Node
	Left  *KdTree_Node
	Right *KdTree_Node
	IsLeaf  bool
}

func (node *KdTree_Node) print(){
	fmt.Println("layer:",node.Layer)
	fmt.Println("isleaf:",node.IsLeaf)

	for _,v := range node.Value{
		if len(v.Corordinate) > 0 {
			v.Print()
		}
	}
}

type KdTree struct{
	Height int
	Root *KdTree_Node
}


type  HeapNode struct{
	Distance float64
	Dot         lib.Dot
}

//heap
type DotHeap []HeapNode

func (h DotHeap) Len() int           { return len(h) }
func (h DotHeap) Less(i, j int) bool { return h[i].Distance < h[j].Distance }
func (h DotHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *DotHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(HeapNode))
}

func (h *DotHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var Tree KdTree

/***knn
*   1-kdtree，2-直接计算
*************************************/
func KNN(){
	fmt.Println("\nKNN:")
	n := 11
	//coordinate特征维，value分类
	var dots []lib.Dot = make([]lib.Dot,n)
	dots[0].Corordinate = []int{3,1,4}
	dots[0].Value = 1
	dots[0].Dim = 3
	dots[1].Corordinate = []int{2,3,7}
	dots[1].Value = 1
	dots[1].Dim = 3
	dots[2].Corordinate = []int{2,1,3}
	dots[2].Value = 2
	dots[2].Dim = 3
	dots[3].Corordinate = []int{2,4,5}
	dots[3].Value = 2
	dots[3].Dim = 3
	dots[4].Corordinate = []int{1,4,4}
	dots[4].Value = 3
	dots[4].Dim = 3
	dots[5].Corordinate = []int{4,3,4}
	dots[5].Value = 3
	dots[5].Dim = 3
	dots[6].Corordinate = []int{6,1,4}
	dots[6].Value = 1
	dots[6].Dim = 3
	dots[7].Corordinate = []int{0,5,7}
	dots[7].Value = 2
	dots[7].Dim = 3
	dots[8].Corordinate = []int{5,2,5}
	dots[8].Value = 2
	dots[8].Dim = 3
	dots[9].Corordinate = []int{4,0,6}
	dots[9].Value = 3
	dots[9].Dim = 3
	dots[10].Corordinate = []int{7,1,6}
	dots[10].Value = 3
	dots[10].Dim = 3

	lib.PrintDots(dots)

	//1-kdtree
	//获取按方差降序排列的特征维下标
	dim := 3   //维数
	attributes := sort_by_variance(dots,dim)
	attributes = []int{0,1,2}
	fmt.Println("attributes:",attributes)

	//创建kdtree
	attributeindex := 0
	KdTree_Create(dots,attributes,attributeindex,nil)
	fmt.Println("\nprint_tree:")
	print_dttree(Tree)

	//搜索kdtree
	var testdot lib.Dot
	cord := []int{6,5,5}
	testdot.Dim = 3
	testdot.Set(cord,0)
	fmt.Println("KdTree_Search:")
	fmt.Print("testdata:")
	testdot.Print()
	k := 3
	Knn_Test(Tree.Root,testdot,k) //k=3


	//2-直接计算k近邻
	c := Knn_Direct_Get(dots,testdot,k)
	fmt.Println("Knn_Direct_Get,c:",c)

}

/** 创建KD-Tree
*    每次挑选方差最大的特征作为分裂特征维，取每一个特征维的中位数作为该维的分裂点，分裂点作为父节点，
*    余下节点作为左右子节点，对左右子节点再递归，直到所有特征维都用完，树构建完毕.
*    这里的中位数跟数学上稍微有点差别，如果点数为奇数，取中间节点，如果为偶数，取中间
*    两节点中的较大者
*/
func KdTree_Create(dots []lib.Dot, attributes []int,curdimindex int,parent *KdTree_Node){
	var x []int = make([]int,len(dots))
	x = lib.GetDimValue(dots,attributes[curdimindex])
	fmt.Println("x:",x)
	m := median(x)  //用中位数作为分裂点
	
	var ldots []lib.Dot
	var rdots []lib.Dot
	var splitdot lib.Dot
	flag := 0
	for i:=0; i<len(dots); i++{
		if dots[i].Corordinate[attributes[curdimindex]] < m{
			ldots = append(ldots,dots[i])
		}else if dots[i].Corordinate[attributes[curdimindex]] >= m{
			if  flag == 0 && dots[i].Corordinate[attributes[curdimindex]] == m{//可能存在值相同的情况,只取第一个作为分裂节点
				splitdot = dots[i]
				flag = 1
			}else {
				rdots = append(rdots, dots[i])
			}
		}
	}
	
	var newnode KdTree_Node
	newnode.Layer = curdimindex+1
	newnode.Left  = nil
	newnode.Right = nil
	newnode.Value = append(newnode.Value,splitdot)
	newnode.Split = attributes[curdimindex]
	newnode.Parent = parent
	if parent == nil{
		Tree.Root = &newnode
	}else{
		if newnode.Value[0].Corordinate[attributes[curdimindex-1]] >= parent.Value[0].Corordinate[attributes[curdimindex-1]]{//这里维度必须减掉1(上一次分裂的维度)，判断是父节点的左或右子节点
			parent.Right = &newnode
		}else{
			parent.Left  = &newnode
		}
	}
	Tree.Height++

	if len(ldots) <= 0 && len(rdots)<=0{
		newnode.IsLeaf = true
	}
	
	fmt.Println("curdimindex:",curdimindex,"curdim:",attributes[curdimindex])
	fmt.Println("splitdot:")
	newnode.Value[0].Print()
	fmt.Println("ldots:")
	lib.PrintDots(ldots)
	fmt.Println("rdots:")
	lib.PrintDots(rdots)

	
	if curdimindex != len(attributes)-1{
		if len(ldots) > 0 {
			KdTree_Create(ldots, attributes, curdimindex + 1, &newnode)
		}

		if len(rdots) > 0 {
			KdTree_Create(rdots, attributes, curdimindex + 1, &newnode)
		}
	}else{ //已经遍历所有特征，剩下的作为叶子节点
		if len(ldots) > 0 {
			var lastnode KdTree_Node
			lastnode.Parent = &newnode
			newnode.Left = &lastnode
			lastnode.Layer = curdimindex+1+1
			lastnode.Split = -1
			lastnode.Left = nil
			lastnode.Right = nil
			lastnode.IsLeaf = true
			for _, v := range ldots {
				lastnode.Value = append(lastnode.Value,v)
			}
		}

		if len(rdots) > 0 {
			var lastnode KdTree_Node
			lastnode.Parent = &newnode
			newnode.Right= &lastnode
			lastnode.Layer = curdimindex+1+1
			lastnode.Split = -1
			lastnode.Left = nil
			lastnode.Right = nil
			lastnode.IsLeaf = true
			for _, v := range rdots {
				lastnode.Value = append(lastnode.Value,v)
			}
		}
	}
}

func Knn_Test(root *KdTree_Node, dot lib.Dot, k int){
	if root == nil{
		panic("the tree is nil")
	}

	//找到叶子节点，并将路径节点压栈
	SearchStack := lib.NewStack()
	KnnHeap := &DotHeap{}
	heap.Init(KnnHeap)

	p := root
	for i:=0; i<dot.Dim; i++{
		fmt.Println("p:")
		p.print()
		if  p == nil {
			break
		}
		SearchStack.Push(p)
		if p.IsLeaf {
			break
		}
		if dot.Corordinate[i] < p.Value[0].Corordinate[i]{
			p = p.Left
		}else{
			p = p.Right
		}
	}

	//递归搜索
	KdTree_Search(root, dot , k,0,SearchStack,KnnHeap)
}

/** KD-Tree 搜索
*   找k个近邻
*/
func KdTree_Search(root *KdTree_Node, dot lib.Dot, k int,radius float64,searchstack *lib.Stack, dotheap *DotHeap){
	for searchstack.Len() > 0 {
		node := searchstack.Pop().(*KdTree_Node)
		maxdist := float64(0)
		if node.IsLeaf {
			fmt.Println("values:")
			lib.PrintDots(node.Value)
			for _,v := range node.Value{
				dist := Get_distance(v,dot)
				var heapnode HeapNode
				heapnode.Distance = dist
				heapnode.Dot = v
				dotheap.Push(heapnode)
				if dist > maxdist{
					maxdist = dist
				}
			}
		}

		fmt.Println("maxdist:",maxdist)
		radius =  maxdist
	}
}

func median(data lib.IntSlice) int{
	if len(data) < 1{
		panic("input error!")
	}
	
	if !sort.IsSorted(data){
		sort.Sort(data)
	}
	
	if len(data) % 2 == 0{
		//return (data[len(data)/2] + data[(len(data)-1)/2])/2
		return data[len(data)/2] //返回大的
	}else{
		return data[len(data)/2]
	}
}


func print_dttree(t KdTree){
	fmt.Println("height:",t.Height)
	print_tree_bylayer(t.Root)
}

func print_tree_bylayer(root *KdTree_Node){
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
		for _,v := range n.Value.(*KdTree_Node).Value{
			fmt.Print(v.Corordinate,"\t")
		}
		cur_level_num--

		if n.Value.(*KdTree_Node).Left != nil {
			l.PushBack(n.Value.(*KdTree_Node).Left)
			next_level_num++
		}

		if n.Value.(*KdTree_Node).Right != nil {
			l.PushBack(n.Value.(*KdTree_Node).Right)
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

/**按特征维的方差降序排序
*   返回维度值(从0开始)
****/
func sort_by_variance(dots [] lib.Dot, dim int)[]int {
	var variance []float64
	for j:=0; j<dim; j++{
		variance = append(variance,get_variance(lib.GetDimValue(dots,j)))
	}

	//获取按variance降序排序后的下标序列
	var count []int = make([]int,dim)
	for i:=0; i<dim; i++{
		for j:=i+1; j<dim; j++{
			if variance[i] >= variance[j]{
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

//方差
func get_variance(values []int) float64{
	sum := Get_sum(values)
	mean := float64(sum) / float64(len(values))

	var variance float64
	for v,_ := range values{
		variance += math.Pow(float64(v)-mean,2)
	}

	return variance
}

func Get_sum(values []int) int{
	sum := 0
	for i:=0; i<len(values);i++{
		sum += values[i]
	}
	return sum
}

/***
*	欧氏距离
*/
func Get_distance(dot1 lib.Dot, dot2 lib.Dot) float64{
	var distance float64 = 0
	for i:=0; i<dot2.Dim; i++{
		distance += (float64(dot1.Corordinate[i]) - float64(dot2.Corordinate[i])) *  (float64(dot1.Corordinate[i]) - float64(dot2.Corordinate[i]))
	}

	return math.Sqrt(distance)
}

func Knn_Direct_Get(dots []lib.Dot,dot lib.Dot,k int) int{
	if len(dots) < 1{
		return -1
	}

	//计算距离并保存到堆结构中
	KnnHeap := &DotHeap{}
	heap.Init(KnnHeap)
	for i:=0; i<len(dots); i++{
		d := Get_distance(dots[i],dot)
		var node HeapNode
		node.Distance = d
		node.Dot = dots[i]
		KnnHeap.Push(node)
	}

	//取k个最近邻点
	var class []int
	classdim := len(dots[0].Corordinate) - 1
	for i:=0; i<k; i++{
		class = append(class,KnnHeap.Pop().(HeapNode).Dot.Corordinate[classdim])
	}

	//多数表决
	pairs := decisiontree.Getdiff_count(class)
	c := -1
	count := 0
	for i:=0; i<len(pairs); i++{
		if pairs[i].Count > count{
			c = pairs[i].Value
		}
	}
	return c
}
