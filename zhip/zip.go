/*
 * zip.go - Hip unzip (and zip) boilerplates for Google Go's archive/zip
 *
 * Copyright (c) 2023-2025 Luiz Antônio Rangel
 *
 * SPDX-Licence-Identifier: BSD 3-Clause
 */

package zhip

import "archive/zip"

var zipent int = 0;

func GetZipEntries(f *zip.ReadCloser) (*zip.FileHeader) {
	var finfo *zip.FileHeader
	if len(f.File) < (zipent + 1)  {
		zipent = 0
		return nil
	}
	finfo = &f.File[zipent].FileHeader
	zipent += 1
	return finfo
}

func GetZipESlice(f *zip.ReadCloser) ([]*zip.FileHeader) {
	var finfo *zip.FileHeader
	var zentries []*zip.FileHeader
	for ;;  {
		if finfo = GetZipEntries(f); (finfo == nil) {
			break
		}
		zentries = append(zentries, finfo)
	}
	return zentries
}

func GetZipLargestEntry(f *zip.ReadCloser) (uint32) {
	var finfo *zip.FileHeader
	/* Get the largest file size. */
	longlen := uint32(0)
	for ;;  {
		if finfo = GetZipEntries(f); (finfo == nil) {
			break
		}
		curlen := finfo.UncompressedSize
		if curlen > longlen {
			longlen = curlen
		}
	}
	return longlen
}

func GetCompressionMethod(f *zip.FileHeader) (string) {
	switch (f.Method) {
	case zip.Deflate:
		return "DEFLATE"
	case zip.Store:
		return "Stored"
	default:
		return ""
	}
}

func GetCompressionRatio(f *zip.FileHeader) (float32) {
	if m := GetCompressionMethod(f); m == "Stored" {
		return float32(0)
	} else {
		return float32(100 - ((f.CompressedSize * 100) / f.UncompressedSize))
	}
}
