/**
 * Package factory
 * @Author fengfeng.mei <Biophiliam@protonmail.com>
 * @Date 2024/12/22 16:28
 */

package factory

import (
	"fm/product"
)

type Factory interface {
	CreatProduct() product.Product
}

type (
	F1 struct {
		Name string
	}

	F2 struct {
		Name string
	}
)

func (f *F1) CreatProduct() product.Product {
	return &product.Product1{
		Name:      "p1",
		CreatDate: "2024-12-22",
	}
}

func (f *F2) CreatProduct() product.Product {
	return &product.Product2{
		Name:      "p2",
		CreatDate: "2024-12-23",
		FixData:   "no fix",
	}
}
