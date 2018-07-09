package efmt

import (
	"fmt"
)

// package to build hierarchial numbering tag and Stringify them

// for proper numbering
type Ntag struct {
	id   []int
	lvl  int
	iStr string
}

// create and setup @ level 1
func NewNtag() (ret *Ntag) {
	nt := &Ntag{make([]int, 1), 0, "  "}
	nt.id[0] = 1
	return nt
}

func (nt *Ntag) Clone() *Ntag {
	rNt := &Ntag{make([]int, nt.lvl+1), nt.lvl, nt.iStr}
	copy(rNt.id, nt.id)
	return rNt
}

func (nt *Ntag) SetIndent(indentStr string) {
	nt.iStr = indentStr
}

func (nt *Ntag) Push() {
	// level zero never used.
	nt.lvl++
	nt.id = append(nt.id, 0)
	nt.id[nt.lvl] = 1
}
func (nt *Ntag) Pop() {
	if nt.lvl == 0 {
		fmt.Println("Panicking: Ntag.lvl can't Decr @ 0")
		panic(fmt.Sprintf("%+v", nt))
	}
	nt.id[nt.lvl] = 0
	nt.lvl--
}
func (nt *Ntag) Next() {
	nt.id[nt.lvl] += 1
}

// Output functions
// String() and Indent()
//
func (nt *Ntag) String() string {
	ret := ""
	for i := 0; i <= nt.lvl; i++ {
		tmpStr := fmt.Sprintf("%d", nt.id[i])
		ret += tmpStr
		if i < nt.lvl {
			ret += "."
		}
	}
	return ret
}

func (nt *Ntag) Indent() string {
	retString := ""
	for i := 0; i < nt.lvl; i++ {
		retString += nt.iStr
	}
	return retString
}
