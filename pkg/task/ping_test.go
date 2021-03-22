package task

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"time"
)

func TestExecuteICMPPing(t *testing.T) {
	param := ICMPPingTaskParam{
		DstAddr: 	"1.1.1.1",
		Interval: 	100 * time.Microsecond,
		Count: 		5,
	}
	_, err := ExecuteICMPPing(&param)
	assert.Equal(t, err, nil)
}
