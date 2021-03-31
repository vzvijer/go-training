package listops

import (
	"fmt"
)

type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

func (il IntList) Length() int {
	return len(il)
}

func (il IntList) Foldr(f binFunc, init int) int {
	fmt.Println("Foldr:", il, init)
	//for i := len(il) - 1; i >= 0; i-- {
	for i := 0; i < len(il); i++ {
		init = f(init, il[i])
		//init = f(il[i], init)
	}
	return init
}

func (il IntList) Foldl(f binFunc, init int) int {
	fmt.Println("Foldl:", il, init)
	for i := 0; i < len(il); i++ {
		//init = f(init, il[i])
		init = f(il[i], init)
	}
	return init
}

func (il IntList) Filter(f predFunc) IntList {
	res := IntList{}
	for _, val := range il {
		if f(val) {
			res = append(res, val)
		}
	}
	return res
}

func (il IntList) Map(f unaryFunc) IntList {
	res := IntList{}
	for _, val := range il {
		res = append(res, f(val))
	}
	return res
}

func (il IntList) Reverse() IntList {
	n := len(il)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = il[n-1-i]
	}
	return IntList(res)
}

func (il IntList) Append(ail IntList) IntList {
	res := IntList{}
	for _, val := range il {
		res = append(res, val)
	}
	for _, val := range ail {
		res = append(res, val)
	}
	return res
}

func (il IntList) Concat(cill []IntList) IntList {
	res := IntList{}
	res = res.Append(il)
	for _, cil := range cill {
		res = res.Append(cil)
	}
	return res
}
