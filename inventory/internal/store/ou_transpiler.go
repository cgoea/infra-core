// SPDX-FileCopyrightText: (C) 2025 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-filters, DO NOT EDIT.

package store

import (
	"github.com/open-edge-platform/infra-core/inventory/v2/internal/ent/ouresource"
	inv_v1 "github.com/open-edge-platform/infra-core/inventory/v2/pkg/api/inventory/v1"
)

func (r *registry) RegisterOuResource() {
	r.Register(
		newResourceTranspiler(
			inv_v1.ResourceKind_RESOURCE_KIND_OU,
			ouresource.ValidColumn,
			map[string]edgeHandler{
				ouresource.EdgeChildren: {
					func(p sqlPredicate) sqlPredicate { return ouresource.HasChildrenWith(p) },
					inv_v1.ResourceKind_RESOURCE_KIND_OU,
				},

				ouresource.EdgeParentOu: {
					func(p sqlPredicate) sqlPredicate { return ouresource.HasParentOuWith(p) },
					inv_v1.ResourceKind_RESOURCE_KIND_OU,
				},
			},
			map[string]sqlPredicate{
				ouresource.EdgeChildren: ouresource.HasChildren(),
				ouresource.EdgeParentOu: ouresource.HasParentOu(),
			},
		))

}
