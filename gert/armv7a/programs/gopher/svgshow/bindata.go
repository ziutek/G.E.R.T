// Copyright 2017 Yanni Coroneos. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func data_bindata_gob() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8c, 0x97,
		0x09, 0x6c, 0x16, 0x45, 0x1b, 0xc7, 0xf7, 0x3f, 0xdb, 0xc2, 0x17, 0xe0,
		0x2b, 0xfd, 0x3e, 0xc1, 0x20, 0x28, 0xa8, 0x80, 0x07, 0x28, 0x51, 0x4c,
		0xb8, 0x24, 0x44, 0x2e, 0x15, 0x04, 0x14, 0x49, 0x00, 0x5b, 0xe5, 0x56,
		0x0e, 0x81, 0x22, 0x69, 0xf0, 0x56, 0x04, 0x24, 0x42, 0x10, 0xa4, 0xdc,
		0x10, 0x40, 0x28, 0x87, 0xa0, 0x56, 0x41, 0x2b, 0x55, 0x90, 0xe0, 0x81,
		0x80, 0x60, 0xb8, 0xc1, 0x70, 0x29, 0x22, 0x87, 0x41, 0x05, 0x6d, 0x8b,
		0xe0, 0xb1, 0xe3, 0xff, 0xd9, 0xd9, 0xd9, 0x77, 0xa7, 0xaf, 0x24, 0xa6,
		0x79, 0xf6, 0xf7, 0xeb, 0xee, 0xcc, 0xec, 0x3b, 0xd7, 0xee, 0xb3, 0xff,
		0xd5, 0xe3, 0x15, 0x94, 0x9e, 0xe0, 0x41, 0x8f, 0xf3, 0xbc, 0xdb, 0xf4,
		0x4b, 0x3e, 0x50, 0xad, 0x7d, 0xde, 0x88, 0x51, 0xfd, 0x07, 0xe6, 0x3f,
		0x90, 0x37, 0x74, 0x64, 0xbe, 0x9c, 0x07, 0x4f, 0xf6, 0x46, 0x25, 0x0f,
		0x78, 0x48, 0x8e, 0x99, 0xed, 0xf3, 0x86, 0xe7, 0x8d, 0xa6, 0x79, 0x5e,
		0x90, 0x75, 0x82, 0x75, 0x03, 0xec, 0xc2, 0x18, 0xe8, 0xbd, 0x1e, 0x1e,
		0x85, 0xde, 0xe2, 0x21, 0x17, 0x7a, 0x9d, 0x87, 0x7e, 0xd0, 0x4b, 0xed,
		0x71, 0x24, 0xf4, 0x4c, 0xde, 0x62, 0x2c, 0x74, 0x01, 0x31, 0x0b, 0x7a,
		0x5e, 0x0c, 0x53, 0x31, 0x79, 0x0c, 0xfc, 0x5e, 0xd0, 0x1f, 0x0b, 0x87,
		0x40, 0xaf, 0x11, 0x8e, 0x86, 0x5e, 0x22, 0x64, 0x33, 0x53, 0x92, 0xec,
		0x8b, 0x17, 0x04, 0x5d, 0xf1, 0x9c, 0xa0, 0x05, 0xa4, 0x1f, 0x29, 0xda,
		0x56, 0x2a, 0x72, 0x05, 0x9e, 0x11, 0xac, 0x85, 0x5e, 0x25, 0x5c, 0x0a,
		0xbd, 0x4f, 0xc8, 0xae, 0xfe, 0x92, 0xe4, 0xb3, 0xd0, 0x65, 0x49, 0x5e,
		0x87, 0xde, 0x09, 0xd4, 0x43, 0xbf, 0x04, 0x7a, 0xa0, 0x8b, 0x60, 0x3c,
		0x7a, 0xc6, 0x37, 0x68, 0x06, 0x3d, 0xd5, 0x43, 0x0e, 0x86, 0xb2, 0xa7,
		0x93, 0xd1, 0x87, 0xc7, 0x4d, 0x18, 0x15, 0x1f, 0x37, 0x62, 0x88, 0x3d,
		0x0e, 0x45, 0x80, 0xca, 0x31, 0x46, 0x11, 0x19, 0x31, 0xda, 0x40, 0x9f,
		0xf1, 0xd0, 0x88, 0x25, 0x6d, 0x83, 0x8d, 0xd8, 0x7e, 0xa6, 0xfc, 0x13,
		0xf8, 0xca, 0x43, 0x07, 0x04, 0xea, 0x4f, 0x0f, 0xdd, 0x89, 0xb2, 0x18,
		0x3d, 0x88, 0x52, 0x0f, 0x3d, 0x09, 0x76, 0xa4, 0x37, 0x71, 0xde, 0xe3,
		0xcc, 0x25, 0x91, 0x4b, 0x9c, 0xf3, 0xf0, 0x70, 0x1a, 0x7e, 0xf6, 0xd0,
		0xc7, 0xa0, 0xaf, 0x83, 0xc7, 0x88, 0x1f, 0xc2, 0x1f, 0xa9, 0xbe, 0xf3,
		0xf0, 0x38, 0x71, 0xc0, 0x01, 0xa7, 0x6d, 0x38, 0xb1, 0xcb, 0x60, 0x47,
		0x8c, 0x11, 0xc4, 0x06, 0x83, 0x35, 0x06, 0x2b, 0x63, 0x54, 0x25, 0x4e,
		0xc7, 0xa8, 0x4e, 0x5c, 0xf4, 0x50, 0xdb, 0x74, 0x2c, 0xea, 0x66, 0xe0,
		0x0f, 0x44, 0x90, 0x59, 0x45, 0x84, 0x77, 0xc9, 0xbc, 0x46, 0x84, 0xbf,
		0x21, 0xf3, 0x56, 0x91, 0x47, 0x28, 0xad, 0x92, 0xa2, 0x2e, 0x46, 0x85,
		0x63, 0xf1, 0xab, 0x21, 0xc8, 0x28, 0x35, 0xd3, 0x16, 0x64, 0x9c, 0x30,
		0xab, 0x23, 0xc8, 0x38, 0xec, 0x48, 0x74, 0x0b, 0xbd, 0x8b, 0xbc, 0x3d,
		0x62, 0xd3, 0x04, 0x77, 0x56, 0xe0, 0x67, 0x64, 0x6b, 0xf2, 0x23, 0xb2,
		0x3d, 0xb9, 0x86, 0xbc, 0x37, 0xc1, 0xe5, 0x64, 0x37, 0x72, 0x01, 0xd9,
		0x95, 0x9c, 0x19, 0x9d, 0xb7, 0x9c, 0x41, 0xde, 0x4d, 0x16, 0x90, 0x1d,
		0xc9, 0xe9, 0x64, 0x87, 0x0a, 0x94, 0x76, 0x5f, 0x4b, 0xb0, 0xdd, 0x65,
		0xd8, 0x96, 0x9c, 0x56, 0x81, 0x25, 0xec, 0xd8, 0xe9, 0x04, 0xd9, 0x8f,
		0x0c, 0xae, 0x22, 0x7d, 0x8a, 0x2c, 0x27, 0xff, 0x62, 0xb9, 0xac, 0x88,
		0xd5, 0x49, 0x4d, 0x66, 0x73, 0x20, 0xc0, 0xc8, 0xfc, 0xbf, 0x88, 0xa2,
		0xd4, 0x70, 0xc4, 0xa7, 0x5c, 0x29, 0x92, 0x49, 0xb9, 0x4a, 0xa4, 0x12,
		0xa5, 0x76, 0x42, 0xec, 0xf0, 0x05, 0xfe, 0x76, 0x4e, 0x40, 0xb9, 0xc8,
		0x26, 0xb3, 0x10, 0x03, 0xbf, 0xd8, 0x2c, 0xa2, 0xc0, 0x5f, 0x65, 0xe6,
		0x3a, 0x21, 0x2b, 0x28, 0x27, 0x45, 0x0a, 0x29, 0xe1, 0x04, 0x2d, 0xa6,
		0x1c, 0x77, 0x64, 0x91, 0x95, 0x85, 0x94, 0x6f, 0x45, 0x16, 0x5c, 0x5e,
		0xe6, 0x53, 0xbe, 0x11, 0x99, 0x97, 0x26, 0xb3, 0x29, 0x47, 0x45, 0x66,
		0x51, 0x0e, 0x89, 0x70, 0x4a, 0xd4, 0x9e, 0x34, 0xd9, 0x69, 0x65, 0x87,
		0x08, 0xa7, 0x4b, 0x6d, 0x71, 0x84, 0x13, 0xa7, 0x4a, 0xac, 0xbc, 0x6d,
		0x65, 0x99, 0x23, 0x47, 0xcc, 0x4e, 0x49, 0xc8, 0x21, 0x3b, 0x08, 0xfb,
		0xcc, 0x36, 0x4d, 0x0d, 0x94, 0x7a, 0x9d, 0xc3, 0x78, 0x97, 0xc8, 0x71,
		0x4a, 0x43, 0xb9, 0x54, 0x19, 0xfa, 0xbc, 0xb0, 0x0a, 0xf4, 0xea, 0x04,
		0xd5, 0xfb, 0xb8, 0x5f, 0xd0, 0x0a, 0x83, 0x04, 0x0a, 0xfa, 0x60, 0x05,
		0x7e, 0x9d, 0xe4, 0xb5, 0x6c, 0xae, 0x96, 0x08, 0x1f, 0x02, 0x68, 0x9d,
		0xbc, 0x13, 0x76, 0x40, 0x7f, 0x25, 0x2c, 0x86, 0x9e, 0x24, 0x7c, 0x11,
		0x03, 0x04, 0xcd, 0x91, 0x13, 0x43, 0x97, 0x22, 0x9f, 0xc7, 0x2f, 0xc3,
		0xe7, 0xa4, 0xde, 0xcc, 0xff, 0x2f, 0x83, 0x00, 0x57, 0x33, 0xee, 0x13,
		0xe1, 0xa8, 0xa1, 0x65, 0xa2, 0x7d, 0x75, 0x3d, 0x7d, 0xad, 0x48, 0x7d,
		0xca, 0x7a, 0x2b, 0x9f, 0x5a, 0xd9, 0x9a, 0x26, 0x3b, 0xad, 0x1c, 0xb0,
		0xd5, 0x8f, 0xa6, 0xc9, 0x61, 0x47, 0xea, 0x51, 0x64, 0x7e, 0x55, 0x1d,
		0x8a, 0x0c, 0xb5, 0xaa, 0x41, 0x39, 0xe9, 0xc8, 0x15, 0xff, 0x42, 0xb2,
		0x29, 0xa7, 0x44, 0xb2, 0x28, 0xb2, 0x38, 0x15, 0x9f, 0x1e, 0x38, 0xe3,
		0x48, 0xb6, 0xbd, 0x14, 0x0b, 0x97, 0x3f, 0xbe, 0x97, 0x1e, 0xcb, 0xdf,
		0x31, 0x11, 0x3e, 0x7c, 0xb0, 0xdf, 0x91, 0x4b, 0x51, 0x77, 0x52, 0x22,
		0x97, 0xb6, 0x5b, 0xd9, 0x64, 0x65, 0x6d, 0x9a, 0x14, 0xd9, 0x5a, 0xab,
		0xad, 0x2c, 0x73, 0xe4, 0x77, 0xc6, 0x7c, 0x2b, 0x33, 0x44, 0xfe, 0x60,
		0x4c, 0x71, 0xa4, 0x94, 0x31, 0x3d, 0x29, 0xaa, 0x21, 0x65, 0x89, 0x23,
		0x7c, 0x3c, 0x62, 0x41, 0x9a, 0xc4, 0x97, 0xde, 0x48, 0x4c, 0xa5, 0x3e,
		0x01, 0x7d, 0x84, 0xb8, 0xc0, 0x27, 0x87, 0xb4, 0x5a, 0x27, 0x62, 0x93,
		0xf0, 0x74, 0x8a, 0x3c, 0x5f, 0x62, 0x8a, 0x95, 0xc4, 0x95, 0x02, 0xfc,
		0x0f, 0x5a, 0x06, 0x1b, 0x35, 0xa1, 0x65, 0x33, 0xa0, 0x6e, 0xc4, 0xfa,
		0xd1, 0x79, 0xcb, 0xba, 0x51, 0xf9, 0x9a, 0x6e, 0x3d, 0x2e, 0x69, 0xbd,
		0x5b, 0xd8, 0x14, 0xba, 0xdc, 0xac, 0x6c, 0xc3, 0x61, 0xd1, 0x79, 0x4b,
		0x9e, 0x2f, 0x8a, 0xca, 0x15, 0x25, 0xeb, 0x75, 0x85, 0x96, 0x9d, 0xaf,
		0xba, 0xf3, 0x27, 0x09, 0x73, 0x22, 0xf6, 0x8d, 0xce, 0x5b, 0xe6, 0x44,
		0xe5, 0xbb, 0xbb, 0xf5, 0x50, 0x12, 0xad, 0x43, 0x6c, 0x65, 0xfc, 0x24,
		0x72, 0x94, 0x71, 0x56, 0xa4, 0x8c, 0x71, 0xd0, 0x91, 0x72, 0xbb, 0x00,
		0x1c, 0xd9, 0xeb, 0x48, 0x59, 0xb4, 0xf8, 0x43, 0xd9, 0x60, 0xa5, 0xc8,
		0x11, 0x29, 0xfc, 0xa6, 0x95, 0x95, 0x56, 0x96, 0x38, 0x72, 0x81, 0x31,
		0x47, 0xe4, 0x37, 0xc6, 0x34, 0x2b, 0xaf, 0xa4, 0xc9, 0x58, 0x47, 0xce,
		0x31, 0xc6, 0x39, 0xb2, 0xce, 0xae, 0x96, 0x58, 0x8a, 0xd3, 0xe4, 0x3d,
		0xc6, 0x4c, 0x47, 0x8a, 0x18, 0x6f, 0x89, 0xbc, 0xcb, 0xf8, 0xc2, 0x19,
		0x28, 0xbe, 0xce, 0x50, 0x20, 0xc2, 0xe7, 0x30, 0x16, 0x8a, 0x14, 0x4a,
		0x88, 0xbc, 0xc3, 0x58, 0x6e, 0x65, 0x85, 0xc8, 0x87, 0xf6, 0xcc, 0xe7,
		0xb6, 0xcc, 0x6e, 0xc6, 0x62, 0x47, 0xf6, 0x30, 0x16, 0x89, 0xec, 0xb3,
		0x72, 0xc0, 0xb6, 0x1c, 0xcb, 0x49, 0xc6, 0x5c, 0x91, 0xf3, 0x8c, 0xd9,
		0x76, 0x37, 0xb9, 0x72, 0xa9, 0xa2, 0x28, 0x44, 0x63, 0xa8, 0xaa, 0xca,
		0xcf, 0x15, 0xa9, 0x15, 0xdd, 0x22, 0x25, 0xb2, 0x65, 0x64, 0xeb, 0xa9,
		0x26, 0xd1, 0x5c, 0xa8, 0x66, 0x76, 0x83, 0xc4, 0xd2, 0xce, 0x5e, 0xea,
		0x42, 0x59, 0x2a, 0xc2, 0x64, 0x2c, 0xdc, 0xa7, 0x29, 0xe9, 0x1f, 0x4d,
		0x4a, 0x28, 0x63, 0x44, 0x7a, 0x51, 0x06, 0x39, 0xd2, 0x99, 0xd2, 0x4b,
		0xa4, 0x25, 0xcc, 0xf3, 0x35, 0xbc, 0x7b, 0x5b, 0x47, 0x1a, 0xa4, 0x49,
		0xfd, 0x8a, 0x82, 0x8d, 0x8c, 0x8e, 0x8e, 0xf0, 0xdd, 0x8a, 0x1e, 0x22,
		0xaf, 0x32, 0x9e, 0x70, 0xa6, 0x69, 0x3d, 0xa3, 0x93, 0x1d, 0xde, 0x06,
		0xd2, 0x4e, 0x65, 0x4a, 0x5d, 0x11, 0x66, 0x65, 0x68, 0xe5, 0x88, 0x3c,
		0x60, 0x73, 0xed, 0x2e, 0xe8, 0xef, 0x54, 0x8f, 0x45, 0xe6, 0xf4, 0x4e,
		0x11, 0x26, 0x4c, 0x68, 0xf7, 0xcf, 0x65, 0xac, 0xa8, 0x4f, 0xd0, 0x98,
		0xf0, 0x3d, 0x74, 0x16, 0xdc, 0x81, 0xa7, 0x04, 0xfc, 0xa0, 0x28, 0x4e,
		0x32, 0x0f, 0x5a, 0xb6, 0x9e, 0xff, 0x3c, 0x6b, 0x35, 0x12, 0x99, 0x48,
		0x19, 0xee, 0xc8, 0xe4, 0x68, 0xb5, 0xf8, 0x53, 0x28, 0xe1, 0xbb, 0x9e,
		0x3d, 0x55, 0x8d, 0xd3, 0xa4, 0xb9, 0x95, 0xb6, 0x56, 0x3a, 0x39, 0x32,
		0x95, 0x92, 0x2f, 0xc2, 0xec, 0x4a, 0xc9, 0x92, 0xf0, 0x99, 0x76, 0x29,
		0x79, 0x91, 0xf9, 0xd3, 0x93, 0x12, 0x27, 0x06, 0xb2, 0xcd, 0xc3, 0xe4,
		0xe1, 0x47, 0x47, 0xe6, 0x30, 0x6e, 0xf0, 0x4c, 0x36, 0xe2, 0x3f, 0xe8,
		0x99, 0x44, 0xc5, 0x7f, 0xda, 0x11, 0xa6, 0x25, 0x7e, 0x98, 0xe7, 0xbc,
		0xcc, 0x4c, 0x2d, 0xcc, 0x6c, 0x87, 0x50, 0x26, 0x39, 0xd2, 0x86, 0xb2,
		0x59, 0x24, 0xc3, 0x64, 0x73, 0x4c, 0xaf, 0x99, 0x77, 0xd5, 0x72, 0x24,
		0x9f, 0x72, 0x8b, 0xc8, 0x8d, 0x94, 0x66, 0x32, 0xbc, 0xfb, 0x29, 0x2d,
		0x1c, 0xe9, 0x46, 0xe1, 0xe4, 0xea, 0x63, 0x6c, 0xe6, 0x57, 0x93, 0x4c,
		0x66, 0x14, 0xa6, 0x38, 0x9a, 0xe0, 0x37, 0xcb, 0x00, 0xa2, 0xaa, 0x7c,
		0xc4, 0x31, 0x51, 0x8b, 0xc1, 0x05, 0xeb, 0x73, 0xe5, 0x32, 0x6d, 0xf6,
		0x07, 0x7b, 0x18, 0x4c, 0xb4, 0x8e, 0x31, 0x8c, 0xc8, 0x36, 0xd9, 0x7f,
		0xa9, 0x49, 0xf4, 0x0f, 0xc7, 0x18, 0x49, 0x6c, 0x33, 0xf8, 0xc0, 0xa0,
		0x30, 0x86, 0x54, 0x18, 0x17, 0x56, 0x57, 0xb9, 0xe6, 0x6b, 0xe3, 0xa6,
		0x18, 0xbc, 0xad, 0xbc, 0x61, 0xf9, 0x5d, 0x22, 0xaf, 0x1f, 0xc1, 0xc4,
		0x18, 0xb2, 0x79, 0xba, 0x99, 0x8f, 0x27, 0x66, 0xa5, 0x4f, 0x86, 0x9f,
		0x4b, 0xe6, 0xa8, 0xe7, 0x86, 0x1f, 0x8d, 0x7a, 0x1b, 0x86, 0x49, 0x7f,
		0xff, 0x63, 0x76, 0x42, 0x84, 0x7b, 0x70, 0xb3, 0xa0, 0x20, 0xcc, 0x8e,
		0x70, 0x16, 0x59, 0x29, 0xa8, 0x36, 0x90, 0x64, 0x58, 0x4d, 0x40, 0x6d,
		0xbb, 0x1e, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x37, 0x1d, 0xb8,
		0x25, 0x0f, 0x00, 0x00,
	},
		"data/bindata.gob",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	if f, ok := _bindata[name]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"data/bindata.gob": data_bindata_gob,
}
