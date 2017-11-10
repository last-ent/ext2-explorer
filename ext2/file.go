package ext2

import (
	"os"
)

// ReadDirectDataBlocks reads the first 12 Data Blocks as raw bytes.
func ReadDirectDataBlocks(file *os.File, bg *BlockGroup, inode *InodeInfo) []byte {
	inodeBlocks := []byte{}

	for i := 0; i < INODE_DIRECT_BLOCKS_INDEX; i++ {
		from := inode.DataBlockPtrs[i] * bg.BLOCK_SIZE
		if from == 0 {
			break
		}

		rawBytes := ReadNBytes(file, from, int(bg.BLOCK_SIZE))
		inodeBlocks = append(inodeBlocks, rawBytes...)
	}
	return inodeBlocks
}

// ReadFile returns string representation
func ReadFile(file *os.File, bg *BlockGroup, fileEntry *DentryInfo) []byte {
	fileInode := bg.InodeTable[fileEntry.Inode-1]

	rawBytes := ReadDirectDataBlocks(file, bg, fileInode)
	return rawBytes
}
