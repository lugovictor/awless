package awsspec

func init() {
	genTestsParams["createtag"] = map[string]interface{}{
		"key":      "MyKey",
		"resource": "my-vpc-id",
		"value":    "MyValue",
	}
}
