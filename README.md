# EXT2 Explorer
Explore EXT2 File System with the convenience of your web browser!

The project aims to create a program that will take an EXT2 File System Parition and print certain information from it.
Even though the program should work with a proper EXT2 partition, it is **highly recommended** to use this program only with a toy/dummy EXT2 partition. To this end it is recommended to create a parition within a file.

### Steps to create an EXT2 partition within a file.
```bash
$ # The partition will be 128MB in size & minimum unit for data will be 1024 bytes.
$ # Create an "partition"
$ dd if=/dev/zero of=linux.ex2 bs=1024 count=131072

$ # Create the filesystem
$ mke2fs linux.ex2

$ # Mount the File System to a path.
$ sudo mkdir /mnt
$ mount linux.ext2 /mnt

$ # Unmount the File System.
$ sudo umount /mnt
```

### How to use the program
The code starts a web server to which point the location of the file system and then press enter. This will redirect us to a page with details from the FS.

```bash
$ go run main.go
Adding handler functions...
Starting web server at port 8080...
```

Using browser go to http://localhost:8080, and a simple page with a form should display saying `FS Path:______________`. Enter the location of the file system in the box provided and hit enter.

This should create the sample output as show below.



### Sample Output
```

--------------------------------------------------


	EXT2 FILE SYSTEM DETAILS


--------------------------------------------------


File at: /path/to/fs/linux.ex2


--------------------------------------------------

SuperBlock:
 	Inodes Count: 32768
	Blocks Count: 131072
	Blocks Per Group: 8192
	Inodes Per Group: 2048
	Mount Count: 8
	Magic Number: 0XEF53
	Inode Size: 128


--------------------------------------------------

Block Group 0
	 Block Size: 1024


--------------------------------------------------

Files at root dir.
	Directory        ->	.
	Directory        ->	a-dir
	Directory        ->	b
	Directory        ->	lost+found
	Regular File     ->	abc.txt
	Directory        ->	a
	Directory        ->	.


--------------------------------------------------

====================================================================================================

File Name: abc.txt

--------------------------------------------------

Hello this is it!

====================================================================================================
```
