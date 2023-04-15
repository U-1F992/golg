package golg

func formatCell(cc *ConsoleColor, upper float32, lower float32) string {
	return NewConsoleColor(
		getIntermediate(cc.bg, cc.fg, upper),
		getIntermediate(cc.bg, cc.fg, lower)).Apply("\u2580")
}

func Format(lg *LifeGame, cc *ConsoleColor) string {
	f := &lg.field
	ret := ""
	for r := 0; r < len(f); r++ {
		if r%2 == 1 {
			continue
		}
		for c := 0; c < len(f[r]); c++ {
			ret += formatCell(cc, f[r][c], f[r+1][c])
		}
		ret += "\n"
	}
	return ret
}
