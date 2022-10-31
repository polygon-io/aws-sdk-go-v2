// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Describes resources needed to authenticate access to some source repositories.
// The specific resource depends on the repository provider.
type AuthenticationConfiguration struct {

	// The Amazon Resource Name (ARN) of the IAM role that grants the App Runner
	// service access to a source repository. It's required for ECR image repositories
	// (but not for ECR Public repositories).
	AccessRoleArn *string

	// The Amazon Resource Name (ARN) of the App Runner connection that enables the App
	// Runner service to connect to a source repository. It's required for GitHub code
	// repositories.
	ConnectionArn *string

	noSmithyDocumentSerde
}

// Describes an App Runner automatic scaling configuration resource. A higher
// MinSize increases the spread of your App Runner service over more Availability
// Zones in the Amazon Web Services Region. The tradeoff is a higher minimal cost.
// A lower MaxSize controls your cost. The tradeoff is lower responsiveness during
// peak demand. Multiple revisions of a configuration might have the same
// AutoScalingConfigurationName and different AutoScalingConfigurationRevision
// values.
type AutoScalingConfiguration struct {

	// The Amazon Resource Name (ARN) of this auto scaling configuration.
	AutoScalingConfigurationArn *string

	// The customer-provided auto scaling configuration name. It can be used in
	// multiple revisions of a configuration.
	AutoScalingConfigurationName *string

	// The revision of this auto scaling configuration. It's unique among all the
	// active configurations ("Status": "ACTIVE") that share the same
	// AutoScalingConfigurationName.
	AutoScalingConfigurationRevision int32

	// The time when the auto scaling configuration was created. It's in Unix time
	// stamp format.
	CreatedAt *time.Time

	// The time when the auto scaling configuration was deleted. It's in Unix time
	// stamp format.
	DeletedAt *time.Time

	// It's set to true for the configuration with the highest Revision among all
	// configurations that share the same AutoScalingConfigurationName. It's set to
	// false otherwise.
	Latest bool

	// The maximum number of concurrent requests that an instance processes. If the
	// number of concurrent requests exceeds this limit, App Runner scales the service
	// up.
	MaxConcurrency int32

	// The maximum number of instances that a service scales up to. At most MaxSize
	// instances actively serve traffic for your service.
	MaxSize int32

	// The minimum number of instances that App Runner provisions for a service. The
	// service always has at least MinSize provisioned instances. Some of them actively
	// serve traffic. The rest of them (provisioned and inactive instances) are a
	// cost-effective compute capacity reserve and are ready to be quickly activated.
	// You pay for memory usage of all the provisioned instances. You pay for CPU usage
	// of only the active subset. App Runner temporarily doubles the number of
	// provisioned instances during deployments, to maintain the same capacity for both
	// old and new code.
	MinSize int32

	// The current state of the auto scaling configuration. If the status of a
	// configuration revision is INACTIVE, it was deleted and can't be used. Inactive
	// configuration revisions are permanently removed some time after they are
	// deleted.
	Status AutoScalingConfigurationStatus

	noSmithyDocumentSerde
}

// Provides summary information about an App Runner automatic scaling configuration
// resource. This type contains limited information about an auto scaling
// configuration. It includes only identification information, without
// configuration details. It's returned by the ListAutoScalingConfigurations
// action. Complete configuration information is returned by the
// CreateAutoScalingConfiguration, DescribeAutoScalingConfiguration, and
// DeleteAutoScalingConfiguration actions using the AutoScalingConfiguration type.
type AutoScalingConfigurationSummary struct {

	// The Amazon Resource Name (ARN) of this auto scaling configuration.
	AutoScalingConfigurationArn *string

	// The customer-provided auto scaling configuration name. It can be used in
	// multiple revisions of a configuration.
	AutoScalingConfigurationName *string

	// The revision of this auto scaling configuration. It's unique among all the
	// active configurations ("Status": "ACTIVE") with the same
	// AutoScalingConfigurationName.
	AutoScalingConfigurationRevision int32

	noSmithyDocumentSerde
}

