/*
 * zip.go - Hip unzip (and zip) boilerplates for  Google Go's archive/zip
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
	finfo = f.File[zipent].FileHeader
	zipent += 1
	return finfo
}
