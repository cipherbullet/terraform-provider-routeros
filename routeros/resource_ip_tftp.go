package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceIPRoute https://wiki.mikrotik.com/wiki/Manual:IP/TFTP
func ResourceIPTftp() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/tftp"),
		MetaId:           PropId(Id),

		KeyDisabled: PropDisabledRw,

		"ip_addresses": {
			Type:         schema.TypeSet,
			Optional:     true,
			Default: 	  "0.0.0.0/0",
			Description:  "range of IP addresses accepted as clients if empty 0.0.0.0/0 will be used",
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.IsCIDR,
			},
		},
		"req_filename": {
			Type:		 schema.TypeString,
			Optional:	 true,
			Default: 	 ".*",
			Description: "requested filename as regular expression (regex) if field is left empty it defaults to .*",
		},
		"real_filename": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "if req-filename and real-filename values are set and valid, the requested filename will be replaced with matched file",
		},
		"allow": {
			Type:         schema.TypeBool,
			Optional:     true,
			Default:      true,
			Description:  "to allow connection if req-filename and real-filename are set. if no, connection will be interrupted",
		},
		"read_only": {
			Type:         schema.TypeBool,
			Optional:     true,
			Default:      true,
			Description:  "sets if file can be written to, if set to 'no' write attempt will fail with error",
		},
		"allow_rollover": {
			Type:         schema.TypeBool,
			Optional:     true,
			Default:      false,
			Description:  "allow sequence number to roll over when maximum value is reached for large downloads",
		},
		"allow_overwrite": {
			Type:         schema.TypeBool,
			Optional:     true,
			Default:      false,
			Description:  "",
		},
		"reading_window_size": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "none",
			Description:  "",
			ValidateFunc: validation.StringInSlice([]string{"none", "pipelining"}, false),
		},
		"hits": {
			Type:         schema.TypeInt,
			Computed:     true,
			Description:  "how many times this access rule entry has been used (read-only)",
		},
	}
	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