// Describes a certificate CNAME record to add to your DNS. For more information,
// see AssociateCustomDomain
// (https://docs.aws.amazon.com/apprunner/latest/api/API_AssociateCustomDomain.html).
type CertificateValidationRecord struct {

	// The certificate CNAME record name.
	Name *string

	// The current state of the certificate CNAME record validation. It should change
	// to SUCCESS after App Runner completes validation with your DNS.
	Status CertificateValidationRecordStatus

	// The record type, always CNAME.
	Type *string

	// The certificate CNAME record value.
	Value *string

	noSmithyDocumentSerde
}

// Describes the configuration that App Runner uses to build and run an App Runner
// service from a source code repository.
type CodeConfiguration struct {

	// The source of the App Runner configuration. Values are interpreted as
	// follows:
	//
	// * REPOSITORY – App Runner reads configuration values from the
	// apprunner.yaml file in the source code repository and ignores
	// CodeConfigurationValues.
	//
	// * API – App Runner uses configuration values provided
	// in CodeConfigurationValues and ignores the apprunner.yaml file in the source
	// code repository.
	//
	// This member is required.
	ConfigurationSource ConfigurationSource

	// The basic configuration for building and running the App Runner service. Use it
	// to quickly launch an App Runner service without providing a apprunner.yaml file
	// in the source code repository (or ignoring the file if it exists).
	CodeConfigurationValues *CodeConfigurationValues

	noSmithyDocumentSerde
}

// Describes the basic configuration needed for building and running an App Runner
// service. This type doesn't support the full set of possible configuration
// options. Fur full configuration capabilities, use a apprunner.yaml file in the
// source code repository.
type CodeConfigurationValues struct {

	// A runtime environment type for building and running an App Runner service. It
	// represents a programming language runtime.
	//
	// This member is required.
	Runtime Runtime

	// The command App Runner runs to build your application.
	BuildCommand *string

	// The port that your application listens to in the container. Default: 8080
	Port *string

	// The environment variables that are available to your running App Runner service.
	// An array of key-value pairs. Keys with a prefix of AWSAPPRUNNER are reserved for
	// system use and aren't valid.
	RuntimeEnvironmentVariables map[string]string

	// The command App Runner runs to start your application.
	StartCommand *string

	noSmithyDocumentSerde
}

// Describes a source code repository.
type CodeRepository struct {

	// The location of the repository that contains the source code.
	//
	// This member is required.
	RepositoryUrl *string

	// The version that should be used within the source code repository.
	//
	// This member is required.
	SourceCodeVersion *SourceCodeVersion

	// Configuration for building and running the service from a source code
	// repository. CodeConfiguration is required only for CreateService request.
	CodeConfiguration *CodeConfiguration

	noSmithyDocumentSerde
}

// Describes an App Runner connection resource.
type Connection struct {

	// The Amazon Resource Name (ARN) of this connection.
	ConnectionArn *string

	// The customer-provided connection name.
	ConnectionName *string

	// The App Runner connection creation time, expressed as a Unix time stamp.
	CreatedAt *time.Time

	// The source repository provider.
	ProviderType ProviderType

	// The current state of the App Runner connection. When the state is AVAILABLE, you
	// can use the connection to create an App Runner service.
	Status ConnectionStatus

	noSmithyDocumentSerde
}

// Provides summary information about an App Runner connection resource.
type ConnectionSummary struct {

	// The Amazon Resource Name (ARN) of this connection.
	ConnectionArn *string

	// The customer-provided connection name.
	ConnectionName *string

	// The App Runner connection creation time, expressed as a Unix time stamp.
	CreatedAt *time.Time

	// The source repository provider.
	ProviderType ProviderType

	// The current state of the App Runner connection. When the state is AVAILABLE, you
	// can use the connection to create an App Runner service.
	Status ConnectionStatus

	noSmithyDocumentSerde
}

// Describes a custom domain that's associated with an App Runner service.
type CustomDomain struct {

	// An associated custom domain endpoint. It can be a root domain (for example,
	// example.com), a subdomain (for example, login.example.com or
	// admin.login.example.com), or a wildcard (for example, *.example.com).
	//
	// This member is required.
	DomainName *string

	// When true, the subdomain www.DomainName  is associated with the App Runner
	// service in addition to the base domain.
	//
	// This member is required.
	EnableWWWSubdomain *bool

	// The current state of the domain name association.
	//
	// This member is required.
	Status CustomDomainAssociationStatus

	// A list of certificate CNAME records that's used for this domain name.
	CertificateValidationRecords []CertificateValidationRecord

	noSmithyDocumentSerde
}

