package test


import(
	"wh_ml/lib"
)

func Matrix_Test(){
	var m lib.Matrix
	m.Data = []float64{float64(1),float64(2),float64(3),float64(2)}
	m.Row=2
	m.Col=2

	var n lib.Matrix
	n.Data = []float64{float64(1),float64(1),float64(0),float64(1)}
	n.Row = 2
	n.Col   = 2

	m.Print()

	t := m.Transfer()
	t.Print()

	n.Print()

	a,_ := m.DataMultiply(&n)
	a.Print()

	o,_ := m.MatixMultiply(&n)

	o.Print()

	o,_ = o.NumDiv(3)

	o.Print()

	d,_ :=  o.MatixDiv(&m)
	d.Print()

	z := lib.Zeros(5,6)
	z.Print()

	one := lib.Ones(5,6)
	one.Print()

	i := lib.Identity(6)
	i.Print()

	i.Set(6,6,12)
	i.Print()
}
