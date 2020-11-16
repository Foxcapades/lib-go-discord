package js

var (
	OpenObject    byte = '{'
	CloseObject   byte = '}'
	OpenArray     byte = '['
	CloseArray    byte = ']'
	SingleQuote   byte = '"'
	FieldDivider  byte = ','
	PairSeparator byte = ':'

	NullBytesBuf     = []byte("null")
	SingleQuoteBuf   = []byte{SingleQuote}
	OpenObjectBuf    = []byte{OpenObject}
	CloseObjectBuf   = []byte{CloseObject}
	FieldDividerBuf  = []byte{FieldDivider}
	PairSeparatorBuf = []byte{PairSeparator}
)
