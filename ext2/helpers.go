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

/*GetMode parses the number and returns File Type.
socket = 0xC000
symbolic link = 0xA000
regular file = 0x8000
block device = 0x6000
directory = 0x4000
character device = 0x2000
fifo = 0x1000
*/
func GetMode(m uint16) string {
	var ft string

	switch m >> 12 {
	case 0xc:
		ft = "Socket"
	case 0xa:
		ft = "Symbolic Link"
	case 0x8:
		ft = "RegularFile"
	case 0x6:
		ft = "Block Device"
	case 0x4:
		ft = "Directory"
	default:
		ft = "Unknown"
	}
	return ft
}

// GetFileMode is used by Dentry which has a different set of codes for File Types.
func GetFileMode(m uint8) string {
	var ft string

	switch m {
	case 1:
		ft = "Regular File"
	case 2:
		ft = "Directory"
	case 3:
		ft = "Character Device"
	case 4:
		ft = "Block Device"
	case 5:
		ft = "Named Pipe"
	case 6:
		ft = "Socket"
	case 7:
		ft = "Symbolic Link"
	default:
		ft = "Unknown"
	}

	return ft
}

// MaxPadLen is useful for string representation of FSReader.
const MaxPadLen int = len("Character Device")
