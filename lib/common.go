package lib


type IntSlice []int

func (c IntSlice) Len() int {
	return len(c)
}
func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c IntSlice) Less(i, j int) bool {
	return c[i] < c[j]
}



type Float64Slice []float64

func (c Float64Slice) Len() int {
	return len(c)
}
func (c Float64Slice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c Float64Slice) Less(i, j int) bool {
	return c[i] < c[j]
}