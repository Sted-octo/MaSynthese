package main

import "Octoptimist/domain"

type SynthesisLines []domain.SynthesisLine

func (sl *SynthesisLines) Accumulate() []domain.SynthesisLine {
	currentKind := ""

	var newListSynthesisLine []domain.SynthesisLine
	var kindSubTotal domain.SynthesisLine
	var kindTotal domain.SynthesisLine = domain.SynthesisLine{Kind: "total", Title: "Total", TimeSum: 0,
		IsLineTotal: true, RowCount: 0}
	for _, line := range *sl {
		if currentKind != line.Kind {
			if currentKind != "" {
				if kindSubTotal.RowCount > 0 {
					newListSynthesisLine = append(newListSynthesisLine, kindSubTotal)
				}

				currentKind = ""
			}
			newListSynthesisLine = append(newListSynthesisLine, line)
			kindTotal.TimeSum += line.TimeSum
			kindTotal.TimeSumDone += line.TimeSumDone
			kindTotal.TimeSumTodo += line.TimeSumTodo
			kindTotal.RowCount++

			if currentKind == "" {
				currentKind = line.Kind
				kindSubTotal = domain.SynthesisLine{
					Kind:           line.Kind,
					Title:          "Sous total " + domain.KindTranslator(line.Kind),
					TimeSum:        line.TimeSum,
					TimeSumDone:    line.TimeSumDone,
					TimeSumTodo:    line.TimeSumTodo,
					IsLineSubTotal: true,
					RowCount:       1}
			}
			continue
		}
		newListSynthesisLine = append(newListSynthesisLine, line)
		kindSubTotal.TimeSum += line.TimeSum
		kindSubTotal.TimeSumDone += line.TimeSumDone
		kindSubTotal.TimeSumTodo += line.TimeSumTodo
		kindSubTotal.RowCount++
		kindTotal.TimeSum += line.TimeSum
		kindTotal.TimeSumDone += line.TimeSumDone
		kindTotal.TimeSumTodo += line.TimeSumTodo
		kindTotal.RowCount++
	}
	if kindSubTotal.RowCount > 0 {
		newListSynthesisLine = append(newListSynthesisLine, kindSubTotal)
	}

	if kindTotal.RowCount > 0 {
		newListSynthesisLine = append(newListSynthesisLine, kindTotal)
	}

	return newListSynthesisLine
}
