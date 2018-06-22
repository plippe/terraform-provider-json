init:
	git submodule update --init --recursive

	go get github.com/aws/aws-sdk-go/aws
	go get github.com/beevik/etree
	go get github.com/davecgh/go-spew/spew
	go get github.com/hashicorp/errwrap
	go get github.com/hashicorp/go-cleanhttp
	go get github.com/hashicorp/go-multierror
	go get github.com/hashicorp/go-version
	go get github.com/hashicorp/terraform
	go get github.com/jen20/awspolicyequivalence
	go get github.com/mitchellh/copystructure
	go get github.com/mitchellh/go-homedir
	go get gopkg.in/yaml.v2

run:
	cp src/aws.go terraform-provider-aws/aws
	go run src/main.go
