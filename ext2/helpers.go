package ext2

import (
	"os"
)

// ReadNBytes returns a stream of N Bytes from the file.
func ReadNBytes(file *os.File, from int64, size int) []byte {
	block := make([]byte, size)
	file.ReadAt(block, from)
	return block
}

/* GetMode parses the number and returns File Type.
socket = 0xC000
symbolic link = 0xA000
regular file = 0x8000
block device = 0x6000
directory = 0x4000
character device = 0x2000
fifo = 0x1000
*/
func GetMode(m uint16) string {
	val := m >> 12
	return GetFileMode(uint8(val))
}

func GetFileMode(m uint8) string {
	var ft string

	switch m {
	case 0xc:
		ft = "Socket"
	case 0xa:
		ft = "Symbolic Link"
	case 0x8:
		ft = "RegularFile"
	case 0x6:
		ft = "BlockDevice"
	case 0x4:
		ft = "Directory"
	default:
		ft = "Unknown"
	}
	return ft
}
