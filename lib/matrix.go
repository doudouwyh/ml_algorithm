package lib

import (
	"fmt"
	"errors"
)

type Matrix struct{
	Row int
	Col   int
	Data []float64
}

func (m* Matrix)Print(){
	for i:=0; i<m.Row;i++{
		fmt.Print("\t")
		for  j:=0; j<m.Col;j++{
			fmt.Print(m.Data[i*m.Col+j],"\t")
		}
		fmt.Println(" ")
	}
	fmt.Println(" ")
}

func (m* Matrix)Set(r,c int,v float64){
	if r <1 || r > m.Row || c < 1 || c > m.Col{
		panic("the parameter is error!")
	}
	m.Data[(r-1)*m.Col+c-1] = v
}


func (m *Matrix)MatixMultiply(n *Matrix)(o Matrix,err error){
	if  m.Col != n.Row{
		err = errors.New("the format of the matrix is wrong!")
		return
	}

	o = Matrix{m.Row,n.Col,make([]float64,m.Row*n.Col)}
	for i:=0; i<m.Row; i++{
		for k:=0; k<n.Col;k++{
			sum := float64(0)
			for  j:=0; j<m.Col; j++{
				sum += m.Data[i*m.Col+j] * n.Data[j*n.Col+k]
			}
			o.Data[i*n.Col + k] = sum
		}
	}

	return
}

func (m *Matrix)DataMultiply(n *Matrix)(o Matrix,err error){
	if  m.Col != n.Col || m.Row != n.Row{
		err = errors.New("the format of the matrix is not match!")
		return
	}

	o = Matrix{m.Row,n.Col,make([]float64,m.Row*n.Col)}
	for i:=0; i<m.Row; i++{
		for k:=0; k<m.Col;k++{
			o.Data[i*m.Col + k] = m.Data[i*m.Col + k] * n.Data[i*m.Col + k]
		}
	}
	return
}

func (m *Matrix)NumMultiply(num float64)(o Matrix){
	o = Matrix{m.Row,m.Col,make([]float64,m.Row*m.Col)}
	for i:=0; i<m.Row; i++{
		for k:=0; k<m.Col;k++{
			o.Data[i*m.Col + k] = m.Data[i*m.Col + k] * num
		}
	}
	return
}


func (m *Matrix)MatrixAdd(n *Matrix)(o Matrix,err error){
	if  m.Col != n.Col || m.Row != n.Row{
		err = errors.New("the format of the matrix is not match!")
		return
	}

	o = Matrix{m.Row,m.Col,make([]float64,m.Row*m.Col)}
	for i:=0; i<m.Row; i++{
		for k:=0; k<m.Col;k++{
			o.Data[i*m.Col + k] = m.Data[i*m.Col + k] + n.Data[i*m.Col + k]
		}
	}
	return
}

func (m *Matrix)DataAdd(num float64)(o Matrix,err error){
	o = Matrix{m.Row,m.Col,make([]float64,m.Row*m.Col)}
	for i:=0; i<m.Row; i++{
		for k:=0; k<m.Col;k++{
			o.Data[i*m.Col + k] = m.Data[i*m.Col + k] + num
		}
	}
	return
}


func (m *Matrix)NumDiv(num int)(o Matrix,err error){
	if  num == 0{
		err = errors.New("can not divided by zero!")
		return
	}

	o = Matrix{m.Row,m.Col,make([]float64,m.Row*m.Col)}
	for i:=0; i<m.Row; i++{
		for j:=0; j<m.Col;j++{
			o.Data[i*m.Col + j] =  m.Data[i*m.Col + j] / float64(num)
		}
	}
	return
}

func (m *Matrix)MatixDiv(n *Matrix)(o Matrix,err error){
	if  m.Col != n.Col || m.Row != n.Row{
		err = errors.New("the format of the matrix is not match!")
		return
	}

	o = Matrix{m.Row,n.Col,make([]float64,m.Row*n.Col)}
	for i:=0; i<m.Row; i++{
		for k:=0; k<m.Col;k++ {
			if n.Data[i * m.Col + k] != 0 {
				o.Data[i * m.Col + k] = m.Data[i * m.Col + k] / n.Data[i * m.Col + k]
			}else{
				o.Row = 0
				o.Col   = 0
				o.Data = [] float64{}
				err = errors.New("can not devided by zero!")
				return
			}
		}
	}
	return
}


func (m *Matrix)Transfer()(o Matrix){
	o = Matrix{m.Row,m.Col,make([]float64,m.Row*m.Col)}
	for i:=0; i<m.Row; i++{
		for j:=0; j<m.Col;j++ {
			o.Data[j*o.Col + i ] = m.Data[i * m.Col + j]
		}
	}
	return
}

//row start with zero!
func (m *Matrix)GetRow(row int)(out []float64,err error){
	if row <0 || row > m.Row{
		err = errors.New("the input param error!")
		return
	}
	for i:=0; i<m.Col; i++{
		out = append(out,m.Data[row*m.Col+i])
	}
	return
}

//col start with zero!
func (m *Matrix)GetCol(col int)(out []float64,err error){
	if col <0 || col > m.Col{
		err = errors.New("the input param error!")
		return
	}
	for i:=0; i<m.Row; i++{
		out = append(out,m.Data[i*m.Col+col])
	}
	return
}

//row,col start with zero!
func (m *Matrix)Get(row,col int)(out float64,err error){
	if col <0 || col > m.Col || row <0 || row > m.Row{
		err = errors.New("the input param error!")
		return
	}
	out = m.Data[row*m.Col + col]
	return
}

//zero-matrix
func Zeros(r,c int) Matrix{
	return Matrix{r,c,make([]float64,r*c)}
}

//one-matrix
func Ones(r,c int) Matrix{
	o := Matrix{r,c,make([]float64,r*c)}
	for i:=0; i<o.Row; i++{
		for j:=0; j<o.Col; j++{
			o.Data[i*o.Col+j] = 1
		}
	}
	return o
}

//identity-matrix
func Identity(n int) Matrix{
	o := Matrix{n,n,make([]float64,n*n)}
	for i:=0; i<o.Row; i++{
		o.Data[i*o.Col+i] = 1
	}
	return o
}

