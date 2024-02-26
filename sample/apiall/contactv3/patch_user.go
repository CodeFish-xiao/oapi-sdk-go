// Package contact code generated by oapi sdk gen
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
	"github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
)

// PATCH /open-apis/contact/v3/users/:user_id
func main() {
	// 创建 Client
	client := lark.NewClient("appID", "appSecret")
	// 创建请求对象
	req := larkcontact.NewPatchUserReqBuilder().
		UserId("ou_7dab8a3d3cdcc9da365777c7ad535d62").
		UserIdType("open_id").
		DepartmentIdType("open_department_id").
		User(larkcontact.NewUserBuilder().
			Name("张三").
			EnName("San Zhang").
			Nickname("Alex Zhang").
			Email("zhangsan@gmail.com").
			Mobile("+41446681800").
			MobileVisible(false).
			Gender(0).
			AvatarKey("2500c7a9-5fff-4d9a-a2de-3d59614ae28g").
			DepartmentIds([]string{}).
			LeaderUserId("ou_7dab8a3d3cdcc9da365777c7ad535d62").
			City("杭州").
			Country("CN").
			WorkStation("北楼-H34").
			JoinTime(2147483647).
			EmployeeNo("1").
			EmployeeType(1).
			Orders([]*larkcontact.UserOrder{larkcontact.NewUserOrderBuilder().Build()}).
			CustomAttrs([]*larkcontact.UserCustomAttr{larkcontact.NewUserCustomAttrBuilder().Build()}).
			EnterpriseEmail("demo@mail.com").
			JobTitle("xxxxx").
			IsFrozen(false).
			JobLevelId("mga5oa8ayjlp9rb").
			JobFamilyId("mga5oa8ayjlp9rb").
			SubscriptionIds([]string{}).
			DottedLineLeaderUserIds([]string{}).
			Build()).
		Build()
	// 发起请求
	resp, err := client.Contact.V3.User.Patch(context.Background(), req)

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
