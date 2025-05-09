// Code generated by ent, DO NOT EDIT.

package ipaddressresource

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-edge-platform/infra-core/inventory/v2/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldLTE(FieldID, id))
}

// ResourceID applies equality check predicate on the "resource_id" field. It's identical to ResourceIDEQ.
func ResourceID(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldResourceID, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldAddress, v))
}

// StatusDetail applies equality check predicate on the "status_detail" field. It's identical to StatusDetailEQ.
func StatusDetail(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldStatusDetail, v))
}

// TenantID applies equality check predicate on the "tenant_id" field. It's identical to TenantIDEQ.
func TenantID(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldTenantID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldUpdatedAt, v))
}

// ResourceIDEQ applies the EQ predicate on the "resource_id" field.
func ResourceIDEQ(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldResourceID, v))
}

// ResourceIDNEQ applies the NEQ predicate on the "resource_id" field.
func ResourceIDNEQ(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNEQ(FieldResourceID, v))
}

// ResourceIDIn applies the In predicate on the "resource_id" field.
func ResourceIDIn(vs ...string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldIn(FieldResourceID, vs...))
}

// ResourceIDNotIn applies the NotIn predicate on the "resource_id" field.
func ResourceIDNotIn(vs ...string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNotIn(FieldResourceID, vs...))
}

// ResourceIDGT applies the GT predicate on the "resource_id" field.
func ResourceIDGT(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldGT(FieldResourceID, v))
}

// ResourceIDGTE applies the GTE predicate on the "resource_id" field.
func ResourceIDGTE(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldGTE(FieldResourceID, v))
}

// ResourceIDLT applies the LT predicate on the "resource_id" field.
func ResourceIDLT(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldLT(FieldResourceID, v))
}

// ResourceIDLTE applies the LTE predicate on the "resource_id" field.
func ResourceIDLTE(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldLTE(FieldResourceID, v))
}

// ResourceIDContains applies the Contains predicate on the "resource_id" field.
func ResourceIDContains(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldContains(FieldResourceID, v))
}

// ResourceIDHasPrefix applies the HasPrefix predicate on the "resource_id" field.
func ResourceIDHasPrefix(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldHasPrefix(FieldResourceID, v))
}

// ResourceIDHasSuffix applies the HasSuffix predicate on the "resource_id" field.
func ResourceIDHasSuffix(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldHasSuffix(FieldResourceID, v))
}

// ResourceIDEqualFold applies the EqualFold predicate on the "resource_id" field.
func ResourceIDEqualFold(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEqualFold(FieldResourceID, v))
}

// ResourceIDContainsFold applies the ContainsFold predicate on the "resource_id" field.
func ResourceIDContainsFold(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldContainsFold(FieldResourceID, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressIsNil applies the IsNil predicate on the "address" field.
func AddressIsNil() predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldIsNull(FieldAddress))
}

// AddressNotNil applies the NotNil predicate on the "address" field.
func AddressNotNil() predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNotNull(FieldAddress))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldContainsFold(FieldAddress, v))
}

// DesiredStateEQ applies the EQ predicate on the "desired_state" field.
func DesiredStateEQ(v DesiredState) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldDesiredState, v))
}

// DesiredStateNEQ applies the NEQ predicate on the "desired_state" field.
func DesiredStateNEQ(v DesiredState) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNEQ(FieldDesiredState, v))
}

// DesiredStateIn applies the In predicate on the "desired_state" field.
func DesiredStateIn(vs ...DesiredState) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldIn(FieldDesiredState, vs...))
}

// DesiredStateNotIn applies the NotIn predicate on the "desired_state" field.
func DesiredStateNotIn(vs ...DesiredState) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNotIn(FieldDesiredState, vs...))
}

// DesiredStateIsNil applies the IsNil predicate on the "desired_state" field.
func DesiredStateIsNil() predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldIsNull(FieldDesiredState))
}

// DesiredStateNotNil applies the NotNil predicate on the "desired_state" field.
func DesiredStateNotNil() predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNotNull(FieldDesiredState))
}

// CurrentStateEQ applies the EQ predicate on the "current_state" field.
func CurrentStateEQ(v CurrentState) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldCurrentState, v))
}

// CurrentStateNEQ applies the NEQ predicate on the "current_state" field.
func CurrentStateNEQ(v CurrentState) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNEQ(FieldCurrentState, v))
}

