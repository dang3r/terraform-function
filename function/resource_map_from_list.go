package function

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMapFromList() *schema.Resource {
	return &schema.Resource{
		Create: resourceMFLCreate,
		Read:   resourceMFLRead,
		Update: resourceMFLUpdate,
		Delete: resourceMFLDelete,

		Schema: map[string]*schema.Schema{
			"list": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"return": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func resourceMFLCreate(d *schema.ResourceData, m interface{}) error {
	list := d.Get("list").([]interface{})

	var uid string
	dict := map[string]string{}
	for i := 0; i < len(list); i += 2 {
		key := list[i].(string)
		val := list[i+1].(string)
		dict[key] = val
		uid += key + ":" + val
	}

	d.Set("return", dict)
	d.SetId(uid)
	return nil
}

func resourceMFLRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMFLUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMFLDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
