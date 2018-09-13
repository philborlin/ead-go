package generatedsl

import (
	"math/rand"
	"testing"
	"time"
)

type testDSLGenerator struct {
	f map[string]interface{}
}

func (g *testDSLGenerator) functions() map[string]interface{} {
	return g.f
}

func (g *testDSLGenerator) packageName() string {
	return "test"
}

func (g *testDSLGenerator) dslName() string {
	return "Test"
}

func newTwoCommandDSLGenerator() *testDSLGenerator {
	return &testDSLGenerator{
		f: map[string]interface{}{
			"Int":  rand.Int,
			"Intn": rand.Intn,
		},
	}
}

func newOneCommandDSLGenerator() *testDSLGenerator {
	return &testDSLGenerator{
		f: map[string]interface{}{
			"Now": time.Now,
		},
	}
}

func TestGenerateOneFunction(t *testing.T) {
	file, err := createDSL(newOneCommandDSLGenerator())
	if err != nil {
		t.Error(err)
	}

	if file.GoString() != randFile {
		t.Errorf("File was incorrect, expected: %s, actual: %s.", randFile, file.GoString())
	}

}

var randFile = `package test

import (
	"time"

	"github.com/philborlin/ead/stack"
)

type NowCmd struct {
	t *time.Time
}

func Now() time.Time {
	t := new(time.Time)
	stack.Add(&NowCmd{t})
	return *t
}

func (c *NowCmd) Interpret() error {
	*c.t = time.Now()
	return nil
}
`
