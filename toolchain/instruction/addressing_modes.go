package instruction

const (
	AM_IMM byte = iota
	AM_DIR
	AM_IND
	AM_INC
)

var AddressingStringMap = map[byte]string{
	AM_IMM: "imm",
	AM_DIR: "dir",
	AM_IND: "ind",
	AM_INC: "inc",
}

var StringAddressingMap = map[string]byte{
	"imm": AM_IMM,
	"dir": AM_DIR,
	"ind": AM_IND,
	"inc": AM_INC,
}
