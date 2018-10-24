package protocol

type MessageType uint8

const (
	FullBuffer MessageType = 0x01
	KeyPress   MessageType = 0x10
)

type KeyCode uint8

const (
	NoKey          KeyCode = 0
	Break          KeyCode = 3
	Backspace      KeyCode = 8
	Tab            KeyCode = 9
	Clear          KeyCode = 12
	Enter          KeyCode = 13
	Shift          KeyCode = 16
	Ctrl           KeyCode = 17
	Alt            KeyCode = 18
	Pause          KeyCode = 19
	CapsLock       KeyCode = 20
	Hangul         KeyCode = 21
	Hanja          KeyCode = 25
	Escape         KeyCode = 27
	Conversion     KeyCode = 28
	NonConversion  KeyCode = 29
	SpaceBar       KeyCode = 32
	PageUp         KeyCode = 33
	PageDown       KeyCode = 34
	End            KeyCode = 35
	Home           KeyCode = 36
	LeftArrow      KeyCode = 37
	UpArrow        KeyCode = 38
	RightArrow     KeyCode = 39
	DownArrow      KeyCode = 40
	Select         KeyCode = 41
	Print          KeyCode = 42
	Execute        KeyCode = 43
	PrintScreen    KeyCode = 44
	Insert         KeyCode = 45
	Delete         KeyCode = 46
	Help           KeyCode = 47
	Key0           KeyCode = 48
	Key1           KeyCode = 49
	Key2           KeyCode = 50
	Key3           KeyCode = 51
	Key4           KeyCode = 52
	Key5           KeyCode = 53
	Key6           KeyCode = 54
	Key7           KeyCode = 55
	Key8           KeyCode = 56
	Key9           KeyCode = 57
	Colon          KeyCode = 58
	Equals         KeyCode = 59
	LessThan       KeyCode = 60
	Beta           KeyCode = 63
	KeyA           KeyCode = 65
	KeyB           KeyCode = 66
	KeyC           KeyCode = 67
	KeyD           KeyCode = 68
	KeyE           KeyCode = 69
	KeyF           KeyCode = 70
	KeyG           KeyCode = 71
	KeyH           KeyCode = 72
	KeyI           KeyCode = 73
	KeyJ           KeyCode = 74
	KeyK           KeyCode = 75
	KeyL           KeyCode = 76
	KeyM           KeyCode = 77
	KeyN           KeyCode = 78
	KeyO           KeyCode = 79
	KeyP           KeyCode = 80
	KeyQ           KeyCode = 81
	KeyR           KeyCode = 82
	KeyS           KeyCode = 83
	KeyT           KeyCode = 84
	KeyU           KeyCode = 85
	KeyV           KeyCode = 86
	KeyW           KeyCode = 87
	KeyX           KeyCode = 88
	KeyY           KeyCode = 89
	KeyZ           KeyCode = 90
	WindowLeft     KeyCode = 91
	WindowRight    KeyCode = 92
	WindowMenu     KeyCode = 93
	Sleep          KeyCode = 95
	NumPad0        KeyCode = 96
	NumPad1        KeyCode = 97
	NumPad2        KeyCode = 98
	NumPad3        KeyCode = 99
	NumPad4        KeyCode = 100
	NumPad5        KeyCode = 101
	NumPad6        KeyCode = 102
	NumPad7        KeyCode = 103
	NumPad8        KeyCode = 104
	NumPad9        KeyCode = 105
	NumPadAsterisk KeyCode = 106
	NumPadPlus     KeyCode = 107
	NumPadPeriod   KeyCode = 108
	NumPadMinus    KeyCode = 109
	NumPadDecimal  KeyCode = 110
	NumPadSlash    KeyCode = 111
	F1             KeyCode = 112
	F2             KeyCode = 113
	F3             KeyCode = 114
	F4             KeyCode = 115
	F5             KeyCode = 116
	F6             KeyCode = 117
	F7             KeyCode = 118
	F8             KeyCode = 119
	F9             KeyCode = 120
	F10            KeyCode = 121
	F11            KeyCode = 122
	F12            KeyCode = 123
	F13            KeyCode = 124
	F14            KeyCode = 125
	F15            KeyCode = 126
	F16            KeyCode = 127
	F17            KeyCode = 128
	F18            KeyCode = 129
	F19            KeyCode = 130
	F20            KeyCode = 131
	F21            KeyCode = 132
	F22            KeyCode = 133
	F23            KeyCode = 134
	F24            KeyCode = 135
	NumLock        KeyCode = 144
	ScrollLock     KeyCode = 145
)

func ParseMessage(msg []byte) interface{} {
	switch MessageType(msg[0]) {
	case KeyPress:
		if len(msg) == 2 {
			return KeyCode(msg[1])
		}
	}
	return nil
}
