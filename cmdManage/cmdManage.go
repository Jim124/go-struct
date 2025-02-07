package cmdManage

import "fmt"

type CMDManage struct {
}

func New() CMDManage {
	return CMDManage{}
}

func (cm CMDManage) ReadLine() ([]string, error) {

	var lines []string
	for {
		var price float64
		fmt.Print("enter price:")
		fmt.Scan(&price)
		if price == 1 {
			break
		}
		fmt.Println(price)
		lines = append(lines, fmt.Sprintf("%.0f", price))
	}
	fmt.Println(lines)
	return lines, nil
}
func (cm CMDManage) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}
