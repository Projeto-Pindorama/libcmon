/*
 * zip.go - Hip unzip (and zip) boilerplates for Google Go's archive/zip
 *
 * Copyright (c) 2023-2025 Luiz Antônio Rangel
 *
 * SPDX-Licence-Identifier: BSD 3-Clause
 */

package zhip

import (
	"archive/zip"
	"errors"
	"os"
)

var zipent int = 0
var CompressionMethod uint16 = zip.Store
var EntryNo = make(map[string]int)

func GetZipEntries(f *zip.ReadCloser) *zip.FileHeader {
	var finfo *zip.FileHeader
	if len(f.File) < (zipent + 1) {
		zipent = 0
		return nil
	}
	finfo = &f.File[zipent].FileHeader
	EntryNo[finfo.Name] = zipent
	zipent += 1
	return finfo
}

func GetZipESlice(f *zip.ReadCloser) []*zip.FileHeader {
	var finfo *zip.FileHeader
	var zentries []*zip.FileHeader
	for ;; {
		if finfo = GetZipEntries(f); finfo == nil {
			break
		}
		zentries = append(zentries, finfo)
	}
	return zentries
}

func GetZipLargestEntry(f *zip.ReadCloser) uint32 {
	var finfo *zip.FileHeader
	/* Get the largest file size. */
	longlen := uint32(0)
	for ;; {
		if finfo = GetZipEntries(f); finfo == nil {
			break
		}
		curlen := finfo.UncompressedSize
		if curlen > longlen {
			longlen = curlen
		}
	}
	return longlen
}

func GetCompressionMethod(f *zip.FileHeader) string {
	switch f.Method {
		case zip.Deflate:
			return "Deflt"
		case zip.Store:
			return "Store"
		default:
			return ""
	}
}

func GetCompressionRatio(f *zip.FileHeader) float32 {
	if m := GetCompressionMethod(f); m == "Store" ||
		f.UncompressedSize64 == uint64(0) {
		return float32(0)
	}
	return float32(100 - float32((f.CompressedSize64 * 100) / f.UncompressedSize64))
}

func RecordNewEntry(awriter *zip.Writer, name string) (int, error) {
	var wbytes int
	file, err := os.Stat(name)
	if err != nil {
		return 0, err
	}

	if file.IsDir() {
		/*
		 * From 'go doc zip.Create':
		 * "To create a directory instead of a file,
		 * add a trailing slash to the name."
		 */
		if name[(len(name)-1):] != "/" {
			name += "/"
		}
	}
	entfhdr, err_fhdr := zip.FileInfoHeader(file)
	/*
	 * Set the name. Per the contrary, we will not be
	 * having files on subdirectories.
	 */
	entfhdr.Name = name
	if !file.IsDir() {
		/*
		 * Can be set outside with
		 * 'zhip.CompressionMethod = [...]'
		 * now.
		 */
		entfhdr.Method = CompressionMethod
	}
	ent, err_creat := awriter.CreateHeader(entfhdr)
	err = errors.Join(err_fhdr, err_creat)
	if err != nil {
		return 0, err
	}

	if !file.IsDir() {
		data, err := os.ReadFile(name)
		if err != nil {
			return 0, err
		}
		wbytes, err = ent.Write(data)
		if err != nil {
			return 0, err
		}
	}

	EntryNo[name] = zipent
	zipent += 1
	return wbytes, nil
}

func LocateZipEntry(name string) int {
	ent, ok := EntryNo[name]
	if !ok {
		/*
		 * This would probably panic().
		 * Perhaps think of a better option later?
		 */
		ent = -1
	}
	return ent
}
