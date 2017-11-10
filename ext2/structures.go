/* Package ext2 consists of the data structures required to work with EXT2 file system.

An EXT2 File System can pictorially be seen as follows:


 Complete FS
=============

| Boot Block | Block Group 0 | Block Group 1 | Block Group 2 | ... | Block Group n |

* Boot Block is ignored when reading File Systems, it is used by Bootloader.


 A Block Group
===============

| Super Block (1 block) | Group Descriptors (n blocks) | Data Block Bitmap (1 block) | I-Node Bitmap (1 block) | I-Node Table (n blocks) | Data Blocks (n blocks) |

*/
package ext2

const (
	// DEFAULT_BLOCK_SIZE will be used be used as Block Size unless alternative value is specified.
	DEFAULT_BLOCK_SIZE = 1024

	// EXT2_N_BLOCKS denotes the number of Data Blocks used by EXT2. It defaults to 15.
	EXT2_N_BLOCKS = 15

	// INODE_SIZE is always constant. 128 Byts
	INODE_SIZE = 128

	// These Block Range/indices are 15 pointers to different types of
	// Data Blocks in an Inode's Block array.
	INODE_DIRECT_BLOCKS_INDEX         = 12
	INODE_INDIRECT_BLOCK_INDEX        = 13
	INODE_DOUBLE_INDIRECT_BLOCK_INDEX = 14
	INODE_TRIPLE_INDIRECT_BLOCK_INDEX = 15

	// EXTFS_MAGIC_NUMBER is used to determine the type of FS while mounting.
	EXTFS_MAGIC_NUMBER = 0xEF53

	// ROOT_DIR_INDEX is the root dir's index in Inode table.
	ROOT_DIR_INDEX = 1
)

// BlockGroup is a set of blocks that is used to contain information & data related to files, directories and other FS entities.
type BlockGroup struct {
	BLOCK_SIZE int64 // This should be treated as a constant value.

	SuperBlock       *SuperBlockInfo
	GroupDescriptors []*GroupDescriptorInfo

	InodeTable []*InodeInfo
}

// SuperBlock is a block that consists of all metadata regarding the complete file system.
// This is struct which corresponds to how data is stored on the disk.
type SuperBlock struct {
	InodesCount          uint32
	BlocksCount          uint32
	RBlocksCount         uint32
	FreeBlocksCount      uint32
	FreeInodesCount      uint32
	FirstDataBlock       uint32
	LogBlockSize         uint32
	LogClusterSize       uint32
	BlocksPerGroup       uint32
	ClustersPerGroup     uint32
	InodesPerGroup       uint32
	Mtime                uint32
	Wtime                uint32
	MntCount             uint16
	MaxMntCount          uint16
	Magic                uint16
	State                uint16
	Errors               uint16
	MinorRevLevel        uint16
	Lastcheck            uint32
	Checkinterval        uint32
	CreatorOs            uint32
	RevLevel             uint32
	DefResUID            uint16
	DefResGID            uint16
	FirstIno             uint32
	InodeSize            uint16
	BlockGroupNr         uint16
	FeatureCompat        uint32
	FeatureIncompat      uint32
	FeatureROCompat      uint32
	UUID                 [16]byte
	VolumeName           [16]byte
	LastMounted          [64]byte
	AlgorithmUsageBitmap uint32
	PreallocBlocks       uint8
	PreallocDirBlocks    uint8
	ReservedGdtBlocks    uint16
	JournalUUID          [16]byte
	JournalInum          uint32
	JournalDev           uint32
	LastOrphan           uint32
	HashSeed             [4]uint32
	DefHashVersion       byte
	JnlBackupType        byte
	DefaultMountOpts     uint32
	FirstMetaBg          uint32
	MkfsTime             uint32
	JnlBlocks            [17]uint32
	BlocksCountHi        uint32
	RBlocksCountHi       uint32
	FreeBlocksCountHi    uint32
	MinExtraIsize        uint16
	WantExtraIsize       uint16
	Flags                uint32
	RaidStride           uint16
	MmpInterval          uint16
	MmpBlock             uint64
	RaidStripeWidth      uint32
	LogGroupsPerFlex     byte
	ChecksumType         byte
	ReservedPad          uint16
	KbytesWritten        uint64
	SnapshotInum         uint32
	SnapshotId           uint32
	SnapshotRBlocksCount uint64
	SnapshotList         uint32
	ErrorCount           uint32
	FirstErrorTime       uint32
	FirstErrorIno        uint32
	FirstErrorBlock      uint64
	FirstErrorFunc       [32]byte
	FirstErrorLine       uint32
	LastErrorTime        uint32
	LastErrorIno         uint32
	LastErrorLine        uint32
	LastErrorBlock       uint64
	LastErrorFunc        [32]byte
	MountOpts            [64]byte
	UsrQuotaInum         uint32
	GrpQuotaInum         uint32
	OverheadBlocks       uint32
	BackupBgs            [2]uint32
	Reserved             [106]uint32
	Checksum             uint32
}

// GroupDescriptor contains the metadata of a given Block Group.
// This is struct which corresponds to how data is stored on the disk.
type GroupDescriptor struct {
	BlockBitmap     uint32
	InodeBitmap     uint32
	InodeTable      uint32 // This is a pointer to actual table.
	FreeBlocksCount uint16
	FreeInodesCount uint16
	UsedDirsCount   uint16

	Pad      uint16
	Reserved [3]uint32
}

// InodeTable is just an array of Inodes.
// But since they work on index numbers,
// (within Inode Table known as Inode Numbers).
// We can safely think of them as a table with integer keys.
type InodeTable []*InodeInfo

// Inode help the File System & OS to interpret actual data related to files, directories etc.
// Inode has a constant size: 128 bytes.
// This is struct which corresponds to how data is stored on the disk.
type Inode struct {
	Mode       uint16
	UID        uint16
	Size       uint32
	Atime      uint32
	Ctime      uint32
	Mtime      uint32
	Dtime      uint32
	GID        uint16
	LinksCount uint16
	Blocks     uint32 // Blocks # of data blocks allocated for the file. (In units of 512B)
	Flags      uint32
	OSD1       uint32
	Block      [15]uint32 // Pointers to Data Blocks for the Inode.
	Generation uint32
	FileACL    uint32
	DirACL     uint32
	FAddr      uint32
	OSD2       [96]byte
}

// Dentry represents a Directoy Entry.
// This is File System's on disk representation.
type Dentry struct {
	Inode    uint32
	RecLen   uint16
	NameLen  uint8
	FileType uint8
	Name     [255]byte
}
