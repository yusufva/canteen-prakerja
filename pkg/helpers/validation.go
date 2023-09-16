package helpers

import (
	"canteen-prakerja/pkg/custerrs"

	"github.com/asaskevich/govalidator"
)

func ValidateStruct(payload interface{}) custerrs.MessageErr {
	_, err := govalidator.ValidateStruct(payload)

	if err != nil {
		return custerrs.NewBadRequest(err.Error())
	}

	return nil
}
