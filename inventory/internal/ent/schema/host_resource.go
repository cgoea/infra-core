// File updated by protoc-gen-ent.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type HostResource struct {
	ent.Schema
}

func (HostResource) Fields() []ent.Field {
	return []ent.Field{field.String("resource_id").Unique(), field.String("kind").Optional(), field.String("name").Optional(), field.Enum("desired_state").Optional().Values("HOST_STATE_UNSPECIFIED", "HOST_STATE_DELETING", "HOST_STATE_DELETED", "HOST_STATE_ONBOARDED", "HOST_STATE_UNTRUSTED", "HOST_STATE_REGISTERED"), field.Enum("current_state").Optional().Values("HOST_STATE_UNSPECIFIED", "HOST_STATE_DELETING", "HOST_STATE_DELETED", "HOST_STATE_ONBOARDED", "HOST_STATE_UNTRUSTED", "HOST_STATE_REGISTERED"), field.String("note").Optional(), field.String("hardware_kind").Optional(), field.String("serial_number").Optional(), field.String("uuid").Optional().Unique(), field.Uint64("memory_bytes").Optional(), field.String("cpu_model").Optional(), field.Uint32("cpu_sockets").Optional(), field.Uint32("cpu_cores").Optional(), field.String("cpu_capabilities").Optional(), field.String("cpu_architecture").Optional(), field.Uint32("cpu_threads").Optional(), field.String("cpu_topology").Optional(), field.String("mgmt_ip").Optional(), field.Enum("bmc_kind").Optional().Values("BAREMETAL_CONTROLLER_KIND_UNSPECIFIED", "BAREMETAL_CONTROLLER_KIND_NONE", "BAREMETAL_CONTROLLER_KIND_IPMI", "BAREMETAL_CONTROLLER_KIND_VPRO", "BAREMETAL_CONTROLLER_KIND_PDU"), field.String("bmc_ip").Optional(), field.String("bmc_username").Optional(), field.String("bmc_password").Optional(), field.String("pxe_mac").Optional(), field.String("hostname").Optional(), field.String("product_name").Optional(), field.String("bios_version").Optional(), field.String("bios_release_date").Optional(), field.String("bios_vendor").Optional(), field.String("metadata").Optional(), field.Enum("current_power_state").Optional().Values("POWER_STATE_UNSPECIFIED", "POWER_STATE_ERROR", "POWER_STATE_ON", "POWER_STATE_OFF"), field.Enum("desired_power_state").Optional().Values("POWER_STATE_UNSPECIFIED", "POWER_STATE_ERROR", "POWER_STATE_ON", "POWER_STATE_OFF"), field.String("host_status").Optional(), field.Enum("host_status_indicator").Optional().Values("STATUS_INDICATION_UNSPECIFIED", "STATUS_INDICATION_ERROR", "STATUS_INDICATION_IN_PROGRESS", "STATUS_INDICATION_IDLE"), field.Uint64("host_status_timestamp").Optional(), field.String("onboarding_status").Optional(), field.Enum("onboarding_status_indicator").Optional().Values("STATUS_INDICATION_UNSPECIFIED", "STATUS_INDICATION_ERROR", "STATUS_INDICATION_IN_PROGRESS", "STATUS_INDICATION_IDLE"), field.Uint64("onboarding_status_timestamp").Optional(), field.String("registration_status").Optional(), field.Enum("registration_status_indicator").Optional().Values("STATUS_INDICATION_UNSPECIFIED", "STATUS_INDICATION_ERROR", "STATUS_INDICATION_IN_PROGRESS", "STATUS_INDICATION_IDLE"), field.Uint64("registration_status_timestamp").Optional(), field.String("tenant_id").Immutable(), field.String("created_at").Immutable().SchemaType(map[string]string{"postgres": "TIMESTAMP"}), field.String("updated_at").SchemaType(map[string]string{"postgres": "TIMESTAMP"})}
}
func (HostResource) Edges() []ent.Edge {
	return []ent.Edge{edge.To("site", SiteResource.Type).Unique(), edge.To("provider", ProviderResource.Type).Unique(), edge.From("host_storages", HoststorageResource.Type).Ref("host"), edge.From("host_nics", HostnicResource.Type).Ref("host"), edge.From("host_usbs", HostusbResource.Type).Ref("host"), edge.From("host_gpus", HostgpuResource.Type).Ref("host"), edge.From("instance", InstanceResource.Type).Ref("host").Unique()}
}
func (HostResource) Annotations() []schema.Annotation {
	return nil
}
func (HostResource) Indexes() []ent.Index {
	return []ent.Index{index.Fields("uuid").Unique().Annotations(entsql.IndexAnnotation{Where: "uuid IS NOT NULL"}), index.Fields("serial_number").Unique().Annotations(entsql.IndexAnnotation{Where: "uuid IS NULL"}), index.Fields("tenant_id")}
}
