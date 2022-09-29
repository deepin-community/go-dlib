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

package utils

import (
	"fmt"
	"os"
	"path"
)

var DefaultCachePrefix = os.Getenv("HOME") + "/.cache/deepin"

func GenerateCacheFilePath(keyword string) (dstfile string) {
	cachePathFormat := DefaultCachePrefix + "/%s"
	md5, _ := SumStrMd5(keyword)
	dstfile = fmt.Sprintf(cachePathFormat, md5)
	_ = EnsureDirExist(path.Dir(dstfile))
	return
}

func GenerateCacheFilePathWithPrefix(prefix, keyword string) (dstfile string) {
	graphicCacheFormat := DefaultCachePrefix + "/%s/%s"
	md5, _ := SumStrMd5(keyword)
	dstfile = fmt.Sprintf(graphicCacheFormat, prefix, md5)
	_ = EnsureDirExist(path.Dir(dstfile))
	return
}
