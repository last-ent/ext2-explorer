package ext2

import (
	"os"
	"unsafe"
)

// SuperBlockInfo is an in-memory representation of the complete Super Block.
// The version used by Linux is different to this one.
// This one is tailor made for the current project.
type SuperBlockInfo struct {
	InodesCount     int64
	BlocksCount     int64
	FreeBlocksCount int64
	FreeInodesCount int64
	BlocksPerGroup  int64
	InodesPerGroup  int64
	MountCount      int
	MaxMountCount   int

	MagicNumber      int // Magic Number helps detect which FS format is being used.
	State            int // File system state
	InodeSize        int
	BlockGroupNumber int
	VolumeName       string
	LastMountedDir   string

	GroupDescriptorsCount int
}

func newSuperBlock(sb *SuperBlock) *SuperBlockInfo {
	sbi := &SuperBlockInfo{
		InodesCount:     int64(sb.InodesCount),
		BlocksCount:     int64(sb.BlocksCount),
		FreeBlocksCount: int64(sb.FreeBlocksCount),
		BlocksPerGroup:  int64(sb.BlocksPerGroup),
		InodesPerGroup:  int64(sb.InodesPerGroup),
		MountCount:      int(sb.MntCount),
		MaxMountCount:   int(sb.MaxMntCount),

		MagicNumber: int(sb.Magic),

		State:            int(sb.State),
		InodeSize:        int(sb.InodeSize),
		BlockGroupNumber: int(sb.BlockGroupNr),

		// FIXME: Convert complete byte stream into a string.
		VolumeName:     string(sb.VolumeName[0]),
		LastMountedDir: string(sb.LastMounted[0]),
	}

	gdCount := 1 + (sbi.BlocksCount-1)/sbi.BlocksPerGroup
	sbi.GroupDescriptorsCount = int(gdCount)

	return sbi
}

// ReadSuperBlock returns Super Block from given location in the file/disk.
func ReadSuperBlock(file *os.File, from int64, blockSize int) *SuperBlockInfo {
	blockBytes := ReadNBytes(file, from, blockSize)
	rawSB := (*SuperBlock)(unsafe.Pointer(&blockBytes[0]))

	return newSuperBlock(rawSB)
}
