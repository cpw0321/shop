// Copyright 2020 The shop Authors

// Package response implements response body.
package response

// CatelogIndexRtnJSON ...
type CatelogIndexRtnJSON struct {
	CategoryList       []CategoryRtnJSON `json:"categoryList"`
	CurrentCategory    CategoryRtnJSON   `json:"currentCategory"`
	CurrentSubCategory []CategoryRtnJSON `json:"currentSubCategory"`
}

// CatalogCurrentRtnJSON ...
type CatalogCurrentRtnJSON struct {
	CurrentCategory    CategoryRtnJSON   `json:"currentCategory"`
	CurrentSubCategory []CategoryRtnJSON `json:"currentSubCategory"`
}
