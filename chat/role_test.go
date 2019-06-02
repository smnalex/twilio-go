package chat

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
)

type optionals interface {
	encode() io.Reader
}

var optionalsFn = func(m optionals, exp []byte) func(*testing.T) {
	return func(t *testing.T) {
		got, err := ioutil.ReadAll(m.encode())
		if err != nil {
			t.Errorf("exp parsing err, got %v", err)
		}
		if !bytes.Equal(got, exp) {
			t.Errorf("exp %s, got %s", exp, got)
		}
	}
}

func TestRoleParamsOptionals(t *testing.T) {
	exp := []byte("FriendlyName=&Type=")
	t.Run("CreateParams", optionalsFn(RoleCreateParams{}, exp))
	exp = []byte("")
	t.Run("UpdateParams", optionalsFn(RoleUpdateParams{}, exp))
}
