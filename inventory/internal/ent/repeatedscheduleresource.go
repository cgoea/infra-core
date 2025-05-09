// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-edge-platform/infra-core/inventory/v2/internal/ent/hostresource"
	"github.com/open-edge-platform/infra-core/inventory/v2/internal/ent/regionresource"
	"github.com/open-edge-platform/infra-core/inventory/v2/internal/ent/repeatedscheduleresource"
	"github.com/open-edge-platform/infra-core/inventory/v2/internal/ent/siteresource"
	"github.com/open-edge-platform/infra-core/inventory/v2/internal/ent/workloadresource"
)

// RepeatedScheduleResource is the model entity for the RepeatedScheduleResource schema.
type RepeatedScheduleResource struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ResourceID holds the value of the "resource_id" field.
	ResourceID string `json:"resource_id,omitempty"`
	// ScheduleStatus holds the value of the "schedule_status" field.
	ScheduleStatus repeatedscheduleresource.ScheduleStatus `json:"schedule_status,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// DurationSeconds holds the value of the "duration_seconds" field.
	DurationSeconds uint32 `json:"duration_seconds,omitempty"`
	// CronMinutes holds the value of the "cron_minutes" field.
	CronMinutes string `json:"cron_minutes,omitempty"`
	// CronHours holds the value of the "cron_hours" field.
	CronHours string `json:"cron_hours,omitempty"`
	// CronDayMonth holds the value of the "cron_day_month" field.
	CronDayMonth string `json:"cron_day_month,omitempty"`
	// CronMonth holds the value of the "cron_month" field.
	CronMonth string `json:"cron_month,omitempty"`
	// CronDayWeek holds the value of the "cron_day_week" field.
	CronDayWeek string `json:"cron_day_week,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID string `json:"tenant_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt string `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RepeatedScheduleResourceQuery when eager-loading is set.
	Edges                                      RepeatedScheduleResourceEdges `json:"edges"`
	repeated_schedule_resource_target_site     *int
	repeated_schedule_resource_target_host     *int
	repeated_schedule_resource_target_workload *int
	repeated_schedule_resource_target_region   *int
	selectValues                               sql.SelectValues
}

// RepeatedScheduleResourceEdges holds the relations/edges for other nodes in the graph.
type RepeatedScheduleResourceEdges struct {
	// TargetSite holds the value of the target_site edge.
	TargetSite *SiteResource `json:"target_site,omitempty"`
	// TargetHost holds the value of the target_host edge.
	TargetHost *HostResource `json:"target_host,omitempty"`
	// TargetWorkload holds the value of the target_workload edge.
	TargetWorkload *WorkloadResource `json:"target_workload,omitempty"`
	// TargetRegion holds the value of the target_region edge.
	TargetRegion *RegionResource `json:"target_region,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// TargetSiteOrErr returns the TargetSite value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RepeatedScheduleResourceEdges) TargetSiteOrErr() (*SiteResource, error) {
	if e.TargetSite != nil {
		return e.TargetSite, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: siteresource.Label}
	}
	return nil, &NotLoadedError{edge: "target_site"}
}

// TargetHostOrErr returns the TargetHost value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RepeatedScheduleResourceEdges) TargetHostOrErr() (*HostResource, error) {
	if e.TargetHost != nil {
		return e.TargetHost, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: hostresource.Label}
	}
	return nil, &NotLoadedError{edge: "target_host"}
}

// TargetWorkloadOrErr returns the TargetWorkload value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RepeatedScheduleResourceEdges) TargetWorkloadOrErr() (*WorkloadResource, error) {
	if e.TargetWorkload != nil {
		return e.TargetWorkload, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: workloadresource.Label}
	}
	return nil, &NotLoadedError{edge: "target_workload"}
}

