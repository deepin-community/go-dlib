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

package util

import (
	"testing"
)

func Test_GetDateFromJulianDay(t *testing.T) {
	y, m, d := GetDateFromJulianDay(2457438.0)
	t.Log(y, m, d)
	if y == 2016 && m == 2 && d == 19 {
		t.Log("ok")
	} else {
		t.Error("fail")
	}

	y, m, d = GetDateFromJulianDay(2248528.0)
	t.Log(y, m, d)
	if y == 1444 && m == 2 && d == 19 {
		t.Log("ok")
	} else {
		t.Error("fail")
	}
}

func Test_GetTimeFromJulianDay(t *testing.T) {
	h, m, s := GetTimeFromJulianDay(2457438.09546)
	t.Log(h, m, s)
	if h == 14 && m == 17 && s == 27 {
		t.Log("ok")
	} else {
		t.Error("fail")
	}

	h, m, s = GetTimeFromJulianDay(2457438.09851)
	t.Log(h, m, s)
	if h == 14 && m == 21 && s == 51 {
		t.Log("ok")
	} else {
		t.Error("fail")
	}
}

func Test_GetDateTimeFromJulianDay(t *testing.T) {
	dt := GetDateTimeFromJulianDay(2457438.10454)
	t.Log(dt)
	if dt.String() == "2016-02-19 14:29:22 +0000 UTC" {
		t.Log("ok")
	} else {
		t.Error("fail")
	}
}
