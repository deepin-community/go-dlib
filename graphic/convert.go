/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package graphic

import (
	"fmt"
)

// ConvertImage converts from any recognized format to target format image.
func ConvertImage(srcfile, dstfile string, f Format) (err error) {
	srcimg, err := LoadImage(srcfile)
	if err != nil {
		return
	}
	return SaveImage(dstfile, srcimg, f)
}

// ConvertImageCache converts from any recognized format to cache
// directory, if already exists, just return it.
func ConvertImageCache(srcfile string, f Format) (dstfile string, useCache bool, err error) {
	dstfile = generateCacheFilePath(fmt.Sprintf("ConvertImageCache%s%s", srcfile, f))
	if isFileExists(dstfile) {
		useCache = true
		return
	}
	err = ConvertImage(srcfile, dstfile, f)
	return
}