// CurrentStateIn applies the In predicate on the "current_state" field.
func CurrentStateIn(vs ...CurrentState) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldIn(FieldCurrentState, vs...))
}

// CurrentStateNotIn applies the NotIn predicate on the "current_state" field.
func CurrentStateNotIn(vs ...CurrentState) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNotIn(FieldCurrentState, vs...))
}

// CurrentStateIsNil applies the IsNil predicate on the "current_state" field.
func CurrentStateIsNil() predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldIsNull(FieldCurrentState))
}

// CurrentStateNotNil applies the NotNil predicate on the "current_state" field.
func CurrentStateNotNil() predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNotNull(FieldCurrentState))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNotNull(FieldStatus))
}

// StatusDetailEQ applies the EQ predicate on the "status_detail" field.
func StatusDetailEQ(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldStatusDetail, v))
}

// StatusDetailNEQ applies the NEQ predicate on the "status_detail" field.
func StatusDetailNEQ(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNEQ(FieldStatusDetail, v))
}

// StatusDetailIn applies the In predicate on the "status_detail" field.
func StatusDetailIn(vs ...string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldIn(FieldStatusDetail, vs...))
}

// StatusDetailNotIn applies the NotIn predicate on the "status_detail" field.
func StatusDetailNotIn(vs ...string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNotIn(FieldStatusDetail, vs...))
}

// StatusDetailGT applies the GT predicate on the "status_detail" field.
func StatusDetailGT(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldGT(FieldStatusDetail, v))
}

// StatusDetailGTE applies the GTE predicate on the "status_detail" field.
func StatusDetailGTE(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldGTE(FieldStatusDetail, v))
}

// StatusDetailLT applies the LT predicate on the "status_detail" field.
func StatusDetailLT(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldLT(FieldStatusDetail, v))
}

// StatusDetailLTE applies the LTE predicate on the "status_detail" field.
func StatusDetailLTE(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldLTE(FieldStatusDetail, v))
}

// StatusDetailContains applies the Contains predicate on the "status_detail" field.
func StatusDetailContains(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldContains(FieldStatusDetail, v))
}

// StatusDetailHasPrefix applies the HasPrefix predicate on the "status_detail" field.
func StatusDetailHasPrefix(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldHasPrefix(FieldStatusDetail, v))
}

// StatusDetailHasSuffix applies the HasSuffix predicate on the "status_detail" field.
func StatusDetailHasSuffix(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldHasSuffix(FieldStatusDetail, v))
}

// StatusDetailIsNil applies the IsNil predicate on the "status_detail" field.
func StatusDetailIsNil() predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldIsNull(FieldStatusDetail))
}

// StatusDetailNotNil applies the NotNil predicate on the "status_detail" field.
func StatusDetailNotNil() predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNotNull(FieldStatusDetail))
}

// StatusDetailEqualFold applies the EqualFold predicate on the "status_detail" field.
func StatusDetailEqualFold(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEqualFold(FieldStatusDetail, v))
}

// StatusDetailContainsFold applies the ContainsFold predicate on the "status_detail" field.
func StatusDetailContainsFold(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldContainsFold(FieldStatusDetail, v))
}

// ConfigMethodEQ applies the EQ predicate on the "config_method" field.
func ConfigMethodEQ(v ConfigMethod) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldConfigMethod, v))
}

// ConfigMethodNEQ applies the NEQ predicate on the "config_method" field.
func ConfigMethodNEQ(v ConfigMethod) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNEQ(FieldConfigMethod, v))
}

// ConfigMethodIn applies the In predicate on the "config_method" field.
func ConfigMethodIn(vs ...ConfigMethod) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldIn(FieldConfigMethod, vs...))
}

// ConfigMethodNotIn applies the NotIn predicate on the "config_method" field.
func ConfigMethodNotIn(vs ...ConfigMethod) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNotIn(FieldConfigMethod, vs...))
}

// ConfigMethodIsNil applies the IsNil predicate on the "config_method" field.
func ConfigMethodIsNil() predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldIsNull(FieldConfigMethod))
}

// ConfigMethodNotNil applies the NotNil predicate on the "config_method" field.
func ConfigMethodNotNil() predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNotNull(FieldConfigMethod))
}

