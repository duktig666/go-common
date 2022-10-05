package res

// code
const (
	SUCCESS_CODE = "I2000"
	ERROR_CODE   = "E3000"
)

// msg
const (
	SUCCESS_MSG   = "Success"
	ERROR_MSG     = "Server Error"
	PARAM_MISSING = "param missing"
)

type resCode struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}
