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

package memory

import (
	"time"

	"bkauth/pkg/cache/memory/backend"
)

// NewCache create a memory cache
func NewCache(name string, disabled bool,
	retrieveFunc RetrieveFunc,
	expiration time.Duration,
	randomDurationFunc backend.RandomExpirationDurationFunc,
) Cache {
	be := backend.NewMemoryBackend(name, expiration, randomDurationFunc)
	return NewBaseCache(disabled, retrieveFunc, be)
}

// NewMockCache create a memory cache for mock
func NewMockCache(retrieveFunc RetrieveFunc) Cache {
	be := backend.NewMemoryBackend("mockCache", 5*time.Minute, nil)

	return NewBaseCache(false, retrieveFunc, be)
}
