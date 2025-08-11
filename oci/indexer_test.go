// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

package oci

import (
	"context"
	"testing"

	"go.uber.org/zap"

	"github.com/elastic/package-registry/packages"
)

func TestOCIIndexer_Init(t *testing.T) {
	logger := zap.NewNop()
	
	tests := []struct {
		name        string
		opts        IndexerOptions
		expectError bool
	}{
		{
			name: "valid configuration",
			opts: IndexerOptions{
				Registry:   "registry.example.com",
				Repository: "packages",
			},
			expectError: false,
		},
		{
			name: "missing registry",
			opts: IndexerOptions{
				Repository: "packages",
			},
			expectError: true,
		},
		{
			name: "missing repository",
			opts: IndexerOptions{
				Registry: "registry.example.com",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indexer := NewIndexer(logger, tt.opts)
			err := indexer.Init(context.Background())
			
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestOCIIndexer_Get(t *testing.T) {
	logger := zap.NewNop()
	opts := IndexerOptions{
		Registry:   "registry.example.com",
		Repository: "packages",
	}
	
	indexer := NewIndexer(logger, opts)
	err := indexer.Init(context.Background())
	if err != nil {
		t.Fatalf("failed to initialize indexer: %v", err)
	}

	packages, err := indexer.Get(context.Background(), nil)
	if err != nil {
		t.Fatalf("failed to get packages: %v", err)
	}

	if len(packages) != 1 {
		t.Errorf("expected 1 package, got %d", len(packages))
	}

	pkg := packages[0]
	if pkg.Name != "oci-mock-package" {
		t.Errorf("expected package name 'oci-mock-package', got '%s'", pkg.Name)
	}

	if pkg.Version != "1.0.0" {
		t.Errorf("expected package version '1.0.0', got '%s'", pkg.Version)
	}

	if pkg.Type != "integration" {
		t.Errorf("expected package type 'integration', got '%s'", pkg.Type)
	}
}

func TestOCIIndexer_GetWithFilter(t *testing.T) {
	logger := zap.NewNop()
	opts := IndexerOptions{
		Registry:   "registry.example.com",
		Repository: "packages",
	}
	
	indexer := NewIndexer(logger, opts)
	err := indexer.Init(context.Background())
	if err != nil {
		t.Fatalf("failed to initialize indexer: %v", err)
	}

	filter := &packages.Filter{}
	getOpts := &packages.GetOptions{Filter: filter}
	
	packages, err := indexer.Get(context.Background(), getOpts)
	if err != nil {
		t.Fatalf("failed to get packages with filter: %v", err)
	}

	if len(packages) != 1 {
		t.Errorf("expected 1 package with filter, got %d", len(packages))
	}
}

func TestOCIIndexer_Close(t *testing.T) {
	logger := zap.NewNop()
	opts := IndexerOptions{
		Registry:   "registry.example.com",
		Repository: "packages",
	}
	
	indexer := NewIndexer(logger, opts)
	err := indexer.Close(context.Background())
	if err != nil {
		t.Errorf("unexpected error closing indexer: %v", err)
	}
}