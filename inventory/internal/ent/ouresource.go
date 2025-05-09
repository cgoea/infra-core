// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-edge-platform/infra-core/inventory/v2/internal/ent/ouresource"
)

// OuResource is the model entity for the OuResource schema.
type OuResource struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ResourceID holds the value of the "resource_id" field.
	ResourceID string `json:"resource_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// OuKind holds the value of the "ou_kind" field.
	OuKind string `json:"ou_kind,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata string `json:"metadata,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID string `json:"tenant_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt string `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OuResourceQuery when eager-loading is set.
	Edges                 OuResourceEdges `json:"edges"`
	ou_resource_parent_ou *int
	selectValues          sql.SelectValues
}

// OuResourceEdges holds the relations/edges for other nodes in the graph.
type OuResourceEdges struct {
	// ParentOu holds the value of the parent_ou edge.
	ParentOu *OuResource `json:"parent_ou,omitempty"`
	// Children holds the value of the children edge.
	Children []*OuResource `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOuOrErr returns the ParentOu value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OuResourceEdges) ParentOuOrErr() (*OuResource, error) {
	if e.ParentOu != nil {
		return e.ParentOu, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: ouresource.Label}
	}
	return nil, &NotLoadedError{edge: "parent_ou"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e OuResourceEdges) ChildrenOrErr() ([]*OuResource, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OuResource) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ouresource.FieldID:
			values[i] = new(sql.NullInt64)
		case ouresource.FieldResourceID, ouresource.FieldName, ouresource.FieldOuKind, ouresource.FieldMetadata, ouresource.FieldTenantID, ouresource.FieldCreatedAt, ouresource.FieldUpdatedAt:
			values[i] = new(sql.NullString)
		case ouresource.ForeignKeys[0]: // ou_resource_parent_ou
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OuResource fields.
func (or *OuResource) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ouresource.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			or.ID = int(value.Int64)
		case ouresource.FieldResourceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resource_id", values[i])
			} else if value.Valid {
				or.ResourceID = value.String
			}
		case ouresource.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				or.Name = value.String
			}
		case ouresource.FieldOuKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ou_kind", values[i])
			} else if value.Valid {
				or.OuKind = value.String
			}
		case ouresource.FieldMetadata:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value.Valid {
				or.Metadata = value.String
			}
		case ouresource.FieldTenantID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				or.TenantID = value.String
			}
		case ouresource.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				or.CreatedAt = value.String
			}
		case ouresource.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				or.UpdatedAt = value.String
			}
		case ouresource.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field ou_resource_parent_ou", value)
			} else if value.Valid {
				or.ou_resource_parent_ou = new(int)
				*or.ou_resource_parent_ou = int(value.Int64)
			}
		default:
			or.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OuResource.
// This includes values selected through modifiers, order, etc.
func (or *OuResource) Value(name string) (ent.Value, error) {
	return or.selectValues.Get(name)
}

// QueryParentOu queries the "parent_ou" edge of the OuResource entity.
func (or *OuResource) QueryParentOu() *OuResourceQuery {
	return NewOuResourceClient(or.config).QueryParentOu(or)
}

// QueryChildren queries the "children" edge of the OuResource entity.
func (or *OuResource) QueryChildren() *OuResourceQuery {
	return NewOuResourceClient(or.config).QueryChildren(or)
}

// Update returns a builder for updating this OuResource.
// Note that you need to call OuResource.Unwrap() before calling this method if this OuResource
// was returned from a transaction, and the transaction was committed or rolled back.
func (or *OuResource) Update() *OuResourceUpdateOne {
	return NewOuResourceClient(or.config).UpdateOne(or)
}

// Unwrap unwraps the OuResource entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (or *OuResource) Unwrap() *OuResource {
	_tx, ok := or.config.driver.(*txDriver)
	if !ok {
		panic("ent: OuResource is not a transactional entity")
	}
	or.config.driver = _tx.drv
	return or
}

// String implements the fmt.Stringer.
func (or *OuResource) String() string {
	var builder strings.Builder
	builder.WriteString("OuResource(")
	builder.WriteString(fmt.Sprintf("id=%v, ", or.ID))
	builder.WriteString("resource_id=")
	builder.WriteString(or.ResourceID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(or.Name)
	builder.WriteString(", ")
	builder.WriteString("ou_kind=")
	builder.WriteString(or.OuKind)
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(or.Metadata)
	builder.WriteString(", ")
	builder.WriteString("tenant_id=")
	builder.WriteString(or.TenantID)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(or.CreatedAt)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(or.UpdatedAt)
	builder.WriteByte(')')
	return builder.String()
}

// OuResources is a parsable slice of OuResource.
type OuResources []*OuResource
