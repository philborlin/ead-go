package generatedsl

import (
	"os"
	"reflect"

	"github.com/dave/jennifer/jen"
)

// DSLGenerator is an interface that returns functions that should be generated
type DSLGenerator interface {
	// This must return a slice of functions or an error condition will occur
	functions() map[string]interface{}
	packageName() string
	dslName() string
}

func createDSLAndWriteToFile(g DSLGenerator) (*os.File, error) {
	return nil, nil
}

func createDSL(g DSLGenerator) (*jen.File, error) {
	file := jen.NewFile(g.packageName())

	file.Type().Id(g.dslName()).Struct(
		jen.Id("stack").Op("*").Qual("github.com/philborlin/ead/stack", "Stack"),
	)

	for fName, f := range g.functions() {
		t := reflect.TypeOf(f)

		var args []jen.Code
		for i := 0; i < t.NumIn(); i++ {
			inV := t.In(i)

			if inV.Kind().String() == "struct" {
				qual := jen.Qual(inV.PkgPath(), inV.Name())
				args = append(args, qual)
			}
		}

		var returns []jen.Code
		for o := 0; o < t.NumOut(); o++ {
			returnV := t.Out(o)

			if returnV.Kind().String() == "struct" {
				qual := jen.Qual(returnV.PkgPath(), returnV.Name())
				returns = append(returns, qual)
			}
		}

		// type NowCmd struct {
		// 	t *time.Time
		// }

		// func Now() time.Time {
		// 	t := new(time.Time)
		// 	stack.Add(&NowCmd{t})
		// 	return *t
		// }

		// func (c *NowCmd) Interpret() error {
		// 	*c.t = time.Now()
		// 	return nil
		// }

		file.Func().
			Params(jen.Id("c").Op("*").Id(g.dslName())).
			Id(fName).
			Params(jen.List(args...)).
			List(returns...).
			Block(
				// TODO This has to handle multiple return values
				jen.Id("a").Op(":=").New(returns[0]),
				jen.Return().Op("*").Id("a"),
			)
	}

	return file, nil
}

func createFromFunction(function string, f *jen.File) {

}
