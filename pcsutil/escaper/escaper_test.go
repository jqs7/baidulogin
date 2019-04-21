package escaper_test

import (
	"fmt"
	"testing"

	"github.com/jqs7/baidulogin/pcsutil/escaper"
)

func TestEscape(t *testing.T) {
	fmt.Println(escaper.Escape(`asdf'asdfasd[]a[\[][sdf\[d]`, []rune{'[', '\''}))
}
