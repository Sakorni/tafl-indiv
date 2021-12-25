package tafl

import (
	"strings"
	"unicode"
)

type ProductionList []string
const eps = `\e`

func FromString(source, delim string) *ProductionList{
	a := ProductionList(strings.Split(source, delim))
	return &a
}

func (p *ProductionList) GetTerminals()[]rune{
	res := make([]rune,0)
	for _, prod := range *p{
		for _, r := range prod{
		if unicode.IsUpper(r){
			res = append(res, r)
			}
		}
	}
	return res
}

func (p *ProductionList) IsBarren(notBarrenTerminals map[rune]struct{}) bool{
	barren := false
	var buf []rune
	for k := range notBarrenTerminals{
		buf = append(buf, k)
	}
	NBT := string(buf)
	contains := func(r rune)bool{ return strings.ContainsRune(NBT, r)}
	for _, prod := range *p{
		for _, s := range prod{
			if unicode.IsUpper(s) && !contains(s){
				barren = true
				break
			}
		}
		if !barren{
			return false
		}
		barren = false
	}
	return true
}

func (p *ProductionList) DeleteTerminal(terminal rune) {
	newProdList := ProductionList{}
	for _, prod := range *p{
		if strings.ContainsRune(prod, terminal){
			continue
		}
		newProdList = append(newProdList, prod)
	}
	*p = newProdList
}