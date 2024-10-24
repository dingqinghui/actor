/**
 * @Author: dingQingHui
 * @Description:
 * @File: error
 * @Version: 1.0.0
 * @Date: 2024/10/24 10:34
 */

package actor

import "errors"

var (
	ErrMailBoxNil        = errors.New("mailbox is nil")
	ErrActorStopped      = errors.New("actor is stopped")
	ErrActorReceiveIsNil = errors.New("actor receive is nil")
)
