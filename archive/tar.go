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

package archive

import (
	"archive/tar"
	"os"
	"pkg.deepin.io/lib/archive/utils"
)

func tarCompressFiles(files []string, dest string) error {
	dw, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer dw.Close()

	tw := tar.NewWriter(dw)
	defer tw.Close()

	return utils.TarWriterCompressFiles(tw, files)
}

func tarExtracteFile(src, dest string) ([]string, error) {
	sr, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer sr.Close()

	return utils.TarReaderExtracte(tar.NewReader(sr), dest)
}
