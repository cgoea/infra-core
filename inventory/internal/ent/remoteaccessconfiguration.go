// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-edge-platform/infra-core/inventory/v2/internal/ent/instanceresource"
	"github.com/open-edge-platform/infra-core/inventory/v2/internal/ent/remoteaccessconfiguration"
)

// RemoteAccessConfiguration is the model entity for the RemoteAccessConfiguration schema.
type RemoteAccessConfiguration struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ResourceID holds the value of the "resource_id" field.
	ResourceID string `json:"resource_id,omitempty"`
	// ExpirationTimestamp holds the value of the "expiration_timestamp" field.
	ExpirationTimestamp uint64 `json:"expiration_timestamp,omitempty"`
	// LocalPort holds the value of the "local_port" field.
	LocalPort uint32 `json:"local_port,omitempty"`
	// User holds the value of the "user" field.
	User string `json:"user,omitempty"`
	// CurrentState holds the value of the "current_state" field.
	CurrentState remoteaccessconfiguration.CurrentState `json:"current_state,omitempty"`
	// DesiredState holds the value of the "desired_state" field.
	DesiredState remoteaccessconfiguration.DesiredState `json:"desired_state,omitempty"`
	// ConfigurationStatus holds the value of the "configuration_status" field.
	ConfigurationStatus string `json:"configuration_status,omitempty"`
	// ConfigurationStatusIndicator holds the value of the "configuration_status_indicator" field.
	ConfigurationStatusIndicator remoteaccessconfiguration.ConfigurationStatusIndicator `json:"configuration_status_indicator,omitempty"`
	// ConfigurationStatusTimestamp holds the value of the "configuration_status_timestamp" field.
	ConfigurationStatusTimestamp uint64 `json:"configuration_status_timestamp,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID string `json:"tenant_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt string `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RemoteAccessConfigurationQuery when eager-loading is set.
	Edges                                RemoteAccessConfigurationEdges `json:"edges"`
	remote_access_configuration_instance *int
	selectValues                         sql.SelectValues
}

