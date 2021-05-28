package keyvault

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccessPolicyUpdateKind enumerates the values for access policy update kind.
type AccessPolicyUpdateKind string

const (
	// Add ...
	Add AccessPolicyUpdateKind = "add"
	// Remove ...
	Remove AccessPolicyUpdateKind = "remove"
	// Replace ...
	Replace AccessPolicyUpdateKind = "replace"
)

// PossibleAccessPolicyUpdateKindValues returns an array of possible values for the AccessPolicyUpdateKind const type.
func PossibleAccessPolicyUpdateKindValues() []AccessPolicyUpdateKind {
	return []AccessPolicyUpdateKind{Add, Remove, Replace}
}

// CertificatePermissions enumerates the values for certificate permissions.
type CertificatePermissions string

const (
	// Create ...
	Create CertificatePermissions = "create"
	// Delete ...
	Delete CertificatePermissions = "delete"
	// Deleteissuers ...
	Deleteissuers CertificatePermissions = "deleteissuers"
	// Get ...
	Get CertificatePermissions = "get"
	// Getissuers ...
	Getissuers CertificatePermissions = "getissuers"
	// Import ...
	Import CertificatePermissions = "import"
	// List ...
	List CertificatePermissions = "list"
	// Listissuers ...
	Listissuers CertificatePermissions = "listissuers"
	// Managecontacts ...
	Managecontacts CertificatePermissions = "managecontacts"
	// Manageissuers ...
	Manageissuers CertificatePermissions = "manageissuers"
	// Purge ...
	Purge CertificatePermissions = "purge"
	// Recover ...
	Recover CertificatePermissions = "recover"
	// Setissuers ...
	Setissuers CertificatePermissions = "setissuers"
	// Update ...
	Update CertificatePermissions = "update"
)

// PossibleCertificatePermissionsValues returns an array of possible values for the CertificatePermissions const type.
func PossibleCertificatePermissionsValues() []CertificatePermissions {
	return []CertificatePermissions{Create, Delete, Deleteissuers, Get, Getissuers, Import, List, Listissuers, Managecontacts, Manageissuers, Purge, Recover, Setissuers, Update}
}

// CreateMode enumerates the values for create mode.
type CreateMode string

const (
	// CreateModeDefault ...
	CreateModeDefault CreateMode = "default"
	// CreateModeRecover ...
	CreateModeRecover CreateMode = "recover"
)

// PossibleCreateModeValues returns an array of possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{CreateModeDefault, CreateModeRecover}
}

// KeyPermissions enumerates the values for key permissions.
type KeyPermissions string

const (
	// KeyPermissionsBackup ...
	KeyPermissionsBackup KeyPermissions = "backup"
	// KeyPermissionsCreate ...
	KeyPermissionsCreate KeyPermissions = "create"
	// KeyPermissionsDecrypt ...
	KeyPermissionsDecrypt KeyPermissions = "decrypt"
	// KeyPermissionsDelete ...
	KeyPermissionsDelete KeyPermissions = "delete"
	// KeyPermissionsEncrypt ...
	KeyPermissionsEncrypt KeyPermissions = "encrypt"
	// KeyPermissionsGet ...
	KeyPermissionsGet KeyPermissions = "get"
	// KeyPermissionsImport ...
	KeyPermissionsImport KeyPermissions = "import"
	// KeyPermissionsList ...
	KeyPermissionsList KeyPermissions = "list"
	// KeyPermissionsPurge ...
	KeyPermissionsPurge KeyPermissions = "purge"
	// KeyPermissionsRecover ...
	KeyPermissionsRecover KeyPermissions = "recover"
	// KeyPermissionsRestore ...
	KeyPermissionsRestore KeyPermissions = "restore"
	// KeyPermissionsSign ...
	KeyPermissionsSign KeyPermissions = "sign"
	// KeyPermissionsUnwrapKey ...
	KeyPermissionsUnwrapKey KeyPermissions = "unwrapKey"
	// KeyPermissionsUpdate ...
	KeyPermissionsUpdate KeyPermissions = "update"
	// KeyPermissionsVerify ...
	KeyPermissionsVerify KeyPermissions = "verify"
	// KeyPermissionsWrapKey ...
	KeyPermissionsWrapKey KeyPermissions = "wrapKey"
)

