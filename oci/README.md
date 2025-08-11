# OCI Indexer

The OCI indexer is a technical preview feature that allows the Elastic Package Registry to serve packages from an OCI (Open Container Initiative) registry using [ORAS](https://oras.land/) (OCI Registry As Storage).

## Features

- **OCI Registry Integration**: Connect to any OCI-compliant registry
- **Package Discovery**: Automatically discover packages from registry tags
- **Authentication Support**: Username/password authentication
- **Insecure Connections**: Option to skip TLS verification for testing
- **Graceful Fallback**: Falls back to mock packages when registry is unavailable

## Configuration

Enable the OCI indexer with the `-feature-oci-indexer` flag and configure it using the following options:

### Command Line Flags

- `-feature-oci-indexer`: Enable OCI indexer (technical preview)
- `-oci-registry`: OCI registry URL (e.g., `registry.example.com`)
- `-oci-repository`: Repository name within the registry (default: `packages`)
- `-oci-username`: Username for registry authentication
- `-oci-password`: Password for registry authentication
- `-oci-insecure`: Allow insecure connections to the registry

### Environment Variables

Credentials can also be provided via environment variables:

- `EPR_OCI_USERNAME`: Username for OCI registry authentication
- `EPR_OCI_PASSWORD`: Password for OCI registry authentication

## Usage Examples

### Basic Usage

```bash
package-registry -feature-oci-indexer -oci-registry "registry.example.com"
```

### With Authentication

```bash
package-registry -feature-oci-indexer \
  -oci-registry "registry.example.com" \
  -oci-repository "my-packages" \
  -oci-username "myuser" \
  -oci-password "mypassword"
```

### With Environment Variables

```bash
export EPR_OCI_USERNAME="myuser"
export EPR_OCI_PASSWORD="mypassword"
package-registry -feature-oci-indexer -oci-registry "registry.example.com"
```

### Insecure Connections (for testing)

```bash
package-registry -feature-oci-indexer \
  -oci-registry "localhost:5000" \
  -oci-insecure
```

## How It Works

1. **Initialization**: The indexer validates configuration and creates an ORAS client
2. **Package Discovery**: Lists all tags in the configured repository
3. **Package Creation**: For each tag, creates a package based on tag information
4. **Integration**: Packages are served through the standard Package Registry APIs

## Current Limitations

This is a technical preview implementation with the following limitations:

- **Tag-based Packages**: Currently creates packages based on tag names rather than pulling actual package manifests
- **Mock Fallback**: Falls back to mock packages when registry operations fail
- **Basic Authentication**: Only supports username/password authentication
- **No Advanced Search**: Does not yet integrate with Zot search extensions

## Future Enhancements

The following features are planned for future releases:

- **Real Package Parsing**: Pull and parse actual `manifest.yml` files from OCI artifacts
- **Advanced Authentication**: Support for token-based authentication and registry-specific auth methods
- **Zot Search Integration**: Integration with [Zot search extensions](https://github.com/project-zot/zot/blob/main/pkg/extensions/search/search.md) for enhanced search capabilities
- **Package Content Serving**: Serve actual package files and assets from OCI registry
- **Caching**: Implement caching for improved performance

## Examples with Real OCI Registries

### Docker Hub

```bash
package-registry -feature-oci-indexer \
  -oci-registry "docker.io" \
  -oci-repository "library/nginx"
```

### GitHub Container Registry

```bash
package-registry -feature-oci-indexer \
  -oci-registry "ghcr.io" \
  -oci-repository "owner/repo" \
  -oci-username "github-username" \
  -oci-password "github-token"
```

### Azure Container Registry

```bash
package-registry -feature-oci-indexer \
  -oci-registry "myregistry.azurecr.io" \
  -oci-repository "packages" \
  -oci-username "myuser" \
  -oci-password "mypassword"
```

## Troubleshooting

### Common Issues

1. **Registry Connection Fails**: 
   - Verify registry URL and network connectivity
   - Check authentication credentials
   - Use `-oci-insecure` for registries with self-signed certificates

2. **No Packages Found**:
   - Verify the repository exists and contains tags
   - Check repository permissions
   - The indexer will fall back to mock packages if no real packages are found

3. **Authentication Errors**:
   - Verify username and password are correct
   - Some registries require special authentication (e.g., GitHub uses tokens as passwords)

### Debug Logging

Enable debug logging to see detailed OCI indexer operations:

```bash
package-registry -log-level debug -feature-oci-indexer -oci-registry "registry.example.com"
```

## Security Considerations

- **Credentials**: Store credentials securely and avoid passing them as command-line arguments in production
- **TLS Verification**: Only use `-oci-insecure` for testing; always use TLS in production
- **Network Access**: Ensure the package registry can access the OCI registry over the network
- **Registry Permissions**: Use least-privilege access for registry authentication