// Describes configuration settings related to outbound network traffic of an App
// Runner service.
type EgressConfiguration struct {

	// The type of egress configuration. Set to DEFAULT for access to resources hosted
	// on public networks. Set to VPC to associate your service to a custom VPC
	// specified by VpcConnectorArn.
	EgressType EgressType

	// The Amazon Resource Name (ARN) of the App Runner VPC connector that you want to
	// associate with your App Runner service. Only valid when EgressType = VPC.
	VpcConnectorArn *string

	noSmithyDocumentSerde
}

// Describes a custom encryption key that App Runner uses to encrypt copies of the
// source repository and service logs.
type EncryptionConfiguration struct {

	// The ARN of the KMS key that's used for encryption.
	//
	// This member is required.
	KmsKey *string

	noSmithyDocumentSerde
}

// Describes the settings for the health check that App Runner performs to monitor
// the health of a service.
type HealthCheckConfiguration struct {

	// The number of consecutive checks that must succeed before App Runner decides
	// that the service is healthy. Default: 1
	HealthyThreshold *int32

	// The time interval, in seconds, between health checks. Default: 5
	Interval *int32

	// The URL that health check requests are sent to. Path is only applicable when you
	// set Protocol to HTTP. Default: "/"
	Path *string

	// The IP protocol that App Runner uses to perform health checks for your service.
	// If you set Protocol to HTTP, App Runner sends health check requests to the HTTP
	// path specified by Path. Default: TCP
	Protocol HealthCheckProtocol

	// The time, in seconds, to wait for a health check response before deciding it
	// failed. Default: 2
	Timeout *int32

	// The number of consecutive checks that must fail before App Runner decides that
	// the service is unhealthy. Default: 5
	UnhealthyThreshold *int32

	noSmithyDocumentSerde
}

// Describes the configuration that App Runner uses to run an App Runner service
// using an image pulled from a source image repository.
type ImageConfiguration struct {

	// The port that your application listens to in the container. Default: 8080
	Port *string

	// Environment variables that are available to your running App Runner service. An
	// array of key-value pairs. Keys with a prefix of AWSAPPRUNNER are reserved for
	// system use and aren't valid.
	RuntimeEnvironmentVariables map[string]string

	// An optional command that App Runner runs to start the application in the source
	// image. If specified, this command overrides the Docker image’s default start
	// command.
	StartCommand *string

	noSmithyDocumentSerde
}

// Describes a source image repository.
type ImageRepository struct {

	// The identifier of an image. For an image in Amazon Elastic Container Registry
	// (Amazon ECR), this is an image name. For the image name format, see Pulling an
	// image
	// (https://docs.aws.amazon.com/AmazonECR/latest/userguide/docker-pull-ecr-image.html)
	// in the Amazon ECR User Guide.
	//
	// This member is required.
	ImageIdentifier *string

	// The type of the image repository. This reflects the repository provider and
	// whether the repository is private or public.
	//
	// This member is required.
	ImageRepositoryType ImageRepositoryType

	// Configuration for running the identified image.
	ImageConfiguration *ImageConfiguration

	noSmithyDocumentSerde
}

// Network configuration settings for inbound network traffic.
type IngressConfiguration struct {

	// Specifies whether your App Runner service is publicly accessible. To make the
	// service publicly accessible set it to True. To make the service privately
	// accessible, from only within an Amazon VPC set it to False.
	IsPubliclyAccessible bool

	noSmithyDocumentSerde
}

// The configuration of your VPC and the associated VPC endpoint. The VPC endpoint
// is an Amazon Web Services PrivateLink resource that allows access to your App
// Runner services from within an Amazon VPC.
type IngressVpcConfiguration struct {

	// The ID of the VPC endpoint that your App Runner service connects to.
	VpcEndpointId *string

	// The ID of the VPC that is used for the VPC endpoint.
	VpcId *string

	noSmithyDocumentSerde
}

