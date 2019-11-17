package uuid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalJSON(t *testing.T) {
	var zero UUID

	enc, err := zero.MarshalJSON()
	if assert.Nil(t, err, fmt.Sprint(err)) {
		assert.Equal(t, `null`, string(enc))
	}

	p := &zero

	err = p.UnmarshalJSON([]byte(`null`))
	if assert.Nil(t, err, fmt.Sprint(err)) {
		assert.Equal(t, Zero, *p)
	}

	err = p.UnmarshalJSON([]byte(`0`))
	if assert.Nil(t, err, fmt.Sprint(err)) {
		assert.Equal(t, Zero, *p)
	}

	v := New()

	enc, err = v.MarshalJSON()
	if assert.Nil(t, err, fmt.Sprint(err)) {
		assert.Equal(t, `"`+v.String()+`"`, string(enc))
	}

	err = p.UnmarshalJSON([]byte(`"` + v.String() + `"`))
	if assert.Nil(t, err, fmt.Sprint(err)) {
		assert.Equal(t, v, *p)
	}

}
