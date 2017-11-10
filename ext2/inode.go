package ext2

import (
	"os"
	"unsafe"
)

// InodeInfo is an in-memory version of the Inode.
// This one is tailor made for the current project.
type InodeInfo struct {
	Mode            string // For simplicity, this is for now set as string repr. of File Type.
	Size            int64
	BlocksAllocated int64
	DataBlockPtrs   []int64

	Blocks *InodeBlocks
}

// InodeBlocks helps in loading various types of Data Blocks.
type InodeBlocks struct {
	DirectBlocks        []byte  // 0..11
	IndirectBlock       []int64 // 12
	DoubleIndirectBlock []int64 // 13
	TripleIndirectBlock []int64 // 14
}

// NewIndode converts raw Inode to InodeInfo.
func NewInode(file *os.File, i *Inode, bg *BlockGroup) *InodeInfo {
	dbPtrs := []int64{}
	for _, dbp := range i.Block {
		dbPtrs = append(dbPtrs, int64(dbp))
	}

	ii := &InodeInfo{
		DataBlockPtrs:   dbPtrs,
		Size:            int64(i.Size),
		Mode:            GetMode(i.Mode),
		BlocksAllocated: int64(i.Blocks),
	}

	ii.Blocks = &InodeBlocks{
		DirectBlocks: ReadDirectDataBlocks(file, bg, ii),
	}

	return ii
}

// ReadInode returns InodeInfo, which is a more usable form of Inode.
func ReadInode(file *os.File, bg *BlockGroup, from int64) *InodeInfo {
	blockBytes := ReadNBytes(file, from, INODE_SIZE)
	i := (*Inode)(unsafe.Pointer(&blockBytes[0]))

	return NewInode(file, i, bg)
}

// ReadInodeTable reads the complete Inode Table into memory.
func ReadInodeTable(file *os.File, bg *BlockGroup, bgIndex int) []*InodeInfo {
	inodesPerBlock := bg.BLOCK_SIZE / INODE_SIZE
	inodesPerGroup := bg.SuperBlock.InodesPerGroup
	tableLen := int(inodesPerGroup / inodesPerBlock)

	inodeTable := make([]*InodeInfo, tableLen)

	from := bg.GroupDescriptors[bgIndex].InodeTablePtr * bg.BLOCK_SIZE

	for i := 0; i < tableLen; i, from = i+1, from+INODE_SIZE {
		inodeTable[i] = ReadInode(file, bg, from)
	}

	return inodeTable
}
