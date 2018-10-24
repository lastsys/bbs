package screen

import (
	"github.com/lastsys/bbs/internal/protocol"
	"sync"
)

const (
	Width  = 40
	Height = 25
)

type Position struct {
	row uint8
	col uint8
}

type Color uint8

const (
	Black Color = iota
	White
	Red
	Cyan
	Purple
	Green
	Blue
	Yellow
	Orange
	Brown
	LightRed
	DarkGray
	Gray
	LightGreen
	LightBlue
	LightGray
)

var charIndexMap = map[string]uint8{
	"@": 0,
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
	"[": 27,
	"£": 28,
	"]": 29,
	" ": 32,
	"!": 33,
	`"`: 34,
	"#": 35,
	"$": 36,
	"%": 37,
	"&": 38,
	"'": 39,
	"(": 40,
	")": 41,
	"*": 42,
	"+": 43,
	",": 44,
	"-": 45,
	".": 46,
	"/": 47,
	"0": 48,
	"1": 49,
	"2": 50,
	"3": 51,
	"4": 52,
	"5": 53,
	"6": 54,
	"7": 55,
	"8": 56,
	"9": 57,
	":": 58,
	";": 59,
	"<": 60,
	"=": 61,
	">": 62,
	"?": 63,
	"a": 129,
	"b": 130,
	"c": 131,
	"d": 132,
	"e": 133,
	"f": 134,
	"g": 135,
	"h": 136,
	"i": 137,
	"j": 138,
	"k": 139,
	"l": 140,
	"m": 141,
	"n": 142,
	"o": 143,
	"p": 144,
	"q": 145,
	"r": 146,
	"s": 147,
	"t": 148,
	"u": 149,
	"v": 150,
	"w": 151,
	"x": 152,
	"y": 153,
	"z": 154,
}

func CharIndex(r rune) uint8 {
	code, ok := charIndexMap[string(r)]
	if !ok {
		return charIndexMap[" "]
	}
	return code
}

type Character struct {
	Char       uint8
	Foreground Color
	Background Color
	Reverse    bool
}

type Buffer struct {
	chars [Height][Width]Character
	mutex sync.Mutex
}

func (b *Buffer) Clear() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	for row := 0; row < Height; row++ {
		for col := 0; col < Width; col++ {
			b.chars[row][col].Char = 32
			b.chars[row][col].Foreground = White
			b.chars[row][col].Background = Black
			b.chars[row][col].Reverse = false
		}
	}
}

func (b *Buffer) SerializeMessage() []byte {
	// 1 byte for message type.
	// The rest for the buffer data.
	b.mutex.Lock()
	defer b.mutex.Unlock()

	bytes := make([]byte, 1+4*Width*Height)
	bytes[0] = byte(protocol.FullBuffer)
	ptr := 1
	for row := 0; row < Height; row++ {
		for col := 0; col < Width; col++ {
			bytes[ptr] = b.chars[row][col].Char
			bytes[ptr+1] = byte(b.chars[row][col].Foreground)
			bytes[ptr+2] = byte(b.chars[row][col].Background)
			if b.chars[row][col].Reverse {
				bytes[ptr+3] = 1
			} else {
				bytes[ptr+3] = 0
			}
			ptr += 4
		}
	}
	return bytes
}

func (b *Buffer) Write(char Character, row uint8, col uint8) {
	b.chars[row][col] = char
}

func (b *Buffer) Print(text string, row uint8, col uint8, fg Color, bg Color) {
	var y = row
	var x = col

	for _, c := range text {
		b.chars[y][x] = Character{CharIndex(c), fg, bg, false}
		x++
		// Wrapping rules.
		if x >= Width {
			x = 0
			y++
		}
		if y >= Height {
			y = 0
		}
	}
}
