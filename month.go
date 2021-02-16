// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package month

import (
    "strconv"
    "strings"
    "time"
)

// 线程安全的月份数据类型
//  range：197001~999912
type Month uint32

const (
    minYear uint32 = 1970
    maxYear uint32 = 9999
    minDot  uint32 = 197001
    maxDot  uint32 = 999912
)

// 构建Month类型数据
func parse(_year, _month uint32) Month {
    if _year < minYear || _year > maxYear || _month < 1 || _month > 12 {
        panic("month parsing failed")
    }
    return Month(_year*100 + _month)
}

// 最小Month实例
func Min() Month {
    return Month(minDot)
}

// 最大Month实例
func Max() Month {
    return Month(maxDot)
}

// 当前时间的实例，即time.Now()的月份值
func Current() Month {
    return FromTime(time.Now())
}

// 创建一个Month实例
func New(year, month uint32) Month {
    return parse(year, month)
}


// 从time.Time创建一个Month实例
func FromTime(t time.Time) Month {
    return parse(uint32(t.Year()), uint32(t.Month()))
}

// 从dot创建一个Month实例，dot是一个符合月份格式的数字。
//  如：197001、202012等属于合法数据。
//  而 202013、9999901等属于越界的非法数据
func FromDot(dot uint32) Month {
    return parse(dot/100, dot%100)
}

// 从月份戳(197001开始计数)创建一个Month实例
func FromTick(tick uint32) Month {
    _year := minYear + (tick-1)/12
    _month := tick % 12
    if tick%12 == 0 {
        _month = 12
    }
    return parse(_year, _month)
}

// 当前Month实例的年份值
func (m Month) Year() uint32 {
    return uint32(m / 100)
}

// 当前Month实例的月份值
func (m Month) Month() uint32 {
    return uint32(m % 100)
}

// 当前m实例的上个月
func (m Month) Prev() Month {
    return m.Add(-1)
}

// 当前m实例的下个月
func (m Month) Next() Month {
    return m.Add(1)
}

// 从197001开始计数的月份戳
func (m Month) Tick() uint32 {
    return (m.Year()-minYear)*12 + m.Month()
}

// 在当前月份实例上加一个差值，返回一个新Month实例
func (m Month) Add(span int) Month {
    t := int(m.Tick()) + span
    return FromTick(uint32(t))
}

// 当前Month实例与另一个Month实例的差值
func (m Month) Diff(other Month) int {
    return int(m.Tick()) - int(other.Tick())
}

// 当前Month所在季度
func (m Month) Quarter() int {
    return int((m.Month()-1)/3 + 1)
}

// Formatted output,
//  example: YYYY年MM月、mm.YY、yy-mm等
func (m Month) Format(format ...string) string {
    if len(format) > 0 {
        _fmt := strings.ToLower(format[0])
        _result := strings.Replace(_fmt, "yyyy", strconv.Itoa(int(m.Year())), 1)
        if _fmt == _result {
            _result = strings.Replace(_fmt, "yy", strconv.Itoa(int(m.Year()%100)), 1)
        }
        if _fmt == _result {
            _result = strings.Replace(_fmt, "y", strconv.Itoa(int(m.Year()%100)), 1)
        }

        _fmt = _result
        _result = strings.Replace(_fmt, "mm", m.String()[4:], 1)

        if _fmt == _result {
            _result = strings.Replace(_fmt, "m", strconv.Itoa(int(m.Month())), 1)
        }
        return _result
    }
    return m.String()
}


func (m Month) String() string {
    return strconv.Itoa(int(m))
}

// 月份轴,左闭右开区间 :[from,to)  fixme:线程安全
func Span(from, to Month) []Month {
    _result := make([]Month, 0)
    for from.Tick() < to.Tick() {
        _result = append(_result, from)
        from = from.Add(1)
    }
    return _result
}

// 遍历一段月份实例数轴,左闭右开: [from,to)
func Range(from, to Month, f func(m Month)) {
    for from.Tick() < to.Tick() {
        f(from)
        from = from.Add(1)
    }
}
