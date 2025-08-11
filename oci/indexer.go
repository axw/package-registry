// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

package oci

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/elastic/package-registry/packages"
)

// Indexer is an implementation of the Indexer interface that reads packages from an OCI registry using ORAS.
type Indexer struct {
	logger *zap.Logger
	opts   IndexerOptions
}

// IndexerOptions contains configuration options for the OCI indexer.
type IndexerOptions struct {
	// Registry is the URL of the OCI registry (e.g., "registry.example.com")
	Registry string
	// Repository is the repository within the registry (e.g., "packages")
	Repository string
	// Username for registry authentication
	Username string
	// Password for registry authentication  
	Password string
	// Insecure allows insecure connections to the registry
	Insecure bool
}

// NewIndexer creates a new OCI indexer with the given options.
func NewIndexer(logger *zap.Logger, opts IndexerOptions) *Indexer {
	return &Indexer{
		logger: logger,
		opts:   opts,
	}
}

// Init initializes the OCI indexer.
func (i *Indexer) Init(ctx context.Context) error {
	i.logger.Info("Initializing OCI indexer", 
		zap.String("registry", i.opts.Registry),
		zap.String("repository", i.opts.Repository))
	
	// For the initial implementation, we'll just validate the configuration
	if i.opts.Registry == "" {
		return fmt.Errorf("OCI registry URL is required")
	}
	
	if i.opts.Repository == "" {
		return fmt.Errorf("OCI repository name is required")
	}

	i.logger.Info("OCI indexer initialized successfully")
	return nil
}

// Get retrieves packages from the OCI registry.
func (i *Indexer) Get(ctx context.Context, opts *packages.GetOptions) (packages.Packages, error) {
	i.logger.Debug("Getting packages from OCI registry")
	
	// For the initial implementation, return a mock package to demonstrate the integration
	// In a real implementation, this would:
	// 1. Connect to the OCI registry using ORAS
	// 2. List available tags/artifacts
	// 3. Pull each artifact to extract package manifests
	// 4. Parse manifest.yml files into Package structs
	
	mockPackageName := "oci-mock-package"
	mockPackageTitle := "Mock OCI Package"
	
	mockPackage := &packages.Package{
		BasePackage: packages.BasePackage{
			Name:        mockPackageName,
			Version:     "1.0.0",
			Title:       &mockPackageTitle,
			Description: fmt.Sprintf("Mock package from OCI registry %s/%s", i.opts.Registry, i.opts.Repository),
			Type:        "integration",
			Categories:  []string{"web"},
		},
		BasePath: fmt.Sprintf("oci://%s/%s:latest", i.opts.Registry, i.opts.Repository),
	}

	allPackages := packages.Packages{mockPackage}

	i.logger.Info("Retrieved packages from OCI registry", zap.Int("count", len(allPackages)))
	return allPackages, nil
}

// Close closes the OCI indexer and cleans up resources.
func (i *Indexer) Close(ctx context.Context) error {
	i.logger.Debug("Closing OCI indexer")
	return nil
}