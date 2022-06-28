package main

type SynthesisLines []SynthesisLine

func (sl *SynthesisLines) Accumulate() []SynthesisLine {
	currentKind := ""

	var newListSynthesisLine []SynthesisLine
	var kindSubTotal SynthesisLine
	var kindTotal SynthesisLine = SynthesisLine{Kind: "ztotal", Title: "Total", TimeSum: 0, TypeLine: LineTotal, RowCount: 0}
	for _, line := range *sl {
		if currentKind != line.Kind {
			if currentKind != "" {
				if kindSubTotal.RowCount > 1 {
					newListSynthesisLine = append(newListSynthesisLine, kindSubTotal)
				}

				currentKind = ""
			}
			newListSynthesisLine = append(newListSynthesisLine, line)
			kindTotal.TimeSum += line.TimeSum
			kindTotal.RowCount++

			if currentKind == "" {
				currentKind = line.Kind
				kindSubTotal = SynthesisLine{Kind: line.Kind, Title: "Sous total " + line.Kind, TimeSum: line.TimeSum, TypeLine: LineSubTotal, RowCount: 1}
			}
			continue
		}
		newListSynthesisLine = append(newListSynthesisLine, line)
		kindSubTotal.TimeSum += line.TimeSum
		kindSubTotal.RowCount++
		kindTotal.TimeSum += line.TimeSum
		kindTotal.RowCount++
	}
	if kindSubTotal.RowCount > 1 {
		newListSynthesisLine = append(newListSynthesisLine, kindSubTotal)
	}

	if kindTotal.RowCount > 1 {
		newListSynthesisLine = append(newListSynthesisLine, kindTotal)
	}

	return newListSynthesisLine
}
