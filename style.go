package clog

import (
	"strings"

	"github.com/pterm/pterm"
)

type Style struct {
	paddingLeft  int
	paddingRight int
	value        string
	style        pterm.RGBStyle
}

func NewStyle(s pterm.RGBStyle) Style {
	return Style{
		paddingLeft:  0,
		paddingRight: 0,
		value:        "",
		style:        s,
	}
}

func (s Style) PaddingLeft(n int) Style {
	s.paddingLeft = n
	return s
}

func (s Style) PaddingRight(n int) Style {
	s.paddingRight = n
	return s
}

func (s Style) SetString(strs ...string) Style {
	if len(s.value) > 0 {
		s.value = s.value + " " + strings.Join(strs, " ")
	} else {
		s.value = strings.Join(strs, " ")
	}
	return s
}

func (s Style) Foreground(style pterm.RGB) Style {
	s.style.Foreground = style
	return s
}

func (s Style) Background(style pterm.RGB) Style {
	s.style.Background = style
	return s
}

func (s Style) Bold(f bool) Style {
	if f {
		s.style = s.style.AddOptions(pterm.Bold)
	}
	return s
}

func (s Style) String() string {
	return s.Render()
}

func (s Style) Render(strs ...string) string {
	if s.value != "" {
		strs = append([]string{s.value}, strs...)
	}

	text := strings.Join(strs, " ")
	if s.paddingLeft > 0 {
		text = strings.Repeat(" ", s.paddingLeft) + text
	}

	if s.paddingRight > 0 {
		text = text + strings.Repeat(" ", s.paddingRight)
	}

	return s.style.Sprint(text)
}

type Theme []Style

var (
	FgWhith = pterm.NewRGBStyle(pterm.NewRGB(255, 255, 255))

	NoneColorTheme = Theme{
		DebugLevel: NewStyle(FgWhith),
		InfoLevel:  NewStyle(FgWhith),
		WarnLevel:  NewStyle(FgWhith),
		ErrorLevel: NewStyle(FgWhith),
		FatalLevel: NewStyle(FgWhith),
	}

	DefaultTheme = Theme{
		DebugLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(94, 101, 85))).Bold(true),
		InfoLevel:  NewStyle(pterm.NewRGBStyle(pterm.NewRGB(4, 124, 140))).Bold(true),
		WarnLevel:  NewStyle(pterm.NewRGBStyle(pterm.NewRGB(194, 156, 5))).Bold(true),
		ErrorLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(141, 35, 6))).Bold(true),
		FatalLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(246, 60, 9))).Bold(true),
	}

	CoolTheme = Theme{
		DebugLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(120, 144, 156))),
		InfoLevel:  NewStyle(pterm.NewRGBStyle(pterm.NewRGB(38, 198, 218))),
		WarnLevel:  NewStyle(pterm.NewRGBStyle(pterm.NewRGB(255, 179, 0))),
		ErrorLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(244, 67, 54))),
		FatalLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(211, 47, 47))),
	}

	WarmTheme = Theme{
		DebugLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(189, 189, 189))),
		InfoLevel:  NewStyle(pterm.NewRGBStyle(pterm.NewRGB(255, 202, 40))),
		WarnLevel:  NewStyle(pterm.NewRGBStyle(pterm.NewRGB(255, 112, 67))),
		ErrorLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(211, 47, 47))),
		FatalLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(183, 28, 28))),
	}

	DarkContrastTheme = Theme{
		DebugLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(130, 130, 130))),
		InfoLevel:  NewStyle(pterm.NewRGBStyle(pterm.NewRGB(0, 188, 212))),
		WarnLevel:  NewStyle(pterm.NewRGBStyle(pterm.NewRGB(255, 235, 59))),
		ErrorLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(255, 87, 34))),
		FatalLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(244, 67, 54))),
	}

	VividTheme = Theme{
		DebugLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(102, 187, 106))),
		InfoLevel:  NewStyle(pterm.NewRGBStyle(pterm.NewRGB(3, 169, 244))),
		WarnLevel:  NewStyle(pterm.NewRGBStyle(pterm.NewRGB(255, 193, 7))),
		ErrorLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(244, 67, 54))),
		FatalLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(156, 39, 176))),
	}

	BusinessTheme = Theme{
		DebugLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(102, 102, 153))), // muted purple
		InfoLevel:  NewStyle(pterm.NewRGBStyle(pterm.NewRGB(51, 102, 204))),  // calm blue
		WarnLevel:  NewStyle(pterm.NewRGBStyle(pterm.NewRGB(153, 102, 255))), // soft violet
		ErrorLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(204, 51, 102))),  // muted magenta
		FatalLevel: NewStyle(pterm.NewRGBStyle(pterm.NewRGB(102, 0, 153))),   // deep purple
	}
)
