/**
 * Package product
 * @Author fengfeng.mei <Biophiliam@protonmail.com>
 * @Date 2024/12/22 16:55
 */

package product

type (
	// AbstractProduct abstract product interface
	AbstractProduct interface {
		GetName() string
	}

	Product struct {
		Name2 string
	}

	// P1  product1
	P1 struct {
		Product
		Name string
	}

	// P2  product2
	P2 struct {
		Product
		Name string
	}
)

// GetName get product name
func (p *Product) GetName() string {
	return p.Name2
}
