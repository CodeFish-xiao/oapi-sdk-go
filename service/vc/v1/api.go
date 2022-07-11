// Package vc code generated by oapi sdk gen
package larkvc

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *VcService {
	v := &VcService{config: config}
	v.Meeting = &meeting{service: v}
	v.MeetingRecording = &meetingRecording{service: v}
	v.Report = &report{service: v}
	v.Reserve = &reserve{service: v}
	v.RoomConfig = &roomConfig{service: v}
	return v
}

// 业务域服务定义
type VcService struct {
	config           *larkcore.Config
	Meeting          *meeting
	MeetingRecording *meetingRecording
	Report           *report
	Reserve          *reserve
	RoomConfig       *roomConfig
}

// 资源服务定义
type meeting struct {
	service *VcService
}
type meetingRecording struct {
	service *VcService
}
type report struct {
	service *VcService
}
type reserve struct {
	service *VcService
}
type roomConfig struct {
	service *VcService
}

// 资源服务方法定义
func (m *meeting) End(ctx context.Context, req *EndMeetingReq, options ...larkcore.RequestOptionFunc) (*EndMeetingResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, m.service.config, http.MethodPatch,
		"/open-apis/vc/v1/meetings/:meeting_id/end", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &EndMeetingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meeting) Get(ctx context.Context, req *GetMeetingReq, options ...larkcore.RequestOptionFunc) (*GetMeetingResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, m.service.config, http.MethodGet,
		"/open-apis/vc/v1/meetings/:meeting_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetMeetingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meeting) Invite(ctx context.Context, req *InviteMeetingReq, options ...larkcore.RequestOptionFunc) (*InviteMeetingResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, m.service.config, http.MethodPatch,
		"/open-apis/vc/v1/meetings/:meeting_id/invite", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &InviteMeetingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meeting) Kickout(ctx context.Context, req *KickoutMeetingReq, options ...larkcore.RequestOptionFunc) (*KickoutMeetingResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, m.service.config, http.MethodPost,
		"/open-apis/vc/v1/meetings/:meeting_id/kickout", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &KickoutMeetingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meeting) ListByNo(ctx context.Context, req *ListByNoMeetingReq, options ...larkcore.RequestOptionFunc) (*ListByNoMeetingResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, m.service.config, http.MethodGet,
		"/open-apis/vc/v1/meetings/list_by_no", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListByNoMeetingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meeting) ListByNoByIterator(ctx context.Context, req *ListByNoMeetingReq, options ...larkcore.RequestOptionFunc) (*ListByNoMeetingIterator, error) {
	return &ListByNoMeetingIterator{
		ctx:      ctx,
		req:      req,
		listFunc: m.ListByNo,
		options:  options,
		limit:    req.Limit}, nil
}
func (m *meeting) SetHost(ctx context.Context, req *SetHostMeetingReq, options ...larkcore.RequestOptionFunc) (*SetHostMeetingResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, m.service.config, http.MethodPatch,
		"/open-apis/vc/v1/meetings/:meeting_id/set_host", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SetHostMeetingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meetingRecording) Get(ctx context.Context, req *GetMeetingRecordingReq, options ...larkcore.RequestOptionFunc) (*GetMeetingRecordingResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, m.service.config, http.MethodGet,
		"/open-apis/vc/v1/meetings/:meeting_id/recording", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetMeetingRecordingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meetingRecording) SetPermission(ctx context.Context, req *SetPermissionMeetingRecordingReq, options ...larkcore.RequestOptionFunc) (*SetPermissionMeetingRecordingResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, m.service.config, http.MethodPatch,
		"/open-apis/vc/v1/meetings/:meeting_id/recording/set_permission", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SetPermissionMeetingRecordingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meetingRecording) Start(ctx context.Context, req *StartMeetingRecordingReq, options ...larkcore.RequestOptionFunc) (*StartMeetingRecordingResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, m.service.config, http.MethodPatch,
		"/open-apis/vc/v1/meetings/:meeting_id/recording/start", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &StartMeetingRecordingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meetingRecording) Stop(ctx context.Context, req *StopMeetingRecordingReq, options ...larkcore.RequestOptionFunc) (*StopMeetingRecordingResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, m.service.config, http.MethodPatch,
		"/open-apis/vc/v1/meetings/:meeting_id/recording/stop", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &StopMeetingRecordingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *report) GetDaily(ctx context.Context, req *GetDailyReportReq, options ...larkcore.RequestOptionFunc) (*GetDailyReportResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, r.service.config, http.MethodGet,
		"/open-apis/vc/v1/reports/get_daily", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetDailyReportResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *report) GetTopUser(ctx context.Context, req *GetTopUserReportReq, options ...larkcore.RequestOptionFunc) (*GetTopUserReportResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, r.service.config, http.MethodGet,
		"/open-apis/vc/v1/reports/get_top_user", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetTopUserReportResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *reserve) Apply(ctx context.Context, req *ApplyReserveReq, options ...larkcore.RequestOptionFunc) (*ApplyReserveResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, r.service.config, http.MethodPost,
		"/open-apis/vc/v1/reserves/apply", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ApplyReserveResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *reserve) Delete(ctx context.Context, req *DeleteReserveReq, options ...larkcore.RequestOptionFunc) (*DeleteReserveResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, r.service.config, http.MethodDelete,
		"/open-apis/vc/v1/reserves/:reserve_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteReserveResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *reserve) Get(ctx context.Context, req *GetReserveReq, options ...larkcore.RequestOptionFunc) (*GetReserveResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, r.service.config, http.MethodGet,
		"/open-apis/vc/v1/reserves/:reserve_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetReserveResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *reserve) GetActiveMeeting(ctx context.Context, req *GetActiveMeetingReserveReq, options ...larkcore.RequestOptionFunc) (*GetActiveMeetingReserveResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, r.service.config, http.MethodGet,
		"/open-apis/vc/v1/reserves/:reserve_id/get_active_meeting", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetActiveMeetingReserveResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *reserve) Update(ctx context.Context, req *UpdateReserveReq, options ...larkcore.RequestOptionFunc) (*UpdateReserveResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, r.service.config, http.MethodPut,
		"/open-apis/vc/v1/reserves/:reserve_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateReserveResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *roomConfig) Query(ctx context.Context, req *QueryRoomConfigReq, options ...larkcore.RequestOptionFunc) (*QueryRoomConfigResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, r.service.config, http.MethodGet,
		"/open-apis/vc/v1/room_configs/query", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryRoomConfigResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *roomConfig) Set(ctx context.Context, req *SetRoomConfigReq, options ...larkcore.RequestOptionFunc) (*SetRoomConfigResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, r.service.config, http.MethodPost,
		"/open-apis/vc/v1/room_configs/set", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SetRoomConfigResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