// TargetRegionOrErr returns the TargetRegion value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RepeatedScheduleResourceEdges) TargetRegionOrErr() (*RegionResource, error) {
	if e.TargetRegion != nil {
		return e.TargetRegion, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: regionresource.Label}
	}
	return nil, &NotLoadedError{edge: "target_region"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RepeatedScheduleResource) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case repeatedscheduleresource.FieldID, repeatedscheduleresource.FieldDurationSeconds:
			values[i] = new(sql.NullInt64)
		case repeatedscheduleresource.FieldResourceID, repeatedscheduleresource.FieldScheduleStatus, repeatedscheduleresource.FieldName, repeatedscheduleresource.FieldCronMinutes, repeatedscheduleresource.FieldCronHours, repeatedscheduleresource.FieldCronDayMonth, repeatedscheduleresource.FieldCronMonth, repeatedscheduleresource.FieldCronDayWeek, repeatedscheduleresource.FieldTenantID, repeatedscheduleresource.FieldCreatedAt, repeatedscheduleresource.FieldUpdatedAt:
			values[i] = new(sql.NullString)
		case repeatedscheduleresource.ForeignKeys[0]: // repeated_schedule_resource_target_site
			values[i] = new(sql.NullInt64)
		case repeatedscheduleresource.ForeignKeys[1]: // repeated_schedule_resource_target_host
			values[i] = new(sql.NullInt64)
		case repeatedscheduleresource.ForeignKeys[2]: // repeated_schedule_resource_target_workload
			values[i] = new(sql.NullInt64)
		case repeatedscheduleresource.ForeignKeys[3]: // repeated_schedule_resource_target_region
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RepeatedScheduleResource fields.
func (rsr *RepeatedScheduleResource) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case repeatedscheduleresource.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rsr.ID = int(value.Int64)
		case repeatedscheduleresource.FieldResourceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resource_id", values[i])
			} else if value.Valid {
				rsr.ResourceID = value.String
			}
		case repeatedscheduleresource.FieldScheduleStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field schedule_status", values[i])
			} else if value.Valid {
				rsr.ScheduleStatus = repeatedscheduleresource.ScheduleStatus(value.String)
			}
		case repeatedscheduleresource.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				rsr.Name = value.String
			}
		case repeatedscheduleresource.FieldDurationSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration_seconds", values[i])
			} else if value.Valid {
				rsr.DurationSeconds = uint32(value.Int64)
			}
		case repeatedscheduleresource.FieldCronMinutes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cron_minutes", values[i])
			} else if value.Valid {
				rsr.CronMinutes = value.String
			}
		case repeatedscheduleresource.FieldCronHours:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cron_hours", values[i])
			} else if value.Valid {
				rsr.CronHours = value.String
			}
		case repeatedscheduleresource.FieldCronDayMonth:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cron_day_month", values[i])
			} else if value.Valid {
				rsr.CronDayMonth = value.String
			}
		case repeatedscheduleresource.FieldCronMonth:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cron_month", values[i])
			} else if value.Valid {
				rsr.CronMonth = value.String
			}
		case repeatedscheduleresource.FieldCronDayWeek:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cron_day_week", values[i])
			} else if value.Valid {
				rsr.CronDayWeek = value.String
			}
		case repeatedscheduleresource.FieldTenantID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				rsr.TenantID = value.String
			}
		case repeatedscheduleresource.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rsr.CreatedAt = value.String
			}
		case repeatedscheduleresource.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				rsr.UpdatedAt = value.String
			}
		case repeatedscheduleresource.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field repeated_schedule_resource_target_site", value)
			} else if value.Valid {
				rsr.repeated_schedule_resource_target_site = new(int)
				*rsr.repeated_schedule_resource_target_site = int(value.Int64)
			}
		case repeatedscheduleresource.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field repeated_schedule_resource_target_host", value)
			} else if value.Valid {
				rsr.repeated_schedule_resource_target_host = new(int)
				*rsr.repeated_schedule_resource_target_host = int(value.Int64)
			}
		case repeatedscheduleresource.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field repeated_schedule_resource_target_workload", value)
			} else if value.Valid {
				rsr.repeated_schedule_resource_target_workload = new(int)
				*rsr.repeated_schedule_resource_target_workload = int(value.Int64)
			}
		case repeatedscheduleresource.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field repeated_schedule_resource_target_region", value)
			} else if value.Valid {
				rsr.repeated_schedule_resource_target_region = new(int)
				*rsr.repeated_schedule_resource_target_region = int(value.Int64)
			}
		default:
			rsr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RepeatedScheduleResource.
