package test

import (
	"fmt"
	"hataip2/core"
	"testing"
)

func TestCal(t *testing.T) {
	p := &core.LogicInput{
		HoldingCost:     1.0160,
		FundValue:       1.0160,
		OriInvestAmount: 50,
	}
	base := p.FundValue
	for i := base - base*5.0/100; i < base+base*5.0/100; i += 0.001 {
		p.FundValue = i
		res := core.LogicCal(p)
		fmt.Println(fmt.Sprintf("%.4f %.4f %.4f %.4f ", p.FundValue, res.Rate, res.TotalRate, res.Total))
	}
}
