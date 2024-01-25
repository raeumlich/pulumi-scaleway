# Scaleway Resource Provider

The Scaleway Resource Provider lets you manage [Scaleway](https://scaleway.com)
resources.

**Please note that the provider is not stable yet. The SDKs might change!**

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

**!!! The Node.js package hasn't been published yet !!!**

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumi/foo
```

or `yarn`:

```bash
yarn add @pulumi/foo
```

### Python

**!!! The Python package hasn't been published yet !!!**

To use from Python, install using `pip`:

```bash
pip install pulumi_foo
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumi/pulumi-foo/sdk/go/...
```

### .NET

**!!! The .NET package hasn't been published yet !!!**

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Foo
```

## Configuration

The following configuration points are available for the `scaleway` provider:

| Configuration Key         | Environment Variable          | Description                                                                                            |
|:--------------------------|:------------------------------|:-------------------------------------------------------------------------------------------------------|
| `scaleway:accessKey`      | `SCW_ACCESS_KEY`              | The access key for authentication (required)                                                           |
| `scaleway:secretKey`      | `SCW_SECRET_KEY`              | The secret key for authentication (required)                                                           |
| `scaleway:projectId`      | `SCW_DEFAULT_PROJECT_ID`      | The project ID that will be used as default value for project-scoped resources (required)              |
| `scaleway:organizationId` | `SCW_DEFAULT_ORGANIZATION_ID` | The organization ID that will be used as default value for organization-scoped resources               |
| `scaleway:region`         | `SCW_DEFAULT_REGION`          | The region that will be used as default value for all resources. (fr-par if none specified)            |
| `scaleway:zone`           | `SCW_DEFAULT_ZONE`            | The zone that will be used as default value for all resources. (fr-par-1 if none specified)            |
| `scaleway:apiUrl`         |                               | The Scaleway API URL to use                                                                            |
| `scaleway:profile`        |                               | Name of the profile to use when you want to use the shared config file of the Scaleway developer tools |

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/foo/api-docs/).
