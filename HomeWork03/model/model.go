package model

import "fmt"

type Manage interface {
	CheckIsOver(i int) bool
	Add(i int)
	Reduce(i int)
	PrintInfo()
}
type Goods struct {
	Name      string
	Price     float64
	Inventory int
}
type Electronics struct {
	Goods
	Brand string
	Type  string
}

func (g *Goods) CheckIsOver(i int) bool {
	return g.Inventory >= i
}
func (g *Goods) Add(i int) {
	g.Inventory += i
}
func (g *Goods) Reduce(i int) {
	g.Inventory -= i
}
func (g *Goods) PrintInfo() {
	fmt.Println("商品名: ", g.Name)
	fmt.Println("商品价格: ", g.Price)
	fmt.Println("商品库存: ", g.Inventory)
}
func (e *Electronics) PrintMore() {
	e.PrintInfo()
	fmt.Println("商品品牌: ", e.Brand)
	fmt.Println("商品型号: ", e.Type)
}