// TenantIDEQ applies the EQ predicate on the "tenant_id" field.
func TenantIDEQ(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldTenantID, v))
}

// TenantIDNEQ applies the NEQ predicate on the "tenant_id" field.
func TenantIDNEQ(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNEQ(FieldTenantID, v))
}

// TenantIDIn applies the In predicate on the "tenant_id" field.
func TenantIDIn(vs ...string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldIn(FieldTenantID, vs...))
}

// TenantIDNotIn applies the NotIn predicate on the "tenant_id" field.
func TenantIDNotIn(vs ...string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNotIn(FieldTenantID, vs...))
}

// TenantIDGT applies the GT predicate on the "tenant_id" field.
func TenantIDGT(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldGT(FieldTenantID, v))
}

// TenantIDGTE applies the GTE predicate on the "tenant_id" field.
func TenantIDGTE(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldGTE(FieldTenantID, v))
}

// TenantIDLT applies the LT predicate on the "tenant_id" field.
func TenantIDLT(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldLT(FieldTenantID, v))
}

// TenantIDLTE applies the LTE predicate on the "tenant_id" field.
func TenantIDLTE(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldLTE(FieldTenantID, v))
}

// TenantIDContains applies the Contains predicate on the "tenant_id" field.
func TenantIDContains(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldContains(FieldTenantID, v))
}

// TenantIDHasPrefix applies the HasPrefix predicate on the "tenant_id" field.
func TenantIDHasPrefix(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldHasPrefix(FieldTenantID, v))
}

// TenantIDHasSuffix applies the HasSuffix predicate on the "tenant_id" field.
func TenantIDHasSuffix(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldHasSuffix(FieldTenantID, v))
}

// TenantIDEqualFold applies the EqualFold predicate on the "tenant_id" field.
func TenantIDEqualFold(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEqualFold(FieldTenantID, v))
}

// TenantIDContainsFold applies the ContainsFold predicate on the "tenant_id" field.
func TenantIDContainsFold(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldContainsFold(FieldTenantID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtContains applies the Contains predicate on the "created_at" field.
func CreatedAtContains(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldContains(FieldCreatedAt, v))
}

// CreatedAtHasPrefix applies the HasPrefix predicate on the "created_at" field.
func CreatedAtHasPrefix(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldHasPrefix(FieldCreatedAt, v))
}

// CreatedAtHasSuffix applies the HasSuffix predicate on the "created_at" field.
func CreatedAtHasSuffix(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldHasSuffix(FieldCreatedAt, v))
}

// CreatedAtEqualFold applies the EqualFold predicate on the "created_at" field.
func CreatedAtEqualFold(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEqualFold(FieldCreatedAt, v))
}

// CreatedAtContainsFold applies the ContainsFold predicate on the "created_at" field.
func CreatedAtContainsFold(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldContainsFold(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtContains applies the Contains predicate on the "updated_at" field.
func UpdatedAtContains(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldContains(FieldUpdatedAt, v))
}

// UpdatedAtHasPrefix applies the HasPrefix predicate on the "updated_at" field.
func UpdatedAtHasPrefix(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldHasPrefix(FieldUpdatedAt, v))
}

// UpdatedAtHasSuffix applies the HasSuffix predicate on the "updated_at" field.
func UpdatedAtHasSuffix(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldHasSuffix(FieldUpdatedAt, v))
}

// UpdatedAtEqualFold applies the EqualFold predicate on the "updated_at" field.
func UpdatedAtEqualFold(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldEqualFold(FieldUpdatedAt, v))
}

// UpdatedAtContainsFold applies the ContainsFold predicate on the "updated_at" field.
func UpdatedAtContainsFold(v string) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.FieldContainsFold(FieldUpdatedAt, v))
}

// HasNic applies the HasEdge predicate on the "nic" edge.
func HasNic() predicate.IPAddressResource {
	return predicate.IPAddressResource(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, NicTable, NicColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNicWith applies the HasEdge predicate on the "nic" edge with a given conditions (other predicates).
func HasNicWith(preds ...predicate.HostnicResource) predicate.IPAddressResource {
	return predicate.IPAddressResource(func(s *sql.Selector) {
		step := newNicStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.IPAddressResource) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.IPAddressResource) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.IPAddressResource) predicate.IPAddressResource {
	return predicate.IPAddressResource(sql.NotPredicates(p))
}
