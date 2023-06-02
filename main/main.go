// Copyright © 2023 OpenIM SDK. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"open_im_sdk/test"
	"time"
)

func main() {
	APIADDR := "http://203.56.175.233:10002"
	WSADDR := "ws://203.56.175.233:10001"
	REGISTERADDR := APIADDR + "/user_register"
	ACCOUNTCHECK := APIADDR + "/manager/account_check"
	TOKENADDR := APIADDR + "/auth/user_token"
	SECRET := "openIM123"
	SENDINTERVAL := 20
	test.REGISTERADDR = REGISTERADDR
	test.TOKENADDR = TOKENADDR
	test.SECRET = SECRET
	test.SENDINTERVAL = SENDINTERVAL
	test.WSADDR = WSADDR
	test.ACCOUNTCHECK = ACCOUNTCHECK
	strMyUidx := "9169012630"

	tokenx := test.RunGetToken(strMyUidx)
	fmt.Println(tokenx)
	test.InOutDoTest(strMyUidx, tokenx, WSADDR, APIADDR)
	test.DoTestRevoke()

	for {
		time.Sleep(10000 * time.Millisecond)
	}

}
