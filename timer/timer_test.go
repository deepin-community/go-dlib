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

package timer_test

import (
	. "github.com/smartystreets/goconvey/convey"
	. "pkg.deepin.io/lib/timer"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	Convey("just stop", t, func(c C) {
		timer := NewTimer()
		timer.Stop()
		c.So(timer.Elapsed(), ShouldEqual, 0)
	})

	Convey("get elapse without stop", t, func(c C) {
		timer := NewTimer()
		timer.Start()

		time.Sleep(time.Second)
		c.So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

		time.Sleep(time.Second)
		c.So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second*2-time.Millisecond*100, time.Second*3+time.Millisecond*100)
	})

	Convey("stop and elapse", t, func(c C) {
		timer := NewTimer()
		timer.Start()

		time.Sleep(time.Second)
		timer.Stop()
		c.So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

		time.Sleep(time.Second)
		c.So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)
	})

	Convey("stop and continue", t, func(c C) {
		timer := NewTimer()
		timer.Start()

		time.Sleep(time.Second)
		c.So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

		timer.Stop()
		c.So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

		time.Sleep(time.Second)
		timer.Continue()
		c.So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

		time.Sleep(time.Second)
		c.So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second*2-time.Millisecond*100, time.Second*2+time.Millisecond*100)
	})

	Convey("reset", t, func(c C) {
		timer := NewTimer()
		timer.Start()

		time.Sleep(time.Second)
		c.So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

		timer.Reset()
		time.Sleep(time.Second)
		c.So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)
	})
}
