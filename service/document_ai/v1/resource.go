// Code generated by Lark OpenAPI.

package larkdocument_ai

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

type V1 struct {
	BankCard                *bankCard                // bank_card
	BusinessCard            *businessCard            // business_card
	BusinessLicense         *businessLicense         // business_license
	ChinesePassport         *chinesePassport         // chinese_passport
	Contract                *contract                // contract
	DrivingLicense          *drivingLicense          // driving_license
	FoodManageLicense       *foodManageLicense       // food_manage_license
	FoodProduceLicense      *foodProduceLicense      // food_produce_license
	HealthCertificate       *healthCertificate       // health_certificate
	HkmMainlandTravelPermit *hkmMainlandTravelPermit // hkm_mainland_travel_permit
	IdCard                  *idCard                  // id_card
	TaxiInvoice             *taxiInvoice             // taxi_invoice
	TrainInvoice            *trainInvoice            // train_invoice
	TwMainlandTravelPermit  *twMainlandTravelPermit  // tw_mainland_travel_permit
	VatInvoice              *vatInvoice              // vat_invoice
	VehicleInvoice          *vehicleInvoice          // vehicle_invoice
	VehicleLicense          *vehicleLicense          // vehicle_license
}

func New(config *larkcore.Config) *V1 {
	return &V1{
		BankCard:                &bankCard{config: config},
		BusinessCard:            &businessCard{config: config},
		BusinessLicense:         &businessLicense{config: config},
		ChinesePassport:         &chinesePassport{config: config},
		Contract:                &contract{config: config},
		DrivingLicense:          &drivingLicense{config: config},
		FoodManageLicense:       &foodManageLicense{config: config},
		FoodProduceLicense:      &foodProduceLicense{config: config},
		HealthCertificate:       &healthCertificate{config: config},
		HkmMainlandTravelPermit: &hkmMainlandTravelPermit{config: config},
		IdCard:                  &idCard{config: config},
		TaxiInvoice:             &taxiInvoice{config: config},
		TrainInvoice:            &trainInvoice{config: config},
		TwMainlandTravelPermit:  &twMainlandTravelPermit{config: config},
		VatInvoice:              &vatInvoice{config: config},
		VehicleInvoice:          &vehicleInvoice{config: config},
		VehicleLicense:          &vehicleLicense{config: config},
	}
}

type bankCard struct {
	config *larkcore.Config
}
type businessCard struct {
	config *larkcore.Config
}
type businessLicense struct {
	config *larkcore.Config
}
type chinesePassport struct {
	config *larkcore.Config
}
type contract struct {
	config *larkcore.Config
}
type drivingLicense struct {
	config *larkcore.Config
}
type foodManageLicense struct {
	config *larkcore.Config
}
type foodProduceLicense struct {
	config *larkcore.Config
}
type healthCertificate struct {
	config *larkcore.Config
}
type hkmMainlandTravelPermit struct {
	config *larkcore.Config
}
type idCard struct {
	config *larkcore.Config
}
type taxiInvoice struct {
	config *larkcore.Config
}
type trainInvoice struct {
	config *larkcore.Config
}
type twMainlandTravelPermit struct {
	config *larkcore.Config
}
type vatInvoice struct {
	config *larkcore.Config
}
type vehicleInvoice struct {
	config *larkcore.Config
}
type vehicleLicense struct {
	config *larkcore.Config
}

