/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - Auth服务(BlueKing - Auth) available.
 * Copyright (C) 2017 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 *     http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 * to the current version of the project delivered to anyone in the future.
 */

package util

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// ErrNotArray ...
var ErrNotArray = errors.New("validate array fail, only support array")

// ToSlice ...
func ToSlice(array interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(array)
	if v.Kind() != reflect.Slice {
		return nil, ErrNotArray
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret, nil
}

// Int64SliceToString ...
func Int64SliceToString(s []int64, sep string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(s)), sep), "[]")
}

// StringToInt64 ...
func StringToInt64(i string) (int64, error) {
	return strconv.ParseInt(i, 10, 64)
}

// StringToInt64Slice ...
func StringToInt64Slice(s, sep string) ([]int64, error) {
	if s == "" {
		return []int64{}, nil
	}
	parts := strings.Split(s, sep)

	int64Slice := make([]int64, 0, len(parts))
	for _, d := range parts {
		i, err := strconv.ParseInt(d, 10, 64)
		if err != nil {
			return nil, err
		}
		int64Slice = append(int64Slice, i)
	}
	return int64Slice, nil
}
