// Package application code generated by oapi sdk gen
package larkapplication

import (
	"context"
)

// 消息处理器定义
type P2ApplicationCreatedV6Handler struct {
	handler func(context.Context, *P2ApplicationCreatedV6) error
}

func NewP2ApplicationCreatedV6Handler(handler func(context.Context, *P2ApplicationCreatedV6) error) *P2ApplicationCreatedV6Handler {
	h := &P2ApplicationCreatedV6Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ApplicationCreatedV6Handler) Event() interface{} {
	return &P2ApplicationCreatedV6{}
}

// 回调开发者注册的handle
func (h *P2ApplicationCreatedV6Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ApplicationCreatedV6))
}

// 消息处理器定义
type P2ApplicationAppVersionAuditV6Handler struct {
	handler func(context.Context, *P2ApplicationAppVersionAuditV6) error
}

func NewP2ApplicationAppVersionAuditV6Handler(handler func(context.Context, *P2ApplicationAppVersionAuditV6) error) *P2ApplicationAppVersionAuditV6Handler {
	h := &P2ApplicationAppVersionAuditV6Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ApplicationAppVersionAuditV6Handler) Event() interface{} {
	return &P2ApplicationAppVersionAuditV6{}
}

// 回调开发者注册的handle
func (h *P2ApplicationAppVersionAuditV6Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ApplicationAppVersionAuditV6))
}

// 消息处理器定义
type P2ApplicationAppVersionPublishApplyV6Handler struct {
	handler func(context.Context, *P2ApplicationAppVersionPublishApplyV6) error
}

func NewP2ApplicationAppVersionPublishApplyV6Handler(handler func(context.Context, *P2ApplicationAppVersionPublishApplyV6) error) *P2ApplicationAppVersionPublishApplyV6Handler {
	h := &P2ApplicationAppVersionPublishApplyV6Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ApplicationAppVersionPublishApplyV6Handler) Event() interface{} {
	return &P2ApplicationAppVersionPublishApplyV6{}
}

// 回调开发者注册的handle
func (h *P2ApplicationAppVersionPublishApplyV6Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ApplicationAppVersionPublishApplyV6))
}

// 消息处理器定义
type P2ApplicationAppVersionPublishRevokeV6Handler struct {
	handler func(context.Context, *P2ApplicationAppVersionPublishRevokeV6) error
}

func NewP2ApplicationAppVersionPublishRevokeV6Handler(handler func(context.Context, *P2ApplicationAppVersionPublishRevokeV6) error) *P2ApplicationAppVersionPublishRevokeV6Handler {
	h := &P2ApplicationAppVersionPublishRevokeV6Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ApplicationAppVersionPublishRevokeV6Handler) Event() interface{} {
	return &P2ApplicationAppVersionPublishRevokeV6{}
}

// 回调开发者注册的handle
func (h *P2ApplicationAppVersionPublishRevokeV6Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ApplicationAppVersionPublishRevokeV6))
}

// 消息处理器定义
type P2ApplicationFeedbackCreatedV6Handler struct {
	handler func(context.Context, *P2ApplicationFeedbackCreatedV6) error
}

func NewP2ApplicationFeedbackCreatedV6Handler(handler func(context.Context, *P2ApplicationFeedbackCreatedV6) error) *P2ApplicationFeedbackCreatedV6Handler {
	h := &P2ApplicationFeedbackCreatedV6Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ApplicationFeedbackCreatedV6Handler) Event() interface{} {
	return &P2ApplicationFeedbackCreatedV6{}
}

// 回调开发者注册的handle
func (h *P2ApplicationFeedbackCreatedV6Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ApplicationFeedbackCreatedV6))
}

// 消息处理器定义
type P2ApplicationFeedbackUpdatedV6Handler struct {
	handler func(context.Context, *P2ApplicationFeedbackUpdatedV6) error
}

func NewP2ApplicationFeedbackUpdatedV6Handler(handler func(context.Context, *P2ApplicationFeedbackUpdatedV6) error) *P2ApplicationFeedbackUpdatedV6Handler {
	h := &P2ApplicationFeedbackUpdatedV6Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ApplicationFeedbackUpdatedV6Handler) Event() interface{} {
	return &P2ApplicationFeedbackUpdatedV6{}
}

// 回调开发者注册的handle
func (h *P2ApplicationFeedbackUpdatedV6Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ApplicationFeedbackUpdatedV6))
}

// 消息处理器定义
type P2ApplicationVisibilityAddedV6Handler struct {
	handler func(context.Context, *P2ApplicationVisibilityAddedV6) error
}

func NewP2ApplicationVisibilityAddedV6Handler(handler func(context.Context, *P2ApplicationVisibilityAddedV6) error) *P2ApplicationVisibilityAddedV6Handler {
	h := &P2ApplicationVisibilityAddedV6Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ApplicationVisibilityAddedV6Handler) Event() interface{} {
	return &P2ApplicationVisibilityAddedV6{}
}

// 回调开发者注册的handle
func (h *P2ApplicationVisibilityAddedV6Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ApplicationVisibilityAddedV6))
}
