package req

import (
	"errors"
	"fmt"
	"net/http"

	errs "github.com/ddan1l/tega-backend/errors"
	res "github.com/ddan1l/tega-backend/web/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func BindAndValidate(c *gin.Context, obj interface{}) bool {
	validate = validator.New(validator.WithRequiredStructEnabled())

	c.ShouldBindJSON(&obj)

	err := validate.Struct(obj)

	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		var invalidValidationError *validator.InvalidValidationError

		if errors.As(err, &invalidValidationError) {
			return errorWith(c, http.StatusBadRequest, "BAD_REQUEST", err.Error(), nil)
		}

		var validateErrs validator.ValidationErrors

		var details = make(map[string]string)

		if errors.As(err, &validateErrs) {
			for _, e := range validateErrs {
				// e.Namespace()
				// e.Field()
				// e.StructNamespace()
				// e.StructField()
				// e.Tag()
				// e.ActualTag()
				// e.Kind()
				// e.Type()
				// e.Value()
				// e.Param()

				if len(e.Param()) > 0 {
					details[e.Field()] = fmt.Sprintf("Must be %s %s", e.Tag(), e.Param())
				} else {
					details[e.Field()] = fmt.Sprintf("Must be %s", e.Tag())
				}

			}
		}

		return errorWith(c, http.StatusUnprocessableEntity, "VALIDATION_ERROR", "Invalid input", details)
	}

	return true
}

func errorWith(c *gin.Context, status int, code, message string, details interface{}) bool {
	res.ErrorResponse(c, &errs.AppError{
		Status:  status,
		Code:    code,
		Message: message,
		Details: details,
	})
	return false
}
