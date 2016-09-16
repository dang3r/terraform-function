package function

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceZip() *schema.Resource {
	return &schema.Resource{
		Create: resourceZipCreate,
		Read:   resourceZipCreate,
		Update: resourceZipRead,
		Delete: resourceZipRead,

		Schema: map[string]*schema.Schema{
			"list1": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"list2": &schema.Schema{
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

func resourceZipCreate(d *schema.ResourceData, m interface{}) error {
	l1 := d.Get("list1").([]interface{})
	l2 := d.Get("list2").([]interface{})

	if len(l1) != len(l2) {
		return errors.New("Array sizes are not equal.")
	}

	var uid string
	dict := map[string]string{}
	for i, _ := range l1 {
		key := l1[i].(string)
		val := l2[i].(string)
		dict[key] = val
		uid += key + ":" + val
	}

	d.Set("return", dict)
	d.SetId(uid)
	return nil
}

func resourceZipRead(d *schema.ResourceData, m interface{}) error {
	d.SetId("LOL")
	return nil
}