// Describes the runtime configuration of an App Runner service instance (scaling
// unit).
type InstanceConfiguration struct {

	// The number of CPU units reserved for each instance of your App Runner service.
	// Default: 1 vCPU
	Cpu *string

	// The Amazon Resource Name (ARN) of an IAM role that provides permissions to your
	// App Runner service. These are permissions that your code needs when it calls any
	// Amazon Web Services APIs.
	InstanceRoleArn *string

	// The amount of memory, in MB or GB, reserved for each instance of your App Runner
	// service. Default: 2 GB
	Memory *string

	noSmithyDocumentSerde
}

// Returns a list of VPC Ingress Connections based on the filter provided. It can
// return either ServiceArn or VpcEndpointId, or both.
type ListVpcIngressConnectionsFilter struct {

	// The Amazon Resource Name (ARN) of a service to filter by.
	ServiceArn *string

	// The ID of a VPC Endpoint to filter by.
	VpcEndpointId *string

	noSmithyDocumentSerde
}

// Describes configuration settings related to network traffic of an App Runner
// service. Consists of embedded objects for each configurable network feature.
type NetworkConfiguration struct {

	// Network configuration settings for outbound message traffic.
	EgressConfiguration *EgressConfiguration

	// Network configuration settings for inbound message traffic.
	IngressConfiguration *IngressConfiguration

	noSmithyDocumentSerde
}

// Describes an App Runner observability configuration resource. Multiple revisions
// of a configuration have the same ObservabilityConfigurationName and different
// ObservabilityConfigurationRevision values. The resource is designed to configure
// multiple features (currently one feature, tracing). This type contains optional
// members that describe the configuration of these features (currently one member,
// TraceConfiguration). If a feature member isn't specified, the feature isn't
// enabled.
type ObservabilityConfiguration struct {

	// The time when the observability configuration was created. It's in Unix time
	// stamp format.
	CreatedAt *time.Time

	// The time when the observability configuration was deleted. It's in Unix time
	// stamp format.
	DeletedAt *time.Time

	// It's set to true for the configuration with the highest Revision among all
	// configurations that share the same ObservabilityConfigurationName. It's set to
	// false otherwise.
	Latest bool

	// The Amazon Resource Name (ARN) of this observability configuration.
	ObservabilityConfigurationArn *string

	// The customer-provided observability configuration name. It can be used in
	// multiple revisions of a configuration.
	ObservabilityConfigurationName *string

	// The revision of this observability configuration. It's unique among all the
	// active configurations ("Status": "ACTIVE") that share the same
	// ObservabilityConfigurationName.
	ObservabilityConfigurationRevision int32

	// The current state of the observability configuration. If the status of a
	// configuration revision is INACTIVE, it was deleted and can't be used. Inactive
	// configuration revisions are permanently removed some time after they are
	// deleted.
	Status ObservabilityConfigurationStatus

	// The configuration of the tracing feature within this observability
	// configuration. If not specified, tracing isn't enabled.
	TraceConfiguration *TraceConfiguration

	noSmithyDocumentSerde
}

// Provides summary information about an App Runner observability configuration
// resource. This type contains limited information about an observability
// configuration. It includes only identification information, without
// configuration details. It's returned by the ListObservabilityConfigurations
// action. Complete configuration information is returned by the
// CreateObservabilityConfiguration, DescribeObservabilityConfiguration, and
// DeleteObservabilityConfiguration actions using the ObservabilityConfiguration
// type.
type ObservabilityConfigurationSummary struct {

	// The Amazon Resource Name (ARN) of this observability configuration.
	ObservabilityConfigurationArn *string

	// The customer-provided observability configuration name. It can be used in
	// multiple revisions of a configuration.
	ObservabilityConfigurationName *string

	// The revision of this observability configuration. It's unique among all the
	// active configurations ("Status": "ACTIVE") that share the same
	// ObservabilityConfigurationName.
	ObservabilityConfigurationRevision int32

	noSmithyDocumentSerde
}

