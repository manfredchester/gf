// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

package ghttp

import (
    "io/ioutil"
)

// 获取返回的数据
func (r *ClientResponse) ReadAll() []byte {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return nil
    }
    return body
}

// 关闭返回的HTTP链接
func (r *ClientResponse) Close()  {
    r.Response.Close = true
    r.Body.Close()
}