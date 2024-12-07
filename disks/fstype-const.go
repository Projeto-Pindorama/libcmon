/*
 * disks/fstype-const.go - Constant table for MBR partition types.
 *
 * Copyright (C) 2024: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 * Borrowed from util-linux's 'include/pt-mbr.h' file and translated
 * from C to Go via some hip RegEx (and some shell help too).
 * From pt-mbr.h's copyright header:
 *
 * No copyright is claimed.  This code is in the public domain; do with
 * it what you wish.
 */

package disks

const (
	MBR_EMPTY_PARTITION                = byte(0x00)
	MBR_FAT12_PARTITION                = byte(0x01)
	MBR_XENIX_ROOT_PARTITION           = byte(0x02)
	MBR_XENIX_USR_PARTITION            = byte(0x03)
	MBR_FAT16_LESS32M_PARTITION        = byte(0x04)
	MBR_DOS_EXTENDED_PARTITION         = byte(0x05)
	MBR_FAT16_PARTITION                = byte(0x06) /* DOS 16-bit >=32M */
	MBR_HPFS_NTFS_PARTITION            = byte(0x07) /* OS/2 IFS e.g. HPFS or NTFS or QNX or exFAT */
	MBR_AIX_PARTITION                  = byte(0x08) /* AIX boot (AIX -- PS/2 port) or SplitDrive */
	MBR_AIX_BOOTABLE_PARTITION         = byte(0x09) /* AIX data or Coherent */
	MBR_OS2_BOOTMNGR_PARTITION         = byte(0x0a) /* OS/2 Boot Manager */
	MBR_W95_FAT32_PARTITION            = byte(0x0b)
	MBR_W95_FAT32_LBA_PARTITION        = byte(0x0c) /* LBA really is `Extended Int 13h' */
	MBR_W95_FAT16_LBA_PARTITION        = byte(0x0e)
	MBR_W95_EXTENDED_PARTITION         = byte(0x0f)
	MBR_OPUS_PARTITION                 = byte(0x10)
	MBR_HIDDEN_FAT12_PARTITION         = byte(0x11)
	MBR_COMPAQ_DIAGNOSTICS_PARTITION   = byte(0x12)
	MBR_HIDDEN_FAT16_L32M_PARTITION    = byte(0x14)
	MBR_HIDDEN_FAT16_PARTITION         = byte(0x16)
	MBR_HIDDEN_HPFS_NTFS_PARTITION     = byte(0x17)
	MBR_AST_SMARTSLEEP_PARTITION       = byte(0x18)
	MBR_HIDDEN_W95_FAT32_PARTITION     = byte(0x1b)
	MBR_HIDDEN_W95_FAT32LBA_PARTITION  = byte(0x1c)
	MBR_HIDDEN_W95_FAT16LBA_PARTITION  = byte(0x1e)
	MBR_NEC_DOS_PARTITION              = byte(0x24)
	MBR_PLAN9_PARTITION                = byte(0x39)
	MBR_PARTITIONMAGIC_PARTITION       = byte(0x3c)
	MBR_VENIX80286_PARTITION           = byte(0x40)
	MBR_PPC_PREP_BOOT_PARTITION        = byte(0x41)
	MBR_SFS_PARTITION                  = byte(0x42)
	MBR_QNX_4X_PARTITION               = byte(0x4d)
	MBR_QNX_4X_2ND_PARTITION           = byte(0x4e)
	MBR_QNX_4X_3RD_PARTITION           = byte(0x4f)
	MBR_DM_PARTITION                   = byte(0x50)
	MBR_DM6_AUX1_PARTITION             = byte(0x51) /* (or Novell) */
	MBR_CPM_PARTITION                  = byte(0x52) /* CP/M or Microport SysV/AT */
	MBR_DM6_AUX3_PARTITION             = byte(0x53)
	MBR_DM6_PARTITION                  = byte(0x54)
	MBR_EZ_DRIVE_PARTITION             = byte(0x55)
	MBR_GOLDEN_BOW_PARTITION           = byte(0x56)
	MBR_PRIAM_EDISK_PARTITION          = byte(0x5c)
	MBR_SPEEDSTOR_PARTITION            = byte(0x61)
	MBR_GNU_HURD_PARTITION             = byte(0x63) /* GNU HURD or Mach or Sys V/386 (such as ISC UNIX) */
	MBR_UNIXWARE_PARTITION             = MBR_GNU_HURD_PARTITION
	MBR_NETWARE_286_PARTITION          = byte(0x64)
	MBR_NETWARE_386_PARTITION          = byte(0x65)
	MBR_DISKSECURE_MULTIBOOT_PARTITION = byte(0x70)
	MBR_PC_IX_PARTITION                = byte(0x75)
	MBR_OLD_MINIX_PARTITION            = byte(0x80) /* Minix 1.4a and earlier */
	MBR_MINIX_PARTITION                = byte(0x81) /* Minix 1.4b and later */
	MBR_LINUX_SWAP_PARTITION           = byte(0x82)
	MBR_SOLARIS_X86_PARTITION          = MBR_LINUX_SWAP_PARTITION
	MBR_LINUX_DATA_PARTITION           = byte(0x83)
	MBR_OS2_HIDDEN_DRIVE_PARTITION     = byte(0x84) /* also hibernation MS APM Intel Rapid Start */
	MBR_INTEL_HIBERNATION_PARTITION    = MBR_OS2_HIDDEN_DRIVE_PARTITION
	MBR_LINUX_EXTENDED_PARTITION       = byte(0x85)
	MBR_NTFS_VOL_SET1_PARTITION        = byte(0x86)
	MBR_NTFS_VOL_SET2_PARTITION        = byte(0x87)
	MBR_LINUX_PLAINTEXT_PARTITION      = byte(0x88)
	MBR_LINUX_LVM_PARTITION            = byte(0x8e)
	MBR_AMOEBA_PARTITION               = byte(0x93)
	MBR_AMOEBA_BBT_PARTITION           = byte(0x94) /* (bad block table) */
	MBR_BSD_OS_PARTITION               = byte(0x9f) /* BSDI */
	MBR_THINKPAD_HIBERNATION_PARTITION = byte(0xa0)
	MBR_FREEBSD_PARTITION              = byte(0xa5) /* various BSD flavours */
	MBR_OPENBSD_PARTITION              = byte(0xa6)
	MBR_NEXTSTEP_PARTITION             = byte(0xa7)
	MBR_DARWIN_UFS_PARTITION           = byte(0xa8)
	MBR_NETBSD_PARTITION               = byte(0xa9)
	MBR_DARWIN_BOOT_PARTITION          = byte(0xab)
	MBR_HFS_HFS_PARTITION              = byte(0xaf)
	MBR_BSDI_FS_PARTITION              = byte(0xb7)
	MBR_BSDI_SWAP_PARTITION            = byte(0xb8)
	MBR_BOOTWIZARD_HIDDEN_PARTITION    = byte(0xbb)
	MBR_ACRONIS_FAT32LBA_PARTITION     = byte(0xbc) /* Acronis Secure Zone with ipl for loader F11.SYS */
	MBR_SOLARIS_BOOT_PARTITION         = byte(0xbe)
	MBR_SOLARIS_PARTITION              = byte(0xbf)
	MBR_DRDOS_FAT12_PARTITION          = byte(0xc1)
	MBR_DRDOS_FAT16_L32M_PARTITION     = byte(0xc4)
	MBR_DRDOS_FAT16_PARTITION          = byte(0xc6)
	MBR_SYRINX_PARTITION               = byte(0xc7)
	MBR_NONFS_DATA_PARTITION           = byte(0xda)
	MBR_CPM_CTOS_PARTITION             = byte(0xdb) /* CP/M or Concurrent CP/M or Concurrent DOS or CTOS */
	MBR_DELL_UTILITY_PARTITION         = byte(0xde) /* Dell PowerEdge Server utilities */
	MBR_BOOTIT_PARTITION               = byte(0xdf) /* BootIt EMBRM */
	MBR_DOS_ACCESS_PARTITION           = byte(0xe1) /* DOS access or SpeedStor 12-bit FAT extended partition */
	MBR_DOS_RO_PARTITION               = byte(0xe3) /* DOS R/O or SpeedStor */
	MBR_SPEEDSTOR_EXTENDED_PARTITION   = byte(0xe4) /* SpeedStor 16-bit FAT extended partition < 1024 cyl. */
	MBR_RUFUS_EXTRA_PARTITION          = byte(0xea) /* Rufus extra partition for alignment */
	MBR_BEOS_FS_PARTITION              = byte(0xeb)
	MBR_GPT_PARTITION                  = byte(0xee) /* Intel EFI GUID Partition Table */
	MBR_EFI_SYSTEM_PARTITION           = byte(0xef) /* Intel EFI System Partition */
	MBR_LINUX_PARISC_BOOT_PARTITION    = byte(0xf0) /* Linux/PA-RISC boot loader */
	MBR_SPEEDSTOR1_PARTITION           = byte(0xf1)
	MBR_SPEEDSTOR2_PARTITION           = byte(0xf4) /* SpeedStor large partition */
	MBR_DOS_SECONDARY_PARTITION        = byte(0xf2) /* DOS 3.3+ secondary */
	MBR_EBBR_PROTECTIVE_PARTITION      = byte(0xf8) /* Arm EBBR firmware protective partition */
	MBR_VMWARE_VMFS_PARTITION          = byte(0xfb)
	MBR_VMWARE_VMKCORE_PARTITION       = byte(0xfc) /* VMware kernel dump partition */
	MBR_LINUX_RAID_PARTITION           = byte(0xfd) /* Linux raid partition with autodetect using persistent superblock */
	MBR_LANSTEP_PARTITION              = byte(0xfe) /* SpeedStor >1024 cyl. or LANstep */
	MBR_XENIX_BBT_PARTITION            = byte(0xff) /* Xenix Bad Block Table */
)
