module github.com/kevinvalk/terraform-provider-transip

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.13.1
	github.com/transip/gotransip/v6 v6.2.0
)

replace github.com/transip/gotransip/v6 => github.com/kevinvalk/gotransip/v6 v6.2.1-0.20200612183621-bbe9d82e6c24
