package hmm

import (
	"wh_ml/lib"
	"fmt"
)

const(
	Sunny int   = 0
	Cloudy int = 1
	Rainny int =2

	Dry     int =0
	Dryish int = 1
	Damp  int = 2
	Soggy  int = 3
)

var  flags []int
/******************************************************
transition matrix:
		sunny	cloudy	rainy
sunny	0.5 		0.375 	0.125
cloudy 	0.25 	0.125 	0.5
rainy   	0.25 	0.5 		0.375

confustion matrix:
		dry		dryish	damp	Soggy
sunny   	0.6 		0.2 		0.15 	0.05
cloudy  	0.25 	0.25 	0.25 	0.25
rainy   	0.05 	0.1 		0.35 	0.5

pai: 		{0.63,0.17,0.20}
obs: 		{Dry, Damp, Soggy,Damp,Dryish}
*****************************************************/
func Viterbi_Test(){
	fmt.Println("\nViterbi:")
	var  weather []string = []string{"Sunny","Cloudy","Rainny"}
	var humidity []string = []string{"Dry","Dryish","Damp","Soggy"}
	var TransMat lib.Matrix
	var ConfusionMat lib.Matrix
	TransMat.Data = []float64{0.5,0.375,0.125,0.25,0.125,0.5,0.25,0.5,0.375}
	TransMat.Col = 3
	TransMat.Row = 3
	ConfusionMat.Data = []float64{0.6,0.2,	0.15,0.05,0.25,0.25,0.25,0.25,0.05,0.10,0.35,0.50}
	ConfusionMat.Col = 4
	ConfusionMat.Row = 3
	fmt.Println("the transfer matrix:")
	TransMat.Print()

	fmt.Println("the confuses matrix:")
	ConfusionMat.Print()

	ObserveStat := []int{Dry, Damp, Soggy,Damp,Dryish}
	fmt.Print("\nthe Observe sequence:\n\t")
	for i:=0; i<len(ObserveStat);i++{
		fmt.Print(humidity[ObserveStat[i]]," ")
	}

	Pai := []float64{0.63,0.17,0.20}
	fmt.Println("\n\nPai:",Pai)

	Viterbi(ObserveStat,Pai,[]float64{},0,TransMat ,ConfusionMat)

	fmt.Print("\nthe Hidden sequence:\n\t")
	for i:=0; i<len(flags);i++{
		fmt.Print(weather[flags[i]]," ")
	}
	fmt.Println(" ")

}

func Viterbi(obs []int, pai []float64,preflags []float64,serials int,TransMat lib.Matrix,ConfusionMat lib.Matrix){
	if len(preflags) <= 0 {
		preflags = make([]float64,len(pai))
		var flag int =0
		var maxp float64 = -1
		for i := 0; i < len(pai); i++ {
			confusion, _ := ConfusionMat.Get(i, obs[serials])
			prop := pai[i] * confusion
			if prop > maxp {
				flag = i
				maxp = prop
			}
			preflags[i] = prop
		}
		flags = append(flags,flag)
		Viterbi(obs,pai,preflags,serials+1,TransMat,ConfusionMat)
	}else {
		if serials > len(obs)-1{
			return
		}

		var preflagsnew []float64 = make([]float64,len(preflags))
		var flag int =0
		var maxp float64 = -1
		for i:=0; i<len(pai); i++{
			var mp float64 = -1
			for j:=0; j<len(preflags); j++{
				trans,_ := TransMat.Get(j,i)
				prop := preflags[j] * trans
				if  prop > mp{
					mp = prop
				}
			}
			confusion,_ := ConfusionMat.Get(i, obs[serials])
			p := mp * confusion
			if p > maxp{
				maxp = p
				flag = i
			}
			preflagsnew[i] = p
		}
		flags = append(flags,flag)
		preflags = preflagsnew
		Viterbi(obs,pai,preflags,serials+1,TransMat,ConfusionMat)
	}
}


func BaumWelch(){

}