// This includes values selected through modifiers, order, etc.
func (rsr *RepeatedScheduleResource) Value(name string) (ent.Value, error) {
	return rsr.selectValues.Get(name)
}

// QueryTargetSite queries the "target_site" edge of the RepeatedScheduleResource entity.
func (rsr *RepeatedScheduleResource) QueryTargetSite() *SiteResourceQuery {
	return NewRepeatedScheduleResourceClient(rsr.config).QueryTargetSite(rsr)
}

// QueryTargetHost queries the "target_host" edge of the RepeatedScheduleResource entity.
func (rsr *RepeatedScheduleResource) QueryTargetHost() *HostResourceQuery {
	return NewRepeatedScheduleResourceClient(rsr.config).QueryTargetHost(rsr)
}

// QueryTargetWorkload queries the "target_workload" edge of the RepeatedScheduleResource entity.
func (rsr *RepeatedScheduleResource) QueryTargetWorkload() *WorkloadResourceQuery {
	return NewRepeatedScheduleResourceClient(rsr.config).QueryTargetWorkload(rsr)
}

// QueryTargetRegion queries the "target_region" edge of the RepeatedScheduleResource entity.
func (rsr *RepeatedScheduleResource) QueryTargetRegion() *RegionResourceQuery {
	return NewRepeatedScheduleResourceClient(rsr.config).QueryTargetRegion(rsr)
}

// Update returns a builder for updating this RepeatedScheduleResource.
// Note that you need to call RepeatedScheduleResource.Unwrap() before calling this method if this RepeatedScheduleResource
// was returned from a transaction, and the transaction was committed or rolled back.
func (rsr *RepeatedScheduleResource) Update() *RepeatedScheduleResourceUpdateOne {
	return NewRepeatedScheduleResourceClient(rsr.config).UpdateOne(rsr)
}

// Unwrap unwraps the RepeatedScheduleResource entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rsr *RepeatedScheduleResource) Unwrap() *RepeatedScheduleResource {
	_tx, ok := rsr.config.driver.(*txDriver)
	if !ok {
		panic("ent: RepeatedScheduleResource is not a transactional entity")
	}
	rsr.config.driver = _tx.drv
	return rsr
}

// String implements the fmt.Stringer.
func (rsr *RepeatedScheduleResource) String() string {
	var builder strings.Builder
	builder.WriteString("RepeatedScheduleResource(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rsr.ID))
	builder.WriteString("resource_id=")
	builder.WriteString(rsr.ResourceID)
	builder.WriteString(", ")
	builder.WriteString("schedule_status=")
	builder.WriteString(fmt.Sprintf("%v", rsr.ScheduleStatus))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(rsr.Name)
	builder.WriteString(", ")
	builder.WriteString("duration_seconds=")
	builder.WriteString(fmt.Sprintf("%v", rsr.DurationSeconds))
	builder.WriteString(", ")
	builder.WriteString("cron_minutes=")
	builder.WriteString(rsr.CronMinutes)
	builder.WriteString(", ")
	builder.WriteString("cron_hours=")
	builder.WriteString(rsr.CronHours)
	builder.WriteString(", ")
	builder.WriteString("cron_day_month=")
	builder.WriteString(rsr.CronDayMonth)
	builder.WriteString(", ")
	builder.WriteString("cron_month=")
	builder.WriteString(rsr.CronMonth)
	builder.WriteString(", ")
	builder.WriteString("cron_day_week=")
	builder.WriteString(rsr.CronDayWeek)
	builder.WriteString(", ")
	builder.WriteString("tenant_id=")
	builder.WriteString(rsr.TenantID)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(rsr.CreatedAt)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(rsr.UpdatedAt)
	builder.WriteByte(')')
	return builder.String()
}

// RepeatedScheduleResources is a parsable slice of RepeatedScheduleResource.
type RepeatedScheduleResources []*RepeatedScheduleResource
