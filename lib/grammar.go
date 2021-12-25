package tafl

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"sync"
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

func (g Grammar) String() string {
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
			g.DeleteProductionAt(i)
		}else{
			i++
			for barren := range barrenTerminals{
				g.productions[key].DeleteTerminal(barren)
			}
		}
	}
}

func (g *Grammar) DeleteProductionAt(index int) {
	key := g.keys[index]
	g.keys = append(g.keys[:index], g.keys[index+1:]...)
	delete(g.productions, key)
}


func (g *Grammar) DeleteUnreachable(){
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	reachable := map[rune]struct{}{}
	visited := map[rune]struct{}{}
	var visit func(rune)
	visit = func(node rune) {
		defer wg.Done()
		if _, seen := visited[node]; seen {
			return
		}
		mu.Lock()
		visited[node] = struct{}{}
		mu.Unlock()
		if prod, ok := g.productions[node]; ok {
			nodes := prod.GetTerminals()
			for _, n := range nodes {
				mu.Lock()
				reachable[n] = struct{}{}
				mu.Unlock()
				wg.Add(1)
				visit(n)
				}
			}
		}
	reachable[g.keys[0]] = struct{}{}
	wg.Add(1)
	go visit(g.keys[0])
	wg.Wait()

	i := 0
	for i < len(g.keys){
		key := g.keys[i]
		if _, isReachable := reachable[key]; !isReachable{
			g.DeleteProductionAt(i)
		}else{
			i++
		}
	}
}
