package pipeglade

import (
	"bufio"
	"fmt"
	"image/color"
	"net/textproto"
	"strconv"
	"strings"
	"time"
)

const eventQuit = "go-pipeglade-quit"

type Event struct {
	UI     *UI
	Type   string
	Widget string
	Data   string
}

func (e Event) IsQuit() bool {
	return e.Data == eventQuit
}

func (e Event) ToBool() bool {
	val, _ := strconv.ParseBool(e.Data)
	return val
}

func (e Event) ToFloat() float32 {
	val, _ := strconv.ParseFloat(e.Data, 32)
	return float32(val)
}

func (e Event) ToColor() color.Color {
	var c color.RGBA
	var a float32
	_, err := fmt.Sscanf(e.Data, "rgba(%d,%d,%d,%f)", &c.R, &c.G, &c.B, &a)
	if err != nil {
		_, err = fmt.Sscanf(e.Data, "rgb(%d,%d,%d)", &c.R, &c.G, &c.B)
		if err != nil {
			return nil
		}
	}
	c.A = uint8(a * 255)
	return c
}

func (e Event) ToTime(loc *time.Location) (t time.Time) {
	var y, m, d int
	_, err := fmt.Sscanf(e.Data, "%d-%d-%d", &y, &m, &d)
	if err == nil {
		t = time.Date(y, time.Month(m), d, 0, 0, 0, 0, loc)
	}
	return
}

func (e Event) ToTableCell() (row, column int, cell string) {
	values := strings.SplitN(e.Data, " ", 3)
	if len(values) == 3 {
		row, _ = strconv.Atoi(values[0])
		column, _ = strconv.Atoi(values[1])
		cell = values[2]
	}
	return
}

func (ui *UI) readRoutine() {
	br := bufio.NewReader(ui.stdout)
	r := textproto.NewReader(br)

	for {
		line, err := r.ReadLine()
		if err != nil {
			break
		}
		first := strings.SplitN(line, " ", 2)
		if len(first) != 2 {
			break
		}
		tag := strings.Split(first[0], ":")
		if len(tag) != 2 {
			break
		}
		ui.C <- Event{
			UI:     ui,
			Widget: tag[0],
			Type:   tag[1],
			Data:   ui.stdoutReplacer.Replace(first[1]),
		}
	}
	ui.MainQuit()
	ui.C <- Event{
		UI:   ui,
		Type: eventQuit,
		Data: ui.cmd.ProcessState.String(),
	}
	close(ui.C)
}
