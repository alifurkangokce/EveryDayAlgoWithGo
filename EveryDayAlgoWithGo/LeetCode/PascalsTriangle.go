package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	res := [][]int{{1}, {1, 1}}
	for i := 2; i < numRows; i++ {
		tmp := []int{1}
		for j := 1; j < i; j++ {
			tmp = append(tmp, res[i-1][j-1]+res[i-1][j])
		}
		tmp = append(tmp, 1)
		res = append(res, tmp)
	}
	return res

}
