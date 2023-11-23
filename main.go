package main

//go:generate oapi-codegen -package=main -generate=client,types -o ./americancloud.gen.go https://app.americancloud.com/docs/api-docs.json

import (
	"context"
	"fmt"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
)

func main() {
	bearerTokenProvider, bearerTokenProviderErr := securityprovider.NewSecurityProviderBearerToken(os.Getenv("AC_TOKEN"))
	if bearerTokenProviderErr != nil {
		panic(bearerTokenProviderErr)
	}

	c, err := NewClientWithResponses("https://app.americancloud.com/api", WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		panic(err)
	}

	resp, err := c.ListProjectsWithResponse(context.Background())
	if err != nil {
		panic(err)
	}
	if !(resp.StatusCode() >= 200 && resp.StatusCode() < 300) {
		panic(resp.Status())
	}

	projects := *resp.JSON200.Data
	if len(projects) > 0 {
		fmt.Printf("Id: %v\n", *projects[0].Id)
		fmt.Printf("AccountID: %v\n", *projects[0].AccountId)
		fmt.Printf("Description: %v\n", *projects[0].Description)
		fmt.Printf("Created At: %v\n", *projects[0].CreatedAt)
	}
}
