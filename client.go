//  code generated by oapi sdk gen
package lark

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/acs/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/admin/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/application/v6"
	"github.com/larksuite/oapi-sdk-go/v3/service/approval/v4"
	"github.com/larksuite/oapi-sdk-go/v3/service/attendance/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/baike/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/calendar/v4"
	"github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
	"github.com/larksuite/oapi-sdk-go/v3/service/docx/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/drive/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/ehr/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/event/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/gray_test_open_sg/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/human_authentication/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/mail/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/optical_char_recognition/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/passport/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/search/v2"
	"github.com/larksuite/oapi-sdk-go/v3/service/sheets/v3"
	"github.com/larksuite/oapi-sdk-go/v3/service/speech_to_text/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/task/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/tenant/v2"
	"github.com/larksuite/oapi-sdk-go/v3/service/translation/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/vc/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/wiki/v2"
)

type Client struct {
	config                 *larkcore.Config
	Acs                    *larkacs.AcsService
	Admin                  *larkadmin.AdminService
	Application            *larkapplication.ApplicationService
	Approval               *larkapproval.ApprovalService
	Attendance             *larkattendance.AttendanceService
	Baike                  *larkbaike.BaikeService
	Bitable                *larkbitable.BitableService
	Calendar               *larkcalendar.CalendarService
	Contact                *larkcontact.ContactService
	Docx                   *larkdocx.DocxService
	Drive                  *larkdrive.DriveService
	Ehr                    *larkehr.EhrService
	Event                  *larkevent.EventService
	GrayTestOpenSg         *larkgray_test_open_sg.GrayTestOpenSgService
	HumanAuthentication    *larkhuman_authentication.HumanAuthenticationService
	Im                     *larkim.ImService
	Mail                   *larkmail.MailService
	OpticalCharRecognition *larkoptical_char_recognition.OpticalCharRecognitionService
	Passport               *larkpassport.PassportService
	Search                 *larksearch.SearchService
	Sheets                 *larksheets.SheetsService
	SpeechToText           *larkspeech_to_text.SpeechToTextService
	Task                   *larktask.TaskService
	Tenant                 *larktenant.TenantService
	Translation            *larktranslation.TranslationService
	Vc                     *larkvc.VcService
	Wiki                   *larkwiki.WikiService
}

type ClientOptionFunc func(config *larkcore.Config)

func WithAppType(appType larkcore.AppType) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.AppType = appType
	}
}

func WithMarketplaceApp() ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.AppType = larkcore.AppTypeMarketplace
	}
}

func WithEnableTokenCache(enableTokenCache bool) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.EnableTokenCache = enableTokenCache
	}
}

func WithLogger(logger larkcore.Logger) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.Logger = logger
	}
}

func WithOpenBaseUrl(baseUrl string) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.BaseUrl = baseUrl
	}
}

func WithLogLevel(logLevel larkcore.LogLevel) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.LogLevel = logLevel
	}
}

func WithTokenCache(cache larkcore.Cache) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.TokenCache = cache
	}
}

func WithLogReqRespInfoAtDebugLevel(printReqRespLog bool) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.LogReqRespInfoAtDebugLevel = printReqRespLog
	}
}

func WithHttpClient(httpClient larkcore.HttpClient) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.HttpClient = httpClient
	}
}

func WithHelpdeskCredential(helpdeskID, helpdeskToken string) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.HelpDeskId = helpdeskID
		config.HelpDeskToken = helpdeskToken
		if helpdeskID != "" && helpdeskToken != "" {
			config.HelpdeskAuthToken = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", helpdeskID, helpdeskToken)))
		}
	}
}

func WithReqTimeout(reqTimeout time.Duration) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.ReqTimeout = reqTimeout
	}
}

