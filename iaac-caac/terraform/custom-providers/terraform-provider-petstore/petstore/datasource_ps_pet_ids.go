package petstore

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	sdk "github.com/terraform-in-action/go-petstore"
	"strings"
)

func dataSourcePSPetIDs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePSPetIDsRead,
		Schema: map[string]*schema.Schema{
			"names": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			}, "ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}
func dataSourcePSPetIDsRead(d *schema.ResourceData, meta interface{}) error {
	names := make(map[string]bool)
	for _, name := range d.Get("names").([]interface{}) {
		names[name.(string)] = true
	}
	conn := meta.(*sdk.Client)
	petList, err := conn.Pets.List(sdk.PetListOptions{})
	if err != nil {
		return err
	}
	var ids []string
	for _, pet := range petList.Items {
		if names["*"] || names[pet.Name] {
			ids = append(ids, pet.ID)
		}
	}

	d.Set("ids", ids)
	id := fmt.Sprintf("%d", schema.HashString(strings.Join(ids, "")))
	d.SetId(id)
	return nil
}
