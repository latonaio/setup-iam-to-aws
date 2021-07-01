package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"pulumi-setup-aws-iam-with-golang/components/resource"
)

func main() {
	pulumi.Run(resource.Setup)
}
