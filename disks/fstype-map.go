/*
 * disks/fstype-map.go - Map for associating MBR partition types with
 * its human-readable names.
 *
 * Copyright (C) 2024: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 * Borrowed from util-linux's 'include/pt-mbr-partnames.h' file and,
 * just like the constant table counterpart, translated from C to Go.
 * From pt-mbr-partnames.h's copyright header:
 *
 * No copyright is claimed.  This code is in the public domain; do with
 * it what you wish.
 */

package disks

var MBRPartNames = map[byte]string{
	MBR_EMPTY_PARTITION:                "Empty",
	MBR_FAT12_PARTITION:                "FAT12",
	MBR_XENIX_ROOT_PARTITION:           "XENIX root",
	MBR_XENIX_USR_PARTITION:            "XENIX usr",
	MBR_FAT16_LESS32M_PARTITION:        "FAT16 <32M",
	MBR_DOS_EXTENDED_PARTITION:         "Extended",          /* DOS 3.3+ extended partition */
	MBR_FAT16_PARTITION:                "FAT16",             /* DOS 16-bit >=32M */
	MBR_HPFS_NTFS_PARTITION:            "HPFS/NTFS/exFAT",   /* OS/2 IFS, e.g., HPFS or NTFS or QNX or exFAT */
	MBR_AIX_PARTITION:                  "AIX",               /* AIX boot (AIX -- PS/2 port) or SplitDrive */
	MBR_AIX_BOOTABLE_PARTITION:         "AIX bootable",      /* AIX data or Coherent */
	MBR_OS2_BOOTMNGR_PARTITION:         "OS/2 Boot Manager", /* OS/2 Boot Manager */
	MBR_W95_FAT32_PARTITION:            "W95 FAT32",
	MBR_W95_FAT32_LBA_PARTITION:        "W95 FAT32 (LBA)", /* LBA really is `Extended Int 13h' */
	MBR_W95_FAT16_LBA_PARTITION:        "W95 FAT16 (LBA)",
	MBR_W95_EXTENDED_PARTITION:         "W95 Ext'd (LBA)",
	MBR_OPUS_PARTITION:                 "OPUS",
	MBR_HIDDEN_FAT12_PARTITION:         "Hidden FAT12",
	MBR_COMPAQ_DIAGNOSTICS_PARTITION:   "Compaq diagnostics",
	MBR_HIDDEN_FAT16_L32M_PARTITION:    "Hidden FAT16 <32M",
	MBR_HIDDEN_FAT16_PARTITION:         "Hidden FAT16",
	MBR_HIDDEN_HPFS_NTFS_PARTITION:     "Hidden HPFS/NTFS",
	MBR_AST_SMARTSLEEP_PARTITION:       "AST SmartSleep",
	MBR_HIDDEN_W95_FAT32_PARTITION:     "Hidden W95 FAT32",
	MBR_HIDDEN_W95_FAT32LBA_PARTITION:  "Hidden W95 FAT32 (LBA)",
	MBR_HIDDEN_W95_FAT16LBA_PARTITION:  "Hidden W95 FAT16 (LBA)",
	MBR_NEC_DOS_PARTITION:              "NEC DOS",
	byte(0x27):                         "Hidden NTFS WinRE",
	MBR_PLAN9_PARTITION:                "Plan 9",
	MBR_PARTITIONMAGIC_PARTITION:       "PartitionMagic recovery",
	MBR_VENIX80286_PARTITION:           "Venix 80286",
	MBR_PPC_PREP_BOOT_PARTITION:        "PPC PReP Boot",
	MBR_SFS_PARTITION:                  "SFS",
	MBR_QNX_4X_PARTITION:               "QNX4.x",
	MBR_QNX_4X_2ND_PARTITION:           "QNX4.x 2nd part",
	MBR_QNX_4X_3RD_PARTITION:           "QNX4.x 3rd part",
	MBR_DM_PARTITION:                   "OnTrack DM",
	MBR_DM6_AUX1_PARTITION:             "OnTrack DM6 Aux1", /* (or Novell) */
	MBR_CPM_PARTITION:                  "CP/M",             /* CP/M or Microport SysV/AT */
	MBR_DM6_AUX3_PARTITION:             "OnTrack DM6 Aux3",
	MBR_DM6_PARTITION:                  "OnTrackDM6",
	MBR_EZ_DRIVE_PARTITION:             "EZ-Drive",
	MBR_GOLDEN_BOW_PARTITION:           "Golden Bow",
	MBR_PRIAM_EDISK_PARTITION:          "Priam Edisk",
	MBR_SPEEDSTOR_PARTITION:            "SpeedStor",
	MBR_UNIXWARE_PARTITION:             "GNU HURD or SysV", /* GNU HURD or Mach or Sys V/386 (such as ISC UNIX) */
	MBR_NETWARE_286_PARTITION:          "Novell Netware 286",
	MBR_NETWARE_386_PARTITION:          "Novell Netware 386",
	MBR_DISKSECURE_MULTIBOOT_PARTITION: "DiskSecure Multi-Boot",
	MBR_PC_IX_PARTITION:                "PC/IX",
	MBR_OLD_MINIX_PARTITION:            "Old Minix",         /* Minix 1.4a and earlier */
	MBR_MINIX_PARTITION:                "Minix / old Linux", /* Minix 1.4b and later */
	MBR_SOLARIS_X86_PARTITION:          "Linux swap / Solaris",
	MBR_LINUX_DATA_PARTITION:           "Linux",
	MBR_INTEL_HIBERNATION_PARTITION:    "OS/2 hidden or Intel hibernation", /* OS/2 hidden C: drive,
										 * hibernation type Microsoft APM
										 * or hibernation Intel Rapid Start */
	MBR_LINUX_EXTENDED_PARTITION:       "Linux extended",
	MBR_NTFS_VOL_SET1_PARTITION:        "NTFS volume set",
	MBR_NTFS_VOL_SET2_PARTITION:        "NTFS volume set",
	MBR_LINUX_PLAINTEXT_PARTITION:      "Linux plaintext",
	MBR_LINUX_LVM_PARTITION:            "Linux LVM",
	MBR_AMOEBA_PARTITION:               "Amoeba",
	MBR_AMOEBA_BBT_PARTITION:           "Amoeba BBT", /* (bad block table) */
	MBR_BSD_OS_PARTITION:               "BSD/OS",     /* BSDI */
	MBR_THINKPAD_HIBERNATION_PARTITION: "IBM Thinkpad hibernation",
	MBR_FREEBSD_PARTITION:              "FreeBSD", /* various BSD flavours */
	MBR_OPENBSD_PARTITION:              "OpenBSD",
	MBR_NEXTSTEP_PARTITION:             "NeXTSTEP",
	MBR_DARWIN_UFS_PARTITION:           "Darwin UFS",
	MBR_NETBSD_PARTITION:               "NetBSD",
	MBR_DARWIN_BOOT_PARTITION:          "Darwin boot",
	MBR_HFS_HFS_PARTITION:              "HFS / HFS+",
	MBR_BSDI_FS_PARTITION:              "BSDI fs",
	MBR_BSDI_SWAP_PARTITION:            "BSDI swap",
	MBR_BOOTWIZARD_HIDDEN_PARTITION:    "Boot Wizard hidden",
	MBR_ACRONIS_FAT32LBA_PARTITION:     "Acronis FAT32 LBA", /* hidden (+0xb0) Acronis Secure Zone (backup software) */
	MBR_SOLARIS_BOOT_PARTITION:         "Solaris boot",
	MBR_SOLARIS_PARTITION:              "Solaris",
	MBR_DRDOS_FAT12_PARTITION:          "DRDOS/sec (FAT-12)",
	MBR_DRDOS_FAT16_L32M_PARTITION:     "DRDOS/sec (FAT-16 < 32M)",
	MBR_DRDOS_FAT16_PARTITION:          "DRDOS/sec (FAT-16)",
	MBR_SYRINX_PARTITION:               "Syrinx",
	MBR_NONFS_DATA_PARTITION:           "Non-FS data",
	MBR_CPM_CTOS_PARTITION:             "CP/M / CTOS / ...", /* CP/M or Concurrent CP/M or Concurrent DOS or CTOS */
	MBR_DELL_UTILITY_PARTITION:         "Dell Utility",      /* Dell PowerEdge Server utilities */
	MBR_BOOTIT_PARTITION:               "BootIt",            /* BootIt EMBRM */
	MBR_DOS_ACCESS_PARTITION:           "DOS access",        /* DOS access or SpeedStor 12-bit FAT  extended partition */
	MBR_DOS_RO_PARTITION:               "DOS R/O",           /* DOS R/O or SpeedStor */
	MBR_SPEEDSTOR_EXTENDED_PARTITION:   "SpeedStor",         /* SpeedStor 16-bit FAT extended
								  * partition < 1024 cyl. */

	/* Linux https://www.freedesktop.org/wiki/Specifications/BootLoaderSpec/ */
	MBR_RUFUS_EXTRA_PARTITION: "Linux extended boot",

	MBR_BEOS_FS_PARTITION:           "BeOS fs",
	MBR_GPT_PARTITION:               "GPT",                /* Intel EFI GUID Partition Table */
	MBR_EFI_SYSTEM_PARTITION:        "EFI (FAT-12/16/32)", /* Intel EFI System Partition */
	MBR_LINUX_PARISC_BOOT_PARTITION: "Linux/PA-RISC boot", /* Linux/PA-RISC boot loader */
	MBR_SPEEDSTOR1_PARTITION:        "SpeedStor",
	MBR_SPEEDSTOR2_PARTITION:        "SpeedStor",       /* SpeedStor large partition */
	MBR_DOS_SECONDARY_PARTITION:     "DOS secondary",   /* DOS 3.3+ secondary */
	MBR_EBBR_PROTECTIVE_PARTITION:   "EBBR protective", /* Arm EBBR firmware protective partition */
	MBR_VMWARE_VMFS_PARTITION:       "VMware VMFS",
	MBR_VMWARE_VMKCORE_PARTITION:    "VMware VMKCORE",        /* VMware kernel dump partition */
	MBR_LINUX_RAID_PARTITION:        "Linux raid autodetect", /* Linux raid partition with
								   * autodetect using persistent
								   * superblock */
	MBR_LANSTEP_PARTITION:   "LANstep", /* SpeedStor >1024 cyl. or LANstep */
	MBR_XENIX_BBT_PARTITION: "BBT",     /* Xenix Bad Block Table */
}
