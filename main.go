// description:
// @author renshiwei
// Date: 2022/10/5 17:59

package main

import (
	"github.com/qiaoshurui/couples-subtotal/cmd/api"
	"github.com/qiaoshurui/couples-subtotal/common/initialize"
)

func main() {
	initialize.InitServer()
	api.Run()
}
