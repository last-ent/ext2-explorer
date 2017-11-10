package ext2

import (
	"os"
	"unsafe"
)

// DentriesMap is a helper type that makes it easier to search for dentry names.
type DentriesMap map[string]*DentryInfo

// Directory represents in-memory representation of a directory and all it's contents.
type Directory struct {
	Name     string
	Dentries []*DentryInfo
}

// DentryInfo represents in-memory Directory Entry.
// This one is tailor made for the current project.
type DentryInfo struct {
	FileType string
	Name     string
	RecLen   int64
	Inode    int64
}

func newDentryInfo(file *os.File, bg *BlockGroup, d *Dentry) *DentryInfo {
	return &DentryInfo{
		Inode:    int64(d.Inode),
		RecLen:   int64(d.RecLen),
		Name:     string(d.Name[:d.NameLen]),
		FileType: GetFileMode(d.FileType),
	}
}

// ReadDentry returns
func ReadDentry(file *os.File, bg *BlockGroup, from int64, size int) *DentryInfo {
	rawBytes := ReadNBytes(file, from, size)

	d := (*Dentry)(unsafe.Pointer(&rawBytes[0]))

	return newDentryInfo(file, bg, d)
}

func ReadDentries(file *os.File, bg *BlockGroup, from int64, size int, inodeSize int64) []*DentryInfo {

	entry := ReadDentry(file, bg, from, size)
	dentries := []*DentryInfo{entry}
	entriesSize := entry.RecLen

	for entriesSize < inodeSize && entry.Inode != 0 {
		rl := entry.RecLen

		from += rl
		entriesSize += rl

		entry = ReadDentry(file, bg, from, size)
		//dentries[entry.Name] = entry
		dentries = append(dentries, entry)
	}

	return dentries
}

func ReadDirectory(file *os.File, bg *BlockGroup, dirName string, bgIndex int) *Directory {
	deSize := int(unsafe.Sizeof(Dentry{}))
	dirInode := bg.InodeTable[bgIndex]

	firstEntryLoc := dirInode.DataBlockPtrs[0] * bg.BLOCK_SIZE
	dirEntries := ReadDentries(file, bg, firstEntryLoc, deSize, dirInode.Size)

	return &Directory{
		Name:     dirName,
		Dentries: dirEntries,
	}
}

func (dir *Directory) AsDentriesMap() DentriesMap {
	dm := DentriesMap{}

	for _, di := range dir.Dentries {
		dm[di.Name] = di
	}

	return dm
}
