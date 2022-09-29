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
	"image"
	"image/draw"
	"os"
	"pkg.deepin.io/lib/utils"
)

func generateCacheFilePath(keyword string) (dstfile string) {
	return utils.GenerateCacheFilePathWithPrefix("graphic", keyword)
}

func openFileOrCreate(file string) (*os.File, error) {
	return os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0644)
}

func isFileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	} else {
		return false
	}
}

// convert image.Image to *image.RGBA
func convertToRGBA(img image.Image) (rgba *image.RGBA) {
	b := img.Bounds()
	r := image.Rect(0, 0, b.Dx(), b.Dy())
	rgba = image.NewRGBA(r)
	draw.Draw(rgba, rgba.Bounds(), img, b.Min, draw.Src)
	return
}
