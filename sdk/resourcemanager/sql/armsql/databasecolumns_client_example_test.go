//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ColumnsListByDatabaseMax.json
func ExampleDatabaseColumnsClient_NewListByDatabasePager_filterDatabaseColumns() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabaseColumnsClient().NewListByDatabasePager("myRG", "serverName", "myDatabase", &armsql.DatabaseColumnsClientListByDatabaseOptions{Schema: []string{
		"dbo"},
		Table: []string{
			"customer",
			"address"},
		Column: []string{
			"username"},
		OrderBy: []string{
			"schema asc",
			"table",
			"column desc"},
		Skiptoken: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DatabaseColumnListResult = armsql.DatabaseColumnListResult{
		// 	Value: []*armsql.DatabaseColumn{
		// 		{
		// 			Name: to.Ptr("username"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/schemas/tables/columns"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/servers/serverName/databases/myDatabase/schemas/dbo/tables/customer/columns/username"),
		// 			Properties: &armsql.DatabaseColumnProperties{
		// 				ColumnType: to.Ptr(armsql.ColumnDataTypeNvarchar),
		// 				IsComputed: to.Ptr(false),
		// 				MemoryOptimized: to.Ptr(false),
		// 				TemporalType: to.Ptr(armsql.TableTemporalTypeNonTemporalTable),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ColumnsListByDatabaseMin.json
func ExampleDatabaseColumnsClient_NewListByDatabasePager_listDatabaseColumns() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabaseColumnsClient().NewListByDatabasePager("myRG", "serverName", "myDatabase", &armsql.DatabaseColumnsClientListByDatabaseOptions{Schema: []string{},
		Table:     []string{},
		Column:    []string{},
		OrderBy:   []string{},
		Skiptoken: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DatabaseColumnListResult = armsql.DatabaseColumnListResult{
		// 	Value: []*armsql.DatabaseColumn{
		// 		{
		// 			Name: to.Ptr("col1"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/schemas/tables/columns"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/servers/serverName/databases/myDatabase/schemas/dbo/tables/table1/columns/col1"),
		// 			Properties: &armsql.DatabaseColumnProperties{
		// 				ColumnType: to.Ptr(armsql.ColumnDataTypeInt),
		// 				IsComputed: to.Ptr(false),
		// 				MemoryOptimized: to.Ptr(false),
		// 				TemporalType: to.Ptr(armsql.TableTemporalTypeNonTemporalTable),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("col2"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/schemas/tables/columns"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/servers/serverName/databases/myDatabase/schemas/dbo/tables/table1/columns/col2"),
		// 			Properties: &armsql.DatabaseColumnProperties{
		// 				ColumnType: to.Ptr(armsql.ColumnDataTypeBit),
		// 				IsComputed: to.Ptr(false),
		// 				MemoryOptimized: to.Ptr(false),
		// 				TemporalType: to.Ptr(armsql.TableTemporalTypeSystemVersionedTemporalTable),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseColumnListByTable.json
func ExampleDatabaseColumnsClient_NewListByTablePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabaseColumnsClient().NewListByTablePager("myRG", "serverName", "myDatabase", "dbo", "table1", &armsql.DatabaseColumnsClientListByTableOptions{Filter: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DatabaseColumnListResult = armsql.DatabaseColumnListResult{
		// 	Value: []*armsql.DatabaseColumn{
		// 		{
		// 			Name: to.Ptr("col1"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/schemas/tables/columns"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/servers/serverName/databases/myDatabase/schemas/dbo/tables/table1/columns/col1"),
		// 			Properties: &armsql.DatabaseColumnProperties{
		// 				ColumnType: to.Ptr(armsql.ColumnDataTypeNvarchar),
		// 				IsComputed: to.Ptr(false),
		// 				MemoryOptimized: to.Ptr(false),
		// 				TemporalType: to.Ptr(armsql.TableTemporalTypeNonTemporalTable),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("col2"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/schemas/tables/columns"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/servers/serverName/databases/myDatabase/schemas/dbo/tables/table1/columns/col2"),
		// 			Properties: &armsql.DatabaseColumnProperties{
		// 				ColumnType: to.Ptr(armsql.ColumnDataTypeBit),
		// 				IsComputed: to.Ptr(false),
		// 				MemoryOptimized: to.Ptr(false),
		// 				TemporalType: to.Ptr(armsql.TableTemporalTypeNonTemporalTable),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseColumnGet.json
func ExampleDatabaseColumnsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabaseColumnsClient().Get(ctx, "myRG", "serverName", "myDatabase", "dbo", "table1", "column1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatabaseColumn = armsql.DatabaseColumn{
	// 	Name: to.Ptr("column1"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/schemas/tables/columns"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/servers/serverName/databases/myDatabase/schemas/dbo/tables/table1/columns/column1"),
	// 	Properties: &armsql.DatabaseColumnProperties{
	// 		ColumnType: to.Ptr(armsql.ColumnDataTypeBit),
	// 		IsComputed: to.Ptr(false),
	// 		MemoryOptimized: to.Ptr(false),
	// 		TemporalType: to.Ptr(armsql.TableTemporalTypeSystemVersionedTemporalTable),
	// 	},
	// }
}
