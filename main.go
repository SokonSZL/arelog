package main

import (
	"context"
	"log"
	//	"reflect"

	cfg "arelog/cfgProv"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/textinput"
)

/*
 * Command line arguments order:
 * arelog [qsodate|"today"] [qsostrt|"now"] [qsoend|"now"] [CALL] [BAND] ->
 * -> [MODE] [PWR] [RX RST] [TX RST] [notes]
 */

type QSOForm struct {
	QSODate   *textinput.TextInput
	QSOStart  *textinput.TextInput
	QSOEnd    *textinput.TextInput
	QSOTXCall *textinput.TextInput
	QSORXCall *textinput.TextInput
	QSOBand   *textinput.TextInput
	QSOMode   *textinput.TextInput
	QSOPower  *textinput.TextInput

	bLog    *button.Button
	bExit   *button.Button
	bRemove *button.Button
}

func (form *QSOForm) prefillForm(c *cfg.Config, cancel context.CancelFunc) error {
	DateInput, err := textinput.New(
		textinput.Label("DATE    "),
		textinput.LabelAlign(align.HorizontalLeft),
		textinput.MaxWidthCells(20),
		textinput.ExclusiveKeyboardOnFocus(),
		textinput.FillColor(cell.ColorTeal),
		textinput.CursorColor(cell.ColorSilver),
		textinput.HighlightedColor(cell.ColorFuchsia),
	)
	if err != nil {
		return err
	}

	StartInput, err := textinput.New(
		textinput.Label("START   "),
		textinput.LabelAlign(align.HorizontalLeft),
		textinput.MaxWidthCells(20),
		textinput.ExclusiveKeyboardOnFocus(),
	)
	if err != nil {
		return err
	}

	EndInput, err := textinput.New(
		textinput.Label("END     "),
		textinput.LabelAlign(align.HorizontalLeft),
		textinput.MaxWidthCells(20),
		textinput.ExclusiveKeyboardOnFocus(),
	)
	if err != nil {
		return err
	}

	txCall := c.GetTXCall()
	TXCallInput, err := textinput.New(
		textinput.Label("TX CALL "),
		textinput.LabelAlign(align.HorizontalLeft),
		textinput.PlaceHolder(txCall),
		textinput.MaxWidthCells(20),
		textinput.ExclusiveKeyboardOnFocus(),
	)
	if err != nil {
		return err
	}

	RXCallInput, err := textinput.New(
		textinput.Label("RX CALL "),
		textinput.LabelAlign(align.HorizontalLeft),
		textinput.MaxWidthCells(20),
		textinput.ExclusiveKeyboardOnFocus(),
	)
	if err != nil {
		return err
	}

	BandInput, err := textinput.New(
		textinput.Label("BAND    "),
		textinput.LabelAlign(align.HorizontalLeft),
		textinput.MaxWidthCells(20),
		textinput.ExclusiveKeyboardOnFocus(),
	)
	if err != nil {
		return err
	}

	ModeInput, err := textinput.New(
		textinput.Label("MODE    "),
		textinput.LabelAlign(align.HorizontalLeft),
		textinput.MaxWidthCells(20),
		textinput.ExclusiveKeyboardOnFocus(),
	)
	if err != nil {
		return err
	}

	PowerInput, err := textinput.New(
		textinput.Label("POWER   "),
		textinput.LabelAlign(align.HorizontalLeft),
		textinput.MaxWidthCells(20),
		textinput.ExclusiveKeyboardOnFocus(),
	)
	if err != nil {
		return err
	}

	logRemoveWidth := 10

	buttonLog, err := button.New(
		"Log",
		nil,
		button.Key(keyboard.KeyEnter),
		button.FillColor(cell.ColorTeal),
		button.FocusedFillColor(cell.ColorMagenta),
		button.DisableShadow(),
		button.Height(1),
		button.Width(logRemoveWidth),
	)
	if err != nil {
		return err
	}

	buttonRemove, err := button.New(
		"Remove",
		nil,
		button.Key(keyboard.KeyEnter),
		button.FillColor(cell.ColorTeal),
		button.FocusedFillColor(cell.ColorMagenta),
		button.DisableShadow(),
		button.Height(1),
		button.Width(logRemoveWidth),
	)

	buttonExit, err := button.New(
		"Exit",
		func() error {
			cancel()
			return nil
		},
		button.Key(keyboard.KeyEnter),
		button.DisableShadow(),
		button.FillColor(cell.ColorTeal),
		button.FocusedFillColor(cell.ColorMagenta),
		button.Height(1),
		button.Width(26),
	)

	form.QSODate = DateInput
	form.QSOStart = StartInput
	form.QSOEnd = EndInput
	form.QSOTXCall = TXCallInput
	form.QSORXCall = RXCallInput
	form.QSOBand = BandInput
	form.QSOMode = ModeInput
	form.QSOPower = PowerInput
	form.bExit = buttonExit
	form.bLog = buttonLog
	form.bRemove = buttonRemove

	return nil
}

