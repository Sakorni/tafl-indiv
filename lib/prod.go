package tafl

import (
	"strings"
	"unicode"
)

type ProductionList []string
const eps = `\e`

func FromString(source, delim string) ProductionList{
	return strings.Split(source, delim)
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

func (p *ProductionList) IsBarren(notBarrenTerminals []rune) bool{
	barren := false
	NBT := string(notBarrenTerminals)
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
