package resources

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AliasPatternType enumerates the values for alias pattern type.
type AliasPatternType string

const (
	// Extract Extract is the only allowed value.
	Extract AliasPatternType = "Extract"
	// NotSpecified NotSpecified is not allowed.
	NotSpecified AliasPatternType = "NotSpecified"
)

// PossibleAliasPatternTypeValues returns an array of possible values for the AliasPatternType const type.
func PossibleAliasPatternTypeValues() []AliasPatternType {
	return []AliasPatternType{Extract, NotSpecified}
}

// AliasType enumerates the values for alias type.
type AliasType string

const (
	// AliasTypeMask Alias value is secret.
	AliasTypeMask AliasType = "Mask"
	// AliasTypeNotSpecified Alias type is unknown (same as not providing alias type).
	AliasTypeNotSpecified AliasType = "NotSpecified"
	// AliasTypePlainText Alias value is not secret.
	AliasTypePlainText AliasType = "PlainText"
)

// PossibleAliasTypeValues returns an array of possible values for the AliasType const type.
func PossibleAliasTypeValues() []AliasType {
	return []AliasType{AliasTypeMask, AliasTypeNotSpecified, AliasTypePlainText}
}

// ChangeType enumerates the values for change type.
type ChangeType string

const (
	// Create The resource does not exist in the current state but is present in the desired state. The
	// resource will be created when the deployment is executed.
	Create ChangeType = "Create"
	// Delete The resource exists in the current state and is missing from the desired state. The resource will
	// be deleted when the deployment is executed.
	Delete ChangeType = "Delete"
	// Deploy The resource exists in the current state and the desired state and will be redeployed when the
	// deployment is executed. The properties of the resource may or may not change.
	Deploy ChangeType = "Deploy"
	// Ignore The resource exists in the current state and is missing from the desired state. The resource will
	// not be deployed or modified when the deployment is executed.
	Ignore ChangeType = "Ignore"
	// Modify The resource exists in the current state and the desired state and will be redeployed when the
	// deployment is executed. The properties of the resource will change.
	Modify ChangeType = "Modify"
	// NoChange The resource exists in the current state and the desired state and will be redeployed when the
	// deployment is executed. The properties of the resource will not change.
	NoChange ChangeType = "NoChange"
)

// PossibleChangeTypeValues returns an array of possible values for the ChangeType const type.
func PossibleChangeTypeValues() []ChangeType {
	return []ChangeType{Create, Delete, Deploy, Ignore, Modify, NoChange}
}

// DeploymentMode enumerates the values for deployment mode.
type DeploymentMode string

const (
	// Complete ...
	Complete DeploymentMode = "Complete"
	// Incremental ...
	Incremental DeploymentMode = "Incremental"
)

// PossibleDeploymentModeValues returns an array of possible values for the DeploymentMode const type.
func PossibleDeploymentModeValues() []DeploymentMode {
	return []DeploymentMode{Complete, Incremental}
}

// OnErrorDeploymentType enumerates the values for on error deployment type.
type OnErrorDeploymentType string

const (
	// LastSuccessful ...
	LastSuccessful OnErrorDeploymentType = "LastSuccessful"
	// SpecificDeployment ...
	SpecificDeployment OnErrorDeploymentType = "SpecificDeployment"
)

// PossibleOnErrorDeploymentTypeValues returns an array of possible values for the OnErrorDeploymentType const type.
func PossibleOnErrorDeploymentTypeValues() []OnErrorDeploymentType {
	return []OnErrorDeploymentType{LastSuccessful, SpecificDeployment}
}

// PropertyChangeType enumerates the values for property change type.
type PropertyChangeType string

const (
	// PropertyChangeTypeArray The property is an array and contains nested changes.
	PropertyChangeTypeArray PropertyChangeType = "Array"
	// PropertyChangeTypeCreate The property does not exist in the current state but is present in the desired
	// state. The property will be created when the deployment is executed.
	PropertyChangeTypeCreate PropertyChangeType = "Create"
	// PropertyChangeTypeDelete The property exists in the current state and is missing from the desired state.
	// It will be deleted when the deployment is executed.
	PropertyChangeTypeDelete PropertyChangeType = "Delete"
	// PropertyChangeTypeModify The property exists in both current and desired state and is different. The
	// value of the property will change when the deployment is executed.
	PropertyChangeTypeModify PropertyChangeType = "Modify"
)

// PossiblePropertyChangeTypeValues returns an array of possible values for the PropertyChangeType const type.
func PossiblePropertyChangeTypeValues() []PropertyChangeType {
	return []PropertyChangeType{PropertyChangeTypeArray, PropertyChangeTypeCreate, PropertyChangeTypeDelete, PropertyChangeTypeModify}
}

