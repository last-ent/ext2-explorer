package ext2

import (
	"fmt"
	"os"
)

// FSReader tries to act like a file system handler.
type FSReader struct {
	PATH        string
	BLOCK_SIZE  int64
	FILE        *os.File
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
		FILE:       file,
		BLOCK_SIZE: blockSize,

		SuperBlock: sb,

		// For now we read a single BlockGroup.
		BlockGroups: []*BlockGroup{ReadBlockGroup(file, sb, blockSize, 0)},
	}

	fs.RootDir = ReadDirectory(fs.FILE, fs.BlockGroups[0], "/", ROOT_DIR_INDEX)
	return fs
}

func (fs *FSReader) Close() error {
	return fs.FILE.Close()
}

func (fs *FSReader) ReadFile(dir string, fName string) ([]byte, error) {
	if dir != "/" {
		return nil, fmt.Errorf("For now only Root dir is supported.")
	}

	file, exists := fs.RootDir.AsDentriesMap()[fName]
	if !exists {
		return nil, fmt.Errorf("File does not exists.")
	}

	return ReadFile(fs.FILE, fs.BlockGroups[0], file), nil
}
