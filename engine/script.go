package engine

import (
	"code.google.com/p/mx3/script"
	"code.google.com/p/mx3/util"
	"io"
	"log"
)

func (f *ScalFn) Eval() interface{} {
	return (*f)()
}

func (f *ScalFn) Assign(e script.Expr) {
	(*f) = func() float64 { return e.Eval().(float64) }
}

func (f *VecFn) Eval() interface{} {
	return (*f)()
}

func (f *ScalFn) Assign(e script.Expr) {
	(*f) = func() [3]float64 {
		log.Println()
		return [3]float64{}
	}
}

func RunScript(src io.Reader) {
	p := script.NewParser()
	p.AddFloat("t", &Time)
	p.AddVar("aex", &Aex)
	p.AddVar("msat", &Msat)
	p.AddVar("alpha", &Alpha)
	p.AddVar("b_ext", &B_ext)
	util.FatalErr(p.Exec(src))

}
