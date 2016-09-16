package function

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"function_zip":           resourceZip(),
			"function_map_from_list": resourceMapFromList(),
		},
	}
}
