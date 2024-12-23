/**
 * Package factoryMethod
 * @Author fengfeng.mei <Biophiliam@protonmail.com>
 * @Date 2024/12/22 16:26
 */

package main

import (
	"fm/factory"
	"fmt"
)

func main() {
	f1 := &factory.F1{Name: "f1"}
	f2 := &factory.F2{Name: "f2"}

	fmt.Println("factory have their product")
	fmt.Println("f1 product:")
	f1.CreatProduct().Use()
	fmt.Println("f2 product:")
	f2.CreatProduct().Use()
}
