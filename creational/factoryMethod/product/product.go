/**
 * Package product
 * @Author fengfeng.mei <Biophiliam@protonmail.com>
 * @Date 2024/12/22 16:28
 */

package product

import "fmt"

type Product interface {
	Use()
}

type (
	Product1 struct {
		Name      string
		CreatDate string
	}

	Product2 struct {
		Name      string
		CreatDate string
		FixData   string
	}
)

func (p *Product1) Use() {
	fmt.Println("name:", p.Name, "creatDate:", p.CreatDate)
}

func (p *Product2) Use() {
	fmt.Println("name:", p.Name, "creatDate:", p.CreatDate, "fixData:", p.FixData)
}
