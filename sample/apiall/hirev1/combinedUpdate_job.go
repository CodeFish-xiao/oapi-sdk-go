// Package hire code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"context"
	"fmt"

	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/hire/v1"
)

// POST /open-apis/hire/v1/jobs/:job_id/combined_update
func main() {
	// 创建 Client
	client := lark.NewClient("appID", "appSecret")
	// 创建请求对象
	req := larkhire.NewCombinedUpdateJobReqBuilder().
		JobId("6960663240925956660").
		UserIdType("open_id").
		DepartmentIdType("open_department_id").
		JobLevelIdType("people_admin_job_level_id").
		JobFamilyIdType("people_admin_job_category_id").
		CombinedJob(larkhire.NewCombinedJobBuilder().
			Id("6960663240925956576").
			Experience(1).
			ExpiryTime(1622484739955).
			CustomizedDataList([]*larkhire.CombinedJobObjectValueMap{larkhire.NewCombinedJobObjectValueMapBuilder().Build()}).
			MinLevelId("6960663240925956547").
			MinSalary(1000).
			Title("后端研发").
			JobManagers(larkhire.NewJobManagerBuilder().Build()).
			JobProcessId("6960663240925956554").
			SubjectId("6960663240925956555").
			JobFunctionId("6960663240925956555").
			DepartmentId("6960663240925956549").
			HeadCount(100).
			IsNeverExpired(false).
			MaxSalary(2000).
			Requirement("熟悉后端研发").
			Description("后端研发岗位描述").
			HighlightList([]string{}).
			JobTypeId("6960663240925956551").
			MaxLevelId("6960663240925956548").
			RequiredDegree(20).
			JobCategoryId("6960663240925956550").
			AddressIdList([]string{}).
			JobAttribute(1).
			ExpiryTimestamp("1622484739955").
			TargetMajorIdList([]string{}).
			Build()).
		Build()
	// 发起请求
	resp, err := client.Hire.V1.Job.CombinedUpdate(context.Background(), req)

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return
	}

	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
}
