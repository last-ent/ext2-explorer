package ext2

import (
	"os"
	"unsafe"
)

// GroupDescriptorInfo is an in-memory version of the Group Descriptor.
// This one is tailor made for the current project.
type GroupDescriptorInfo struct {
	BlockBitmap   int64
	InodeBitmap   int64
	InodeTablePtr int64

	FreeBlocksCount int
	FreeInodesCount int

	InodeTable InodeTable // Table loaded from file addr InodeTablePtr onwards.
}

func newGroupDescriptor(gd *GroupDescriptor) *GroupDescriptorInfo {
	return &GroupDescriptorInfo{
		BlockBitmap: int64(gd.BlockBitmap),
		InodeBitmap: int64(gd.InodeBitmap),

		FreeBlocksCount: int(gd.FreeBlocksCount),
		FreeInodesCount: int(gd.FreeInodesCount),

		InodeTablePtr: int64(gd.InodeTable),
	}
}

// ReadGroupDescriptor returns Super Block from given location in the file/disk.
func ReadGroupDescriptor(file *os.File, gdIndex int, blockSize int64) *GroupDescriptorInfo {

	offset := 2 * blockSize // Boot Block + Super Block
	from := offset + int64(gdIndex)*blockSize

	rawBytes := ReadNBytes(file, from, int(blockSize))
	gd := (*GroupDescriptor)(unsafe.Pointer(&rawBytes[0]))

	return newGroupDescriptor(gd)
}
