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
			Description:  "",
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.IsCIDR,
			},
		},
		"req_filename": {
			Type:		 schema.TypeString,
			Optional:	 true,
			Description: "",
		},
		"real_filename": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "",
		},
		"allow": {
			Type:         schema.TypeBool,
			Optional:     true,
			Default:      true,
			Description:  "",
		},
		"read_only": {
			Type:         schema.TypeBool,
			Optional:     true,
			Default:      true,
			Description:  "",
		},
		"allow_rollover": {
			Type:         schema.TypeBool,
			Optional:     true,
			Default:      false,
			Description:  "",
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
			Description:  "",
			Default:      "none",
			ValidateFunc: validation.StringInSlice([]string{"none", "pipelining"}, false),
		},
		"hits": {
			Type:         schema.TypeInt,
			Computed:     true,
			Description:  "",
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