// Provides summary information for an operation that occurred on an App Runner
// service.
type OperationSummary struct {

	// The time when the operation ended. It's in the Unix time stamp format.
	EndedAt *time.Time

	// A unique ID of this operation. It's unique in the scope of the App Runner
	// service.
	Id *string

	// The time when the operation started. It's in the Unix time stamp format.
	StartedAt *time.Time

	// The current state of the operation.
	Status OperationStatus

	// The Amazon Resource Name (ARN) of the resource that the operation acted on (for
	// example, an App Runner service).
	TargetArn *string

	// The type of operation. It indicates a specific action that occured.
	Type OperationType

	// The time when the operation was last updated. It's in the Unix time stamp
	// format.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// Describes an App Runner service. It can describe a service in any state,
// including deleted services. This type contains the full information about a
// service, including configuration details. It's returned by the CreateService
// (https://docs.aws.amazon.com/apprunner/latest/api/API_CreateService.html),
// DescribeService
// (https://docs.aws.amazon.com/apprunner/latest/api/API_DescribeService.html), and
// DeleteService
// (https://docs.aws.amazon.com/apprunner/latest/api/API_DeleteService.html)
// actions. A subset of this information is returned by the ListServices
// (https://docs.aws.amazon.com/apprunner/latest/api/API_ListServices.html) action
// using the ServiceSummary
// (https://docs.aws.amazon.com/apprunner/latest/api/API_ServiceSummary.html) type.
type Service struct {

	// Summary information for the App Runner automatic scaling configuration resource
	// that's associated with this service.
	//
	// This member is required.
	AutoScalingConfigurationSummary *AutoScalingConfigurationSummary

	// The time when the App Runner service was created. It's in the Unix time stamp
	// format.
	//
	// This member is required.
	CreatedAt *time.Time

	// The runtime configuration of instances (scaling units) of this service.
	//
	// This member is required.
	InstanceConfiguration *InstanceConfiguration

	// Configuration settings related to network traffic of the web application that
	// this service runs.
	//
	// This member is required.
	NetworkConfiguration *NetworkConfiguration

	// The Amazon Resource Name (ARN) of this service.
	//
	// This member is required.
	ServiceArn *string

	// An ID that App Runner generated for this service. It's unique within the Amazon
	// Web Services Region.
	//
	// This member is required.
	ServiceId *string

	// The customer-provided service name.
	//
	// This member is required.
	ServiceName *string

	// The source deployed to the App Runner service. It can be a code or an image
	// repository.
	//
	// This member is required.
	SourceConfiguration *SourceConfiguration

	// The current state of the App Runner service. These particular values mean the
	// following.
	//
	// * CREATE_FAILED – The service failed to create. To troubleshoot this
	// failure, read the failure events and logs, change any parameters that need to be
	// fixed, and retry the call to create the service. The failed service isn't
	// usable, and still counts towards your service quota. When you're done analyzing
	// the failure, delete the service.
	//
	// * DELETE_FAILED – The service failed to delete
	// and can't be successfully recovered. Retry the service deletion call to ensure
	// that all related resources are removed.
	//
	// This member is required.
	Status ServiceStatus

	// The time when the App Runner service was last updated at. It's in the Unix time
	// stamp format.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The time when the App Runner service was deleted. It's in the Unix time stamp
	// format.
	DeletedAt *time.Time

	// The encryption key that App Runner uses to encrypt the service logs and the copy
	// of the source repository that App Runner maintains for the service. It can be
	// either a customer-provided encryption key or an Amazon Web Services managed key.
	EncryptionConfiguration *EncryptionConfiguration

	// The settings for the health check that App Runner performs to monitor the health
	// of this service.
	HealthCheckConfiguration *HealthCheckConfiguration

	// The observability configuration of this service.
	ObservabilityConfiguration *ServiceObservabilityConfiguration

	// A subdomain URL that App Runner generated for this service. You can use this URL
	// to access your service web application.
	ServiceUrl *string

	noSmithyDocumentSerde
}

// Describes the observability configuration of an App Runner service. These are
// additional observability features, like tracing, that you choose to enable.
// They're configured in a separate resource that you associate with your service.
type ServiceObservabilityConfiguration struct {

	// When true, an observability configuration resource is associated with the
	// service, and an ObservabilityConfigurationArn is specified.
	//
	// This member is required.
	ObservabilityEnabled bool

	// The Amazon Resource Name (ARN) of the observability configuration that is
	// associated with the service. Specified only when ObservabilityEnabled is true.
	// Specify an ARN with a name and a revision number to associate that revision. For
	// example:
	// arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing/3
	// Specify just the name to associate the latest revision. For example:
	// arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing
	ObservabilityConfigurationArn *string

	noSmithyDocumentSerde
}

// Provides summary information for an App Runner service. This type contains
// limited information about a service. It doesn't include configuration details.
// It's returned by the ListServices
// (https://docs.aws.amazon.com/apprunner/latest/api/API_ListServices.html) action.
// Complete service information is returned by the CreateService
// (https://docs.aws.amazon.com/apprunner/latest/api/API_CreateService.html),
// DescribeService
// (https://docs.aws.amazon.com/apprunner/latest/api/API_DescribeService.html), and
// DeleteService
// (https://docs.aws.amazon.com/apprunner/latest/api/API_DeleteService.html)
// actions using the Service
// (https://docs.aws.amazon.com/apprunner/latest/api/API_Service.html) type.
type ServiceSummary struct {

	// The time when the App Runner service was created. It's in the Unix time stamp
	// format.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of this service.
	ServiceArn *string

	// An ID that App Runner generated for this service. It's unique within the Amazon
	// Web Services Region.
	ServiceId *string

	// The customer-provided service name.
	ServiceName *string

	// A subdomain URL that App Runner generated for this service. You can use this URL
	// to access your service web application.
	ServiceUrl *string

	// The current state of the App Runner service. These particular values mean the
	// following.
	//
	// * CREATE_FAILED – The service failed to create. Read the failure
	// events and logs, change any parameters that need to be fixed, and retry the call
	// to create the service. The failed service isn't usable, and still counts towards
	// your service quota. When you're done analyzing the failure, delete the
	// service.
	//
	// * DELETE_FAILED – The service failed to delete and can't be
	// successfully recovered. Retry the service deletion call to ensure that all
	// related resources are removed.
	Status ServiceStatus

	// The time when the App Runner service was last updated. It's in theUnix time
	// stamp format.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// Identifies a version of code that App Runner refers to within a source code
// repository.
type SourceCodeVersion struct {

	// The type of version identifier. For a git-based repository, branches represent
	// versions.
	//
	// This member is required.
	Type SourceCodeVersionType

	// A source code version. For a git-based repository, a branch name maps to a
	// specific version. App Runner uses the most recent commit to the branch.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// Describes the source deployed to an App Runner service. It can be a code or an
// image repository.
type SourceConfiguration struct {

	// Describes the resources that are needed to authenticate access to some source
	// repositories.
	AuthenticationConfiguration *AuthenticationConfiguration

	// If true, continuous integration from the source repository is enabled for the
	// App Runner service. Each repository change (including any source code commit or
	// new image version) starts a deployment. Default: App Runner sets to false for a
	// source image that uses an ECR Public repository or an ECR repository that's in
	// an Amazon Web Services account other than the one that the service is in. App
	// Runner sets to true in all other cases (which currently include a source code
	// repository or a source image using a same-account ECR repository).
	AutoDeploymentsEnabled *bool

	// The description of a source code repository. You must provide either this member
	// or ImageRepository (but not both).
	CodeRepository *CodeRepository

	// The description of a source image repository. You must provide either this
	// member or CodeRepository (but not both).
	ImageRepository *ImageRepository

	noSmithyDocumentSerde
}

// Describes a tag that is applied to an App Runner resource. A tag is a metadata
// item consisting of a key-value pair.
type Tag struct {

	// The key of the tag.
	Key *string

	// The value of the tag.
	Value *string

	noSmithyDocumentSerde
}

// Describes the configuration of the tracing feature within an App Runner
// observability configuration.
type TraceConfiguration struct {

	// The implementation provider chosen for tracing App Runner services.
	//
	// This member is required.
	Vendor TracingVendor

	noSmithyDocumentSerde
}

// Describes an App Runner VPC connector resource. A VPC connector describes the
// Amazon Virtual Private Cloud (Amazon VPC) that an App Runner service is
// associated with, and the subnets and security group that are used. Multiple
// revisions of a connector might have the same Name and different Revision values.
// At this time, App Runner supports only one revision per name.
type VpcConnector struct {

	// The time when the VPC connector was created. It's in Unix time stamp format.
	CreatedAt *time.Time

	// The time when the VPC connector was deleted. It's in Unix time stamp format.
	DeletedAt *time.Time

	// A list of IDs of security groups that App Runner uses for access to Amazon Web
	// Services resources under the specified subnets. If not specified, App Runner
	// uses the default security group of the Amazon VPC. The default security group
	// allows all outbound traffic.
	SecurityGroups []string

	// The current state of the VPC connector. If the status of a connector revision is
	// INACTIVE, it was deleted and can't be used. Inactive connector revisions are
	// permanently removed some time after they are deleted.
	Status VpcConnectorStatus

	// A list of IDs of subnets that App Runner uses for your service. All IDs are of
	// subnets of a single Amazon VPC.
	Subnets []string

	// The Amazon Resource Name (ARN) of this VPC connector.
	VpcConnectorArn *string

	// The customer-provided VPC connector name.
	VpcConnectorName *string

	// The revision of this VPC connector. It's unique among all the active connectors
	// ("Status": "ACTIVE") that share the same Name. At this time, App Runner supports
	// only one revision per name.
	VpcConnectorRevision int32

	noSmithyDocumentSerde
}

// DNS Target record for a custom domain of this Amazon VPC.
type VpcDNSTarget struct {

	// The domain name of your target DNS that is associated with the Amazon VPC.
	DomainName *string

	// The ID of the Amazon VPC that is associated with the custom domain name of the
	// target DNS.
	VpcId *string

	// The Amazon Resource Name (ARN) of the VPC Ingress Connection that is associated
	// with your service.
	VpcIngressConnectionArn *string

	noSmithyDocumentSerde
}

// The App Runner resource that specifies an App Runner endpoint for incoming
// traffic. It establishes a connection between a VPC interface endpoint and a App
// Runner service, to make your App Runner service accessible from only within an
// Amazon VPC.
type VpcIngressConnection struct {

	// The Account Id you use to create the VPC Ingress Connection resource.
	AccountId *string

	// The time when the VPC Ingress Connection was created. It's in the Unix time
	// stamp format.
	//
	// * Type: Timestamp
	//
	// * Required: Yes
	CreatedAt *time.Time

	// The time when the App Runner service was deleted. It's in the Unix time stamp
	// format.
	//
	// * Type: Timestamp
	//
	// * Required: No
	DeletedAt *time.Time

	// The domain name associated with the VPC Ingress Connection resource.
	DomainName *string

	// Specifications for the customer’s VPC and related PrivateLink VPC endpoint that
	// are used to associate with the VPC Ingress Connection resource.
	IngressVpcConfiguration *IngressVpcConfiguration

	// The Amazon Resource Name (ARN) of the service associated with the VPC Ingress
	// Connection.
	ServiceArn *string

	// The current status of the VPC Ingress Connection. The VPC Ingress Connection
	// displays one of the following statuses: AVAILABLE, PENDING_CREATION,
	// PENDING_UPDATE, PENDING_DELETION,FAILED_CREATION, FAILED_UPDATE,
	// FAILED_DELETION, and DELETED..
	Status VpcIngressConnectionStatus

	// The Amazon Resource Name (ARN) of the VPC Ingress Connection.
	VpcIngressConnectionArn *string

	// The customer-provided VPC Ingress Connection name.
	VpcIngressConnectionName *string

	noSmithyDocumentSerde
}

// Provides summary information about an VPC Ingress Connection, which includes its
// VPC Ingress Connection ARN and its associated Service ARN.
type VpcIngressConnectionSummary struct {

	// The Amazon Resource Name (ARN) of the service associated with the VPC Ingress
	// Connection.
	ServiceArn *string

	// The Amazon Resource Name (ARN) of the VPC Ingress Connection.
	VpcIngressConnectionArn *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
