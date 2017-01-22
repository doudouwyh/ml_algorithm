package main

import(
	"wh_ml/test"
	"wh_ml/classifier"
	"wh_ml/hmm"
	"wh_ml/decisiontree"
	"wh_ml/em"
)

func main(){
	test.Matrix_Test()
	test.Dot_Test()
	test.TestStack()

	classifier.LeastSquare()
	classifier.Perceptron()
	classifier.NaiveBayes()
	classifier.KNN()

	decisiontree.ID3()
	decisiontree.C45()
	decisiontree.CART()

	em.K_Means()

	hmm.Viterbi_Test()
	hmm.BaumWelch()
}

