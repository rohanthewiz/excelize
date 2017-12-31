![Excelize](./excelize.png "Excelize")

# Excelize

[![Build Status](https://travis-ci.org/360EntSecGroup-Skylar/excelize.svg?branch=master)](https://travis-ci.org/360EntSecGroup-Skylar/excelize)
[![Code Coverage](https://codecov.io/gh/360EntSecGroup-Skylar/excelize/branch/master/graph/badge.svg)](https://codecov.io/gh/360EntSecGroup-Skylar/excelize)
[![Go Report Card](https://goreportcard.com/badge/github.com/360EntSecGroup-Skylar/excelize)](https://goreportcard.com/report/github.com/360EntSecGroup-Skylar/excelize)
[![GoDoc](https://godoc.org/github.com/360EntSecGroup-Skylar/excelize?status.svg)](https://godoc.org/github.com/360EntSecGroup-Skylar/excelize)
[![Licenses](https://img.shields.io/badge/license-bsd-orange.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.me/xuri)

## About this fork
This fork will diverge from 360EntSecGroup-Skylar/excelize mainly because I intend to go a different direction with the API at least for options and formats particularly for charts. The original project provides only a JSON interface to chart options and formatting. I plan to make public all the format structs. This provides explicitness and compiler checks at the caller, and also avoids any JSON marshalling and unmarshalling in the Chart API.
Therefore, the API here *will* change. For production please continue to use 360EntSecGroup-Skylar/excelize. I will be adding a release once things are stable.

## Introduction

Excelize is a library written in pure Golang and providing a set of functions that allow you to write to and read from XLSX files. Support reads and writes XLSX file generated by Microsoft Excel™ 2007 and later. Support save file without losing original charts of XLSX. This library needs Go version 1.8 or later. The full API docs can be seen using go's built-in documentation tool, or online at [godoc.org](https://godoc.org/github.com/360EntSecGroup-Skylar/excelize).

## Basic Usage

### Installation

```go
go get github.com/360EntSecGroup-Skylar/excelize
```

### Create XLSX file

Here is a minimal example usage that will create XLSX file.

```go
package main

import (
    "fmt"

    "github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
    xlsx := excelize.NewFile()
    // Create a new sheet.
    index := xlsx.NewSheet("Sheet2")
    // Set value of a cell.
    xlsx.SetCellValue("Sheet2", "A2", "Hello world.")
    xlsx.SetCellValue("Sheet1", "B2", 100)
    // Set active sheet of the workbook.
    xlsx.SetActiveSheet(index)
    // Save xlsx file by the given path.
    err := xlsx.SaveAs("./Workbook.xlsx")
    if err != nil {
        fmt.Println(err)
    }
}
```

### Reading XLSX file

The following constitutes the bare to read a XLSX document.

```go
package main

import (
    "fmt"

    "github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
    xlsx, err := excelize.OpenFile("./Workbook.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    // Get value from cell by given worksheet name and axis.
    cell := xlsx.GetCellValue("Sheet1", "B2")
    fmt.Println(cell)
    // Get all the rows in the Sheet1.
    rows := xlsx.GetRows("Sheet1")
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Print(colCell, "\t")
        }
        fmt.Println()
    }
}

```

### Add chart to XLSX file

With Excelize chart generation and management is as easy as a few lines of code. You can build charts based off data in your worksheet or generate charts without any data in your worksheet at all.

![Excelize](./test/images/chart.png "Excelize")

```go
package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	categories := map[string]string{"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	xlsx := excelize.NewFile()
	for k, v := range categories {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	xlsx.AddChart("Sheet1", "E1", `{"type":"bar3D","series":[{"name":"=Sheet1!$A$2","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$2:$D$2"},{"name":"=Sheet1!$A$3","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$3:$D$3"},{"name":"=Sheet1!$A$4","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$4:$D$4"}],"title":{"name":"Fruit 3D Line Chart"}}`)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./Workbook.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
```

### Add picture to XLSX file

```go
package main

import (
    "fmt"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"

    "github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
    xlsx, err := excelize.OpenFile("./Workbook.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    // Insert a picture.
    err = xlsx.AddPicture("Sheet1", "A2", "./image1.png", "")
    if err != nil {
        fmt.Println(err)
    }
    // Insert a picture to worksheet with scaling.
    err = xlsx.AddPicture("Sheet1", "D2", "./image2.jpg", `{"x_scale": 0.5, "y_scale": 0.5}`)
    if err != nil {
        fmt.Println(err)
    }
    // Insert a picture offset in the cell with printing support.
    err = xlsx.AddPicture("Sheet1", "H2", "./image3.gif", `{"x_offset": 15, "y_offset": 10, "print_obj": true, "lock_aspect_ratio": false, "locked": false}`)
    if err != nil {
        fmt.Println(err)
    }
    // Save the xlsx file with the origin path.
    err = xlsx.Save()
    if err != nil {
        fmt.Println(err)
    }
}
```

## Contributing

Contributions are welcome! Open a pull request to fix a bug, or open an issue to discuss a new feature or change. XML is compliant with [part 1 of the 5th edition of the ECMA-376 Standard for Office Open XML](http://www.ecma-international.org/publications/standards/Ecma-376.htm).

## Credits

Some struct of XML originally by [tealeg/xlsx](https://github.com/tealeg/xlsx).

## Licenses

This program is under the terms of the BSD 3-Clause License. See [https://opensource.org/licenses/BSD-3-Clause](https://opensource.org/licenses/BSD-3-Clause).
