package main

type ISO9995 struct{}

func NewISO9995() (i ISO9995) {
	return
}

func (i ISO9995) CapsLock() uint16 {
	return 58
}

func (i ISO9995) LeftShift() uint16 {
	return 42
}

func (i ISO9995) RightShift() uint16 {
	return 54
}

func (i ISO9995) Alt() uint16 {
	return 56
}

func (i ISO9995) AltGr() uint16 {
	return 100
}

func (i ISO9995) Mappings() (m []Mapping) {
	m = make([]Mapping, 128)

	m[1] = Mapping{"<escape>", "<escape>", "<escape>", "<escape>"}
	m[59] = Mapping{"<f1>", "<f1>", "<f1>", "<f1>"}
	m[60] = Mapping{"<f2>", "<f2>", "<f2>", "<f2>"}
	m[61] = Mapping{"<f3>", "<f3>", "<f3>", "<f3>"}
	m[62] = Mapping{"<f4>", "<f4>", "<f4>", "<f4>"}
	m[63] = Mapping{"<f5>", "<f5>", "<f5>", "<f5>"}
	m[64] = Mapping{"<f6>", "<f6>", "<f6>", "<f6>"}
	m[65] = Mapping{"<f7>", "<f7>", "<f7>", "<f7>"}
	m[66] = Mapping{"<f8>", "<f8>", "<f8>", "<f8>"}
	m[67] = Mapping{"<f9>", "<f9>", "<f9>", "<f9>"}
	m[68] = Mapping{"<f10>", "<f10>", "<f10>", "<f10>"}
	m[87] = Mapping{"<f11>", "<f11>", "<f11>", "<f11>"}
	m[88] = Mapping{"<f12>", "<f12>", "<f12>", "<f12>"}
	m[99] = Mapping{"<print_screen>", "<print_screen>", "<system_request>", "<print_screen>"}
	m[110] = Mapping{"<insert>", "<insert>", "<scroll_lock>", "<insert>"}
	m[111] = Mapping{"<delete>", "<delete>", "<delete>", "<deletw>"}
	m[41] = Mapping{"<`>", "`>", "<~>", "<¦>"}
	m[2] = Mapping{"1", "1", "!", "±"}
	m[3] = Mapping{"2", "2", "@", "2"}
	m[4] = Mapping{"3", "3", "#", "3"}
	m[5] = Mapping{"4", "4", "$", "4"}
	m[6] = Mapping{"5", "5", "%", "5"}
	m[7] = Mapping{"6", "6", "^", "6"}
	m[8] = Mapping{"7", "7", "&", "7"}
	m[9] = Mapping{"8", "8", "*", "8"}
	m[10] = Mapping{"9", "9", "(", "9"}
	m[11] = Mapping{"10", "10", ")", "10"}
	m[12] = Mapping{"-", "-", "_", "–"}
	m[13] = Mapping{"=", "=", "+", "≠"}
	m[14] = Mapping{"<backspace>", "<backspace>", "<backspace>", "<backspace>"}
	m[15] = Mapping{"\t", "\t", "\t", "\t"}
	m[16] = Mapping{"q", "Q", "Q", "œ"}
	m[17] = Mapping{"w", "W", "W", "∑"}
	m[18] = Mapping{"e", "E", "E", "´"}
	m[19] = Mapping{"r", "R", "R", "®"}
	m[20] = Mapping{"t", "T", "T", "†"}
	m[21] = Mapping{"y", "Y", "Y", "\\"}
	m[22] = Mapping{"u", "U", "U", "¨"}
	m[23] = Mapping{"i", "I", "I", "ˆ"}
	m[24] = Mapping{"o", "O", "O", "ø"}
	m[25] = Mapping{"p", "P", "P", "π"}
	m[26] = Mapping{"[", "[", "{", "“"}
	m[27] = Mapping{"]", "]", "}", "‘"}
	m[28] = Mapping{"\n", "\n", "\n", "\n"}
	m[30] = Mapping{"a", "A", "A", "Å"}
	m[31] = Mapping{"s", "S", "S", "Í"}
	m[32] = Mapping{"d", "D", "D", "Î"}
	m[33] = Mapping{"f", "F", "F", "Ï"}
	m[34] = Mapping{"g", "G", "G", "©"}
	m[35] = Mapping{"h", "H", "H", "Ó"}
	m[36] = Mapping{"j", "J", "J", "Ô"}
	m[37] = Mapping{"k", "K", "K", "˚"}
	m[38] = Mapping{"l", "L", "L", "Ò"}
	m[39] = Mapping{";", ";", ":", "…"}
	m[40] = Mapping{"'", "'", "\"", "Æ"}
	m[43] = Mapping{"\\", "\\", "|", "«"}
	m[86] = Mapping{"§", "§", "±", "§"}
	m[44] = Mapping{"z", "Z", "Z", "Ω"}
	m[45] = Mapping{"x", "X", "X", "≈"}
	m[46] = Mapping{"c", "C", "C", "Ç"}
	m[47] = Mapping{"v", "V", "V", "√"}
	m[48] = Mapping{"b", "B", "B", "ı"}
	m[49] = Mapping{"n", "N", "N", "˜"}
	m[50] = Mapping{"m", "M", "M", "Â"}
	m[51] = Mapping{",", ",", "<", "≤"}
	m[52] = Mapping{".", ".", ">", "≥"}
	m[53] = Mapping{"/", "/", "?", "÷"}
	m[29] = Mapping{"<control>", "<control>", "<control>", "<control>"}
	m[125] = Mapping{"<windows>", "<windows>", "<windows>", "<windows>"}
	m[57] = Mapping{" ", " ", " ", " "}
	m[127] = Mapping{"<menu>", "<menu>", "<menu>", "<menu>"}
	m[97] = Mapping{"<control>", "<control>", "<control>", "<control>"}

	return
}
