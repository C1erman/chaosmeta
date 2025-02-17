/*
 * Copyright 2022-2023 Chaos Meta Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package handler

import (
	"context"
	"github.com/traas-stack/chaosmeta/chaosmetad/pkg/version"
	"github.com/traas-stack/chaosmeta/chaosmetad/pkg/web/model"
	"net/http"
)

func VersionGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	ctx := context.Background()
	WriteResponse(ctx, w, &model.VersionResponse{
		Code:    0,
		Message: "success",
		Data:    version.GetVersion(),
	})
}