func NewClient(appId, appSecret string, options ...ClientOptionFunc) *Client {
	// 构建配置
	config := &larkcore.Config{
		BaseUrl:          FeishuBaseUrl,
		AppId:            appId,
		AppSecret:        appSecret,
		EnableTokenCache: true,
		AppType:          larkcore.AppTypeSelfBuilt,
	}
	for _, option := range options {
		option(config)
	}

	// 构建日志器
	larkcore.NewLogger(config)

	// 构建缓存
	larkcore.NewCache(config)

	// 创建httpclient
	larkcore.NewHttpClient(config)

	// 创建sdk-client，并初始化服务
	client := &Client{config: config}
	initService(client, config)
	return client
}

func initService(client *Client, config *larkcore.Config) {
	client.Acs = larkacs.NewService(config)
	client.Admin = larkadmin.NewService(config)
	client.Application = larkapplication.NewService(config)
	client.Approval = larkapproval.NewService(config)
	client.Attendance = larkattendance.NewService(config)
	client.Baike = larkbaike.NewService(config)
	client.Bitable = larkbitable.NewService(config)
	client.Calendar = larkcalendar.NewService(config)
	client.Contact = larkcontact.NewService(config)
	client.Docx = larkdocx.NewService(config)
	client.Drive = larkdrive.NewService(config)
	client.Ehr = larkehr.NewService(config)
	client.Event = larkevent.NewService(config)
	client.GrayTestOpenSg = larkgray_test_open_sg.NewService(config)
	client.HumanAuthentication = larkhuman_authentication.NewService(config)
	client.Im = larkim.NewService(config)
	client.Mail = larkmail.NewService(config)
	client.OpticalCharRecognition = larkoptical_char_recognition.NewService(config)
	client.Passport = larkpassport.NewService(config)
	client.Search = larksearch.NewService(config)
	client.Sheets = larksheets.NewService(config)
	client.SpeechToText = larkspeech_to_text.NewService(config)
	client.Task = larktask.NewService(config)
	client.Tenant = larktenant.NewService(config)
	client.Translation = larktranslation.NewService(config)
	client.Vc = larkvc.NewService(config)
	client.Wiki = larkwiki.NewService(config)

}

func (cli *Client) Post(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.RawResponse, error) {
	return larkcore.SendRequest(ctx, cli.config, http.MethodPost, httpPath, []larkcore.AccessTokenType{accessTokeType}, body, options...)
}

func (cli *Client) Get(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.RawResponse, error) {
	return larkcore.SendRequest(ctx, cli.config, http.MethodGet, httpPath, []larkcore.AccessTokenType{accessTokeType}, body, options...)
}

func (cli *Client) Delete(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.RawResponse, error) {
	return larkcore.SendRequest(ctx, cli.config, http.MethodDelete, httpPath, []larkcore.AccessTokenType{accessTokeType}, body, options...)
}

func (cli *Client) Put(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.RawResponse, error) {
	return larkcore.SendRequest(ctx, cli.config, http.MethodPut, httpPath, []larkcore.AccessTokenType{accessTokeType}, body, options...)
}

func (cli *Client) Patch(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.RawResponse, error) {
	return larkcore.SendRequest(ctx, cli.config, http.MethodPatch, httpPath, []larkcore.AccessTokenType{accessTokeType}, body, options...)
}

func (cli *Client) Head(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.RawResponse, error) {
	return larkcore.SendRequest(ctx, cli.config, http.MethodHead, httpPath, []larkcore.AccessTokenType{accessTokeType}, body, options...)
}

func (cli *Client) Options(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.RawResponse, error) {
	return larkcore.SendRequest(ctx, cli.config, http.MethodOptions, httpPath, []larkcore.AccessTokenType{accessTokeType}, body, options...)
}

func (cli *Client) Trace(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.RawResponse, error) {
	return larkcore.SendRequest(ctx, cli.config, http.MethodTrace, httpPath, []larkcore.AccessTokenType{accessTokeType}, body, options...)
}

var FeishuBaseUrl = "https://open.feishu.cn"
var LarkBaseUrl = "https://open.larksuite.com"
