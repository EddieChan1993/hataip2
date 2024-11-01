package core

import (
	"math"
	"strconv"
)

/**
https://zhuanlan.zhihu.com/p/689510512
*/

type LogicInput struct {
	HoldingCost     float64 //持仓成本
	FundValue       float64 //基金净值
	OriInvestAmount float64 //原始定投金额
}

type LogicRes struct {
	Input     *LogicInput
	Rate      float64 //(持仓成本-基金净值）/基金净值=浮动差
	TotalRate float64 //实际定投比例(80~140)
	Total     float64 //建议定投金额
}

const minRate = 50.0    //最小定投原始比例
const maxRate = 140.0   //最大定投原始比例
const FundCaLine = 25.0 //(持仓成本-基金净值)/持仓成本 5%浮动

func LogicCal(params *LogicInput) *LogicRes {
	res := &LogicRes{
		Input: params,
	}
	rateCa := (params.HoldingCost - params.FundValue) / params.FundValue * 100
	res.Rate = rateCa
	if rateCa <= -FundCaLine {
		//持仓成本<=基金净值 增幅小于fundCaLine%
		res.Total = minRate / 100 * params.OriInvestAmount
		res.TotalRate = minRate
	} else if rateCa > -FundCaLine && rateCa < FundCaLine {
		totalRate := 0.0
		rateSure := 0.0
		if rateCa >= 0 {
			//持仓成本>=基金净值
			rateSure = math.Pow(rateCa/FundCaLine, 2)
			totalRate = 100 + rateSure*(maxRate-100)
		} else {
			rateSure = math.Pow(math.Abs(FundCaLine+rateCa)/FundCaLine, 3)
			totalRate = minRate + rateSure*(100-minRate)
		}
		res.Rate = rateCa
		res.TotalRate = totalRate
		res.Total = totalRate / 100 * params.OriInvestAmount
	} else {
		//持仓成本>=基金净值 增幅超过fundCaLine%
		res.TotalRate = maxRate
		res.Total = maxRate / 100 * params.OriInvestAmount
	}
	return res
}

func Float642str(f float64, prec int) string {
	return strconv.FormatFloat(f, 'f', prec, 64)
}
