/**
 * Package abstractFactory
 * @Author fengfeng.mei <Biophiliam@protonmail.com>
 * @Date 2024/12/22 16:54
 */

package main

import "af/factory"

func main() {
	f := factory.NewConFn("ConF1")
	product := f.CreateProduct()
	product.GetName()
}
