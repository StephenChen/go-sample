package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type RegisterReq struct {
	// 字符串的 gt=0 表示长度必须 > 0，gt = greater than
	Username    string `validate:"gt=0"`
	PasswordNew string `validate:"gt=0"`
	// eqfield 跨字段相等校验
	PasswordRepeat string `validate:"eqfield=PasswordNew"`
	// 合法 email 格式校验
	Email string `validate:"email"`
}

var validate = validator.New()

func validateReq(req RegisterReq) error {
	err := validate.Struct(req)
	if err != nil {
		// doSomething
		return err
	}
	return nil
}

func Validate() {
	req := RegisterReq{
		Username:       "ranxi",
		PasswordNew:    "haha",
		PasswordRepeat: "hah",
		Email:          "ranxi@abc.com",
	}

	err := validateReq(req)
	fmt.Println(err)
}
