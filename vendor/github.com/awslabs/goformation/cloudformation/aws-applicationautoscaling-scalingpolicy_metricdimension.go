package cloudformation

import (
	"encoding/json"
)

// AWSApplicationAutoScalingScalingPolicy_MetricDimension AWS CloudFormation Resource (AWS::ApplicationAutoScaling::ScalingPolicy.MetricDimension)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-metricdimension.html
type AWSApplicationAutoScalingScalingPolicy_MetricDimension struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-metricdimension.html#cfn-applicationautoscaling-scalingpolicy-metricdimension-name
	Name *Value `json:"Name,omitempty"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-metricdimension.html#cfn-applicationautoscaling-scalingpolicy-metricdimension-value
	Value *Value `json:"Value,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApplicationAutoScalingScalingPolicy_MetricDimension) AWSCloudFormationType() string {
	return "AWS::ApplicationAutoScaling::ScalingPolicy.MetricDimension"
}

func (r *AWSApplicationAutoScalingScalingPolicy_MetricDimension) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
