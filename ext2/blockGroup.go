package ext2

import (
	"os"
)

// ReadBlockGroup returns Block Group at index: `gdIndex`.
// However, the Super Block & Group Descriptors are only read for Block Group 0.
// This is because SB & GD are reliably maintained on GB0 while other BG might not have latest or most up-to-date info.
func ReadBlockGroup(file *os.File, sb *SuperBlockInfo, blockSize int64, gdIndex int) *BlockGroup {
	bg := &BlockGroup{
		SuperBlock: sb,
		BLOCK_SIZE: blockSize,
	}

	// Super Block & Group Descriptors are available across all the Block Groups but only BG0's SB & GD are relevant.
	if gdIndex == 0 {
		gds := []*GroupDescriptorInfo{}
		for i := 0; i < sb.GroupDescriptorsCount; i++ {
			gd := ReadGroupDescriptor(file, i, blockSize)

			gds = append(gds, gd)
		}

		bg.GroupDescriptors = gds
	}

	bg.InodeTable = ReadInodeTable(file, bg, gdIndex)

	return bg
}
