package data_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) []int32 {
	slice := []int32{}
	for _,value := range(vals){
		slice = append(slice, op(value))
	}
	
	return slice

}

func filterInts(op filterOperation, vals []int32) []int32 {
	slice := []int32{}
	for _,value := range(vals){
		if op(value){
			slice = append(slice, value)
		}
	}
	
	return slice
}

func concatenate(dest []string, newValues ...string) []string {
	for  _, value := range newValues {
		dest = append(dest, value)
	}
	return dest
}

func equals(list1 []string, list2 []string) bool {
	return true
}

func partialReverse(src []int, from, to int) []int {
	src = src[from:to+1]
	slice := []int{}
	for index ,_ := range src {
		slice = append(slice, src[len(src)-1-index])
	}
	
	return slice
}
