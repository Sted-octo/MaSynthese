package main

type SynthesisLines []SynthesisLine

func (sl *SynthesisLines) Accumulate() []SynthesisLine {
	currentKind := ""

	var newListSynthesisLine []SynthesisLine
	var kindSubTotal SynthesisLine
	var kindTotal SynthesisLine = SynthesisLine{Kind: "total", Title: "Total", TimeSum: 0,
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
				kindSubTotal = SynthesisLine{
					Kind:           line.Kind,
					Title:          "Sous total " + line.Kind,
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
