package ext2

import (
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
