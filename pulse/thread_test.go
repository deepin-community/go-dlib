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

package pulse

import (
	"math/rand"
	"testing"
)

func BenchmarkTheadSafe(b *testing.B) {
	ctx := GetContext()
	sinks := ctx.GetSinkList()
	for _, s := range sinks {
		old := s.Volume
		for i := 0; i < b.N; i++ {
			v := s.Volume.SetAvg(rand.Float64())
			ctx.SetSinkVolumeByIndex(s.Index, v)
		}
		ctx.SetSinkVolumeByIndex(s.Index, old)
	}
}
