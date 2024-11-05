package utils

type AnsiControlSequence string

const (
	ArrowUp         AnsiControlSequence = "\033[A"
	ArrowDown       AnsiControlSequence = "\033[B"
	ArrowRight      AnsiControlSequence = "\033[C"
	ArrowLeft       AnsiControlSequence = "\033[D"
	CtrlArrowUp     AnsiControlSequence = "\033[1;5A"
	CtrlArrowDown   AnsiControlSequence = "\033[1;5B"
	CtrlArrowRight  AnsiControlSequence = "\033[1;5C"
	CtrlArrowLeft   AnsiControlSequence = "\033[1;5D"
	ShiftArrowUp    AnsiControlSequence = "\033[1;2A"
	ShiftArrowDown  AnsiControlSequence = "\033[1;2B"
	ShiftArrowRight AnsiControlSequence = "\033[1;2C"
	ShiftArrowLeft  AnsiControlSequence = "\033[1;2D"
	EndKey          AnsiControlSequence = "\033[F"
	KeyPad5         AnsiControlSequence = "\033[G"
	HomeKey         AnsiControlSequence = "\033[H"
	F1              AnsiControlSequence = "\033OP"
	F2              AnsiControlSequence = "\033OQ"
	F3              AnsiControlSequence = "\033OR"
	F4              AnsiControlSequence = "\033OS"
	Escape          AnsiControlSequence = "\033"
	Tab             AnsiControlSequence = "\011"
	Enter           AnsiControlSequence = "\012"
)
