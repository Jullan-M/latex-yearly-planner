package header

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/components/hyper"
)

type MonthItem struct {
	Val     time.Month
	bold	bool
	ref     bool
	shorten bool
}

func (m MonthItem) Display() string {
	ref := m.Val.String()
	text := ref

	if m.shorten {
		text = text[:3]
	}

	if m.bold {
		text = "\\textbf{" + text + "}"
	}

	if m.ref {
		return hyper.Target(ref, text)
	}

	return hyper.Link(ref, text)
}

func (m MonthItem) Ref() MonthItem {
	m.ref = true

	return m
}

func (m MonthItem) Bold(f bool) MonthItem {
	m.bold = f

	return m
}

func (m MonthItem) Shorten(f bool) MonthItem {
	m.shorten = f

	return m
}

func NewMonthItem(m time.Month) MonthItem {
	return MonthItem{Val: m}
}
