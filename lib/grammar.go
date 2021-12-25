package tafl

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Grammar struct{
	productions map[rune]*ProductionList
	keys []rune
}

func NewGrammar(r io.Reader) (Grammar, error){
	res := Grammar{
		make(map[rune]*ProductionList),
		make([]rune, 0),
	}
	scan := bufio.NewScanner(r)
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		line := scan.Text()
		line = strings.TrimSpace(line)
		t := rune(line[0])
		if unicode.IsLetter(t) && unicode.IsUpper(t){
			res.keys = append(res.keys, t)
			res.productions[t] = FromString(line[3:], "|")
		}else{
			return Grammar{}, fmt.Errorf("invalid format of input data, \"%s\" is not a letter, or a terminal", string(t))
		}
	}
	return res, nil
}

func (g *Grammar) String() string {
	buf := strings.Builder{}

	for _, k := range g.keys{
		fmt.Fprintf(&buf, "%s -> %s\n", string(k), strings.Join(*g.productions[k],"|"))
		}
	return buf.String()
	}

func (g *Grammar) DeleteBarren() {
	barrenTerminals := make(map[rune]struct{})
	oldLen := len(barrenTerminals)
	for true{
		for _, k := range g.keys{
			if (*g).productions[k].IsBarren(barrenTerminals){
				barrenTerminals[k] = struct{}{}
			}
		}
		if len(barrenTerminals) == oldLen{
			break
		}else{
			oldLen = len(barrenTerminals)
		}
	}
	i := 0
	for i < len(g.keys){
		key := g.keys[i]
		if _, isBarren := barrenTerminals[key]; isBarren{
			g.keys = append(g.keys[:i], g.keys[i+1:]...)
			delete(g.productions, key)
		}else{
			i++
			for barren := range barrenTerminals{
				g.productions[key].DeleteTerminal(barren)
			}
		}
	}
}


