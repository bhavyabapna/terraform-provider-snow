package main

import (
    "github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
    return &schema.Provider{
        ResourcesMap: map[string]*schema.Resource{
          "instance_name": &schema.Schema{
    				Type:     schema.TypeString,
    				Optional: true,
    				ForceNew: true,
    				Default:  "",
    			},
          "user_name": &schema.Schema{
    				Type:     schema.TypeString,
    				Optional: true,
    				ForceNew: true,
    				Default:  "",
    			},
          "user_password": &schema.Schema{
    				Type:     schema.TypeString,
    				Optional: true,
    				ForceNew: true,
    				Default:  "",
    			},
          "access_token": &schema.Schema{
    				Type:        schema.TypeString,
    				Description: "The access token",
    				Optional:    true,
    				ForceNew:    true,
    				Default:     "",
    			},
        },
    }
}
