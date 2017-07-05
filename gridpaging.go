package gridpaging

import(
	"linewrap"
)

type Datarow struct {
	Columns []string
}
type Dataline struct {
	Cols []string
	Bg string
	IRow int
	NLinesInRow int
	LineNOInRow int
}
type Cellline struct {
	Lines []int
}
/*func measureTextWidth(fontname,text string,fontsize int) (w float64){		//parameter func sample
	canvas.SetFont(fontname, "", fontsize)
	w = canvas.MeasureTextWidth( text )
	return
}*/
func SplitGrid2Line( w float64,fields []string, cells []Datarow,fontname string,fontsize int, 
						measureTextWidth func(txt,fontname string,fontsize int) float64 )(widths []float64,lines []Dataline,clines[] Cellline){
	ncols := len(fields)
	nrows := len(cells)
	widths = make( []float64,ncols )
	heights := make( []int,nrows )
	clines = make( []Cellline,nrows )	//lines in one row
	var totalWidth float64 = 0.0
	var totalSmall float64 = 0.0
	var totalLarge float64 = 0.0
	var nLarges int= 0
	for i:=0;i<ncols;i++ {
		ww := measureTextWidth(fields[i],fontname,fontsize)
		widths[i] = ww+3
		for j:=0;j<nrows;j++ {
			ww = measureTextWidth(cells[j].Columns[i],fontname,fontsize)
			ww += 3
			if ww>widths[i] {
				widths[i] = ww
			}
		}
		totalWidth += widths[i]
		if widths[i]<=150 {
			totalSmall += widths[i]
		}else{
			totalLarge += widths[i]
			nLarges ++
		}
	}
	if totalWidth>w {
		totalRemains := w - totalSmall
		j := 0
		totalWidth = 0.0
		for i:=0;i<ncols;i++ {
			w := widths[i]
			if w>150 {
				j++
				w = float64(w/totalLarge)*float64(totalRemains)
				if w<80 { w = 80 }
				if j==nLarges && totalRemains>w {
					w = totalRemains
				}
				totalRemains -= w
				widths[i] = w
			}
			totalWidth += widths[i]
		}
	}
	totalRows := 0
	for j:=0;j<nrows;j++ {
		nn := 1
		clines[j].Lines = make([]int,ncols,ncols)
		for i:=0;i<ncols;i++ {
			m := len( linewrap.Split2MultiLine( cells[j].Columns[i],fontname,widths[i],fontsize,measureTextWidth) )
			clines[j].Lines[i] = m
			if m>nn { nn = m }
		}
		heights[j] = nn
		totalRows += nn
	}
	lines = make([]Dataline,totalRows)
	r := 0
	for j:=0;j<nrows;j++ {
		nn := heights[j]
		for i:=0;i<nn;i++ {
			lines[r+i].Cols = make([]string,ncols,ncols)
			lines[r+i].NLinesInRow = nn
			lines[r+i].LineNOInRow = i
			lines[r+i].IRow = j
		}
		for i:=0;i<ncols;i++ {
			lns := linewrap.Split2MultiLine( cells[j].Columns[i],fontname,widths[i],fontsize,measureTextWidth)
			for k,ln := range lns {
				lines[r+k].Cols[i] = ln
			}
		}
		r += nn
	}
	return
}
