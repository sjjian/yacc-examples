package yacc_parseJson

import (
	"fmt"
	"testing"
)

func TestParseJson(t *testing.T) {
	d, err := ParseJson(`{"a":1,"b":"str","c":[1,"str"],"d":{"d1":1, "d2":"str"},"e":[],"f":{}}`, true)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(d) // map[a:1 b:str c:[1 str] d:map[d1:1 d2:str] e:[] f:map[]]
}