// ProvisioningOperation enumerates the values for provisioning operation.
type ProvisioningOperation string

const (
	// ProvisioningOperationAction The provisioning operation is action.
	ProvisioningOperationAction ProvisioningOperation = "Action"
	// ProvisioningOperationAzureAsyncOperationWaiting The provisioning operation is waiting Azure async
	// operation.
	ProvisioningOperationAzureAsyncOperationWaiting ProvisioningOperation = "AzureAsyncOperationWaiting"
	// ProvisioningOperationCreate The provisioning operation is create.
	ProvisioningOperationCreate ProvisioningOperation = "Create"
	// ProvisioningOperationDelete The provisioning operation is delete.
	ProvisioningOperationDelete ProvisioningOperation = "Delete"
	// ProvisioningOperationDeploymentCleanup The provisioning operation is cleanup. This operation is part of
	// the 'complete' mode deployment.
	ProvisioningOperationDeploymentCleanup ProvisioningOperation = "DeploymentCleanup"
	// ProvisioningOperationEvaluateDeploymentOutput The provisioning operation is evaluate output.
	ProvisioningOperationEvaluateDeploymentOutput ProvisioningOperation = "EvaluateDeploymentOutput"
	// ProvisioningOperationNotSpecified The provisioning operation is not specified.
	ProvisioningOperationNotSpecified ProvisioningOperation = "NotSpecified"
	// ProvisioningOperationRead The provisioning operation is read.
	ProvisioningOperationRead ProvisioningOperation = "Read"
	// ProvisioningOperationResourceCacheWaiting The provisioning operation is waiting for resource cache.
	ProvisioningOperationResourceCacheWaiting ProvisioningOperation = "ResourceCacheWaiting"
	// ProvisioningOperationWaiting The provisioning operation is waiting.
	ProvisioningOperationWaiting ProvisioningOperation = "Waiting"
)

// PossibleProvisioningOperationValues returns an array of possible values for the ProvisioningOperation const type.
func PossibleProvisioningOperationValues() []ProvisioningOperation {
	return []ProvisioningOperation{ProvisioningOperationAction, ProvisioningOperationAzureAsyncOperationWaiting, ProvisioningOperationCreate, ProvisioningOperationDelete, ProvisioningOperationDeploymentCleanup, ProvisioningOperationEvaluateDeploymentOutput, ProvisioningOperationNotSpecified, ProvisioningOperationRead, ProvisioningOperationResourceCacheWaiting, ProvisioningOperationWaiting}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// None ...
	None ResourceIdentityType = "None"
	// SystemAssigned ...
	SystemAssigned ResourceIdentityType = "SystemAssigned"
	// SystemAssignedUserAssigned ...
	SystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	// UserAssigned ...
	UserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{None, SystemAssigned, SystemAssignedUserAssigned, UserAssigned}
}

// TagsPatchOperation enumerates the values for tags patch operation.
type TagsPatchOperation string

const (
	// TagsPatchOperationDelete The 'delete' option allows selectively deleting tags based on given names or
	// name/value pairs.
	TagsPatchOperationDelete TagsPatchOperation = "Delete"
	// TagsPatchOperationMerge The 'merge' option allows adding tags with new names and updating the values of
	// tags with existing names.
	TagsPatchOperationMerge TagsPatchOperation = "Merge"
	// TagsPatchOperationReplace The 'replace' option replaces the entire set of existing tags with a new set.
	TagsPatchOperationReplace TagsPatchOperation = "Replace"
)

// PossibleTagsPatchOperationValues returns an array of possible values for the TagsPatchOperation const type.
func PossibleTagsPatchOperationValues() []TagsPatchOperation {
	return []TagsPatchOperation{TagsPatchOperationDelete, TagsPatchOperationMerge, TagsPatchOperationReplace}
}

// WhatIfResultFormat enumerates the values for what if result format.
type WhatIfResultFormat string

const (
	// FullResourcePayloads ...
	FullResourcePayloads WhatIfResultFormat = "FullResourcePayloads"
	// ResourceIDOnly ...
	ResourceIDOnly WhatIfResultFormat = "ResourceIdOnly"
)

// PossibleWhatIfResultFormatValues returns an array of possible values for the WhatIfResultFormat const type.
func PossibleWhatIfResultFormatValues() []WhatIfResultFormat {
	return []WhatIfResultFormat{FullResourcePayloads, ResourceIDOnly}
}