func (form *QSOForm) makeMainLayout(c *container.Container) error {
	return c.Update("ROOT",
		container.KeyFocusNext(keyboard.KeyTab),
		container.KeyFocusGroupsNext(keyboard.KeyArrowDown, 1),
		container.KeyFocusGroupsPrevious(keyboard.KeyArrowUp, 1),
		container.KeyFocusGroupsNext(keyboard.KeyArrowRight, 2),
		container.KeyFocusGroupsPrevious(keyboard.KeyArrowLeft, 2),
		container.Border(linestyle.Light),
		container.SplitVertical(
			container.Left(
				container.SplitHorizontal(
					container.Top(
						container.SplitHorizontal(
							container.Top(
								container.SplitHorizontal(
									container.Top(
										container.SplitHorizontal(
											container.Top(
												container.SplitHorizontal(
													container.Top(
														container.Focused(),
														container.PlaceWidget(form.QSODate),
														container.AlignHorizontal(align.HorizontalLeft),
														container.AlignVertical(align.VerticalTop),
														container.PaddingLeft(2),
														container.PaddingRight(2),
													),
													container.Bottom(
														container.PlaceWidget(form.QSOStart),
														container.AlignHorizontal(align.HorizontalLeft),
														container.AlignVertical(align.VerticalTop),
														container.PaddingLeft(2),
														container.PaddingRight(2),
													),
												),
											),
											container.Bottom(
												container.SplitHorizontal(
													container.Top(
														container.PlaceWidget(form.QSOEnd),
														container.AlignHorizontal(align.HorizontalLeft),
														container.AlignVertical(align.VerticalTop),
														container.PaddingLeft(2),
														container.PaddingRight(2),
													),
													container.Bottom(
														container.PlaceWidget(form.QSOTXCall),
														container.AlignHorizontal(align.HorizontalLeft),
														container.AlignVertical(align.VerticalTop),
														container.PaddingLeft(2),
														container.PaddingRight(2),
													),
												),
											),
										),
									),
									container.Bottom(
										container.SplitHorizontal(
											container.Top(
												container.SplitHorizontal(
													container.Top(
														container.PlaceWidget(form.QSORXCall),
														container.AlignHorizontal(align.HorizontalLeft),
														container.AlignVertical(align.VerticalTop),
														container.PaddingLeft(2),
														container.PaddingRight(2),
													),
													container.Bottom(
														container.PlaceWidget(form.QSOBand),
														container.AlignHorizontal(align.HorizontalLeft),
														container.AlignVertical(align.VerticalTop),
														container.PaddingLeft(2),
														container.PaddingRight(2),
													),
												),
											),
											container.Bottom(
												container.SplitHorizontal(
													container.Top(
														container.PlaceWidget(form.QSOMode),
														container.AlignHorizontal(align.HorizontalLeft),
														container.AlignVertical(align.VerticalTop),
														container.PaddingLeft(2),
														container.PaddingRight(2),
													),
													container.Bottom(
														container.PlaceWidget(form.QSOPower),
														container.AlignHorizontal(align.HorizontalLeft),
														container.AlignVertical(align.VerticalTop),
														container.PaddingLeft(2),
														container.PaddingRight(2),
													),
												),
											),
										),
									),
								),
								container.PaddingTop(1),
							),
							container.Bottom(
								container.SplitHorizontal(
									container.Top(
										container.SplitVertical(
											container.Left(
												container.PlaceWidget(form.bLog),
												container.AlignVertical(align.VerticalTop),
												container.PaddingLeft(2),
												container.PaddingRight(2),
											),
											container.Right(
												container.PlaceWidget(form.bRemove),
												container.AlignVertical(align.VerticalTop),
												container.PaddingLeft(2),
												container.PaddingRight(2),
											),
										),
									),
									container.Bottom(
										container.PlaceWidget(form.bExit),
										container.AlignVertical(align.VerticalTop),
										container.PaddingBottom(1),
										container.PaddingLeft(2),
										container.PaddingRight(2),
									),
								),
							),
							container.SplitFixed(18),
						),
						container.Border(linestyle.Light),
					),
					container.Bottom(
						container.KeyFocusSkip(),
					),
					container.SplitFixed(24),
				),
			),
			container.Right(
				container.Border(linestyle.Light),
			),
			container.SplitFixed(35),
		),
	)
}

func main() {
	conf := &cfg.Config{
		LogSavePath:          "",
		TXCall:               "N0CALL",
		ButtonPrimaryColor:   "Teal",
		ButtonSecondaryColor: "Fuchsia",
	}

	t, err := tcell.New()
	if err != nil {
		log.Panicln(err)
	}
	defer t.Close()

	ctx, cancel := context.WithCancel(context.Background())

	c, err := container.New(t, container.ID("ROOT"))
	if err != nil {
		log.Panicln(err)
	}

	form := new(QSOForm)
	if err = form.prefillForm(conf, cancel); err != nil {
		log.Panicln(err)
	}

	if err = form.makeMainLayout(c); err != nil {
		log.Panicln(err)
	}

	err = termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(func(k *terminalapi.Keyboard) {
		switch k.Key {
		case 'q':
			cancel()
		}
	}))

	if err != nil {
		log.Panicln(err)
	}

}
