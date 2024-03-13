// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rdb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway Database Instances.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
//
// ## Example Usage
//
// ### Example Basic
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/rdb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rdb.NewInstance(ctx, "main", &rdb.InstanceArgs{
//				DisableBackup: pulumi.Bool(true),
//				Engine:        pulumi.String("PostgreSQL-11"),
//				IsHaCluster:   pulumi.Bool(true),
//				NodeType:      pulumi.String("DB-DEV-S"),
//				Password:      pulumi.String("thiZ_is_v&ry_s3cret"),
//				UserName:      pulumi.String("my_initial_user"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Example with Settings
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/rdb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rdb.NewInstance(ctx, "main", &rdb.InstanceArgs{
//				DisableBackup: pulumi.Bool(true),
//				Engine:        pulumi.String("MySQL-8"),
//				InitSettings: pulumi.StringMap{
//					"lower_case_table_names": pulumi.String("1"),
//				},
//				NodeType: pulumi.String("db-dev-s"),
//				Password: pulumi.String("thiZ_is_v&ry_s3cret"),
//				Settings: pulumi.StringMap{
//					"max_connections": pulumi.String("350"),
//				},
//				UserName: pulumi.String("my_initial_user"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Example with backup schedule
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/rdb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rdb.NewInstance(ctx, "main", &rdb.InstanceArgs{
//				BackupScheduleFrequency: pulumi.Int(24),
//				BackupScheduleRetention: pulumi.Int(7),
//				DisableBackup:           pulumi.Bool(false),
//				Engine:                  pulumi.String("PostgreSQL-11"),
//				IsHaCluster:             pulumi.Bool(true),
//				NodeType:                pulumi.String("DB-DEV-S"),
//				Password:                pulumi.String("thiZ_is_v&ry_s3cret"),
//				UserName:                pulumi.String("my_initial_user"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Examples of endpoints configuration
//
// RDB Instances can have a maximum of 1 public endpoint and 1 private endpoint. It can have both, or none.
//
// ### 1 static private network endpoint
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/rdb"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			pn, err := vpc.NewPrivateNetwork(ctx, "pn", &vpc.PrivateNetworkArgs{
//				Ipv4Subnet: &vpc.PrivateNetworkIpv4SubnetArgs{
//					Subnet: pulumi.String("172.16.20.0/22"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rdb.NewInstance(ctx, "main", &rdb.InstanceArgs{
//				NodeType: pulumi.String("db-dev-s"),
//				Engine:   pulumi.String("PostgreSQL-11"),
//				PrivateNetwork: &rdb.InstancePrivateNetworkArgs{
//					PnId:  pn.ID(),
//					IpNet: pulumi.String("172.16.20.4/22"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### 1 IPAM private network endpoint + 1 public endpoint
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/rdb"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			pn, err := vpc.NewPrivateNetwork(ctx, "pn", nil)
//			if err != nil {
//				return err
//			}
//			_, err = rdb.NewInstance(ctx, "main", &rdb.InstanceArgs{
//				NodeType: pulumi.String("DB-DEV-S"),
//				Engine:   pulumi.String("PostgreSQL-11"),
//				PrivateNetwork: &rdb.InstancePrivateNetworkArgs{
//					PnId:       pn.ID(),
//					EnableIpam: pulumi.Bool(true),
//				},
//				LoadBalancers: rdb.InstanceLoadBalancerArray{
//					nil,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Default: 1 public endpoint
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/rdb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rdb.NewInstance(ctx, "main", &rdb.InstanceArgs{
//				Engine:   pulumi.String("PostgreSQL-11"),
//				NodeType: pulumi.String("db-dev-s"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// > If nothing is defined, your instance will have a default public load-balancer endpoint
//
// ## Limitations
//
// The Managed Database product is only compliant with the private network in the default availability zone (AZ).
// i.e. `fr-par-1`, `nl-ams-1`, `pl-waw-1`. To learn more, read our
// section [How to connect a PostgreSQL and MySQL Database Instance to a Private Network](https://www.scaleway.com/en/docs/managed-databases/postgresql-and-mysql/how-to/connect-database-private-network/)
//
// ## Import
//
// Database Instance can be imported using the `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:rdb/instance:Instance rdb01 fr-par/11111111-1111-1111-1111-111111111111
// ```
type Instance struct {
	pulumi.CustomResourceState

	// Boolean to store logical backups in the same region as the database instance.
	BackupSameRegion pulumi.BoolOutput `pulumi:"backupSameRegion"`
	// Backup schedule frequency in hours.
	BackupScheduleFrequency pulumi.IntOutput `pulumi:"backupScheduleFrequency"`
	// Backup schedule retention in days.
	BackupScheduleRetention pulumi.IntOutput `pulumi:"backupScheduleRetention"`
	// Certificate of the database instance.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Disable automated backup for the database instance.
	DisableBackup pulumi.BoolPtrOutput `pulumi:"disableBackup"`
	// (Deprecated) The IP of the Database Instance.
	//
	// Deprecated: Please use the private_network or the load_balancer attribute
	EndpointIp pulumi.StringOutput `pulumi:"endpointIp"`
	// (Deprecated) The port of the Database Instance.
	EndpointPort pulumi.IntOutput `pulumi:"endpointPort"`
	// Database Instance's engine version (e.g. `PostgreSQL-11`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Map of engine settings to be set at database initialisation.
	//
	// > **Important:** Updates to `initSettings` will recreate the Database Instance.
	//
	// Please consult the [GoDoc](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@v1.0.0-beta.9/api/rdb/v1#EngineVersion) to list all available `settings` and `initSettings` for the `nodeType` of your convenience.
	InitSettings pulumi.StringMapOutput `pulumi:"initSettings"`
	// Enable or disable high availability for the database instance.
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	IsHaCluster pulumi.BoolPtrOutput `pulumi:"isHaCluster"`
	// List of load balancer endpoints of the database instance. A load-balancer endpoint will be set by default if no private network is.
	// This block must be defined if you want a public endpoint in addition to your private endpoint.
	LoadBalancers InstanceLoadBalancerArrayOutput `pulumi:"loadBalancers"`
	// The name of the Database Instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of database instance you want to create (e.g. `db-dev-s`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	//
	// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
	// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
	NodeType pulumi.StringOutput `pulumi:"nodeType"`
	// The organization ID the Database Instance is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Password for the first user of the database instance.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// List of private networks endpoints of the database instance.
	PrivateNetwork InstancePrivateNetworkPtrOutput `pulumi:"privateNetwork"`
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// List of read replicas of the database instance.
	ReadReplicas InstanceReadReplicaArrayOutput `pulumi:"readReplicas"`
	// `region`) The region
	// in which the Database Instance should be created.
	Region pulumi.StringOutput `pulumi:"region"`
	// Map of engine settings to be set. Using this option will override default config.
	Settings pulumi.StringMapOutput `pulumi:"settings"`
	// The tags associated with the Database Instance.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName pulumi.StringOutput `pulumi:"userName"`
	// Volume size (in GB). Cannot be used when `volumeType` is set to `lssd`.
	//
	// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
	VolumeSizeInGb pulumi.IntOutput `pulumi:"volumeSizeInGb"`
	// Type of volume where data are stored (`bssd`, `lssd` or `sbs5k`).
	VolumeType pulumi.StringPtrOutput `pulumi:"volumeType"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Engine == nil {
		return nil, errors.New("invalid value for required argument 'Engine'")
	}
	if args.NodeType == nil {
		return nil, errors.New("invalid value for required argument 'NodeType'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("scaleway:rdb/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("scaleway:rdb/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Boolean to store logical backups in the same region as the database instance.
	BackupSameRegion *bool `pulumi:"backupSameRegion"`
	// Backup schedule frequency in hours.
	BackupScheduleFrequency *int `pulumi:"backupScheduleFrequency"`
	// Backup schedule retention in days.
	BackupScheduleRetention *int `pulumi:"backupScheduleRetention"`
	// Certificate of the database instance.
	Certificate *string `pulumi:"certificate"`
	// Disable automated backup for the database instance.
	DisableBackup *bool `pulumi:"disableBackup"`
	// (Deprecated) The IP of the Database Instance.
	//
	// Deprecated: Please use the private_network or the load_balancer attribute
	EndpointIp *string `pulumi:"endpointIp"`
	// (Deprecated) The port of the Database Instance.
	EndpointPort *int `pulumi:"endpointPort"`
	// Database Instance's engine version (e.g. `PostgreSQL-11`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine *string `pulumi:"engine"`
	// Map of engine settings to be set at database initialisation.
	//
	// > **Important:** Updates to `initSettings` will recreate the Database Instance.
	//
	// Please consult the [GoDoc](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@v1.0.0-beta.9/api/rdb/v1#EngineVersion) to list all available `settings` and `initSettings` for the `nodeType` of your convenience.
	InitSettings map[string]string `pulumi:"initSettings"`
	// Enable or disable high availability for the database instance.
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	IsHaCluster *bool `pulumi:"isHaCluster"`
	// List of load balancer endpoints of the database instance. A load-balancer endpoint will be set by default if no private network is.
	// This block must be defined if you want a public endpoint in addition to your private endpoint.
	LoadBalancers []InstanceLoadBalancer `pulumi:"loadBalancers"`
	// The name of the Database Instance.
	Name *string `pulumi:"name"`
	// The type of database instance you want to create (e.g. `db-dev-s`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	//
	// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
	// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
	NodeType *string `pulumi:"nodeType"`
	// The organization ID the Database Instance is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// Password for the first user of the database instance.
	Password *string `pulumi:"password"`
	// List of private networks endpoints of the database instance.
	PrivateNetwork *InstancePrivateNetwork `pulumi:"privateNetwork"`
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId *string `pulumi:"projectId"`
	// List of read replicas of the database instance.
	ReadReplicas []InstanceReadReplica `pulumi:"readReplicas"`
	// `region`) The region
	// in which the Database Instance should be created.
	Region *string `pulumi:"region"`
	// Map of engine settings to be set. Using this option will override default config.
	Settings map[string]string `pulumi:"settings"`
	// The tags associated with the Database Instance.
	Tags []string `pulumi:"tags"`
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName *string `pulumi:"userName"`
	// Volume size (in GB). Cannot be used when `volumeType` is set to `lssd`.
	//
	// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
	VolumeSizeInGb *int `pulumi:"volumeSizeInGb"`
	// Type of volume where data are stored (`bssd`, `lssd` or `sbs5k`).
	VolumeType *string `pulumi:"volumeType"`
}

type InstanceState struct {
	// Boolean to store logical backups in the same region as the database instance.
	BackupSameRegion pulumi.BoolPtrInput
	// Backup schedule frequency in hours.
	BackupScheduleFrequency pulumi.IntPtrInput
	// Backup schedule retention in days.
	BackupScheduleRetention pulumi.IntPtrInput
	// Certificate of the database instance.
	Certificate pulumi.StringPtrInput
	// Disable automated backup for the database instance.
	DisableBackup pulumi.BoolPtrInput
	// (Deprecated) The IP of the Database Instance.
	//
	// Deprecated: Please use the private_network or the load_balancer attribute
	EndpointIp pulumi.StringPtrInput
	// (Deprecated) The port of the Database Instance.
	EndpointPort pulumi.IntPtrInput
	// Database Instance's engine version (e.g. `PostgreSQL-11`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine pulumi.StringPtrInput
	// Map of engine settings to be set at database initialisation.
	//
	// > **Important:** Updates to `initSettings` will recreate the Database Instance.
	//
	// Please consult the [GoDoc](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@v1.0.0-beta.9/api/rdb/v1#EngineVersion) to list all available `settings` and `initSettings` for the `nodeType` of your convenience.
	InitSettings pulumi.StringMapInput
	// Enable or disable high availability for the database instance.
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	IsHaCluster pulumi.BoolPtrInput
	// List of load balancer endpoints of the database instance. A load-balancer endpoint will be set by default if no private network is.
	// This block must be defined if you want a public endpoint in addition to your private endpoint.
	LoadBalancers InstanceLoadBalancerArrayInput
	// The name of the Database Instance.
	Name pulumi.StringPtrInput
	// The type of database instance you want to create (e.g. `db-dev-s`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	//
	// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
	// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
	NodeType pulumi.StringPtrInput
	// The organization ID the Database Instance is associated with.
	OrganizationId pulumi.StringPtrInput
	// Password for the first user of the database instance.
	Password pulumi.StringPtrInput
	// List of private networks endpoints of the database instance.
	PrivateNetwork InstancePrivateNetworkPtrInput
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId pulumi.StringPtrInput
	// List of read replicas of the database instance.
	ReadReplicas InstanceReadReplicaArrayInput
	// `region`) The region
	// in which the Database Instance should be created.
	Region pulumi.StringPtrInput
	// Map of engine settings to be set. Using this option will override default config.
	Settings pulumi.StringMapInput
	// The tags associated with the Database Instance.
	Tags pulumi.StringArrayInput
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName pulumi.StringPtrInput
	// Volume size (in GB). Cannot be used when `volumeType` is set to `lssd`.
	//
	// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
	VolumeSizeInGb pulumi.IntPtrInput
	// Type of volume where data are stored (`bssd`, `lssd` or `sbs5k`).
	VolumeType pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Boolean to store logical backups in the same region as the database instance.
	BackupSameRegion *bool `pulumi:"backupSameRegion"`
	// Backup schedule frequency in hours.
	BackupScheduleFrequency *int `pulumi:"backupScheduleFrequency"`
	// Backup schedule retention in days.
	BackupScheduleRetention *int `pulumi:"backupScheduleRetention"`
	// Disable automated backup for the database instance.
	DisableBackup *bool `pulumi:"disableBackup"`
	// Database Instance's engine version (e.g. `PostgreSQL-11`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine string `pulumi:"engine"`
	// Map of engine settings to be set at database initialisation.
	//
	// > **Important:** Updates to `initSettings` will recreate the Database Instance.
	//
	// Please consult the [GoDoc](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@v1.0.0-beta.9/api/rdb/v1#EngineVersion) to list all available `settings` and `initSettings` for the `nodeType` of your convenience.
	InitSettings map[string]string `pulumi:"initSettings"`
	// Enable or disable high availability for the database instance.
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	IsHaCluster *bool `pulumi:"isHaCluster"`
	// List of load balancer endpoints of the database instance. A load-balancer endpoint will be set by default if no private network is.
	// This block must be defined if you want a public endpoint in addition to your private endpoint.
	LoadBalancers []InstanceLoadBalancer `pulumi:"loadBalancers"`
	// The name of the Database Instance.
	Name *string `pulumi:"name"`
	// The type of database instance you want to create (e.g. `db-dev-s`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	//
	// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
	// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
	NodeType string `pulumi:"nodeType"`
	// Password for the first user of the database instance.
	Password *string `pulumi:"password"`
	// List of private networks endpoints of the database instance.
	PrivateNetwork *InstancePrivateNetwork `pulumi:"privateNetwork"`
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region
	// in which the Database Instance should be created.
	Region *string `pulumi:"region"`
	// Map of engine settings to be set. Using this option will override default config.
	Settings map[string]string `pulumi:"settings"`
	// The tags associated with the Database Instance.
	Tags []string `pulumi:"tags"`
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName *string `pulumi:"userName"`
	// Volume size (in GB). Cannot be used when `volumeType` is set to `lssd`.
	//
	// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
	VolumeSizeInGb *int `pulumi:"volumeSizeInGb"`
	// Type of volume where data are stored (`bssd`, `lssd` or `sbs5k`).
	VolumeType *string `pulumi:"volumeType"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Boolean to store logical backups in the same region as the database instance.
	BackupSameRegion pulumi.BoolPtrInput
	// Backup schedule frequency in hours.
	BackupScheduleFrequency pulumi.IntPtrInput
	// Backup schedule retention in days.
	BackupScheduleRetention pulumi.IntPtrInput
	// Disable automated backup for the database instance.
	DisableBackup pulumi.BoolPtrInput
	// Database Instance's engine version (e.g. `PostgreSQL-11`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine pulumi.StringInput
	// Map of engine settings to be set at database initialisation.
	//
	// > **Important:** Updates to `initSettings` will recreate the Database Instance.
	//
	// Please consult the [GoDoc](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@v1.0.0-beta.9/api/rdb/v1#EngineVersion) to list all available `settings` and `initSettings` for the `nodeType` of your convenience.
	InitSettings pulumi.StringMapInput
	// Enable or disable high availability for the database instance.
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	IsHaCluster pulumi.BoolPtrInput
	// List of load balancer endpoints of the database instance. A load-balancer endpoint will be set by default if no private network is.
	// This block must be defined if you want a public endpoint in addition to your private endpoint.
	LoadBalancers InstanceLoadBalancerArrayInput
	// The name of the Database Instance.
	Name pulumi.StringPtrInput
	// The type of database instance you want to create (e.g. `db-dev-s`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	//
	// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
	// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
	NodeType pulumi.StringInput
	// Password for the first user of the database instance.
	Password pulumi.StringPtrInput
	// List of private networks endpoints of the database instance.
	PrivateNetwork InstancePrivateNetworkPtrInput
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region
	// in which the Database Instance should be created.
	Region pulumi.StringPtrInput
	// Map of engine settings to be set. Using this option will override default config.
	Settings pulumi.StringMapInput
	// The tags associated with the Database Instance.
	Tags pulumi.StringArrayInput
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName pulumi.StringPtrInput
	// Volume size (in GB). Cannot be used when `volumeType` is set to `lssd`.
	//
	// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
	VolumeSizeInGb pulumi.IntPtrInput
	// Type of volume where data are stored (`bssd`, `lssd` or `sbs5k`).
	VolumeType pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// Boolean to store logical backups in the same region as the database instance.
func (o InstanceOutput) BackupSameRegion() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.BackupSameRegion }).(pulumi.BoolOutput)
}

// Backup schedule frequency in hours.
func (o InstanceOutput) BackupScheduleFrequency() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.BackupScheduleFrequency }).(pulumi.IntOutput)
}

// Backup schedule retention in days.
func (o InstanceOutput) BackupScheduleRetention() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.BackupScheduleRetention }).(pulumi.IntOutput)
}

