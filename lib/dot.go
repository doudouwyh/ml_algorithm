package lib

import (
	"fmt"
)

type Dot struct {
	Corordinate []int
	Dim   int
	Value float64
}

func (d *Dot)Set(coordiate []int,value float64){
	if d.Dim != len(coordiate){
		panic("the coordinate and the dim is not match!")
	}
	d.Corordinate = coordiate
	d.Value = value
}

func (d *Dot)Print(){
	//fmt.Println("\tdim:",d.Dim)
	fmt.Println("\tcor:",d.Corordinate)
	//fmt.Println("\tvalue:",d.Value)
}

func NewDot(dim int) Dot{
	var o Dot
	o.Corordinate = make([]int,dim)
	o.Dim = dim
	return o
}

func PrintDots(dots []Dot){
	for i:=0; i<len(dots); i++{
		dots[i].Print()
	}
}

/**获取dot数组某个维度的值
* dim 从0开始
***/
func GetDimValue(dots[] Dot, dim int) []int{
	var values []int
	for i:=0; i<len(dots);i++{
		values = append(values,dots[i].Corordinate[dim])
	}
	return values
}