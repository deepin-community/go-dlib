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

package pinyin

import (
	"strconv"
	"strings"
	"unicode"
)

func HansToPinyin(hans string) []string {
	return getPinyinFromKey(hans)
}

func getPinyinFromKey(key string) []string {
	rets := []string{}
	for _, c := range key {
		if unicode.Is(unicode.Scripts["Han"], c) {
			array := getPinyinByHan(int64(c))
			if len(rets) == 0 {
				rets = array
				continue
			}
			rets = rangeArray(rets, array)
		} else {
			array := []string{string(c)}
			if len(rets) == 0 {
				rets = array
			} else {
				rets = rangeArray(rets, array)
			}
		}
	}

	return rets
}

func getPinyinByHan(han int64) []string {
	code := strconv.FormatInt(han, 16)
	value := PinyinDataMap[strings.ToUpper(code)]
	array := strings.Split(value, ";")
	return array
}

func rangeArray(a1, a2 []string) []string {
	rets := []string{}
	for _, v := range a1 {
		for _, r := range a2 {
			rets = append(rets, v+r)
		}
	}

	return rets
}
