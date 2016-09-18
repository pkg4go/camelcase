package camelcase

import . "github.com/pkg4go/assert"
import "testing"

func TestReverse(t *testing.T) {
	a := A{t}

	a.Equal(Reverse("CAMELCASE"), "camelcase")
	a.Equal(Reverse("CAMElCASE"), "cam_el_case")
	a.Equal(Reverse("CAmelCAse"), "c_amel_c_ase")
	a.Equal(Reverse("Camel0case"), "camel0case")
	a.Equal(Reverse("Camel0Case"), "camel0_case")
	a.Equal(Reverse("Camelcase0"), "camelcase0")
	a.Equal(Reverse("Camelcase"), "camelcase")
	a.Equal(Reverse("CamelCase"), "camel_case")
	a.Equal(Reverse("camelCase"), "camel_case")
}
