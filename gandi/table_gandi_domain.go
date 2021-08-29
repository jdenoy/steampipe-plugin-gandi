package gandi

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableGandiDomain() *plugin.Table {
	return &plugin.Table{
		Name:        "gandi_domain",
		Description: "List gandi domains",
		List: &plugin.ListConfig{
			Hydrate: listDomain,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "unique id of the domain"},
			{Name: "fqdn", Type: proto.ColumnType_STRING, Transform: transform.FromField("FQDN"), Description: "Fully qualified domain name, written in its native alphabet (IDN)"},
			{Name: "fqdn_unicode", Type: proto.ColumnType_STRING, Transform: transform.FromField("FQDNUnicode"), Description: "Fully qualified domain name, written in unicode"},
			{Name: "tld", Type: proto.ColumnType_STRING, Transform: transform.FromField("TLD"), Description: "The top-level domain"},
			{Name: "domain_owner", Type: proto.ColumnType_STRING, Description: "The full name of the owner"},
			{Name: "orga_owner", Type: proto.ColumnType_STRING, Description: "The username of the organization owner"},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "The username of the owner"},
			{Name: "nameserver_current", Type: proto.ColumnType_STRING, Transform: transform.FromField("NameServerConfig.Current"), Description: ""},
			{Name: "auto_renew", Type: proto.ColumnType_BOOL, Description: "Automatic renewal status"},
			{Name: "sharing_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("SharingID"), Description: "The id of the organization"},
			{Name: "status", Type: proto.ColumnType_JSON, Description: "The status of the domain"},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "Tags associated to this domain"},
			{Name: "dates_registry_created_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.RegistryCreatedAt"), Description: ""},
			{Name: "dates_created_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.CreatedAt"), Description: ""},
			{Name: "dates_updated_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.UpdatedAt"), Description: ""},
			{Name: "dates_deletes_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.DeletesAt"), Description: ""},
			{Name: "dates_auth_info_expires_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.AuthInfoExpiresAt"), Description: ""},
			{Name: "dates_hold_begins_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.HoldBeginsAt"), Description: ""},
			{Name: "dates_hold_ends_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.HoldEndsAt"), Description: ""},
			{Name: "dates_pending_delete_ends_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.PendingDeleteEndsAt"), Description: ""},
			{Name: "dates_registry_ends_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.RegistryEndsAt"), Description: ""},
			{Name: "dates_renew_begins_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.RenewBeginsAt"), Description: ""},
			{Name: "dates_renew_ends_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.RenewEndsAt"), Description: ""},
		},
	}
}

func listDomain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	domains, err := client.ListDomains()
	if err != nil {
		return nil, err
	}
	for _, domain := range domains {
		d.StreamListItem(ctx, domain)
	}
	return nil, nil
}
