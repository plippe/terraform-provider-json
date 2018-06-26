# Terraform Provider JSON

Terraform is a great tool, crippled by a horrible interface, HashiCorp Configuration Language (HCL). People suggest
using templating engines, or full blown code generators to avoid HCL. ​

This project attempts to facilitate ​building new Terraform libraries​. It reads in the official source code to create
JSON representations of data sources, and resources. These can easily be used to generate a rough prototype of your
HCL abstraction.

### Getting started
`make init run`

### Supported Providers
- [terraform-provider-aws](https://github.com/terraform-providers/terraform-provider-aws) - [JSON](https://plippe.github.io/terraform-provider-json/terraform-provider-aws.json)