// RemoteAccessConfigurationEdges holds the relations/edges for other nodes in the graph.
type RemoteAccessConfigurationEdges struct {
	// Instance holds the value of the instance edge.
	Instance *InstanceResource `json:"instance,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// InstanceOrErr returns the Instance value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RemoteAccessConfigurationEdges) InstanceOrErr() (*InstanceResource, error) {
	if e.Instance != nil {
		return e.Instance, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: instanceresource.Label}
	}
	return nil, &NotLoadedError{edge: "instance"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RemoteAccessConfiguration) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case remoteaccessconfiguration.FieldID, remoteaccessconfiguration.FieldExpirationTimestamp, remoteaccessconfiguration.FieldLocalPort, remoteaccessconfiguration.FieldConfigurationStatusTimestamp:
			values[i] = new(sql.NullInt64)
		case remoteaccessconfiguration.FieldResourceID, remoteaccessconfiguration.FieldUser, remoteaccessconfiguration.FieldCurrentState, remoteaccessconfiguration.FieldDesiredState, remoteaccessconfiguration.FieldConfigurationStatus, remoteaccessconfiguration.FieldConfigurationStatusIndicator, remoteaccessconfiguration.FieldTenantID, remoteaccessconfiguration.FieldCreatedAt, remoteaccessconfiguration.FieldUpdatedAt:
			values[i] = new(sql.NullString)
		case remoteaccessconfiguration.ForeignKeys[0]: // remote_access_configuration_instance
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RemoteAccessConfiguration fields.
func (rac *RemoteAccessConfiguration) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case remoteaccessconfiguration.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rac.ID = int(value.Int64)
		case remoteaccessconfiguration.FieldResourceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resource_id", values[i])
			} else if value.Valid {
				rac.ResourceID = value.String
			}
		case remoteaccessconfiguration.FieldExpirationTimestamp:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field expiration_timestamp", values[i])
			} else if value.Valid {
				rac.ExpirationTimestamp = uint64(value.Int64)
			}
		case remoteaccessconfiguration.FieldLocalPort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field local_port", values[i])
			} else if value.Valid {
				rac.LocalPort = uint32(value.Int64)
			}
		case remoteaccessconfiguration.FieldUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user", values[i])
			} else if value.Valid {
				rac.User = value.String
			}
		case remoteaccessconfiguration.FieldCurrentState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field current_state", values[i])
			} else if value.Valid {
				rac.CurrentState = remoteaccessconfiguration.CurrentState(value.String)
			}
		case remoteaccessconfiguration.FieldDesiredState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desired_state", values[i])
			} else if value.Valid {
				rac.DesiredState = remoteaccessconfiguration.DesiredState(value.String)
			}
		case remoteaccessconfiguration.FieldConfigurationStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field configuration_status", values[i])
			} else if value.Valid {
				rac.ConfigurationStatus = value.String
			}
		case remoteaccessconfiguration.FieldConfigurationStatusIndicator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field configuration_status_indicator", values[i])
			} else if value.Valid {
				rac.ConfigurationStatusIndicator = remoteaccessconfiguration.ConfigurationStatusIndicator(value.String)
			}
		case remoteaccessconfiguration.FieldConfigurationStatusTimestamp:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field configuration_status_timestamp", values[i])
			} else if value.Valid {
				rac.ConfigurationStatusTimestamp = uint64(value.Int64)
			}
		case remoteaccessconfiguration.FieldTenantID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				rac.TenantID = value.String
			}
		case remoteaccessconfiguration.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rac.CreatedAt = value.String
			}
		case remoteaccessconfiguration.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				rac.UpdatedAt = value.String
			}
		case remoteaccessconfiguration.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field remote_access_configuration_instance", value)
			} else if value.Valid {
				rac.remote_access_configuration_instance = new(int)
				*rac.remote_access_configuration_instance = int(value.Int64)
			}
		default:
			rac.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RemoteAccessConfiguration.
// This includes values selected through modifiers, order, etc.
func (rac *RemoteAccessConfiguration) Value(name string) (ent.Value, error) {
	return rac.selectValues.Get(name)
}

// QueryInstance queries the "instance" edge of the RemoteAccessConfiguration entity.
func (rac *RemoteAccessConfiguration) QueryInstance() *InstanceResourceQuery {
	return NewRemoteAccessConfigurationClient(rac.config).QueryInstance(rac)
}

// Update returns a builder for updating this RemoteAccessConfiguration.
// Note that you need to call RemoteAccessConfiguration.Unwrap() before calling this method if this RemoteAccessConfiguration
// was returned from a transaction, and the transaction was committed or rolled back.
func (rac *RemoteAccessConfiguration) Update() *RemoteAccessConfigurationUpdateOne {
	return NewRemoteAccessConfigurationClient(rac.config).UpdateOne(rac)
}

// Unwrap unwraps the RemoteAccessConfiguration entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rac *RemoteAccessConfiguration) Unwrap() *RemoteAccessConfiguration {
	_tx, ok := rac.config.driver.(*txDriver)
	if !ok {
		panic("ent: RemoteAccessConfiguration is not a transactional entity")
	}
	rac.config.driver = _tx.drv
	return rac
}

// String implements the fmt.Stringer.
func (rac *RemoteAccessConfiguration) String() string {
	var builder strings.Builder
	builder.WriteString("RemoteAccessConfiguration(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rac.ID))
	builder.WriteString("resource_id=")
	builder.WriteString(rac.ResourceID)
	builder.WriteString(", ")
	builder.WriteString("expiration_timestamp=")
	builder.WriteString(fmt.Sprintf("%v", rac.ExpirationTimestamp))
	builder.WriteString(", ")
	builder.WriteString("local_port=")
	builder.WriteString(fmt.Sprintf("%v", rac.LocalPort))
	builder.WriteString(", ")
	builder.WriteString("user=")
	builder.WriteString(rac.User)
	builder.WriteString(", ")
	builder.WriteString("current_state=")
	builder.WriteString(fmt.Sprintf("%v", rac.CurrentState))
	builder.WriteString(", ")
	builder.WriteString("desired_state=")
	builder.WriteString(fmt.Sprintf("%v", rac.DesiredState))
	builder.WriteString(", ")
	builder.WriteString("configuration_status=")
	builder.WriteString(rac.ConfigurationStatus)
	builder.WriteString(", ")
	builder.WriteString("configuration_status_indicator=")
	builder.WriteString(fmt.Sprintf("%v", rac.ConfigurationStatusIndicator))
	builder.WriteString(", ")
	builder.WriteString("configuration_status_timestamp=")
	builder.WriteString(fmt.Sprintf("%v", rac.ConfigurationStatusTimestamp))
	builder.WriteString(", ")
	builder.WriteString("tenant_id=")
	builder.WriteString(rac.TenantID)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(rac.CreatedAt)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(rac.UpdatedAt)
	builder.WriteByte(')')
	return builder.String()
}

// RemoteAccessConfigurations is a parsable slice of RemoteAccessConfiguration.
type RemoteAccessConfigurations []*RemoteAccessConfiguration