// Recognize
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=recognize&project=document_ai&resource=bank_card&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/document_aiv1/recognize_bankCard.go
func (b *bankCard) Recognize(ctx context.Context, req *RecognizeBankCardReq, options ...larkcore.RequestOptionFunc) (*RecognizeBankCardResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/document_ai/v1/bank_card/recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, b.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RecognizeBankCardResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, b.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Recognize
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=recognize&project=document_ai&resource=business_card&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/document_aiv1/recognize_businessCard.go
func (b *businessCard) Recognize(ctx context.Context, req *RecognizeBusinessCardReq, options ...larkcore.RequestOptionFunc) (*RecognizeBusinessCardResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/document_ai/v1/business_card/recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, b.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RecognizeBusinessCardResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, b.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Recognize
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=recognize&project=document_ai&resource=business_license&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/document_aiv1/recognize_businessLicense.go
func (b *businessLicense) Recognize(ctx context.Context, req *RecognizeBusinessLicenseReq, options ...larkcore.RequestOptionFunc) (*RecognizeBusinessLicenseResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/document_ai/v1/business_license/recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, b.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RecognizeBusinessLicenseResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, b.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Recognize
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=recognize&project=document_ai&resource=chinese_passport&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/document_aiv1/recognize_chinesePassport.go
func (c *chinesePassport) Recognize(ctx context.Context, req *RecognizeChinesePassportReq, options ...larkcore.RequestOptionFunc) (*RecognizeChinesePassportResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/document_ai/v1/chinese_passport/recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, c.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RecognizeChinesePassportResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, c.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// FieldExtraction
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=field_extraction&project=document_ai&resource=contract&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/document_aiv1/fieldExtraction_contract.go
func (c *contract) FieldExtraction(ctx context.Context, req *FieldExtractionContractReq, options ...larkcore.RequestOptionFunc) (*FieldExtractionContractResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/document_ai/v1/contract/field_extraction"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, c.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &FieldExtractionContractResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, c.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Recognize
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=recognize&project=document_ai&resource=driving_license&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/document_aiv1/recognize_drivingLicense.go
func (d *drivingLicense) Recognize(ctx context.Context, req *RecognizeDrivingLicenseReq, options ...larkcore.RequestOptionFunc) (*RecognizeDrivingLicenseResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/document_ai/v1/driving_license/recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, d.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RecognizeDrivingLicenseResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, d.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Recognize
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=recognize&project=document_ai&resource=food_manage_license&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/document_aiv1/recognize_foodManageLicense.go
func (f *foodManageLicense) Recognize(ctx context.Context, req *RecognizeFoodManageLicenseReq, options ...larkcore.RequestOptionFunc) (*RecognizeFoodManageLicenseResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/document_ai/v1/food_manage_license/recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, f.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RecognizeFoodManageLicenseResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, f.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Recognize
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=recognize&project=document_ai&resource=food_produce_license&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/document_aiv1/recognize_foodProduceLicense.go
func (f *foodProduceLicense) Recognize(ctx context.Context, req *RecognizeFoodProduceLicenseReq, options ...larkcore.RequestOptionFunc) (*RecognizeFoodProduceLicenseResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/document_ai/v1/food_produce_license/recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, f.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RecognizeFoodProduceLicenseResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, f.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Recognize
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=recognize&project=document_ai&resource=health_certificate&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/document_aiv1/recognize_healthCertificate.go
func (h *healthCertificate) Recognize(ctx context.Context, req *RecognizeHealthCertificateReq, options ...larkcore.RequestOptionFunc) (*RecognizeHealthCertificateResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/document_ai/v1/health_certificate/recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, h.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RecognizeHealthCertificateResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, h.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Recognize
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=recognize&project=document_ai&resource=hkm_mainland_travel_permit&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/document_aiv1/recognize_hkmMainlandTravelPermit.go
func (h *hkmMainlandTravelPermit) Recognize(ctx context.Context, req *RecognizeHkmMainlandTravelPermitReq, options ...larkcore.RequestOptionFunc) (*RecognizeHkmMainlandTravelPermitResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/document_ai/v1/hkm_mainland_travel_permit/recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, h.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RecognizeHkmMainlandTravelPermitResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, h.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Recognize
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=recognize&project=document_ai&resource=id_card&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/document_aiv1/recognize_idCard.go
func (i *idCard) Recognize(ctx context.Context, req *RecognizeIdCardReq, options ...larkcore.RequestOptionFunc) (*RecognizeIdCardResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/document_ai/v1/id_card/recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RecognizeIdCardResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Recognize
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=recognize&project=document_ai&resource=taxi_invoice&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/document_aiv1/recognize_taxiInvoice.go
func (t *taxiInvoice) Recognize(ctx context.Context, req *RecognizeTaxiInvoiceReq, options ...larkcore.RequestOptionFunc) (*RecognizeTaxiInvoiceResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/document_ai/v1/taxi_invoice/recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, t.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RecognizeTaxiInvoiceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Recognize
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=recognize&project=document_ai&resource=train_invoice&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/document_aiv1/recognize_trainInvoice.go
func (t *trainInvoice) Recognize(ctx context.Context, req *RecognizeTrainInvoiceReq, options ...larkcore.RequestOptionFunc) (*RecognizeTrainInvoiceResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/document_ai/v1/train_invoice/recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, t.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RecognizeTrainInvoiceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Recognize
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=recognize&project=document_ai&resource=tw_mainland_travel_permit&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/document_aiv1/recognize_twMainlandTravelPermit.go
func (t *twMainlandTravelPermit) Recognize(ctx context.Context, req *RecognizeTwMainlandTravelPermitReq, options ...larkcore.RequestOptionFunc) (*RecognizeTwMainlandTravelPermitResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/document_ai/v1/tw_mainland_travel_permit/recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, t.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RecognizeTwMainlandTravelPermitResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Recognize
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=recognize&project=document_ai&resource=vat_invoice&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/document_aiv1/recognize_vatInvoice.go
func (v *vatInvoice) Recognize(ctx context.Context, req *RecognizeVatInvoiceReq, options ...larkcore.RequestOptionFunc) (*RecognizeVatInvoiceResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/document_ai/v1/vat_invoice/recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, v.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RecognizeVatInvoiceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, v.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Recognize
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=recognize&project=document_ai&resource=vehicle_invoice&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/document_aiv1/recognize_vehicleInvoice.go
func (v *vehicleInvoice) Recognize(ctx context.Context, req *RecognizeVehicleInvoiceReq, options ...larkcore.RequestOptionFunc) (*RecognizeVehicleInvoiceResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/document_ai/v1/vehicle_invoice/recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, v.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RecognizeVehicleInvoiceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, v.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Recognize
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=recognize&project=document_ai&resource=vehicle_license&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/document_aiv1/recognize_vehicleLicense.go
func (v *vehicleLicense) Recognize(ctx context.Context, req *RecognizeVehicleLicenseReq, options ...larkcore.RequestOptionFunc) (*RecognizeVehicleLicenseResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/document_ai/v1/vehicle_license/recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, v.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RecognizeVehicleLicenseResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, v.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
