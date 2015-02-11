package pipeglade

import (
	"fmt"
	"image/color"
	"io"
	"time"
)

func (ui *UI) do(line string) {
	ui.outLock.Lock()
	ui.stdinReplacer.WriteString(ui.stdin, line)
	io.WriteString(ui.stdin, "\n")
	ui.outLock.Unlock()
}

// Widgets:
//  GtkLabel
//  GtkEntry
//  GtkTextView
//  GtkSpinButton
//  GtkProgressBar
func (ui *UI) SetText(widget, text string) {
	ui.do(fmt.Sprintf("%s:set_text %s", widget, text))
}

// Widgets:
//  GtkImage
func (ui *UI) SetFromIconName(widget, iconName string) {
	ui.do(fmt.Sprintf("%s:set_from_icon_name %s", widget, iconName))
}

// Widgets:
//  GtkImage
func (ui *UI) SetFromFile(widget, path string) {
	ui.do(fmt.Sprintf("%s:set_from_file %s", widget, path))
}

// Widgets:
//  GtkTextView
func (ui *UI) Delete(widget string) {
	ui.do(fmt.Sprintf("%s:delete %s", widget))
}

// Widgets:
//  GtkTextView
func (ui *UI) InsertAtCursor(widget, text string) {
	ui.do(fmt.Sprintf("%s:insert_at_cursor %s", widget, text))
}

// Widgets:
//  GtkTextView
func (ui *UI) PlaceCursor(widget string, position int) {
	ui.do(fmt.Sprintf("%s:place_cursor %d", widget, position))
}

// Widgets:
//  GtkTextView
func (ui *UI) PlaceCursorAtEnd(widget string) {
	ui.do(fmt.Sprintf("%s:place_cursor end", widget))
}

// Widgets:
//  GtkTextView
func (ui *UI) PlaceCursorAtLine(widget string, line int) {
	ui.do(fmt.Sprintf("%s:place_cursor_at_line %d", widget, line))
}

// Widgets:
//  GtkTextView
func (ui *UI) ScrollToCursor(widget string) {
	ui.do(fmt.Sprintf("%s:scroll_to_cursor", widget))
}

// Widgets:
//  GtkButton
//  GtkToggleButton
//  GtkCheckButton
func (ui *UI) SetLabel(widget, text string) {
	ui.do(fmt.Sprintf("%s:set_label %s", widget, text))
}

// Widgets:
//  GtkToggleButton
//  GtkCheckButton
//  GtkRadioButton
func (ui *UI) SetActive(widget string, active bool) {
	var activeInt int
	if active {
		activeInt = 1
	}
	ui.do(fmt.Sprintf("%s:set_active %d", widget, activeInt))
}

// Widgets:
//  GtkScale
func (ui *UI) SetValue(widget string, value float64) {
	ui.do(fmt.Sprintf("%s:set_value %f", widget, value))
}

// Widgets:
//  GtkProgressBar
func (ui *UI) SetFraction(widget string, value float64) {
	ui.do(fmt.Sprintf("%s:set_fraction %f", widget, value))
}

// Widgets:
//  GtkSpinner
func (ui *UI) Start(widget string) {
	ui.do(fmt.Sprintf("%s:start", widget))
}

// Widgets:
//  GtkSpinner
func (ui *UI) Stop(widget string) {
	ui.do(fmt.Sprintf("%s:stop", widget))
}

// Widgets:
//  GtkStatusbar
func (ui *UI) Push(widget, text string) {
	ui.do(fmt.Sprintf("%s:push %s", widget, text))
}

// Widgets:
//  GtkStatusbar
func (ui *UI) Pop(widget string) {
	ui.do(fmt.Sprintf("%s:pop", widget))
}

// Widgets:
//  GtkComboBoxText
func (ui *UI) PrependText(widget, text string) {
	ui.do(fmt.Sprintf("%s:prepend_text %s", widget, text))
}

// Widgets:
//  GtkComboBoxText
func (ui *UI) AppendText(widget, text string) {
	ui.do(fmt.Sprintf("%s:append_text %s", widget, text))
}

