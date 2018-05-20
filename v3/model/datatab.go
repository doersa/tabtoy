package model

import (
	"github.com/davyxu/tabtoy/v3/table"
)

type DataRow []Cell

func (self DataRow) IsEmpty() bool {
	return len(self) == 0
}

type DataTable struct {
	HeaderType string // 表名，Index表里定义的类型

	OriginalHeaderType string // HeaderFields对应的ObjectType，KV表为TableField

	FileName string

	SheetName string

	Rows []DataRow

	RawHeader    DataRow
	HeaderFields []*table.TableField // 列索引
}

// 代码生成专用
func (self *DataTable) GetValue(row, col int) string {

	return self.Rows[row][col].Value
}

// 代码生成专用
func (self *DataTable) GetType(col int) *table.TableField {
	return self.HeaderFields[col]
}

func (self *DataTable) GetTypeByName(name string) (*table.TableField, int) {

	if name == "" {
		return nil, -1
	}

	for col, f := range self.HeaderFields {
		if f.Name == name || f.FieldName == name {
			return f, col
		}
	}

	return nil, -1
}

// 根据列头找到该行对应的值
func (self *DataTable) GetValueByName(row int, name string) (Cell, *table.TableField) {

	hf, col := self.GetTypeByName(name)

	if hf == nil {
		return Cell{}, nil
	}

	return self.Rows[row][col], hf

}

// 添加表头类型
func (self *DataTable) AddHeaderField(types *table.TableField) {
	self.HeaderFields = append(self.HeaderFields, types)
}

// 添加行数据
func (self *DataTable) AddRow(row DataRow) {

	self.Rows = append(self.Rows, row)
}

func NewDataTable() *DataTable {
	return &DataTable{}
}
