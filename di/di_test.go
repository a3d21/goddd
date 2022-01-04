package di

import (
	"bytes"
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
)

var gen = flag.Bool("gen", false, "generates dot & svg")

func TestGenProdDIGraph(t *testing.T) {
	c := newProdContainer()
	genDot(t, c, "deps-prod", *gen)
}
func TestGenMemDIGraph(t *testing.T) {
	c := newProdContainer()
	genDot(t, c, "deps-mem", *gen)
}
func TestGenTestDIGraph(t *testing.T) {
	c := newProdContainer()
	genDot(t, c, "deps-test", *gen)
}

func genDot(t *testing.T, c *dig.Container, name string, gen bool) {

	var b bytes.Buffer
	err := dig.Visualize(c, &b)
	assert.Nil(t, err)

	if gen {
		t.Log("generates dot * svg")

		dotFile := filepath.Join(".", name+".dot")
		err = ioutil.WriteFile(dotFile, b.Bytes(), 0644)
		assert.Nil(t, err)

		g := graphviz.New()
		graph, err := graphviz.ParseBytes(b.Bytes())
		assert.Nil(t, err)
		graph.SetRankDir(cgraph.LRRank)

		err = g.RenderFilename(graph, graphviz.SVG, name+".svg")
		assert.Nil(t, err)
	}
}