// Widgets:
//  GtkComboBoxText
func (ui *UI) InsertText(widget string, position int, text string) {
	ui.do(fmt.Sprintf("%s:insert_text %d %s", widget, position, text))
}

// Widgets:
//  GtkTreeView
func (ui *UI) Set(widget string, row, column int, data string) {
	ui.do(fmt.Sprintf("%s:set %d %d %s", widget, row, column, data))
}

// Widgets:
//  GtkTreeView
func (ui *UI) InsertRow(widget string, position int) {
	ui.do(fmt.Sprintf("%s:insert_row %d", widget, position))
}

// Widgets:
//  GtkTreeView
func (ui *UI) InsertRowAtEnd(widget string) {
	ui.do(fmt.Sprintf("%s:insert_row end", widget))
}

// Widgets:
//  GtkTreeView
func (ui *UI) MoveRow(widget string, origin, destination int) {
	ui.do(fmt.Sprintf("%s:move_row %d %d", widget, origin, destination))
}

// Widgets:
//  GtkTreeView
func (ui *UI) MoveRowToEnd(widget string, origin int) {
	ui.do(fmt.Sprintf("%s:move_row %d end", widget, origin))
}

// Widgets:
//  GtkTreeView
func (ui *UI) RemoveRow(widget string, position int) {
	ui.do(fmt.Sprintf("%s:remove_row %d", widget, position))
}

// Widgets:
//  GtkTreeView
func (ui *UI) Scroll(widget string, row, column int) {
	ui.do(fmt.Sprintf("%s:scroll %d %d", widget, row, column))
}

// Widgets:
//  GtkColorButton
func (ui *UI) SetColor(widget string, color color.Color) {
	r, g, b, a := color.RGBA()
	r >>= 8
	g >>= 8
	b >>= 8
	a >>= 8
	ui.do(fmt.Sprintf("%s:set_color rgba(%d,%d,%d,%d)", widget, r&0xFF, g&0xFF, b&0xFF, a&0xFF))
}

// Widgets:
//  GtkColorButton
func (ui *UI) SetColorName(widget, color string) {
	ui.do(fmt.Sprintf("%s:set_color %s", widget, color))
}

// Widgets:
//  GtkFontButton
func (ui *UI) SetFontName(widget, font string) {
	ui.do(fmt.Sprintf("%s:set_font_name %s", widget, font))
}

// Widgets:
//  GtkFileChooserButton
//  GtkFileChooserDialog
func (ui *UI) SetFileName(widget, path string) {
	ui.do(fmt.Sprintf("%s:set_filename %s", widget, path))
}

// Widgets:
//  GtkFileChooserDialog
func (ui *UI) SetCurrentName(widget, name string) {
	ui.do(fmt.Sprintf("%s:set_current_name %s", widget, name))
}

// Widgets:
//  GtkCalendar
func (ui *UI) SelectDate(widget string, date time.Time) {
	year, month, day := date.Date()
	ui.do(fmt.Sprintf("%s:select_date %d-%02d-%02d", widget, year, month, day))
}

// Widgets:
//  GtkCalendar
func (ui *UI) MarkDay(widget string, day int) {
	ui.do(fmt.Sprintf("%s:mark_day %d", widget, day))
}

// Widgets:
//  GtkCalendar
func (ui *UI) ClearMarks(widget string) {
	ui.do(fmt.Sprintf("%s:clear_marks", widget))
}

func (ui *UI) MainQuit() {
	ui.do(fmt.Sprintf("pg:main_quit"))
	ui.cmd.Wait()
}

// Widgets:
//  Any Widget
func (ui *UI) SetSensitive(widget string, sensitive bool) {
	var sensitiveInt int
	if sensitive {
		sensitiveInt = 1
	}
	ui.do(fmt.Sprintf("%s:set_sensitive %d", widget, sensitiveInt))
}

// Widgets:
//  Any Widget
func (ui *UI) SetVisible(widget string, visible bool) {
	var visibleInt int
	if visible {
		visibleInt = 1
	}
	ui.do(fmt.Sprintf("%s:set_visible %d", widget, visibleInt))
}

// Widgets:
//  Any Widget
func (ui *UI) ForceCB(widget string) {
	ui.do(fmt.Sprintf("%s:force_cb", widget))
}
