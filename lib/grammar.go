package tafl

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Grammar map[rune]ProductionList

func NewGrammar(r io.Reader) (Grammar, error){
	res := make(Grammar)
	scan := bufio.NewScanner(r)
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		line := scan.Text()
		line = strings.TrimSpace(line)
		t := rune(line[0])
		if unicode.IsLetter(t) && unicode.IsUpper(t){
			res[t] = FromString(line[3:], "|")
		}else{
			return nil, fmt.Errorf("invalid format of input data, \"%s\" is not a letter, or a terminal", string(t))
		}
	}
	return res, nil
}

func (g Grammar) String() string {
	buf := strings.Builder{}
	for k,v := range g{
		fmt.Fprintf(&buf, "%s -> %s\n", string(k), strings.Join(v,"|"))
		}
	return buf.String()
	}



