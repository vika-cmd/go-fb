package validator

import (
	"strings"

	"github.com/gobuffalo/validate"
)

func FormatErrors(errs *validate.Errors) string{
	var resp string
	for k,v := range errs.Errors{
		resp = resp +k + ": " + strings.Join(v,",") + ", "
	}
	return resp
}