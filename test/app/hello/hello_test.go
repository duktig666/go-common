// description:
// @author renshiwei
// Date: 2022/10/5 22:42

package hello

import (
	"fmt"
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	//initialize.InitServer()
	//api.Run()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	global.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println(w.Body.String())
}

func TestFmt(t *testing.T) {
	fmt.Printf("name:%s\n", "Q")
}
