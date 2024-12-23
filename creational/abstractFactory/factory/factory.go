/**
 * Package factory
 * @Author fengfeng.mei <Biophiliam@protonmail.com>
 * @Date 2024/12/22 16:55
 */

package factory

import "af/product"

type (
	AbstractFactory interface {
		CreateProduct() product.AbstractProduct
	}

	ConF1 struct {
		Name string
	}

	ConF2 struct {
		Name string
	}
)

func NewConFn(fnName string) AbstractFactory {
	switch fnName {
	case "ConF1":
		return &ConF1{
			Name: "ConF1",
		}
	case "ConF2":
		return &ConF2{
			Name: "ConF2",
		}
	}
	return nil
}

func (c *ConF1) CreateProduct() product.AbstractProduct {
	return &product.P1{
		Name: "p1",
	}
}

func (c *ConF2) CreateProduct() product.AbstractProduct {
	return &product.P2{
		Name: "p2",
	}
}
