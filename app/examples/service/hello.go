// description:
// @author renshiwei
// Date: 2022/10/5 21:16

package service

import "fmt"

type Hello struct{}

func (h *Hello) SayHello(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
