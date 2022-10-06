// description:
// @author renshiwei
// Date: 2022/10/5 21:16

package service

import (
	"fmt"
	"github.com/pkg/errors"
)

type Hello struct{}

func (h *Hello) SayHello(name string) (string, error) {
	if name == "" {
		return "", errors.Errorf("param err:%s.", "name is empty")
	}
	return fmt.Sprintf("Hello %s!", name), nil
}
