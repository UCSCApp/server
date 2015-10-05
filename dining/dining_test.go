package dining

import (
	"fmt"
	"testing"
)

func TestSelectC8(t *testing.T) {
	fmt.Printf("%#v", c8Doc().selectMenuTable().parseDinnerMenu())
}
