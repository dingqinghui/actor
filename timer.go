/**
 * @Author: dingQingHui
 * @Description:
 * @File: timer
 * @Version: 1.0.0
 * @Date: 2024/11/8 18:18
 */

package actor

import (
	"github.com/RussellLuo/timingwheel"
	"time"
)

var (
	tw = timingwheel.NewTimingWheel(10*time.Millisecond, 3600)
)

func init() {
	tw.Start()
}
