package header

import (
	"strconv"

	"github.com/kudrykv/latex-yearly-planner/app/components/hyper"
)

type IntItem struct {
	Val int
	bold bool
	ref bool
}

func (i IntItem) Display() string {
	var out string
	ref := strconv.Itoa(i.Val)
	text := ref

	if i.bold {
		text = "\\textbf{" + text + "}"
	}

	if i.ref {
		out = hyper.Target(ref, text)
	} else {
		out = hyper.Link(ref, text)
	}

	return out
}

func (i IntItem) Ref() IntItem {
	i.ref = true

	return i
}

func (i IntItem) Bold(f bool) IntItem {
	i.bold = f

	return i
}

func NewIntItem(val int) IntItem {
	return IntItem{Val: val}
}