// Certificate of the database instance.
func (o InstanceOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// Disable automated backup for the database instance.
func (o InstanceOutput) DisableBackup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.DisableBackup }).(pulumi.BoolPtrOutput)
}

// (Deprecated) The IP of the Database Instance.
//
// Deprecated: Please use the private_network or the load_balancer attribute
func (o InstanceOutput) EndpointIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.EndpointIp }).(pulumi.StringOutput)
}

// (Deprecated) The port of the Database Instance.
func (o InstanceOutput) EndpointPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.EndpointPort }).(pulumi.IntOutput)
}

// Database Instance's engine version (e.g. `PostgreSQL-11`).
//
// > **Important:** Updates to `engine` will recreate the Database Instance.
func (o InstanceOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// Map of engine settings to be set at database initialisation.
//
// > **Important:** Updates to `initSettings` will recreate the Database Instance.
//
// Please consult the [GoDoc](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@v1.0.0-beta.9/api/rdb/v1#EngineVersion) to list all available `settings` and `initSettings` for the `nodeType` of your convenience.
func (o InstanceOutput) InitSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringMapOutput { return v.InitSettings }).(pulumi.StringMapOutput)
}

// Enable or disable high availability for the database instance.
//
// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
func (o InstanceOutput) IsHaCluster() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.IsHaCluster }).(pulumi.BoolPtrOutput)
}

