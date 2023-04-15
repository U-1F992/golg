package golg

import (
	"math"
	"strconv"
)

type Color struct {
	r uint8
	g uint8
	b uint8

	h int // 0-360
	s int // 0-100
	v int // 0-100
}

func NewColorHSV(h int, s int, v int) *Color {
	if h < 0 {
		return NewColorHSV(h+360, s, v)
	} else if 360 <= h {
		return NewColorHSV(h-360, s, v)
	}
	if s < 0 {
		return NewColorHSV(h, 0, v)
	} else if 100 < s {
		return NewColorHSV(h, 100, v)
	}
	if v < 0 {
		return NewColorHSV(h, s, 0)
	} else if 100 < s {
		return NewColorHSV(h, s, 100)
	}
	h_f := float64(h)
	s_f := float64(s) / 100 * 255
	v_f := float64(v) / 100 * 255

	max := v_f
	min := max - ((s_f / 255) * max)
	var r float64 = 0
	var g float64 = 0
	var b float64 = 0
	switch {
	case 0 <= h && h <= 60:
		r = max
		g = (h_f/60)*(max-min) + min
		b = min
	case 60 <= h && h <= 120:
		r = ((120-h_f)/60)*(max-min) + min
		g = max
		b = min
	case 120 <= h && h <= 180:
		r = min
		g = max
		b = ((h_f-120)/60)*(max-min) + min
	case 180 <= h && h <= 240:
		r = min
		g = ((240-h_f)/60)*(max-min) + min
		b = max
	case 240 <= h && h <= 300:
		r = ((h_f-240)/60)*(max-min) + min
		g = min
		b = max
	case 300 <= h && h <= 360:
		r = max
		g = min
		b = ((360-h_f)/60)*(max-min) + min
	}

	return &Color{
		r: uint8(math.Round(r)),
		g: uint8(math.Round(g)),
		b: uint8(math.Round(b)),

		h: h,
		s: s,
		v: v,
	}
}

func NewColorRGB(r uint8, g uint8, b uint8) *Color {
	r_f := float64(r)
	g_f := float64(g)
	b_f := float64(b)

	min := math.Min(r_f, math.Min(g_f, b_f))
	max := math.Max(r_f, math.Max(g_f, b_f))
	var h float64 = 0
	switch {
	case r == g && g == b:
		h = 0
	case g <= r && b <= r:
		h = 60 * ((g_f - b_f) / (max - min))
	case r <= g && b <= g:
		h = 60*((b_f-r_f)/(max-min)) + 120
	case r <= b && g <= b:
		h = 60*((r_f-g_f)/(max-min)) + 240
	}
	if h < 0 {
		h += 360
	}
	s := (max - min) / max // 0-1
	v := max               // 0-255

	return &Color{
		r: r,
		g: g,
		b: b,

		h: int(math.Round(h)),
		s: int(math.Round(s * 100)),
		v: int(math.Round(v / 255 * 100)),
	}
}

func (c *Color) AddHue(hue int) *Color {
	return NewColorHSV(c.h+hue, c.s, c.v)
}

func getIntermediate(from *Color, to *Color, val float32) *Color {
	switch {
	case val < 0:
		return getIntermediate(from, to, 0)
	case 1 < val:
		return getIntermediate(from, to, 1)
	}

	r := uint8(float32(from.r) + float32(to.r-from.r)*val)
	g := uint8(float32(from.g) + float32(to.g-from.g)*val)
	b := uint8(float32(from.b) + float32(to.b-from.b)*val)
	return NewColorRGB(r, g, b)

	// from + (to - from) * val

	// e.g. #808080 -> #000000, 0.5
	// 0x80 + (0x00 - 0x80) * 0.5
}

type ConsoleColor struct {
	fg *Color
	bg *Color
}

func NewConsoleColor(fg *Color, bg *Color) *ConsoleColor {
	return &ConsoleColor{
		fg: fg,
		bg: bg,
	}
}

const Esc = "\033"
const Reset = Esc + "[0m"
const Clear = Esc + "[2J"
const Top = Esc + "[H"

func (cc *ConsoleColor) Apply(s string) string {
	fore := Esc + "[38;2;" +
		strconv.FormatInt(int64(cc.fg.r), 10) + ";" +
		strconv.FormatInt(int64(cc.fg.g), 10) + ";" +
		strconv.FormatInt(int64(cc.fg.b), 10) + "m"
	back := Esc + "[48;2;" +
		strconv.FormatInt(int64(cc.bg.r), 10) + ";" +
		strconv.FormatInt(int64(cc.bg.g), 10) + ";" +
		strconv.FormatInt(int64(cc.bg.b), 10) + "m"

	return fore + back + s + Reset
}
