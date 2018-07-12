/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package operation

type opcondition struct {
	InstID []int64 `json:"inst_ids"`
}

type deleteCondition struct {
	opcondition `json:",inline"`
}

type updateCondition struct {
	InstID   int64                  `json:"inst_id"`
	InstInfo map[string]interface{} `json:"datas"`
}

// OpCondition the condition operation
type OpCondition struct {
	Delete deleteCondition   `json:"delete"`
	Update []updateCondition `json:"update"`
}

type instBatchInfo struct {
	// BatchInfo batch info
	BatchInfo *map[int]map[string]interface{} `json:"BatchInfo"`
	InputType string                          `json:"input_type"`
}
