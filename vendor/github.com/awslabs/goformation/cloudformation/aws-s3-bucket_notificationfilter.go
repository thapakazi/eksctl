package cloudformation

import (
	"encoding/json"
)

// AWSS3Bucket_NotificationFilter AWS CloudFormation Resource (AWS::S3::Bucket.NotificationFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter.html
type AWSS3Bucket_NotificationFilter struct {

	// S3Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter.html#cfn-s3-bucket-notificationconfiguraiton-config-filter-s3key
	S3Key *AWSS3Bucket_S3KeyFilter `json:"S3Key,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_NotificationFilter) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.NotificationFilter"
}

func (r *AWSS3Bucket_NotificationFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