// PossibleKeyPermissionsValues returns an array of possible values for the KeyPermissions const type.
func PossibleKeyPermissionsValues() []KeyPermissions {
	return []KeyPermissions{KeyPermissionsBackup, KeyPermissionsCreate, KeyPermissionsDecrypt, KeyPermissionsDelete, KeyPermissionsEncrypt, KeyPermissionsGet, KeyPermissionsImport, KeyPermissionsList, KeyPermissionsPurge, KeyPermissionsRecover, KeyPermissionsRestore, KeyPermissionsSign, KeyPermissionsUnwrapKey, KeyPermissionsUpdate, KeyPermissionsVerify, KeyPermissionsWrapKey}
}

// Reason enumerates the values for reason.
type Reason string

const (
	// AccountNameInvalid ...
	AccountNameInvalid Reason = "AccountNameInvalid"
	// AlreadyExists ...
	AlreadyExists Reason = "AlreadyExists"
)

// PossibleReasonValues returns an array of possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{AccountNameInvalid, AlreadyExists}
}

// SecretPermissions enumerates the values for secret permissions.
type SecretPermissions string

const (
	// SecretPermissionsBackup ...
	SecretPermissionsBackup SecretPermissions = "backup"
	// SecretPermissionsDelete ...
	SecretPermissionsDelete SecretPermissions = "delete"
	// SecretPermissionsGet ...
	SecretPermissionsGet SecretPermissions = "get"
	// SecretPermissionsList ...
	SecretPermissionsList SecretPermissions = "list"
	// SecretPermissionsPurge ...
	SecretPermissionsPurge SecretPermissions = "purge"
	// SecretPermissionsRecover ...
	SecretPermissionsRecover SecretPermissions = "recover"
	// SecretPermissionsRestore ...
	SecretPermissionsRestore SecretPermissions = "restore"
	// SecretPermissionsSet ...
	SecretPermissionsSet SecretPermissions = "set"
)

// PossibleSecretPermissionsValues returns an array of possible values for the SecretPermissions const type.
func PossibleSecretPermissionsValues() []SecretPermissions {
	return []SecretPermissions{SecretPermissionsBackup, SecretPermissionsDelete, SecretPermissionsGet, SecretPermissionsList, SecretPermissionsPurge, SecretPermissionsRecover, SecretPermissionsRestore, SecretPermissionsSet}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Premium ...
	Premium SkuName = "premium"
	// Standard ...
	Standard SkuName = "standard"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{Premium, Standard}
}

// StoragePermissions enumerates the values for storage permissions.
type StoragePermissions string

const (
	// StoragePermissionsBackup ...
	StoragePermissionsBackup StoragePermissions = "backup"
	// StoragePermissionsDelete ...
	StoragePermissionsDelete StoragePermissions = "delete"
	// StoragePermissionsDeletesas ...
	StoragePermissionsDeletesas StoragePermissions = "deletesas"
	// StoragePermissionsGet ...
	StoragePermissionsGet StoragePermissions = "get"
	// StoragePermissionsGetsas ...
	StoragePermissionsGetsas StoragePermissions = "getsas"
	// StoragePermissionsList ...
	StoragePermissionsList StoragePermissions = "list"
	// StoragePermissionsListsas ...
	StoragePermissionsListsas StoragePermissions = "listsas"
	// StoragePermissionsPurge ...
	StoragePermissionsPurge StoragePermissions = "purge"
	// StoragePermissionsRecover ...
	StoragePermissionsRecover StoragePermissions = "recover"
	// StoragePermissionsRegeneratekey ...
	StoragePermissionsRegeneratekey StoragePermissions = "regeneratekey"
	// StoragePermissionsRestore ...
	StoragePermissionsRestore StoragePermissions = "restore"
	// StoragePermissionsSet ...
	StoragePermissionsSet StoragePermissions = "set"
	// StoragePermissionsSetsas ...
	StoragePermissionsSetsas StoragePermissions = "setsas"
	// StoragePermissionsUpdate ...
	StoragePermissionsUpdate StoragePermissions = "update"
)

// PossibleStoragePermissionsValues returns an array of possible values for the StoragePermissions const type.
func PossibleStoragePermissionsValues() []StoragePermissions {
	return []StoragePermissions{StoragePermissionsBackup, StoragePermissionsDelete, StoragePermissionsDeletesas, StoragePermissionsGet, StoragePermissionsGetsas, StoragePermissionsList, StoragePermissionsListsas, StoragePermissionsPurge, StoragePermissionsRecover, StoragePermissionsRegeneratekey, StoragePermissionsRestore, StoragePermissionsSet, StoragePermissionsSetsas, StoragePermissionsUpdate}
}
