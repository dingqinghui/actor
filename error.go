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
	ErrMailBoxNil              = errors.New("mailbox is nil")
	ErrActorStopped            = errors.New("actor is stopped")
	ErrActorRespondEnvIsNil    = errors.New("actor respond env is nil")
	ErrActorRespondSenderIsNil = errors.New("actor respond sender is nil")
	ErrActorCallTimeout        = errors.New("actor call timeout")
)
