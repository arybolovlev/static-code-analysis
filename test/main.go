package main

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func main() {
	// // first("hello")
	fn := fullName{firstName: "Sasha"}
	fn.printFirstName()
	p := Print
	p(fn.firstName)
	Print(fn.secondName)
	fmt.Println(fn.getFirstName())
}

func Print(s string) {
	fmt.Println(s)
}

type fullName struct {
	firstName  string
	secondName string
}

func (fn *fullName) printFirstName() {
	fmt.Println(fn.firstName)
}

func (fn *fullName) getFirstName() string {
	return fn.firstName
}

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"data_source": dataSource(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"resource": resource(),
		},
	}
}

func dataSource() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,
	}
}

func dataSourceRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resource() *schema.Resource {
	return &schema.Resource{
		Read: resourceRead,
	}
}

func resourceRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
