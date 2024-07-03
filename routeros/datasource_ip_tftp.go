package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceIPTftp() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceIPTftpRead,
		Schema: map[string]*schema.Schema{
			MetaResourcePath: PropResourcePath("/ip/tftp"),
			MetaId:           PropId(Id),

			KeyFilter: PropFilterRw,
			"tftp": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"disabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"ip_addresses": {
							Type:         schema.TypeSet,
							Computed: 	 true,
						},
						"req_filename": {
							Type:		 schema.TypeString,
							Computed: 	 true,
						},
						"real_filename": {
							Type:         schema.TypeString,
							Computed: 	 true,
						},
						"allow": {
							Type:         schema.TypeBool,
							Computed: 	 true,
						},
						"read_only": {
							Type:         schema.TypeBool,
							Computed: 	 true,
						},
						"allow_rollover": {
							Type:         schema.TypeBool,
							Optional:     true,
						},
						"hits": {
							Type:         schema.TypeInt,
							Optional:     true,
						},
						"reading_window_size": {
							Type:         schema.TypeString,
							Optional:     true,
						},
						"allow_overwrite": {
							Type:         schema.TypeBool,
							Optional:     true,
						},
					},
				},
			},
		},
	}
}

func datasourceIPTftpRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	s := DatasourceIPTftp().Schema
	path := s[MetaResourcePath].Default.(string)

	res, err := ReadItemsFiltered(buildReadFilter(d.Get(KeyFilter).(map[string]interface{})), path, m.(Client))
	if err != nil {
		return diag.FromErr(err)
	}

	return MikrotikResourceDataToTerraformDatasource(res, "tftp", s, d)
}