// List of load balancer endpoints of the database instance. A load-balancer endpoint will be set by default if no private network is.
// This block must be defined if you want a public endpoint in addition to your private endpoint.
func (o InstanceOutput) LoadBalancers() InstanceLoadBalancerArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceLoadBalancerArrayOutput { return v.LoadBalancers }).(InstanceLoadBalancerArrayOutput)
}

// The name of the Database Instance.
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of database instance you want to create (e.g. `db-dev-s`).
//
// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
// interruption. Keep in mind that you cannot downgrade a Database Instance.
//
// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
func (o InstanceOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.NodeType }).(pulumi.StringOutput)
}

// The organization ID the Database Instance is associated with.
func (o InstanceOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// Password for the first user of the database instance.
func (o InstanceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// List of private networks endpoints of the database instance.
func (o InstanceOutput) PrivateNetwork() InstancePrivateNetworkPtrOutput {
	return o.ApplyT(func(v *Instance) InstancePrivateNetworkPtrOutput { return v.PrivateNetwork }).(InstancePrivateNetworkPtrOutput)
}

// `projectId`) The ID of the project the Database
// Instance is associated with.
func (o InstanceOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// List of read replicas of the database instance.
func (o InstanceOutput) ReadReplicas() InstanceReadReplicaArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceReadReplicaArrayOutput { return v.ReadReplicas }).(InstanceReadReplicaArrayOutput)
}

// `region`) The region
// in which the Database Instance should be created.
func (o InstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Map of engine settings to be set. Using this option will override default config.
func (o InstanceOutput) Settings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringMapOutput { return v.Settings }).(pulumi.StringMapOutput)
}

// The tags associated with the Database Instance.
func (o InstanceOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Identifier for the first user of the database instance.
//
// > **Important:** Updates to `userName` will recreate the Database Instance.
func (o InstanceOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

// Volume size (in GB). Cannot be used when `volumeType` is set to `lssd`.
//
// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
func (o InstanceOutput) VolumeSizeInGb() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.VolumeSizeInGb }).(pulumi.IntOutput)
}

// Type of volume where data are stored (`bssd`, `lssd` or `sbs5k`).
func (o InstanceOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.VolumeType }).(pulumi.StringPtrOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
