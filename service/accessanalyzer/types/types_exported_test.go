// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
)

func ExampleAclGrantee_outputUsage() {
	var union types.AclGrantee
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AclGranteeMemberId:
		_ = v.Value // Value is string

	case *types.AclGranteeMemberUri:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *string

func ExampleAnalyzerConfiguration_outputUsage() {
	var union types.AnalyzerConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AnalyzerConfigurationMemberUnusedAccess:
		_ = v.Value // Value is types.UnusedAccessConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.UnusedAccessConfiguration

func ExampleConfiguration_outputUsage() {
	var union types.Configuration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ConfigurationMemberEbsSnapshot:
		_ = v.Value // Value is types.EbsSnapshotConfiguration

	case *types.ConfigurationMemberEcrRepository:
		_ = v.Value // Value is types.EcrRepositoryConfiguration

	case *types.ConfigurationMemberEfsFileSystem:
		_ = v.Value // Value is types.EfsFileSystemConfiguration

	case *types.ConfigurationMemberIamRole:
		_ = v.Value // Value is types.IamRoleConfiguration

	case *types.ConfigurationMemberKmsKey:
		_ = v.Value // Value is types.KmsKeyConfiguration

	case *types.ConfigurationMemberRdsDbClusterSnapshot:
		_ = v.Value // Value is types.RdsDbClusterSnapshotConfiguration

	case *types.ConfigurationMemberRdsDbSnapshot:
		_ = v.Value // Value is types.RdsDbSnapshotConfiguration

	case *types.ConfigurationMemberS3Bucket:
		_ = v.Value // Value is types.S3BucketConfiguration

	case *types.ConfigurationMemberS3ExpressDirectoryBucket:
		_ = v.Value // Value is types.S3ExpressDirectoryBucketConfiguration

	case *types.ConfigurationMemberSecretsManagerSecret:
		_ = v.Value // Value is types.SecretsManagerSecretConfiguration

	case *types.ConfigurationMemberSnsTopic:
		_ = v.Value // Value is types.SnsTopicConfiguration

	case *types.ConfigurationMemberSqsQueue:
		_ = v.Value // Value is types.SqsQueueConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.S3ExpressDirectoryBucketConfiguration
var _ *types.SnsTopicConfiguration
var _ *types.SqsQueueConfiguration
var _ *types.EcrRepositoryConfiguration
var _ *types.IamRoleConfiguration
var _ *types.RdsDbClusterSnapshotConfiguration
var _ *types.RdsDbSnapshotConfiguration
var _ *types.SecretsManagerSecretConfiguration
var _ *types.EfsFileSystemConfiguration
var _ *types.S3BucketConfiguration
var _ *types.KmsKeyConfiguration
var _ *types.EbsSnapshotConfiguration

func ExampleFindingDetails_outputUsage() {
	var union types.FindingDetails
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FindingDetailsMemberExternalAccessDetails:
		_ = v.Value // Value is types.ExternalAccessDetails

	case *types.FindingDetailsMemberUnusedIamRoleDetails:
		_ = v.Value // Value is types.UnusedIamRoleDetails

	case *types.FindingDetailsMemberUnusedIamUserAccessKeyDetails:
		_ = v.Value // Value is types.UnusedIamUserAccessKeyDetails

	case *types.FindingDetailsMemberUnusedIamUserPasswordDetails:
		_ = v.Value // Value is types.UnusedIamUserPasswordDetails

	case *types.FindingDetailsMemberUnusedPermissionDetails:
		_ = v.Value // Value is types.UnusedPermissionDetails

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ExternalAccessDetails
var _ *types.UnusedIamRoleDetails
var _ *types.UnusedPermissionDetails
var _ *types.UnusedIamUserAccessKeyDetails
var _ *types.UnusedIamUserPasswordDetails

func ExampleNetworkOriginConfiguration_outputUsage() {
	var union types.NetworkOriginConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.NetworkOriginConfigurationMemberInternetConfiguration:
		_ = v.Value // Value is types.InternetConfiguration

	case *types.NetworkOriginConfigurationMemberVpcConfiguration:
		_ = v.Value // Value is types.VpcConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.VpcConfiguration
var _ *types.InternetConfiguration

func ExamplePathElement_outputUsage() {
	var union types.PathElement
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PathElementMemberIndex:
		_ = v.Value // Value is int32

	case *types.PathElementMemberKey:
		_ = v.Value // Value is string

	case *types.PathElementMemberSubstring:
		_ = v.Value // Value is types.Substring

	case *types.PathElementMemberValue:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *int32
var _ *types.Substring

func ExampleRdsDbClusterSnapshotAttributeValue_outputUsage() {
	var union types.RdsDbClusterSnapshotAttributeValue
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RdsDbClusterSnapshotAttributeValueMemberAccountIds:
		_ = v.Value // Value is []string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ []string

func ExampleRdsDbSnapshotAttributeValue_outputUsage() {
	var union types.RdsDbSnapshotAttributeValue
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RdsDbSnapshotAttributeValueMemberAccountIds:
		_ = v.Value // Value is []string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ []string
