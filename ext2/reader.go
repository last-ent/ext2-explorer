package ext2

import (
	"bytes"
	"fmt"
	"os"
)

// FSReader tries to act like a file system handler.
type FSReader struct {
	Path        string // Should not be changed once set.
	BlockSize   int64  // Should not be changed once set.
	File        *os.File
	SuperBlock  *SuperBlockInfo
	BlockGroups []*BlockGroup
	RootDir     *Directory
}

// NewFSReader returns an FSReader instance.
func NewFSReader(path string, blockSize int64) *FSReader {
	file, err := os.Open(path)
	if err != nil {
		file.Close()
		panic(err)
	}

	sb := ReadSuperBlock(file, blockSize, int(blockSize))
	fs := &FSReader{
		File:      file,
		Path:      path,
		BlockSize: blockSize,

		SuperBlock: sb,

		// For now we read a single BlockGroup.
		BlockGroups: []*BlockGroup{ReadBlockGroup(file, sb, blockSize, 0)},
	}

	fs.RootDir = ReadDirectory(fs.File, fs.BlockGroups[0], "/", ROOT_DIR_INDEX)
	return fs
}

// Close should be once all work with the FSReader is completed.
func (fs *FSReader) Close() error {
	return fs.File.Close()
}

// ReadFile returns a stream of bytes that consists of all the data for the given file.
// For now, only Root dir & reading only direct data blocks is supported.
func (fs *FSReader) ReadFile(dir string, fName string) ([]byte, error) {
	if dir != "/" {
		return nil, fmt.Errorf("for now only Root dir is supported")
	}

	file, exists := fs.RootDir.AsDentriesMap()[fName]
	if !exists {
		return nil, fmt.Errorf("file does not exists")
	}

	return ReadFile(fs.File, fs.BlockGroups[0], file), nil
}

// BufferedString allows us to combine the string representation efficiently.
func (fs *FSReader) BufferedString() bytes.Buffer {
	var buff bytes.Buffer

	buff.WriteString("\n\n\n--------------------------------------------------\n\n\n")
	buff.WriteString("\tEXT2 FILE SYSTEM DETAILS")

	buff.WriteString("\n\n\n--------------------------------------------------\n\n\n")
	buff.WriteString(fmt.Sprintf("File at: %s\n", fs.Path))
	buff.WriteString("\n\n--------------------------------------------------")
	buff.WriteString(fmt.Sprintf("\n\nSuperBlock:\n "))

	sb := fs.SuperBlock
	buff.WriteString(fmt.Sprintf("\tInodes Count: %d\n", sb.InodesCount))
	buff.WriteString(fmt.Sprintf("\tBlocks Count: %d\n", sb.BlocksCount))
	buff.WriteString(fmt.Sprintf("\tBlocks Per Group: %d\n", sb.BlocksPerGroup))
	buff.WriteString(fmt.Sprintf("\tInodes Per Group: %d\n", sb.InodesPerGroup))
	buff.WriteString(fmt.Sprintf("\tMount Count: %d\n", sb.MountCount))
	buff.WriteString(fmt.Sprintf("\tMagic Number: %#X\n", sb.MagicNumber))
	buff.WriteString(fmt.Sprintf("\tInode Size: %d\n", sb.InodeSize))
	buff.WriteString("\n\n--------------------------------------------------")

	bg0 := fs.BlockGroups[0]
	buff.WriteString("\n\nBlock Group 0\n")
	buff.WriteString(fmt.Sprintf("\t Block Size: %d\n", bg0.BLOCK_SIZE))
	buff.WriteString("\n\n--------------------------------------------------")

	buff.WriteString("\n\nFiles at root dir.\n")
	for _, de := range fs.RootDir.Dentries {
		buff.WriteString(fmt.Sprintf("\t%s ->\t%s\n", padRight(de.FileType, MaxPadLen), de.Name))
	}
	buff.WriteString("\n\n--------------------------------------------------")
	return buff
}

func padRight(s string, l int) string {
	for i := len(s); i < l; i++ {
		s += " "
	}
	return s
}

// Details is a map of values for various FS entities.
type Details map[string]string

// ReprMap returns a nested map version of Buffered String.
// FS Block -> Block Entity -> Entity Value
func (fs *FSReader) ReprMap() map[string]Details {
	sb := fs.SuperBlock
	bg0 := fs.BlockGroups[0]

	repr := map[string]Details{
		"File at": Details{"Location": fs.Path},

		"Block Group 0": Details{
			"Block Size": fmt.Sprintf("%d", bg0.BLOCK_SIZE),
		},

		"Super Block": Details{
			"Inodes Count":     fmt.Sprintf("%d", sb.InodesCount),
			"Blocks Count":     fmt.Sprintf("%d", sb.BlocksCount),
			"Blocks Per Group": fmt.Sprintf("%d", sb.BlocksPerGroup),
			"Inodes Per Group": fmt.Sprintf("%d", sb.InodesPerGroup),
			"Magic Number":     fmt.Sprintf("%x", sb.MagicNumber),
			"Inode Size":       fmt.Sprintf("%d", sb.InodeSize),
		},
	}

	files := Details{}
	for _, de := range fs.RootDir.Dentries {
		files[de.Name] = de.FileType
	}
	repr["Files at root dir"] = files

	return repr
}

func (fs *FSReader) String() string {
	b := fs.BufferedString()
	return b.String()
}
