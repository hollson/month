# Month数据类型
`month.Month`是一个跟`time.Time`类似的**月期类型**，适用于表达年月数据，如账单、税期、月刊等信息。

<br/>

## 安装
```go
go get github.com/hollson/month
```
<br/>

## 使用示例
```go
package main

import (
    "fmt"
    "time"

    "github.com/hollson/month"
)

func main() {
    Instance()
    Methods()
    Ranges()
}

// 创建Month实例
func Instance() {
    m1 := month.Current()
    m2 := month.Month(202101)
    m3 := month.New(2021, 12)
    m4 := month.FromTime(time.Now())
    m5 := month.FromDot(201801)
    m6 := month.FromTick(3)
    m7 := month.Min()
    m8 := month.Max()
    fmt.Printf(" %v\n %v\n %v\n %v\n %v\n %v\n %v\n %v\n",
        m1, m2, m3, m4, m5, m6, m7, m8)
}

// 实例方法
func Methods() {
    m := month.FromDot(201801)
    fmt.Println(m.Year())
    fmt.Println(m.Month())
    fmt.Println(m.Add(-2))
    fmt.Println(m.Add(13))
    fmt.Println(m.Prev())
    fmt.Println(m.Next())
    fmt.Println(m.Tick())
    fmt.Println(m.Diff(month.Month(201812)))
    fmt.Println(m.Diff(month.Month(201711)))
    fmt.Println(m.Quarter())
    fmt.Println(m.String())
    fmt.Println(m.Format("YY年MM月"))
    fmt.Println(m.Format("yyyy年mm月"))
    fmt.Println(m.Format("mm月yyyy年"))
    fmt.Println(m.Format("yyyy-mm yyyy-mm"))
}

// 集合遍历
func Ranges() {
    month.Range(month.Month(202002), month.Month(202103), func(m month.Month) {
        fmt.Printf("%d - %d\n", m.Year(), m.Month())
    })
    fmt.Println(month.Span(month.Month(202002), month.Month(202103)))
}

```
