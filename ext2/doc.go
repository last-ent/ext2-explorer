/*Package ext2 consists of the data structures required to work with EXT2 file system.

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
