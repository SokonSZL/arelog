package cfgProv

import (
	"errors"
	"github.com/mum4k/termdash/cell"
)

var colorList = map[string]cell.Color{
	"Black":   cell.ColorBlack,
	"Maroon":  cell.ColorMaroon,
	"Green":   cell.ColorGreen,
	"Olive":   cell.ColorOlive,
	"Navy":    cell.ColorNavy,
	"Purple":  cell.ColorPurple,
	"Teal":    cell.ColorTeal,
	"Silver":  cell.ColorSilver,
	"Gray":    cell.ColorGray,
	"Red":     cell.ColorRed,
	"Lime":    cell.ColorLime,
	"Yellow":  cell.ColorYellow,
	"Blue":    cell.ColorBlue,
	"Fuchsia": cell.ColorFuchsia,
	"Aqua":    cell.ColorAqua,
	"White":   cell.ColorWhite,
}

type UIElem int

const (
	ButtonPrimary UIElem = iota
	ButtonSecondary
	TextFieldCursor
	TextFieldFill
	TextFieldHighlight
	TextFieldText
)

type Config struct {
	LogSavePath             string //if "" use the same path as config
	TXCall                  string
	ButtonPrimaryColor      string
	ButtonSecondaryColor    string
	TextFieldCursorColor    string
	TextFieldFillColor      string
	TextFieldTextColor      string
	TextFieldHighlightColor string
}

func (c *Config) GetTXCall() string {
	return c.TXCall
}

func getColorFromName(s string) (cell.Color, error) {
	for key, val := range colorList {
		if key == s {
			return val, nil
		}
	}

	return cell.ColorBlack, errors.New("No color of that name found")
}

func (c *Config) GetColorOf(e UIElem) cell.Color {
	switch e {
	case ButtonPrimary:
		col, err := getColorFromName(c.ButtonPrimaryColor)
		if err != nil {
			return cell.ColorTeal
		} else {
			return col
		}
	case ButtonSecondary:
		col, err := getColorFromName(c.ButtonSecondaryColor)
		if err != nil {
			return cell.ColorFuchsia
		} else {
			return col
		}
	case TextFieldCursor:
		col, err := getColorFromName(c.TextFieldCursorColor)
		if err != nil {
			return cell.ColorFuchsia
		} else {
			return col
		}
	case TextFieldFill:
		col, err := getColorFromName(c.TextFieldFillColor)
		if err != nil {
			return cell.ColorFuchsia
		} else {
			return col
		}
	case TextFieldHighlight:
		col, err := getColorFromName(c.TextFieldHighlightColor)
		if err != nil {
			return cell.ColorFuchsia
		} else {
			return col
		}
	case TextFieldText:
		col, err := getColorFromName(c.TextFieldTextColor)
		if err != nil {
			return cell.ColorFuchsia
		} else {
			return col
		}
	default:
		return cell.ColorNavy
	}
}
