package utils

import "fmt"

var ANSI_MAP = map[string]string{
	"SET_TITLE":          "\033]0;",
	"HOME":               "\033[H",
	"CLEAR_TO_END":       "\033[0J",
	"CLEAR_TO_BEGINNING": "\033[1J",
	"CLEAR":              "\033[2J",
	"SHOW_CURSOR":        "\033[?25h",
	"HIDE_CURSOR":        "\033[?25l",
	"MOVE_CURSOR":        "\033[%d;%dH",
	"MOD_BOLD":           "\033[1m",
	"MOD_ITALIC":         "\033[3m",
	"MOD_UNDERLINE":      "\033[4m",
	"MOD_SLOW_BLINK":     "\033[5m",
	"MOD_RAPID_BLINK":    "\033[6m",
	"MOD_STRIKETHROUGH":  "\033[9m",
	"MOD_END":            "\033[0m",
	"SCREEN_MODE_SET":    "\033[=19h",
	"SCREEN_MODE_UNSET":  "\033[=19l",
	"SCREEN_RESTORE":     "\033[?47l",
	"SCREEN_SAVE":        "\033[?47h",
	"XTERM_ARROW_UP":     "\033[A",
	"XTERM_ARROW_DOWN":   "\033[B",
	"XTERM_ARROW_RIGHT":  "\033[C",
	"XTERM_ARROW_LEFT":   "\033[D",
}

func MoveCursorTo(x int, y int) string {
	return fmt.Sprintf(ANSI_MAP["MOVE_CURSOR"], x, y)
}

func MakeBold(content string) string {
	return ANSI_MAP["MOD_BOLD"] + content + ANSI_MAP["MOD_END"]
}

func MakeItalic(content string) string {
	return ANSI_MAP["MOD_ITALIC"] + content + ANSI_MAP["MOD_END"]
}

func MakeUndeline(content string) string {
	return ANSI_MAP["MOD_UNDERLINE"] + content + ANSI_MAP["MOD_END"]
}

func MakeStrikethrough(content string) string {
	return ANSI_MAP["MOD_STRIKETHROUGH"] + content + ANSI_MAP["MOD_END"]
}

func SetWindowTitle(content string) string {
	return ANSI_MAP["SET_TITLE"] + content
}
