package functions

type Encoding int

const (
	Binary Encoding = iota
	Decimal
)

// 'encoding_type' has binary and decimal
func InitPopulation(size int, length int, encoding_type Encoding) {
	if encoding_type == Binary {
		return
	}

	var pop [size][length]bool
	for i := range size {
		pop[][]
	}
}
