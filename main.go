package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.ibm.com/Orpheus/terraform-provider-snow/snow"
)

/*func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: snow.Provider,
	})
}*/

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: func() terraform.ResourceProvider {
            return Provider()
        },
    })
}
