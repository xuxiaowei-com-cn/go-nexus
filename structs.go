package nexus

type AbstractApiRepository struct {
	Type   string `json:"type,omitempty" url:"type,omitempty"`     //	description: Controls if deployments of and updates to artifacts are allowed, example: hosted, pattern:
	Url    string `json:"url,omitempty" url:"url,omitempty"`       //	description: URL to the repository, example: , pattern:
	Format string `json:"format,omitempty" url:"format,omitempty"` //	description: Component format held in this repository, example: npm, pattern:
	Name   string `json:"name,omitempty" url:"name,omitempty"`     //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online bool   `json:"online,omitempty" url:"online,omitempty"` //	description: Whether this repository accepts incoming requests, example: true, pattern:
}

type AnonymousAccessSettingsXO struct {
	Enabled   bool   `json:"enabled,omitempty" url:"enabled,omitempty"`     //	description: Whether or not Anonymous Access is enabled, example: , pattern:
	RealmName string `json:"realmName,omitempty" url:"realmName,omitempty"` //	description: The name of the authentication realm for the anonymous account, example: , pattern:
	UserId    string `json:"userId,omitempty" url:"userId,omitempty"`       //	description: The username of the anonymous account, example: , pattern:
}

type ApiCertificate struct {
	ExpiresOn                 int64  `json:"expiresOn,omitempty" url:"expiresOn,omitempty"`                                 //	description: , example: , pattern:
	Fingerprint               string `json:"fingerprint,omitempty" url:"fingerprint,omitempty"`                             //	description: , example: , pattern:
	SerialNumber              string `json:"serialNumber,omitempty" url:"serialNumber,omitempty"`                           //	description: , example: , pattern:
	SubjectCommonName         string `json:"subjectCommonName,omitempty" url:"subjectCommonName,omitempty"`                 //	description: , example: , pattern:
	SubjectOrganization       string `json:"subjectOrganization,omitempty" url:"subjectOrganization,omitempty"`             //	description: , example: , pattern:
	SubjectOrganizationalUnit string `json:"subjectOrganizationalUnit,omitempty" url:"subjectOrganizationalUnit,omitempty"` //	description: , example: , pattern:
	Id                        string `json:"id,omitempty" url:"id,omitempty"`                                               //	description: , example: , pattern:
	IssuedOn                  int64  `json:"issuedOn,omitempty" url:"issuedOn,omitempty"`                                   //	description: , example: , pattern:
	IssuerCommonName          string `json:"issuerCommonName,omitempty" url:"issuerCommonName,omitempty"`                   //	description: , example: , pattern:
	IssuerOrganization        string `json:"issuerOrganization,omitempty" url:"issuerOrganization,omitempty"`               //	description: , example: , pattern:
	IssuerOrganizationalUnit  string `json:"issuerOrganizationalUnit,omitempty" url:"issuerOrganizationalUnit,omitempty"`   //	description: , example: , pattern:
	Pem                       string `json:"pem,omitempty" url:"pem,omitempty"`                                             //	description: , example: , pattern:
}

type ApiCreateUser struct {
	Roles        []string `json:"roles,omitempty" url:"roles,omitempty"`               //	description: The roles which the user has been assigned within Nexus., example: , pattern:
	Status       string   `json:"status,omitempty" url:"status,omitempty"`             //	description: The user's status, e.g. active or disabled., example: , pattern:
	UserId       string   `json:"userId,omitempty" url:"userId,omitempty"`             //	description: The userid which is required for login. This value cannot be changed., example: , pattern:
	EmailAddress string   `json:"emailAddress,omitempty" url:"emailAddress,omitempty"` //	description: The email address associated with the user., example: , pattern:
	FirstName    string   `json:"firstName,omitempty" url:"firstName,omitempty"`       //	description: The first name of the user., example: , pattern:
	LastName     string   `json:"lastName,omitempty" url:"lastName,omitempty"`         //	description: The last name of the user., example: , pattern:
	Password     string   `json:"password,omitempty" url:"password,omitempty"`         //	description: The password for the new user., example: , pattern:
}

type ApiEmailConfiguration struct {
	StartTlsRequired              bool   `json:"startTlsRequired,omitempty" url:"startTlsRequired,omitempty"`                           //	description: Require STARTTLS Support, example: , pattern:
	Username                      string `json:"username,omitempty" url:"username,omitempty"`                                           //	description: , example: , pattern:
	Password                      string `json:"password,omitempty" url:"password,omitempty"`                                           //	description: , example: , pattern:
	Port                          int64  `json:"port,omitempty" url:"port,omitempty"`                                                   //	description: , example: , pattern:
	StartTlsEnabled               bool   `json:"startTlsEnabled,omitempty" url:"startTlsEnabled,omitempty"`                             //	description: Enable STARTTLS Support for Insecure Connections, example: , pattern:
	NexusTrustStoreEnabled        bool   `json:"nexusTrustStoreEnabled,omitempty" url:"nexusTrustStoreEnabled,omitempty"`               //	description: Use the Nexus Repository Manager's certificate truststore, example: , pattern:
	SslOnConnectEnabled           bool   `json:"sslOnConnectEnabled,omitempty" url:"sslOnConnectEnabled,omitempty"`                     //	description: Enable SSL/TLS Encryption upon Connection, example: , pattern:
	SslServerIdentityCheckEnabled bool   `json:"sslServerIdentityCheckEnabled,omitempty" url:"sslServerIdentityCheckEnabled,omitempty"` //	description: Verify the server certificate when using TLS or SSL, example: , pattern:

	SubjectPrefix string `json:"subjectPrefix,omitempty" url:"subjectPrefix,omitempty"` //	description: A prefix to add to all email subjects to aid in identifying automated emails, example: , pattern:
	Enabled       bool   `json:"enabled,omitempty" url:"enabled,omitempty"`             //	description: , example: , pattern:
	FromAddress   string `json:"fromAddress,omitempty" url:"fromAddress,omitempty"`     //	description: , example: nexus@example.org, pattern:
	Host          string `json:"host,omitempty" url:"host,omitempty"`                   //	description: , example: , pattern:
}

type ApiEmailValidation struct {
	Reason  string `json:"reason,omitempty" url:"reason,omitempty"`   //	description: , example: , pattern:
	Success bool   `json:"success,omitempty" url:"success,omitempty"` //	description: , example: , pattern:
}

type ApiLicenseDetailsXO struct {
	LicenseType    string `json:"licenseType,omitempty" url:"licenseType,omitempty"`       //	description: , example: , pattern:
	LicensedUsers  string `json:"licensedUsers,omitempty" url:"licensedUsers,omitempty"`   //	description: , example: , pattern:
	ContactCompany string `json:"contactCompany,omitempty" url:"contactCompany,omitempty"` //	description: , example: , pattern:
	ContactName    string `json:"contactName,omitempty" url:"contactName,omitempty"`       //	description: , example: , pattern:
	EffectiveDate  string `json:"effectiveDate,omitempty" url:"effectiveDate,omitempty"`   //	description: , example: , pattern:
	ExpirationDate string `json:"expirationDate,omitempty" url:"expirationDate,omitempty"` //	description: , example: , pattern:
	ContactEmail   string `json:"contactEmail,omitempty" url:"contactEmail,omitempty"`     //	description: , example: , pattern:
	Features       string `json:"features,omitempty" url:"features,omitempty"`             //	description: , example: , pattern:
	Fingerprint    string `json:"fingerprint,omitempty" url:"fingerprint,omitempty"`       //	description: , example: , pattern:
}

type ApiPrivilege struct {
	Name        string `json:"name,omitempty" url:"name,omitempty"`               //	description: The name of the privilege.  This value cannot be changed., example: , pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	ReadOnly    bool   `json:"readOnly,omitempty" url:"readOnly,omitempty"`       //	description: Indicates whether the privilege can be changed. External values supplied to this will be ignored by the system., example: , pattern:
	Type        string `json:"type,omitempty" url:"type,omitempty"`               //	description: The type of privilege, each type covers different portions of the system. External values supplied to this will be ignored by the system., example: , pattern:
	Description string `json:"description,omitempty" url:"description,omitempty"` //	description: , example: , pattern:
}

type ApiPrivilegeApplicationRequest struct {
	Actions     []string `json:"actions,omitempty" url:"actions,omitempty"`         //	description: A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as 'run' for script privileges., example: , pattern:
	Description string   `json:"description,omitempty" url:"description,omitempty"` //	description: , example: , pattern:
	Domain      string   `json:"domain,omitempty" url:"domain,omitempty"`           //	description: The domain (i.e. 'blobstores', 'capabilities' or even '*' for all) that this privilege is granting access to.  Note that creating new privileges with a domain is only necessary when using plugins that define their own domain(s)., example: , pattern:
	Name        string   `json:"name,omitempty" url:"name,omitempty"`               //	description: The name of the privilege.  This value cannot be changed., example: , pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
}

type ApiPrivilegeRepositoryAdminRequest struct {
	Actions     []string `json:"actions,omitempty" url:"actions,omitempty"`         //	description: A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as 'run' for script privileges., example: , pattern:
	Description string   `json:"description,omitempty" url:"description,omitempty"` //	description: , example: , pattern:
	Format      string   `json:"format,omitempty" url:"format,omitempty"`           //	description: The repository format (i.e 'nuget', 'npm') this privilege will grant access to (or * for all)., example: , pattern:
	Name        string   `json:"name,omitempty" url:"name,omitempty"`               //	description: The name of the privilege.  This value cannot be changed., example: , pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Repository  string   `json:"repository,omitempty" url:"repository,omitempty"`   //	description: The name of the repository this privilege will grant access to (or * for all)., example: , pattern:
}

type ApiPrivilegeRepositoryContentSelectorRequest struct {
	Format          string   `json:"format,omitempty" url:"format,omitempty"`                   //	description: The repository format (i.e 'nuget', 'npm') this privilege will grant access to (or * for all)., example: , pattern:
	Name            string   `json:"name,omitempty" url:"name,omitempty"`                       //	description: The name of the privilege.  This value cannot be changed., example: , pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Repository      string   `json:"repository,omitempty" url:"repository,omitempty"`           //	description: The name of the repository this privilege will grant access to (or * for all)., example: , pattern:
	Actions         []string `json:"actions,omitempty" url:"actions,omitempty"`                 //	description: A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as 'run' for script privileges., example: , pattern:
	ContentSelector string   `json:"contentSelector,omitempty" url:"contentSelector,omitempty"` //	description: The name of a content selector that will be used to grant access to content via this privilege., example: , pattern:
	Description     string   `json:"description,omitempty" url:"description,omitempty"`         //	description: , example: , pattern:
}

type ApiPrivilegeRepositoryViewRequest struct {
	Actions     []string `json:"actions,omitempty" url:"actions,omitempty"`         //	description: A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as 'run' for script privileges., example: , pattern:
	Description string   `json:"description,omitempty" url:"description,omitempty"` //	description: , example: , pattern:
	Format      string   `json:"format,omitempty" url:"format,omitempty"`           //	description: The repository format (i.e 'nuget', 'npm') this privilege will grant access to (or * for all)., example: , pattern:
	Name        string   `json:"name,omitempty" url:"name,omitempty"`               //	description: The name of the privilege.  This value cannot be changed., example: , pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Repository  string   `json:"repository,omitempty" url:"repository,omitempty"`   //	description: The name of the repository this privilege will grant access to (or * for all)., example: , pattern:
}

type ApiPrivilegeScriptRequest struct {
	Actions     []string `json:"actions,omitempty" url:"actions,omitempty"`         //	description: A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as 'run' for script privileges., example: , pattern:
	Description string   `json:"description,omitempty" url:"description,omitempty"` //	description: , example: , pattern:
	Name        string   `json:"name,omitempty" url:"name,omitempty"`               //	description: The name of the privilege.  This value cannot be changed., example: , pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	ScriptName  string   `json:"scriptName,omitempty" url:"scriptName,omitempty"`   //	description: The name of a script to give access to., example: , pattern:
}

type ApiPrivilegeWildcardRequest struct {
	Description string `json:"description,omitempty" url:"description,omitempty"` //	description: , example: , pattern:
	Name        string `json:"name,omitempty" url:"name,omitempty"`               //	description: The name of the privilege.  This value cannot be changed., example: , pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Pattern     string `json:"pattern,omitempty" url:"pattern,omitempty"`         //	description: A colon separated list of parts that create a permission string., example: , pattern:
}

type ApiUser struct {
	ReadOnly      bool     `json:"readOnly,omitempty" url:"readOnly,omitempty"`           //	description: Indicates whether the user's properties could be modified by the Nexus Repository Manager. When false only roles are considered during update., example: , pattern:
	Roles         []string `json:"roles,omitempty" url:"roles,omitempty"`                 //	description: The roles which the user has been assigned within Nexus., example: , pattern:
	Source        string   `json:"source,omitempty" url:"source,omitempty"`               //	description: The user source which is the origin of this user. This value cannot be changed., example: , pattern:
	UserId        string   `json:"userId,omitempty" url:"userId,omitempty"`               //	description: The userid which is required for login. This value cannot be changed., example: , pattern:
	FirstName     string   `json:"firstName,omitempty" url:"firstName,omitempty"`         //	description: The first name of the user., example: , pattern:
	ExternalRoles []string `json:"externalRoles,omitempty" url:"externalRoles,omitempty"` //	description: The roles which the user has been assigned in an external source, e.g. LDAP group. These cannot be changed within the Nexus Repository Manager., example: , pattern:
	LastName      string   `json:"lastName,omitempty" url:"lastName,omitempty"`           //	description: The last name of the user., example: , pattern:
	Status        string   `json:"status,omitempty" url:"status,omitempty"`               //	description: The user's status, e.g. active or disabled., example: , pattern:
	EmailAddress  string   `json:"emailAddress,omitempty" url:"emailAddress,omitempty"`   //	description: The email address associated with the user., example: , pattern:
}

type ApiUserSource struct {
	Id   string `json:"id,omitempty" url:"id,omitempty"`     //	description: , example: , pattern:
	Name string `json:"name,omitempty" url:"name,omitempty"` //	description: , example: , pattern:
}

type AptHostedApiRepository struct {
	Apt        AptHostedRepositoriesAttributes  `json:"apt,omitempty" url:"apt,omitempty"`               //	description: , example: , pattern:
	AptSigning AptSigningRepositoriesAttributes `json:"aptSigning,omitempty" url:"aptSigning,omitempty"` //	description: , example: , pattern:
	Cleanup    CleanupPolicyAttributes          `json:"cleanup,omitempty" url:"cleanup,omitempty"`       //	description: , example: , pattern:
	Component  ComponentAttributes              `json:"component,omitempty" url:"component,omitempty"`   //	description: , example: , pattern:
	Name       string                           `json:"name,omitempty" url:"name,omitempty"`             //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online     bool                             `json:"online,omitempty" url:"online,omitempty"`         //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage    HostedStorageAttributes          `json:"storage,omitempty" url:"storage,omitempty"`       //	description: , example: , pattern:
}

type AptHostedRepositoriesAttributes struct {
	Distribution string `json:"distribution,omitempty" url:"distribution,omitempty"` //	description: Distribution to fetch, example: bionic, pattern:
}

type AptHostedRepositoryApiRequest struct {
	Apt        AptHostedRepositoriesAttributes  `json:"apt,omitempty" url:"apt,omitempty"`               //	description: , example: , pattern:
	AptSigning AptSigningRepositoriesAttributes `json:"aptSigning,omitempty" url:"aptSigning,omitempty"` //	description: , example: , pattern:
	Cleanup    CleanupPolicyAttributes          `json:"cleanup,omitempty" url:"cleanup,omitempty"`       //	description: , example: , pattern:
	Component  ComponentAttributes              `json:"component,omitempty" url:"component,omitempty"`   //	description: , example: , pattern:
	Name       string                           `json:"name,omitempty" url:"name,omitempty"`             //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online     bool                             `json:"online,omitempty" url:"online,omitempty"`         //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage    HostedStorageAttributes          `json:"storage,omitempty" url:"storage,omitempty"`       //	description: , example: , pattern:
}

type AptProxyApiRepository struct {
	NegativeCache   NegativeCacheAttributes        `json:"negativeCache,omitempty" url:"negativeCache,omitempty"`     //	description: , example: , pattern:
	Proxy           ProxyAttributes                `json:"proxy,omitempty" url:"proxy,omitempty"`                     //	description: , example: , pattern:
	Storage         StorageAttributes              `json:"storage,omitempty" url:"storage,omitempty"`                 //	description: , example: , pattern:
	Name            string                         `json:"name,omitempty" url:"name,omitempty"`                       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online          bool                           `json:"online,omitempty" url:"online,omitempty"`                   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Replication     ReplicationAttributes          `json:"replication,omitempty" url:"replication,omitempty"`         //	description: , example: , pattern:
	RoutingRuleName string                         `json:"routingRuleName,omitempty" url:"routingRuleName,omitempty"` //	description: The name of the routing rule assigned to this repository, example: , pattern:
	Apt             AptProxyRepositoriesAttributes `json:"apt,omitempty" url:"apt,omitempty"`                         //	description: , example: , pattern:
	Cleanup         CleanupPolicyAttributes        `json:"cleanup,omitempty" url:"cleanup,omitempty"`                 //	description: , example: , pattern:
	HttpClient      HttpClientAttributes           `json:"httpClient,omitempty" url:"httpClient,omitempty"`           //	description: , example: , pattern:
}

type AptProxyRepositoriesAttributes struct {
	Distribution string `json:"distribution,omitempty" url:"distribution,omitempty"` //	description: Distribution to fetch, example: bionic, pattern:
	Flat         bool   `json:"flat,omitempty" url:"flat,omitempty"`                 //	description: Whether this repository is flat, example: false, pattern:
}

type AptProxyRepositoryApiRequest struct {
	NegativeCache NegativeCacheAttributes        `json:"negativeCache,omitempty" url:"negativeCache,omitempty"` //	description: , example: , pattern:
	Online        bool                           `json:"online,omitempty" url:"online,omitempty"`               //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Proxy         ProxyAttributes                `json:"proxy,omitempty" url:"proxy,omitempty"`                 //	description: , example: , pattern:
	Replication   ReplicationAttributes          `json:"replication,omitempty" url:"replication,omitempty"`     //	description: , example: , pattern:
	Storage       StorageAttributes              `json:"storage,omitempty" url:"storage,omitempty"`             //	description: , example: , pattern:
	Apt           AptProxyRepositoriesAttributes `json:"apt,omitempty" url:"apt,omitempty"`                     //	description: , example: , pattern:
	Cleanup       CleanupPolicyAttributes        `json:"cleanup,omitempty" url:"cleanup,omitempty"`             //	description: , example: , pattern:
	Name          string                         `json:"name,omitempty" url:"name,omitempty"`                   //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	HttpClient    HttpClientAttributes           `json:"httpClient,omitempty" url:"httpClient,omitempty"`       //	description: , example: , pattern:
	RoutingRule   string                         `json:"routingRule,omitempty" url:"routingRule,omitempty"`     //	description: , example: , pattern:
}

type AptSigningRepositoriesAttributes struct {
	Keypair    string `json:"keypair,omitempty" url:"keypair,omitempty"`       //	description: PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor), example: , pattern:
	Passphrase string `json:"passphrase,omitempty" url:"passphrase,omitempty"` //	description: Passphrase to access PGP signing key, example: , pattern:
}

type AssetXO struct {
	DownloadUrl    string      `json:"downloadUrl,omitempty" url:"downloadUrl,omitempty"`       //	description: , example: , pattern:
	LastDownloaded string      `json:"lastDownloaded,omitempty" url:"lastDownloaded,omitempty"` //	description: , example: , pattern:
	LastModified   string      `json:"lastModified,omitempty" url:"lastModified,omitempty"`     //	description: , example: , pattern:
	Path           string      `json:"path,omitempty" url:"path,omitempty"`                     //	description: , example: , pattern:
	BlobCreated    string      `json:"blobCreated,omitempty" url:"blobCreated,omitempty"`       //	description: , example: , pattern:
	ContentType    string      `json:"contentType,omitempty" url:"contentType,omitempty"`       //	description: , example: , pattern:
	FileSize       int64       `json:"fileSize,omitempty" url:"fileSize,omitempty"`             //	description: , example: , pattern:
	Id             string      `json:"id,omitempty" url:"id,omitempty"`                         //	description: , example: , pattern:
	UploaderIp     string      `json:"uploaderIp,omitempty" url:"uploaderIp,omitempty"`         //	description: , example: , pattern:
	Repository     string      `json:"repository,omitempty" url:"repository,omitempty"`         //	description: , example: , pattern:
	Uploader       string      `json:"uploader,omitempty" url:"uploader,omitempty"`             //	description: , example: , pattern:
	Checksum       interface{} `json:"checksum,omitempty" url:"checksum,omitempty"`             //	description: , example: , pattern:
	Format         string      `json:"format,omitempty" url:"format,omitempty"`                 //	description: , example: , pattern:
}

type AzureBlobStoreApiAuthentication struct {
	AccountKey           string `json:"accountKey,omitempty" url:"accountKey,omitempty"`                     //	description: The account key., example: , pattern:
	AuthenticationMethod string `json:"authenticationMethod,omitempty" url:"authenticationMethod,omitempty"` //	description: The type of Azure authentication to use., example: , pattern:
}

type AzureBlobStoreApiBucketConfiguration struct {
	AccountName    string                          `json:"accountName,omitempty" url:"accountName,omitempty"`       //	description: Account name found under Access keys for the storage account., example: , pattern:
	Authentication AzureBlobStoreApiAuthentication `json:"authentication,omitempty" url:"authentication,omitempty"` //	description: The Azure specific authentication details., example: , pattern:
	ContainerName  string                          `json:"containerName,omitempty" url:"containerName,omitempty"`   //	description: The name of an existing container to be used for storage., example: , pattern: ^[a-z0-9][a-z0-9-]{2,62}$
}

type AzureBlobStoreApiModel struct {
	BucketConfiguration AzureBlobStoreApiBucketConfiguration `json:"bucketConfiguration,omitempty" url:"bucketConfiguration,omitempty"` //	description: The Azure specific configuration details for the Azure object that'll contain the blob store., example: , pattern:
	Name                string                               `json:"name,omitempty" url:"name,omitempty"`                               //	description: The name of the Azure blob store., example: , pattern:
	SoftQuota           BlobStoreApiSoftQuota                `json:"softQuota,omitempty" url:"softQuota,omitempty"`                     //	description: Settings to control the soft quota., example: , pattern:
}

type AzureConnectionXO struct {
	AccountKey           string `json:"accountKey,omitempty" url:"accountKey,omitempty"`                     //	description: , example: , pattern:
	AccountName          string `json:"accountName,omitempty" url:"accountName,omitempty"`                   //	description: , example: , pattern:
	AuthenticationMethod string `json:"authenticationMethod,omitempty" url:"authenticationMethod,omitempty"` //	description: , example: , pattern:
	ContainerName        string `json:"containerName,omitempty" url:"containerName,omitempty"`               //	description: , example: , pattern:
}

type BlobStoreApiSoftQuota struct {
	Limit int64  `json:"limit,omitempty" url:"limit,omitempty"` //	description: The limit in MB., example: , pattern:
	Type  string `json:"type,omitempty" url:"type,omitempty"`   //	description: The type to use such as spaceRemainingQuota, or spaceUsedQuota, example: , pattern:
}

type BlobStoreQuotaResultXO struct {
	BlobStoreName string `json:"blobStoreName,omitempty" url:"blobStoreName,omitempty"` //	description: , example: , pattern:
	IsViolation   bool   `json:"isViolation,omitempty" url:"isViolation,omitempty"`     //	description: , example: , pattern:
	Message       string `json:"message,omitempty" url:"message,omitempty"`             //	description: , example: , pattern:
}

type BowerAttributes struct {
	RewritePackageUrls bool `json:"rewritePackageUrls,omitempty" url:"rewritePackageUrls,omitempty"` //	description: Whether to force Bower to retrieve packages through this proxy repository, example: true, pattern:
}

type BowerGroupRepositoryApiRequest struct {
	Group   GroupAttributes   `json:"group,omitempty" url:"group,omitempty"`     //	description: , example: , pattern:
	Name    string            `json:"name,omitempty" url:"name,omitempty"`       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online  bool              `json:"online,omitempty" url:"online,omitempty"`   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage StorageAttributes `json:"storage,omitempty" url:"storage,omitempty"` //	description: , example: , pattern:
}

type BowerHostedRepositoryApiRequest struct {
	Cleanup   CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`     //	description: , example: , pattern:
	Component ComponentAttributes     `json:"component,omitempty" url:"component,omitempty"` //	description: , example: , pattern:
	Name      string                  `json:"name,omitempty" url:"name,omitempty"`           //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online    bool                    `json:"online,omitempty" url:"online,omitempty"`       //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage   HostedStorageAttributes `json:"storage,omitempty" url:"storage,omitempty"`     //	description: , example: , pattern:
}

type BowerProxyApiRepository struct {
	Replication     ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`         //	description: , example: , pattern:
	Bower           BowerAttributes         `json:"bower,omitempty" url:"bower,omitempty"`                     //	description: , example: , pattern:
	HttpClient      HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`           //	description: , example: , pattern:
	Online          bool                    `json:"online,omitempty" url:"online,omitempty"`                   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Proxy           ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                     //	description: , example: , pattern:
	RoutingRuleName string                  `json:"routingRuleName,omitempty" url:"routingRuleName,omitempty"` //	description: The name of the routing rule assigned to this repository, example: , pattern:
	Storage         StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`                 //	description: , example: , pattern:
	Cleanup         CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`                 //	description: , example: , pattern:
	Name            string                  `json:"name,omitempty" url:"name,omitempty"`                       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	NegativeCache   NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"`     //	description: , example: , pattern:

}

type BowerProxyRepositoryApiRequest struct {
	NegativeCache NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"` //	description: , example: , pattern:
	Proxy         ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                 //	description: , example: , pattern:
	RoutingRule   string                  `json:"routingRule,omitempty" url:"routingRule,omitempty"`     //	description: , example: , pattern:
	Storage       StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`             //	description: , example: , pattern:
	HttpClient    HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`       //	description: , example: , pattern:
	Name          string                  `json:"name,omitempty" url:"name,omitempty"`                   //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online        bool                    `json:"online,omitempty" url:"online,omitempty"`               //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Replication   ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`     //	description: , example: , pattern:
	Bower         BowerAttributes         `json:"bower,omitempty" url:"bower,omitempty"`                 //	description: , example: , pattern:
	Cleanup       CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`             //	description: , example: , pattern:
}

type CleanupPolicyAttributes struct {
	PolicyNames []string `json:"policyNames,omitempty" url:"policyNames,omitempty"` //	description: Components that match any of the applied policies will be deleted, example: , pattern:
}

type CocoapodsProxyRepositoryApiRequest struct {
	Storage       StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`             //	description: , example: , pattern:
	Cleanup       CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`             //	description: , example: , pattern:
	HttpClient    HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`       //	description: , example: , pattern:
	Proxy         ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                 //	description: , example: , pattern:
	RoutingRule   string                  `json:"routingRule,omitempty" url:"routingRule,omitempty"`     //	description: , example: , pattern:
	Name          string                  `json:"name,omitempty" url:"name,omitempty"`                   //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	NegativeCache NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"` //	description: , example: , pattern:
	Online        bool                    `json:"online,omitempty" url:"online,omitempty"`               //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Replication   ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`     //	description: , example: , pattern:
}

type ComponentAttributes struct {
	ProprietaryComponents bool `json:"proprietaryComponents,omitempty" url:"proprietaryComponents,omitempty"` //	description: Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall), example: , pattern:
}

type ComponentXO struct {
	Repository string    `json:"repository,omitempty" url:"repository,omitempty"` //	description: , example: , pattern:
	Version    string    `json:"version,omitempty" url:"version,omitempty"`       //	description: , example: , pattern:
	Assets     []AssetXO `json:"assets,omitempty" url:"assets,omitempty"`         //	description: , example: , pattern:
	Format     string    `json:"format,omitempty" url:"format,omitempty"`         //	description: , example: , pattern:
	Group      string    `json:"group,omitempty" url:"group,omitempty"`           //	description: , example: , pattern:
	Id         string    `json:"id,omitempty" url:"id,omitempty"`                 //	description: , example: , pattern:
	Name       string    `json:"name,omitempty" url:"name,omitempty"`             //	description: , example: , pattern:
}

type ConanProxyRepositoryApiRequest struct {
	RoutingRule   string                  `json:"routingRule,omitempty" url:"routingRule,omitempty"`     //	description: , example: , pattern:
	Cleanup       CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`             //	description: , example: , pattern:
	Name          string                  `json:"name,omitempty" url:"name,omitempty"`                   //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Proxy         ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                 //	description: , example: , pattern:
	Replication   ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`     //	description: , example: , pattern:
	HttpClient    HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`       //	description: , example: , pattern:
	NegativeCache NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"` //	description: , example: , pattern:
	Online        bool                    `json:"online,omitempty" url:"online,omitempty"`               //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage       StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`             //	description: , example: , pattern:
}

type CondaProxyRepositoryApiRequest struct {
	Name          string                  `json:"name,omitempty" url:"name,omitempty"`                   //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Proxy         ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                 //	description: , example: , pattern:
	Storage       StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`             //	description: , example: , pattern:
	Replication   ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`     //	description: , example: , pattern:
	RoutingRule   string                  `json:"routingRule,omitempty" url:"routingRule,omitempty"`     //	description: , example: , pattern:
	Cleanup       CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`             //	description: , example: , pattern:
	HttpClient    HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`       //	description: , example: , pattern:
	NegativeCache NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"` //	description: , example: , pattern:
	Online        bool                    `json:"online,omitempty" url:"online,omitempty"`               //	description: Whether this repository accepts incoming requests, example: true, pattern:
}

type ContentSelectorApiCreateRequest struct {
	Description string `json:"description,omitempty" url:"description,omitempty"` //	description: A human-readable description, example: , pattern:
	Expression  string `json:"expression,omitempty" url:"expression,omitempty"`   //	description: The expression used to identify content, example: format == "maven2" and path =^ "/org/sonatype/nexus", pattern:
	Name        string `json:"name,omitempty" url:"name,omitempty"`               //	description: The content selector name cannot be changed after creation, example: , pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
}

type ContentSelectorApiResponse struct {
	Description string `json:"description,omitempty" url:"description,omitempty"` //	description: A human-readable description, example: , pattern:
	Expression  string `json:"expression,omitempty" url:"expression,omitempty"`   //	description: The expression used to identify content, example: format == "maven2" and path =^ "/org/sonatype/nexus", pattern:
	Name        string `json:"name,omitempty" url:"name,omitempty"`               //	description: The content selector name cannot be changed after creation, example: , pattern:
	Type        string `json:"type,omitempty" url:"type,omitempty"`               //	description: The type of content selector the backend is using, example: , pattern:
}

type ContentSelectorApiUpdateRequest struct {
	Description string `json:"description,omitempty" url:"description,omitempty"` //	description: An optional description of this content selector, example: , pattern:
	Expression  string `json:"expression,omitempty" url:"expression,omitempty"`   //	description: The expression used to identify content, example: format == "maven2" and path =^ "/org/sonatype/nexus", pattern:
}

type CreateLdapServerXo struct {
	Protocol                    string `json:"protocol,omitempty" url:"protocol,omitempty"`                                       //	description: LDAP server connection Protocol to use, example: , pattern:
	UserBaseDn                  string `json:"userBaseDn,omitempty" url:"userBaseDn,omitempty"`                                   //	description: The relative DN where user objects are found (e.g. ou=people). This value will have the Search base DN value appended to form the full User search base DN., example: ou=people, pattern:
	UserMemberOfAttribute       string `json:"userMemberOfAttribute,omitempty" url:"userMemberOfAttribute,omitempty"`             //	description: Set this to the attribute used to store the attribute which holds groups DN in the user object. Required if groupType is dynamic, example: memberOf, pattern:
	MaxIncidentsCount           int64  `json:"maxIncidentsCount,omitempty" url:"maxIncidentsCount,omitempty"`                     //	description: How many retry attempts, example: , pattern:
	Name                        string `json:"name,omitempty" url:"name,omitempty"`                                               //	description: LDAP server name, example: , pattern:
	UserObjectClass             string `json:"userObjectClass,omitempty" url:"userObjectClass,omitempty"`                         //	description: LDAP class for user objects, example: inetOrgPerson, pattern:
	GroupIdAttribute            string `json:"groupIdAttribute,omitempty" url:"groupIdAttribute,omitempty"`                       //	description: This field specifies the attribute of the Object class that defines the Group ID. Required if groupType is static, example: cn, pattern:
	GroupObjectClass            string `json:"groupObjectClass,omitempty" url:"groupObjectClass,omitempty"`                       //	description: LDAP class for group objects. Required if groupType is static, example: posixGroup, pattern:
	LdapGroupsAsRoles           bool   `json:"ldapGroupsAsRoles,omitempty" url:"ldapGroupsAsRoles,omitempty"`                     //	description: Denotes whether LDAP assigned roles are used as Nexus Repository Manager roles, example: , pattern:
	UserSubtree                 bool   `json:"userSubtree,omitempty" url:"userSubtree,omitempty"`                                 //	description: Are users located in structures below the user base DN?, example: , pattern:
	AuthRealm                   string `json:"authRealm,omitempty" url:"authRealm,omitempty"`                                     //	description: The SASL realm to bind to. Required if authScheme is CRAM_MD5 or DIGEST_MD5, example: example.com, pattern:
	ConnectionTimeoutSeconds    int64  `json:"connectionTimeoutSeconds,omitempty" url:"connectionTimeoutSeconds,omitempty"`       //	description: How long to wait before timeout, example: 1, pattern:
	UserIdAttribute             string `json:"userIdAttribute,omitempty" url:"userIdAttribute,omitempty"`                         //	description: This is used to find a user given its user ID, example: uid, pattern:
	UserLdapFilter              string `json:"userLdapFilter,omitempty" url:"userLdapFilter,omitempty"`                           //	description: LDAP search filter to limit user search, example: (|(mail=*@example.com)(uid=dom*)), pattern:
	UserRealNameAttribute       string `json:"userRealNameAttribute,omitempty" url:"userRealNameAttribute,omitempty"`             //	description: This is used to find a real name given the user ID, example: cn, pattern:
	AuthScheme                  string `json:"authScheme,omitempty" url:"authScheme,omitempty"`                                   //	description: Authentication scheme used for connecting to LDAP server, example: , pattern:
	GroupMemberAttribute        string `json:"groupMemberAttribute,omitempty" url:"groupMemberAttribute,omitempty"`               //	description: LDAP attribute containing the usernames for the group. Required if groupType is static, example: memberUid, pattern:
	AuthPassword                string `json:"authPassword,omitempty" url:"authPassword,omitempty"`                               //	description: The password to bind with. Required if authScheme other than none., example: , pattern:
	ConnectionRetryDelaySeconds int64  `json:"connectionRetryDelaySeconds,omitempty" url:"connectionRetryDelaySeconds,omitempty"` //	description: How long to wait before retrying, example: , pattern:
	GroupBaseDn                 string `json:"groupBaseDn,omitempty" url:"groupBaseDn,omitempty"`                                 //	description: The relative DN where group objects are found (e.g. ou=Group). This value will have the Search base DN value appended to form the full Group search base DN., example: ou=Group, pattern:
	GroupMemberFormat           string `json:"groupMemberFormat,omitempty" url:"groupMemberFormat,omitempty"`                     //	description: The format of user ID stored in the group member attribute. Required if groupType is static, example: uid=${username},ou=people,dc=example,dc=com, pattern:
	GroupSubtree                bool   `json:"groupSubtree,omitempty" url:"groupSubtree,omitempty"`                               //	description: Are groups located in structures below the group base DN, example: , pattern:
	Port                        int64  `json:"port,omitempty" url:"port,omitempty"`                                               //	description: LDAP server connection port to use, example: 636, pattern:
	UseTrustStore               bool   `json:"useTrustStore,omitempty" url:"useTrustStore,omitempty"`                             //	description: Whether to use certificates stored in Nexus Repository Manager's truststore, example: , pattern:
	UserPasswordAttribute       string `json:"userPasswordAttribute,omitempty" url:"userPasswordAttribute,omitempty"`             //	description: If this field is blank the user will be authenticated against a bind with the LDAP server, example: , pattern:
	AuthUsername                string `json:"authUsername,omitempty" url:"authUsername,omitempty"`                               //	description: This must be a fully qualified username if simple authentication is used. Required if authScheme other than none., example: , pattern:
	Host                        string `json:"host,omitempty" url:"host,omitempty"`                                               //	description: LDAP server connection hostname, example: , pattern:
	SearchBase                  string `json:"searchBase,omitempty" url:"searchBase,omitempty"`                                   //	description: LDAP location to be added to the connection URL, example: dc=example,dc=com, pattern:
	GroupType                   string `json:"groupType,omitempty" url:"groupType,omitempty"`                                     //	description: Defines a type of groups used: static (a group contains a list of users) or dynamic (a user contains a list of groups). Required if ldapGroupsAsRoles is true., example: , pattern:
	UserEmailAddressAttribute   string `json:"userEmailAddressAttribute,omitempty" url:"userEmailAddressAttribute,omitempty"`     //	description: This is used to find an email address given the user ID, example: mail, pattern:
}

type DockerAttributes struct {
	ForceBasicAuth bool   `json:"forceBasicAuth,omitempty" url:"forceBasicAuth,omitempty"` //	description: Whether to force authentication (Docker Bearer Token Realm required if false), example: true, pattern:
	HttpPort       int64  `json:"httpPort,omitempty" url:"httpPort,omitempty"`             //	description: Create an HTTP connector at specified port, example: 8082, pattern:
	HttpsPort      int64  `json:"httpsPort,omitempty" url:"httpsPort,omitempty"`           //	description: Create an HTTPS connector at specified port, example: 8083, pattern:
	Subdomain      string `json:"subdomain,omitempty" url:"subdomain,omitempty"`           //	description: Allows to use subdomain, example: docker-a, pattern:
	V1Enabled      bool   `json:"v1Enabled,omitempty" url:"v1Enabled,omitempty"`           //	description: Whether to allow clients to use the V1 API to interact with this repository, example: false, pattern:
}

type DockerGroupApiRepository struct {
	Docker  DockerAttributes      `json:"docker,omitempty" url:"docker,omitempty"`   //	description: , example: , pattern:
	Group   GroupDeployAttributes `json:"group,omitempty" url:"group,omitempty"`     //	description: , example: , pattern:
	Name    string                `json:"name,omitempty" url:"name,omitempty"`       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online  bool                  `json:"online,omitempty" url:"online,omitempty"`   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage StorageAttributes     `json:"storage,omitempty" url:"storage,omitempty"` //	description: , example: , pattern:
}

type DockerGroupRepositoryApiRequest struct {
	Docker  DockerAttributes      `json:"docker,omitempty" url:"docker,omitempty"`   //	description: , example: , pattern:
	Group   GroupDeployAttributes `json:"group,omitempty" url:"group,omitempty"`     //	description: , example: , pattern:
	Name    string                `json:"name,omitempty" url:"name,omitempty"`       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online  bool                  `json:"online,omitempty" url:"online,omitempty"`   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage StorageAttributes     `json:"storage,omitempty" url:"storage,omitempty"` //	description: , example: , pattern:
}

type DockerHostedApiRepository struct {
	Cleanup   CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`     //	description: , example: , pattern:
	Component ComponentAttributes     `json:"component,omitempty" url:"component,omitempty"` //	description: , example: , pattern:
	Docker    DockerAttributes        `json:"docker,omitempty" url:"docker,omitempty"`       //	description: , example: , pattern:
	Name      string                  `json:"name,omitempty" url:"name,omitempty"`           //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online    bool                    `json:"online,omitempty" url:"online,omitempty"`       //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage   HostedStorageAttributes `json:"storage,omitempty" url:"storage,omitempty"`     //	description: , example: , pattern:
}

type DockerHostedRepositoryApiRequest struct {
	Cleanup   CleanupPolicyAttributes       `json:"cleanup,omitempty" url:"cleanup,omitempty"`     //	description: , example: , pattern:
	Component ComponentAttributes           `json:"component,omitempty" url:"component,omitempty"` //	description: , example: , pattern:
	Docker    DockerAttributes              `json:"docker,omitempty" url:"docker,omitempty"`       //	description: , example: , pattern:
	Name      string                        `json:"name,omitempty" url:"name,omitempty"`           //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online    bool                          `json:"online,omitempty" url:"online,omitempty"`       //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage   DockerHostedStorageAttributes `json:"storage,omitempty" url:"storage,omitempty"`     //	description: , example: , pattern:
}

type DockerHostedStorageAttributes struct {
	BlobStoreName               string `json:"blobStoreName,omitempty" url:"blobStoreName,omitempty"`                             //	description: Blob store used to store repository contents, example: default, pattern:
	LatestPolicy                bool   `json:"latestPolicy,omitempty" url:"latestPolicy,omitempty"`                               //	description: Whether to allow redeploying the 'latest' tag but defer to the Deployment Policy for all other tags, example: true, pattern:
	StrictContentTypeValidation bool   `json:"strictContentTypeValidation,omitempty" url:"strictContentTypeValidation,omitempty"` //	description: Whether to validate uploaded content's MIME type appropriate for the repository format, example: true, pattern:
	WritePolicy                 string `json:"writePolicy,omitempty" url:"writePolicy,omitempty"`                                 //	description: Controls if deployments of and updates to assets are allowed, example: allow_once, pattern:
}

type DockerProxyApiRepository struct {
	Cleanup         CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`                 //	description: , example: , pattern:
	Docker          DockerAttributes        `json:"docker,omitempty" url:"docker,omitempty"`                   //	description: , example: , pattern:
	HttpClient      HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`           //	description: , example: , pattern:
	Name            string                  `json:"name,omitempty" url:"name,omitempty"`                       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Proxy           ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                     //	description: , example: , pattern:
	Replication     ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`         //	description: , example: , pattern:
	Storage         StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`                 //	description: , example: , pattern:
	DockerProxy     DockerProxyAttributes   `json:"dockerProxy,omitempty" url:"dockerProxy,omitempty"`         //	description: , example: , pattern:
	NegativeCache   NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"`     //	description: , example: , pattern:
	Online          bool                    `json:"online,omitempty" url:"online,omitempty"`                   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	RoutingRuleName string                  `json:"routingRuleName,omitempty" url:"routingRuleName,omitempty"` //	description: The name of the routing rule assigned to this repository, example: , pattern:
}

type DockerProxyAttributes struct {
	CacheForeignLayers       bool     `json:"cacheForeignLayers,omitempty" url:"cacheForeignLayers,omitempty"`             //	description: Allow Nexus Repository Manager to download and cache foreign layers, example: , pattern:
	ForeignLayerUrlWhitelist []string `json:"foreignLayerUrlWhitelist,omitempty" url:"foreignLayerUrlWhitelist,omitempty"` //	description: Regular expressions used to identify URLs that are allowed for foreign layer requests, example: , pattern:
	IndexType                string   `json:"indexType,omitempty" url:"indexType,omitempty"`                               //	description: Type of Docker Index, example: HUB, pattern:
	IndexUrl                 string   `json:"indexUrl,omitempty" url:"indexUrl,omitempty"`                                 //	description: Url of Docker Index to use, example: , pattern:
}

type DockerProxyRepositoryApiRequest struct {
	Proxy         ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                 //	description: , example: , pattern:
	Storage       StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`             //	description: , example: , pattern:
	Docker        DockerAttributes        `json:"docker,omitempty" url:"docker,omitempty"`               //	description: , example: , pattern:
	DockerProxy   DockerProxyAttributes   `json:"dockerProxy,omitempty" url:"dockerProxy,omitempty"`     //	description: , example: , pattern:
	HttpClient    HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`       //	description: , example: , pattern:
	NegativeCache NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"` //	description: , example: , pattern:
	RoutingRule   string                  `json:"routingRule,omitempty" url:"routingRule,omitempty"`     //	description: , example: , pattern:
	Cleanup       CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`             //	description: , example: , pattern:
	Name          string                  `json:"name,omitempty" url:"name,omitempty"`                   //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online        bool                    `json:"online,omitempty" url:"online,omitempty"`               //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Replication   ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`     //	description: , example: , pattern:
}

type FileBlobStoreApiCreateRequest struct {
	Name      string                `json:"name,omitempty" url:"name,omitempty"`           //	description: , example: , pattern:
	Path      string                `json:"path,omitempty" url:"path,omitempty"`           //	description: The path to the blobstore contents. This can be an absolute path to anywhere on the system Nexus Repository Manager has access to or it can be a path relative to the sonatype-work directory., example: , pattern:
	SoftQuota BlobStoreApiSoftQuota `json:"softQuota,omitempty" url:"softQuota,omitempty"` //	description: Settings to control the soft quota, example: , pattern:
}

type FileBlobStoreApiModel struct {
	Path      string                `json:"path,omitempty" url:"path,omitempty"`           //	description: The path to the blobstore contents. This can be an absolute path to anywhere on the system Nexus Repository Manager has access to or it can be a path relative to the sonatype-work directory., example: , pattern:
	SoftQuota BlobStoreApiSoftQuota `json:"softQuota,omitempty" url:"softQuota,omitempty"` //	description: Settings to control the soft quota, example: , pattern:
}

type FileBlobStoreApiUpdateRequest struct {
	Path      string                `json:"path,omitempty" url:"path,omitempty"`           //	description: The path to the blobstore contents. This can be an absolute path to anywhere on the system Nexus Repository Manager has access to or it can be a path relative to the sonatype-work directory., example: , pattern:
	SoftQuota BlobStoreApiSoftQuota `json:"softQuota,omitempty" url:"softQuota,omitempty"` //	description: Settings to control the soft quota, example: , pattern:
}

type GenericBlobStoreApiResponse struct {
	Name                  string                `json:"name,omitempty" url:"name,omitempty"`                                   //	description: , example: , pattern:
	SoftQuota             BlobStoreApiSoftQuota `json:"softQuota,omitempty" url:"softQuota,omitempty"`                         //	description: Settings to control the soft quota, example: , pattern:
	TotalSizeInBytes      int64                 `json:"totalSizeInBytes,omitempty" url:"totalSizeInBytes,omitempty"`           //	description: , example: , pattern:
	Type                  string                `json:"type,omitempty" url:"type,omitempty"`                                   //	description: , example: , pattern:
	Unavailable           bool                  `json:"unavailable,omitempty" url:"unavailable,omitempty"`                     //	description: , example: , pattern:
	AvailableSpaceInBytes int64                 `json:"availableSpaceInBytes,omitempty" url:"availableSpaceInBytes,omitempty"` //	description: , example: , pattern:
	BlobCount             int64                 `json:"blobCount,omitempty" url:"blobCount,omitempty"`                         //	description: , example: , pattern:
}

type GitLfsHostedRepositoryApiRequest struct {
	Cleanup   CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`     //	description: , example: , pattern:
	Component ComponentAttributes     `json:"component,omitempty" url:"component,omitempty"` //	description: , example: , pattern:
	Name      string                  `json:"name,omitempty" url:"name,omitempty"`           //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online    bool                    `json:"online,omitempty" url:"online,omitempty"`       //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage   HostedStorageAttributes `json:"storage,omitempty" url:"storage,omitempty"`     //	description: , example: , pattern:
}

type GolangGroupRepositoryApiRequest struct {
	Name    string            `json:"name,omitempty" url:"name,omitempty"`       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online  bool              `json:"online,omitempty" url:"online,omitempty"`   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage StorageAttributes `json:"storage,omitempty" url:"storage,omitempty"` //	description: , example: , pattern:
	Group   GroupAttributes   `json:"group,omitempty" url:"group,omitempty"`     //	description: , example: , pattern:
}

type GolangProxyRepositoryApiRequest struct {
	Cleanup       CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`             //	description: , example: , pattern:
	Name          string                  `json:"name,omitempty" url:"name,omitempty"`                   //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	NegativeCache NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"` //	description: , example: , pattern:
	Online        bool                    `json:"online,omitempty" url:"online,omitempty"`               //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Replication   ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`     //	description: , example: , pattern:
	RoutingRule   string                  `json:"routingRule,omitempty" url:"routingRule,omitempty"`     //	description: , example: , pattern:
	HttpClient    HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`       //	description: , example: , pattern:
	Proxy         ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                 //	description: , example: , pattern:
	Storage       StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`             //	description: , example: , pattern:
}

type GroupAttributes struct {
	MemberNames []string `json:"memberNames,omitempty" url:"memberNames,omitempty"` //	description: Member repositories' names, example: , pattern:
}

type GroupDeployAttributes struct {
	WritableMember string   `json:"writableMember,omitempty" url:"writableMember,omitempty"` //	description: Pro-only: This field is for the Group Deployment feature available in NXRM Pro., example: , pattern:
	MemberNames    []string `json:"memberNames,omitempty" url:"memberNames,omitempty"`       //	description: Member repositories' names, example: , pattern:
}

type HelmHostedRepositoryApiRequest struct {
	Cleanup   CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`     //	description: , example: , pattern:
	Component ComponentAttributes     `json:"component,omitempty" url:"component,omitempty"` //	description: , example: , pattern:
	Name      string                  `json:"name,omitempty" url:"name,omitempty"`           //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online    bool                    `json:"online,omitempty" url:"online,omitempty"`       //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage   HostedStorageAttributes `json:"storage,omitempty" url:"storage,omitempty"`     //	description: , example: , pattern:
}

type HelmProxyRepositoryApiRequest struct {
	Cleanup       CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`             //	description: , example: , pattern:
	HttpClient    HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`       //	description: , example: , pattern:
	Name          string                  `json:"name,omitempty" url:"name,omitempty"`                   //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	NegativeCache NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"` //	description: , example: , pattern:
	Online        bool                    `json:"online,omitempty" url:"online,omitempty"`               //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Proxy         ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                 //	description: , example: , pattern:
	Replication   ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`     //	description: , example: , pattern:
	RoutingRule   string                  `json:"routingRule,omitempty" url:"routingRule,omitempty"`     //	description: , example: , pattern:
	Storage       StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`             //	description: , example: , pattern:
}

type HostedStorageAttributes struct {
	BlobStoreName               string `json:"blobStoreName,omitempty" url:"blobStoreName,omitempty"`                             //	description: Blob store used to store repository contents, example: default, pattern:
	StrictContentTypeValidation bool   `json:"strictContentTypeValidation,omitempty" url:"strictContentTypeValidation,omitempty"` //	description: Whether to validate uploaded content's MIME type appropriate for the repository format, example: true, pattern:
	WritePolicy                 string `json:"writePolicy,omitempty" url:"writePolicy,omitempty"`                                 //	description: Controls if deployments of and updates to assets are allowed, example: allow_once, pattern:
}

type HttpClientAttributes struct {
	Authentication HttpClientConnectionAuthenticationAttributes `json:"authentication,omitempty" url:"authentication,omitempty"` //	description: , example: , pattern:
	AutoBlock      bool                                         `json:"autoBlock,omitempty" url:"autoBlock,omitempty"`           //	description: Whether to auto-block outbound connections if remote peer is detected as unreachable/unresponsive, example: true, pattern:
	Blocked        bool                                         `json:"blocked,omitempty" url:"blocked,omitempty"`               //	description: Whether to block outbound connections on the repository, example: false, pattern:
	Connection     HttpClientConnectionAttributes               `json:"connection,omitempty" url:"connection,omitempty"`         //	description: , example: , pattern:
}

type HttpClientAttributesWithPreemptiveAuth struct {
	Authentication HttpClientConnectionAuthenticationAttributesWithPreemptive `json:"authentication,omitempty" url:"authentication,omitempty"` //	description: , example: , pattern:
	AutoBlock      bool                                                       `json:"autoBlock,omitempty" url:"autoBlock,omitempty"`           //	description: Whether to auto-block outbound connections if remote peer is detected as unreachable/unresponsive, example: true, pattern:
	Blocked        bool                                                       `json:"blocked,omitempty" url:"blocked,omitempty"`               //	description: Whether to block outbound connections on the repository, example: false, pattern:
	Connection     HttpClientConnectionAttributes                             `json:"connection,omitempty" url:"connection,omitempty"`         //	description: , example: , pattern:
}

type HttpClientConnectionAttributes struct {
	EnableCircularRedirects bool   `json:"enableCircularRedirects,omitempty" url:"enableCircularRedirects,omitempty"` //	description: Whether to enable redirects to the same location (may be required by some servers), example: false, pattern:
	EnableCookies           bool   `json:"enableCookies,omitempty" url:"enableCookies,omitempty"`                     //	description: Whether to allow cookies to be stored and used, example: false, pattern:
	Retries                 int64  `json:"retries,omitempty" url:"retries,omitempty"`                                 //	description: Total retries if the initial connection attempt suffers a timeout, example: 0, pattern:
	Timeout                 int64  `json:"timeout,omitempty" url:"timeout,omitempty"`                                 //	description: Seconds to wait for activity before stopping and retrying the connection, example: 60, pattern:
	UseTrustStore           bool   `json:"useTrustStore,omitempty" url:"useTrustStore,omitempty"`                     //	description: Use certificates stored in the Nexus Repository Manager truststore to connect to external systems, example: false, pattern:
	UserAgentSuffix         string `json:"userAgentSuffix,omitempty" url:"userAgentSuffix,omitempty"`                 //	description: Custom fragment to append to User-Agent header in HTTP requests, example: , pattern:
}

type HttpClientConnectionAuthenticationAttributes struct {
	Username   string `json:"username,omitempty" url:"username,omitempty"`     //	description: , example: , pattern:
	NtlmDomain string `json:"ntlmDomain,omitempty" url:"ntlmDomain,omitempty"` //	description: , example: , pattern:
	NtlmHost   string `json:"ntlmHost,omitempty" url:"ntlmHost,omitempty"`     //	description: , example: , pattern:
	Password   string `json:"password,omitempty" url:"password,omitempty"`     //	description: , example: , pattern:
	Type       string `json:"type,omitempty" url:"type,omitempty"`             //	description: Authentication type, example: , pattern:
}

type HttpClientConnectionAuthenticationAttributesWithPreemptive struct {
	Password   string `json:"password,omitempty" url:"password,omitempty"`     //	description: , example: , pattern:
	Preemptive bool   `json:"preemptive,omitempty" url:"preemptive,omitempty"` //	description: Whether to use pre-emptive authentication. Use with caution. Defaults to false., example: false, pattern:
	Type       string `json:"type,omitempty" url:"type,omitempty"`             //	description: Authentication type, example: , pattern:
	Username   string `json:"username,omitempty" url:"username,omitempty"`     //	description: , example: , pattern:
	NtlmDomain string `json:"ntlmDomain,omitempty" url:"ntlmDomain,omitempty"` //	description: , example: , pattern:
	NtlmHost   string `json:"ntlmHost,omitempty" url:"ntlmHost,omitempty"`     //	description: , example: , pattern:
}

type InputStream struct {
}

type IqConnectionVerificationXo struct {
	Reason  string `json:"reason,omitempty" url:"reason,omitempty"`   //	description: , example: , pattern:
	Success bool   `json:"success,omitempty" url:"success,omitempty"` //	description: , example: , pattern:
}

type IqConnectionXo struct {
	Url                 string `json:"url,omitempty" url:"url,omitempty"`                                 //	description: The address of your Sonatype Repository Firewall, example: , pattern:
	UseTrustStoreForUrl bool   `json:"useTrustStoreForUrl,omitempty" url:"useTrustStoreForUrl,omitempty"` //	description: Use certificates stored in the Nexus Repository Manager truststore to connect to Sonatype Repository Firewall, example: , pattern:
	FailOpenModeEnabled bool   `json:"failOpenModeEnabled,omitempty" url:"failOpenModeEnabled,omitempty"` //	description: Allow by default when quarantine is enabled and the IQ connection fails, example: , pattern:
	ShowLink            bool   `json:"showLink,omitempty" url:"showLink,omitempty"`                       //	description: Show Sonatype Repository Firewall link in Browse menu when server is enabled, example: , pattern:
	TimeoutSeconds      int64  `json:"timeoutSeconds,omitempty" url:"timeoutSeconds,omitempty"`           //	description: Seconds to wait for activity before stopping and retrying the connection. Leave blank to use the globally defined HTTP timeout., example: , pattern:
	Properties          string `json:"properties,omitempty" url:"properties,omitempty"`                   //	description: Additional properties to configure for Sonatype Repository Firewall, example: , pattern:
	Username            string `json:"username,omitempty" url:"username,omitempty"`                       //	description: User with access to Sonatype Repository Firewall, example: , pattern:
	AuthenticationType  string `json:"authenticationType,omitempty" url:"authenticationType,omitempty"`   //	description: Authentication method, example: , pattern:
	Enabled             bool   `json:"enabled,omitempty" url:"enabled,omitempty"`                         //	description: Whether to use Sonatype Repository Firewall, example: , pattern:
	Password            string `json:"password,omitempty" url:"password,omitempty"`                       //	description: Credentials for the Sonatype Repository Firewall User, example: , pattern:
}

type MavenAttributes struct {
	LayoutPolicy       string `json:"layoutPolicy,omitempty" url:"layoutPolicy,omitempty"`             //	description: Validate that all paths are maven artifact or metadata paths, example: STRICT, pattern:
	VersionPolicy      string `json:"versionPolicy,omitempty" url:"versionPolicy,omitempty"`           //	description: What type of artifacts does this repository store?, example: MIXED, pattern:
	ContentDisposition string `json:"contentDisposition,omitempty" url:"contentDisposition,omitempty"` //	description: Content Disposition, example: ATTACHMENT, pattern:
}

type MavenGroupRepositoryApiRequest struct {
	Online  bool              `json:"online,omitempty" url:"online,omitempty"`   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage StorageAttributes `json:"storage,omitempty" url:"storage,omitempty"` //	description: , example: , pattern:
	Group   GroupAttributes   `json:"group,omitempty" url:"group,omitempty"`     //	description: , example: , pattern:
	Name    string            `json:"name,omitempty" url:"name,omitempty"`       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
}

type MavenHostedApiRepository struct {
	Maven     MavenAttributes         `json:"maven,omitempty" url:"maven,omitempty"`         //	description: , example: , pattern:
	Name      string                  `json:"name,omitempty" url:"name,omitempty"`           //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online    bool                    `json:"online,omitempty" url:"online,omitempty"`       //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage   HostedStorageAttributes `json:"storage,omitempty" url:"storage,omitempty"`     //	description: , example: , pattern:
	Cleanup   CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`     //	description: , example: , pattern:
	Component ComponentAttributes     `json:"component,omitempty" url:"component,omitempty"` //	description: , example: , pattern:
}

type MavenHostedRepositoryApiRequest struct {
	Maven     MavenAttributes         `json:"maven,omitempty" url:"maven,omitempty"`         //	description: , example: , pattern:
	Name      string                  `json:"name,omitempty" url:"name,omitempty"`           //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online    bool                    `json:"online,omitempty" url:"online,omitempty"`       //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage   HostedStorageAttributes `json:"storage,omitempty" url:"storage,omitempty"`     //	description: , example: , pattern:
	Cleanup   CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`     //	description: , example: , pattern:
	Component ComponentAttributes     `json:"component,omitempty" url:"component,omitempty"` //	description: , example: , pattern:
}

type MavenProxyApiRepository struct {
	Name            string                  `json:"name,omitempty" url:"name,omitempty"`                       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	NegativeCache   NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"`     //	description: , example: , pattern:
	RoutingRuleName string                  `json:"routingRuleName,omitempty" url:"routingRuleName,omitempty"` //	description: The name of the routing rule assigned to this repository, example: , pattern:
	Maven           MavenAttributes         `json:"maven,omitempty" url:"maven,omitempty"`                     //	description: , example: , pattern:
	HttpClient      HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`           //	description: , example: , pattern:
	Online          bool                    `json:"online,omitempty" url:"online,omitempty"`                   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Proxy           ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                     //	description: , example: , pattern:
	Replication     ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`         //	description: , example: , pattern:
	Storage         StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`                 //	description: , example: , pattern:
	Cleanup         CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`                 //	description: , example: , pattern:
}

type MavenProxyRepositoryApiRequest struct {
	Name          string                                 `json:"name,omitempty" url:"name,omitempty"`                   //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	NegativeCache NegativeCacheAttributes                `json:"negativeCache,omitempty" url:"negativeCache,omitempty"` //	description: , example: , pattern:
	Online        bool                                   `json:"online,omitempty" url:"online,omitempty"`               //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Proxy         ProxyAttributes                        `json:"proxy,omitempty" url:"proxy,omitempty"`                 //	description: , example: , pattern:
	Cleanup       CleanupPolicyAttributes                `json:"cleanup,omitempty" url:"cleanup,omitempty"`             //	description: , example: , pattern:
	HttpClient    HttpClientAttributesWithPreemptiveAuth `json:"httpClient,omitempty" url:"httpClient,omitempty"`       //	description: , example: , pattern:
	Maven         MavenAttributes                        `json:"maven,omitempty" url:"maven,omitempty"`                 //	description: , example: , pattern:
	Replication   ReplicationAttributes                  `json:"replication,omitempty" url:"replication,omitempty"`     //	description: , example: , pattern:
	RoutingRule   string                                 `json:"routingRule,omitempty" url:"routingRule,omitempty"`     //	description: , example: , pattern:
	Storage       StorageAttributes                      `json:"storage,omitempty" url:"storage,omitempty"`             //	description: , example: , pattern:
}

type NegativeCacheAttributes struct {
	Enabled    bool  `json:"enabled,omitempty" url:"enabled,omitempty"`       //	description: Whether to cache responses for content not present in the proxied repository, example: true, pattern:
	TimeToLive int64 `json:"timeToLive,omitempty" url:"timeToLive,omitempty"` //	description: How long to cache the fact that a file was not found in the repository (in minutes), example: 1440, pattern:
}

type NpmAttributes struct {
	RemoveNonCataloged bool `json:"removeNonCataloged,omitempty" url:"removeNonCataloged,omitempty"` //	description: Remove Non-Cataloged Versions, example: true, pattern:
	RemoveQuarantined  bool `json:"removeQuarantined,omitempty" url:"removeQuarantined,omitempty"`   //	description: Remove Quarantined Versions, example: true, pattern:
}

type NpmGroupRepositoryApiRequest struct {
	Group   GroupDeployAttributes `json:"group,omitempty" url:"group,omitempty"`     //	description: , example: , pattern:
	Name    string                `json:"name,omitempty" url:"name,omitempty"`       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online  bool                  `json:"online,omitempty" url:"online,omitempty"`   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage StorageAttributes     `json:"storage,omitempty" url:"storage,omitempty"` //	description: , example: , pattern:
}

type NpmHostedRepositoryApiRequest struct {
	Storage   HostedStorageAttributes `json:"storage,omitempty" url:"storage,omitempty"`     //	description: , example: , pattern:
	Cleanup   CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`     //	description: , example: , pattern:
	Component ComponentAttributes     `json:"component,omitempty" url:"component,omitempty"` //	description: , example: , pattern:
	Name      string                  `json:"name,omitempty" url:"name,omitempty"`           //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online    bool                    `json:"online,omitempty" url:"online,omitempty"`       //	description: Whether this repository accepts incoming requests, example: true, pattern:
}

type NpmProxyApiRepository struct {
	Online          bool                    `json:"online,omitempty" url:"online,omitempty"`                   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Proxy           ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                     //	description: , example: , pattern:
	Replication     ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`         //	description: , example: , pattern:
	Cleanup         CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`                 //	description: , example: , pattern:
	HttpClient      HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`           //	description: , example: , pattern:
	Name            string                  `json:"name,omitempty" url:"name,omitempty"`                       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	NegativeCache   NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"`     //	description: , example: , pattern:
	Npm             NpmAttributes           `json:"npm,omitempty" url:"npm,omitempty"`                         //	description: , example: , pattern:
	RoutingRuleName string                  `json:"routingRuleName,omitempty" url:"routingRuleName,omitempty"` //	description: The name of the routing rule assigned to this repository, example: , pattern:
	Storage         StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`                 //	description: , example: , pattern:
}

type NpmProxyRepositoryApiRequest struct {
	RoutingRule   string                  `json:"routingRule,omitempty" url:"routingRule,omitempty"`     //	description: , example: , pattern:
	HttpClient    HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`       //	description: , example: , pattern:
	NegativeCache NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"` //	description: , example: , pattern:
	Npm           NpmAttributes           `json:"npm,omitempty" url:"npm,omitempty"`                     //	description: , example: , pattern:
	Replication   ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`     //	description: , example: , pattern:
	Storage       StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`             //	description: , example: , pattern:
	Cleanup       CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`             //	description: , example: , pattern:
	Name          string                  `json:"name,omitempty" url:"name,omitempty"`                   //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online        bool                    `json:"online,omitempty" url:"online,omitempty"`               //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Proxy         ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                 //	description: , example: , pattern:
}

type NugetAttributes struct {
	NugetVersion         string `json:"nugetVersion,omitempty" url:"nugetVersion,omitempty"`                 //	description: Nuget protocol version, example: V3, pattern:
	QueryCacheItemMaxAge int64  `json:"queryCacheItemMaxAge,omitempty" url:"queryCacheItemMaxAge,omitempty"` //	description: How long to cache query results from the proxied repository (in seconds), example: 3600, pattern:

}

type NugetGroupRepositoryApiRequest struct {
	Group   GroupAttributes   `json:"group,omitempty" url:"group,omitempty"`     //	description: , example: , pattern:
	Name    string            `json:"name,omitempty" url:"name,omitempty"`       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online  bool              `json:"online,omitempty" url:"online,omitempty"`   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage StorageAttributes `json:"storage,omitempty" url:"storage,omitempty"` //	description: , example: , pattern:
}

type NugetHostedRepositoryApiRequest struct {
	Cleanup   CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`     //	description: , example: , pattern:
	Component ComponentAttributes     `json:"component,omitempty" url:"component,omitempty"` //	description: , example: , pattern:
	Name      string                  `json:"name,omitempty" url:"name,omitempty"`           //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online    bool                    `json:"online,omitempty" url:"online,omitempty"`       //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage   HostedStorageAttributes `json:"storage,omitempty" url:"storage,omitempty"`     //	description: , example: , pattern:
}

type NugetProxyApiRepository struct {
	RoutingRuleName string                  `json:"routingRuleName,omitempty" url:"routingRuleName,omitempty"` //	description: The name of the routing rule assigned to this repository, example: , pattern:
	Cleanup         CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`                 //	description: , example: , pattern:
	Name            string                  `json:"name,omitempty" url:"name,omitempty"`                       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	NugetProxy      NugetAttributes         `json:"nugetProxy,omitempty" url:"nugetProxy,omitempty"`           //	description: , example: , pattern:
	Online          bool                    `json:"online,omitempty" url:"online,omitempty"`                   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage         StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`                 //	description: , example: , pattern:
	HttpClient      HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`           //	description: , example: , pattern:
	NegativeCache   NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"`     //	description: , example: , pattern:

	Proxy       ProxyAttributes       `json:"proxy,omitempty" url:"proxy,omitempty"`             //	description: , example: , pattern:
	Replication ReplicationAttributes `json:"replication,omitempty" url:"replication,omitempty"` //	description: , example: , pattern:
}

type NugetProxyRepositoryApiRequest struct {
	Name          string                  `json:"name,omitempty" url:"name,omitempty"`                   //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Replication   ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`     //	description: , example: , pattern:
	Storage       StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`             //	description: , example: , pattern:
	Cleanup       CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`             //	description: , example: , pattern:
	HttpClient    HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`       //	description: , example: , pattern:
	NegativeCache NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"` //	description: , example: , pattern:
	NugetProxy    NugetAttributes         `json:"nugetProxy,omitempty" url:"nugetProxy,omitempty"`       //	description: , example: , pattern:
	Online        bool                    `json:"online,omitempty" url:"online,omitempty"`               //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Proxy         ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                 //	description: , example: , pattern:
	RoutingRule   string                  `json:"routingRule,omitempty" url:"routingRule,omitempty"`     //	description: , example: , pattern:
}

type P2ProxyRepositoryApiRequest struct {
	Name          string                  `json:"name,omitempty" url:"name,omitempty"`                   //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Replication   ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`     //	description: , example: , pattern:
	RoutingRule   string                  `json:"routingRule,omitempty" url:"routingRule,omitempty"`     //	description: , example: , pattern:
	Storage       StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`             //	description: , example: , pattern:
	Cleanup       CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`             //	description: , example: , pattern:
	HttpClient    HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`       //	description: , example: , pattern:
	NegativeCache NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"` //	description: , example: , pattern:
	Online        bool                    `json:"online,omitempty" url:"online,omitempty"`               //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Proxy         ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                 //	description: , example: , pattern:
}

type Page struct {
	ContinuationToken string        `json:"continuationToken,omitempty" url:"continuationToken,omitempty"` //	description: , example: , pattern:
	Items             []interface{} `json:"items,omitempty" url:"items,omitempty"`                         //	description: , example: , pattern:
}

type PageAssetXO struct {
	ContinuationToken string    `json:"continuationToken,omitempty" url:"continuationToken,omitempty"` //	description: , example: , pattern:
	Items             []AssetXO `json:"items,omitempty" url:"items,omitempty"`                         //	description: , example: , pattern:
}

type PageComponentXO struct {
	ContinuationToken string        `json:"continuationToken,omitempty" url:"continuationToken,omitempty"` //	description: , example: , pattern:
	Items             []ComponentXO `json:"items,omitempty" url:"items,omitempty"`                         //	description: , example: , pattern:
}

type PageTaskXO struct {
	ContinuationToken string   `json:"continuationToken,omitempty" url:"continuationToken,omitempty"` //	description: , example: , pattern:
	Items             []TaskXO `json:"items,omitempty" url:"items,omitempty"`                         //	description: , example: , pattern:
}

type ProxyAttributes struct {
	ContentMaxAge  int64  `json:"contentMaxAge,omitempty" url:"contentMaxAge,omitempty"`   //	description: How long to cache artifacts before rechecking the remote repository (in minutes), example: 1440, pattern:
	MetadataMaxAge int64  `json:"metadataMaxAge,omitempty" url:"metadataMaxAge,omitempty"` //	description: How long to cache metadata before rechecking the remote repository (in minutes), example: 1440, pattern:
	RemoteUrl      string `json:"remoteUrl,omitempty" url:"remoteUrl,omitempty"`           //	description: Location of the remote repository being proxied, example: https://remote.repository.com, pattern:
}

type PyPiProxyAttributes struct {
	RemoveQuarantined bool `json:"removeQuarantined,omitempty" url:"removeQuarantined,omitempty"` //	description: Remove Quarantined Versions, example: true, pattern:
}

type PypiGroupRepositoryApiRequest struct {
	Storage StorageAttributes `json:"storage,omitempty" url:"storage,omitempty"` //	description: , example: , pattern:
	Group   GroupAttributes   `json:"group,omitempty" url:"group,omitempty"`     //	description: , example: , pattern:
	Name    string            `json:"name,omitempty" url:"name,omitempty"`       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online  bool              `json:"online,omitempty" url:"online,omitempty"`   //	description: Whether this repository accepts incoming requests, example: true, pattern:
}

type PypiHostedRepositoryApiRequest struct {
	Name      string                  `json:"name,omitempty" url:"name,omitempty"`           //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online    bool                    `json:"online,omitempty" url:"online,omitempty"`       //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage   HostedStorageAttributes `json:"storage,omitempty" url:"storage,omitempty"`     //	description: , example: , pattern:
	Cleanup   CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`     //	description: , example: , pattern:
	Component ComponentAttributes     `json:"component,omitempty" url:"component,omitempty"` //	description: , example: , pattern:
}

type PypiProxyRepositoryApiRequest struct {
	Pypi          PyPiProxyAttributes     `json:"pypi,omitempty" url:"pypi,omitempty"`                   //	description: , example: , pattern:
	Replication   ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`     //	description: , example: , pattern:
	Storage       StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`             //	description: , example: , pattern:
	HttpClient    HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`       //	description: , example: , pattern:
	NegativeCache NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"` //	description: , example: , pattern:

	Online      bool                    `json:"online,omitempty" url:"online,omitempty"`           //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Proxy       ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`             //	description: , example: , pattern:
	Cleanup     CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`         //	description: , example: , pattern:
	Name        string                  `json:"name,omitempty" url:"name,omitempty"`               //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	RoutingRule string                  `json:"routingRule,omitempty" url:"routingRule,omitempty"` //	description: , example: , pattern:
}

type RGroupRepositoryApiRequest struct {
	Group   GroupAttributes   `json:"group,omitempty" url:"group,omitempty"`     //	description: , example: , pattern:
	Name    string            `json:"name,omitempty" url:"name,omitempty"`       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online  bool              `json:"online,omitempty" url:"online,omitempty"`   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage StorageAttributes `json:"storage,omitempty" url:"storage,omitempty"` //	description: , example: , pattern:
}

type RHostedRepositoryApiRequest struct {
	Cleanup   CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`     //	description: , example: , pattern:
	Component ComponentAttributes     `json:"component,omitempty" url:"component,omitempty"` //	description: , example: , pattern:
	Name      string                  `json:"name,omitempty" url:"name,omitempty"`           //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online    bool                    `json:"online,omitempty" url:"online,omitempty"`       //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage   HostedStorageAttributes `json:"storage,omitempty" url:"storage,omitempty"`     //	description: , example: , pattern:
}

type RProxyRepositoryApiRequest struct {
	Name          string                  `json:"name,omitempty" url:"name,omitempty"`                   //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	NegativeCache NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"` //	description: , example: , pattern:
	Proxy         ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                 //	description: , example: , pattern:
	Replication   ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`     //	description: , example: , pattern:
	RoutingRule   string                  `json:"routingRule,omitempty" url:"routingRule,omitempty"`     //	description: , example: , pattern:
	Cleanup       CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`             //	description: , example: , pattern:
	HttpClient    HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`       //	description: , example: , pattern:
	Online        bool                    `json:"online,omitempty" url:"online,omitempty"`               //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage       StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`             //	description: , example: , pattern:
}

type RawAttributes struct {
	ContentDisposition string `json:"contentDisposition,omitempty" url:"contentDisposition,omitempty"` //	description: Content Disposition, example: ATTACHMENT, pattern:
}

type RawGroupRepositoryApiRequest struct {
	Group   GroupAttributes   `json:"group,omitempty" url:"group,omitempty"`     //	description: , example: , pattern:
	Name    string            `json:"name,omitempty" url:"name,omitempty"`       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online  bool              `json:"online,omitempty" url:"online,omitempty"`   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Raw     RawAttributes     `json:"raw,omitempty" url:"raw,omitempty"`         //	description: , example: , pattern:
	Storage StorageAttributes `json:"storage,omitempty" url:"storage,omitempty"` //	description: , example: , pattern:
}

type RawHostedRepositoryApiRequest struct {
	Name      string                  `json:"name,omitempty" url:"name,omitempty"`           //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online    bool                    `json:"online,omitempty" url:"online,omitempty"`       //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Raw       RawAttributes           `json:"raw,omitempty" url:"raw,omitempty"`             //	description: , example: , pattern:
	Storage   HostedStorageAttributes `json:"storage,omitempty" url:"storage,omitempty"`     //	description: , example: , pattern:
	Cleanup   CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`     //	description: , example: , pattern:
	Component ComponentAttributes     `json:"component,omitempty" url:"component,omitempty"` //	description: , example: , pattern:
}

type RawProxyRepositoryApiRequest struct {
	Storage       StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`             //	description: , example: , pattern:
	Proxy         ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                 //	description: , example: , pattern:
	Raw           RawAttributes           `json:"raw,omitempty" url:"raw,omitempty"`                     //	description: , example: , pattern:
	Name          string                  `json:"name,omitempty" url:"name,omitempty"`                   //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	NegativeCache NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"` //	description: , example: , pattern:
	Online        bool                    `json:"online,omitempty" url:"online,omitempty"`               //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Replication   ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`     //	description: , example: , pattern:
	RoutingRule   string                  `json:"routingRule,omitempty" url:"routingRule,omitempty"`     //	description: , example: , pattern:
	Cleanup       CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`             //	description: , example: , pattern:
	HttpClient    HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`       //	description: , example: , pattern:
}

type ReadLdapServerXo struct {
	GroupBaseDn                 string `json:"groupBaseDn,omitempty" url:"groupBaseDn,omitempty"`                                 //	description: The relative DN where group objects are found (e.g. ou=Group). This value will have the Search base DN value appended to form the full Group search base DN., example: ou=Group, pattern:
	Port                        int64  `json:"port,omitempty" url:"port,omitempty"`                                               //	description: LDAP server connection port to use, example: 636, pattern:
	UserObjectClass             string `json:"userObjectClass,omitempty" url:"userObjectClass,omitempty"`                         //	description: LDAP class for user objects, example: inetOrgPerson, pattern:
	UserRealNameAttribute       string `json:"userRealNameAttribute,omitempty" url:"userRealNameAttribute,omitempty"`             //	description: This is used to find a real name given the user ID, example: cn, pattern:
	GroupMemberAttribute        string `json:"groupMemberAttribute,omitempty" url:"groupMemberAttribute,omitempty"`               //	description: LDAP attribute containing the usernames for the group. Required if groupType is static, example: memberUid, pattern:
	SearchBase                  string `json:"searchBase,omitempty" url:"searchBase,omitempty"`                                   //	description: LDAP location to be added to the connection URL, example: dc=example,dc=com, pattern:
	UserLdapFilter              string `json:"userLdapFilter,omitempty" url:"userLdapFilter,omitempty"`                           //	description: LDAP search filter to limit user search, example: (|(mail=*@example.com)(uid=dom*)), pattern:
	UserPasswordAttribute       string `json:"userPasswordAttribute,omitempty" url:"userPasswordAttribute,omitempty"`             //	description: If this field is blank the user will be authenticated against a bind with the LDAP server, example: , pattern:
	ConnectionRetryDelaySeconds int64  `json:"connectionRetryDelaySeconds,omitempty" url:"connectionRetryDelaySeconds,omitempty"` //	description: How long to wait before retrying, example: , pattern:
	GroupType                   string `json:"groupType,omitempty" url:"groupType,omitempty"`                                     //	description: Defines a type of groups used: static (a group contains a list of users) or dynamic (a user contains a list of groups). Required if ldapGroupsAsRoles is true., example: , pattern:
	UserIdAttribute             string `json:"userIdAttribute,omitempty" url:"userIdAttribute,omitempty"`                         //	description: This is used to find a user given its user ID, example: uid, pattern:
	UserMemberOfAttribute       string `json:"userMemberOfAttribute,omitempty" url:"userMemberOfAttribute,omitempty"`             //	description: Set this to the attribute used to store the attribute which holds groups DN in the user object. Required if groupType is dynamic, example: memberOf, pattern:
	Name                        string `json:"name,omitempty" url:"name,omitempty"`                                               //	description: LDAP server name, example: , pattern:
	UserSubtree                 bool   `json:"userSubtree,omitempty" url:"userSubtree,omitempty"`                                 //	description: Are users located in structures below the user base DN?, example: , pattern:
	GroupIdAttribute            string `json:"groupIdAttribute,omitempty" url:"groupIdAttribute,omitempty"`                       //	description: This field specifies the attribute of the Object class that defines the Group ID. Required if groupType is static, example: cn, pattern:
	GroupSubtree                bool   `json:"groupSubtree,omitempty" url:"groupSubtree,omitempty"`                               //	description: Are groups located in structures below the group base DN, example: , pattern:
	LdapGroupsAsRoles           bool   `json:"ldapGroupsAsRoles,omitempty" url:"ldapGroupsAsRoles,omitempty"`                     //	description: Denotes whether LDAP assigned roles are used as Nexus Repository Manager roles, example: , pattern:
	MaxIncidentsCount           int64  `json:"maxIncidentsCount,omitempty" url:"maxIncidentsCount,omitempty"`                     //	description: How many retry attempts, example: , pattern:
	AuthUsername                string `json:"authUsername,omitempty" url:"authUsername,omitempty"`                               //	description: This must be a fully qualified username if simple authentication is used. Required if authScheme other than none., example: , pattern:
	GroupMemberFormat           string `json:"groupMemberFormat,omitempty" url:"groupMemberFormat,omitempty"`                     //	description: The format of user ID stored in the group member attribute. Required if groupType is static, example: uid=${username},ou=people,dc=example,dc=com, pattern:
	UseTrustStore               bool   `json:"useTrustStore,omitempty" url:"useTrustStore,omitempty"`                             //	description: Whether to use certificates stored in Nexus Repository Manager's truststore, example: , pattern:
	UserBaseDn                  string `json:"userBaseDn,omitempty" url:"userBaseDn,omitempty"`                                   //	description: The relative DN where user objects are found (e.g. ou=people). This value will have the Search base DN value appended to form the full User search base DN., example: ou=people, pattern:
	AuthRealm                   string `json:"authRealm,omitempty" url:"authRealm,omitempty"`                                     //	description: The SASL realm to bind to. Required if authScheme is CRAM_MD5 or DIGEST_MD5, example: example.com, pattern:
	AuthScheme                  string `json:"authScheme,omitempty" url:"authScheme,omitempty"`                                   //	description: Authentication scheme used for connecting to LDAP server, example: , pattern:
	ConnectionTimeoutSeconds    int64  `json:"connectionTimeoutSeconds,omitempty" url:"connectionTimeoutSeconds,omitempty"`       //	description: How long to wait before timeout, example: 1, pattern:
	Id                          string `json:"id,omitempty" url:"id,omitempty"`                                                   //	description: LDAP server ID, example: , pattern:
	Order                       int64  `json:"order,omitempty" url:"order,omitempty"`                                             //	description: Order number in which the server is being used when looking for a user, example: , pattern:
	Protocol                    string `json:"protocol,omitempty" url:"protocol,omitempty"`                                       //	description: LDAP server connection Protocol to use, example: , pattern:
	UserEmailAddressAttribute   string `json:"userEmailAddressAttribute,omitempty" url:"userEmailAddressAttribute,omitempty"`     //	description: This is used to find an email address given the user ID, example: mail, pattern:
	GroupObjectClass            string `json:"groupObjectClass,omitempty" url:"groupObjectClass,omitempty"`                       //	description: LDAP class for group objects. Required if groupType is static, example: posixGroup, pattern:
	Host                        string `json:"host,omitempty" url:"host,omitempty"`                                               //	description: LDAP server connection hostname, example: , pattern:
}

type ReadOnlyState struct {
	Frozen          bool   `json:"frozen,omitempty" url:"frozen,omitempty"`                   //	description: , example: , pattern:
	SummaryReason   string `json:"summaryReason,omitempty" url:"summaryReason,omitempty"`     //	description: , example: , pattern:
	SystemInitiated bool   `json:"systemInitiated,omitempty" url:"systemInitiated,omitempty"` //	description: , example: , pattern:
}

type RealmApiXO struct {
	Id   string `json:"id,omitempty" url:"id,omitempty"`     //	description: , example: , pattern:
	Name string `json:"name,omitempty" url:"name,omitempty"` //	description: , example: , pattern:
}

type ReplicationAttributes struct {
	AssetPathRegex        string `json:"assetPathRegex,omitempty" url:"assetPathRegex,omitempty"`               //	description: Regular Expression of Asset Paths to pull pre-emptively pull, example: , pattern:
	PreemptivePullEnabled bool   `json:"preemptivePullEnabled,omitempty" url:"preemptivePullEnabled,omitempty"` //	description: Whether pre-emptive pull is enabled, example: false, pattern:
}

type RepositoryXO struct {
	Format     string      `json:"format,omitempty" url:"format,omitempty"`         //	description: , example: , pattern:
	Name       string      `json:"name,omitempty" url:"name,omitempty"`             //	description: , example: , pattern:
	Type       string      `json:"type,omitempty" url:"type,omitempty"`             //	description: , example: , pattern:
	Url        string      `json:"url,omitempty" url:"url,omitempty"`               //	description: , example: , pattern:
	Attributes interface{} `json:"attributes,omitempty" url:"attributes,omitempty"` //	description: , example: , pattern:
}

type Result struct {
	Message   string      `json:"message,omitempty" url:"message,omitempty"`     //	description: , example: , pattern:
	Time      int64       `json:"time,omitempty" url:"time,omitempty"`           //	description: , example: , pattern:
	Timestamp string      `json:"timestamp,omitempty" url:"timestamp,omitempty"` //	description: , example: , pattern:
	Details   interface{} `json:"details,omitempty" url:"details,omitempty"`     //	description: , example: , pattern:
	Duration  int64       `json:"duration,omitempty" url:"duration,omitempty"`   //	description: , example: , pattern:
	Error     Throwable   `json:"error,omitempty" url:"error,omitempty"`         //	description: , example: , pattern:
	Healthy   bool        `json:"healthy,omitempty" url:"healthy,omitempty"`     //	description: , example: , pattern:
}

type RoleXORequest struct {
	Privileges  []string `json:"privileges,omitempty" url:"privileges,omitempty"`   //	description: The list of privileges assigned to this role., example: , pattern:
	Roles       []string `json:"roles,omitempty" url:"roles,omitempty"`             //	description: The list of roles assigned to this role., example: , pattern:
	Description string   `json:"description,omitempty" url:"description,omitempty"` //	description: The description of this role., example: , pattern:
	Id          string   `json:"id,omitempty" url:"id,omitempty"`                   //	description: The id of the role., example: , pattern:
	Name        string   `json:"name,omitempty" url:"name,omitempty"`               //	description: The name of the role., example: , pattern:
}

type RoleXOResponse struct {
	Roles       []string `json:"roles,omitempty" url:"roles,omitempty"`             //	description: The list of roles assigned to this role., example: , pattern:
	Source      string   `json:"source,omitempty" url:"source,omitempty"`           //	description: The user source which is the origin of this role., example: , pattern:
	Description string   `json:"description,omitempty" url:"description,omitempty"` //	description: The description of this role., example: , pattern:
	Id          string   `json:"id,omitempty" url:"id,omitempty"`                   //	description: The id of the role., example: , pattern:
	Name        string   `json:"name,omitempty" url:"name,omitempty"`               //	description: The name of the role., example: , pattern:
	Privileges  []string `json:"privileges,omitempty" url:"privileges,omitempty"`   //	description: The list of privileges assigned to this role., example: , pattern:
	ReadOnly    bool     `json:"readOnly,omitempty" url:"readOnly,omitempty"`       //	description: Indicates whether the role can be changed. The system will ignore any supplied external values., example: , pattern:
}

type RoutingRuleXO struct {
	Matchers    []string `json:"matchers,omitempty" url:"matchers,omitempty"`       //	description: Regular expressions used to identify request paths that are allowed or blocked (depending on mode), example: , pattern:
	Mode        string   `json:"mode,omitempty" url:"mode,omitempty"`               //	description: Determines what should be done with requests when their path matches any of the matchers, example: , pattern:
	Name        string   `json:"name,omitempty" url:"name,omitempty"`               //	description: , example: , pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Description string   `json:"description,omitempty" url:"description,omitempty"` //	description: , example: , pattern:
}

type RubyGemsGroupRepositoryApiRequest struct {
	Storage StorageAttributes `json:"storage,omitempty" url:"storage,omitempty"` //	description: , example: , pattern:
	Group   GroupAttributes   `json:"group,omitempty" url:"group,omitempty"`     //	description: , example: , pattern:
	Name    string            `json:"name,omitempty" url:"name,omitempty"`       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online  bool              `json:"online,omitempty" url:"online,omitempty"`   //	description: Whether this repository accepts incoming requests, example: true, pattern:
}

type RubyGemsHostedRepositoryApiRequest struct {
	Name      string                  `json:"name,omitempty" url:"name,omitempty"`           //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online    bool                    `json:"online,omitempty" url:"online,omitempty"`       //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage   HostedStorageAttributes `json:"storage,omitempty" url:"storage,omitempty"`     //	description: , example: , pattern:
	Cleanup   CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`     //	description: , example: , pattern:
	Component ComponentAttributes     `json:"component,omitempty" url:"component,omitempty"` //	description: , example: , pattern:
}

type RubyGemsProxyRepositoryApiRequest struct {
	Storage       StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`             //	description: , example: , pattern:
	HttpClient    HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`       //	description: , example: , pattern:
	Name          string                  `json:"name,omitempty" url:"name,omitempty"`                   //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	NegativeCache NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"` //	description: , example: , pattern:
	Online        bool                    `json:"online,omitempty" url:"online,omitempty"`               //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Proxy         ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                 //	description: , example: , pattern:
	Replication   ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`     //	description: , example: , pattern:
	RoutingRule   string                  `json:"routingRule,omitempty" url:"routingRule,omitempty"`     //	description: , example: , pattern:
	Cleanup       CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`             //	description: , example: , pattern:
}

type S3BlobStoreApiAdvancedBucketConnection struct {
	Endpoint              string `json:"endpoint,omitempty" url:"endpoint,omitempty"`                           //	description: A custom endpoint URL for third party object stores using the S3 API., example: , pattern:
	ForcePathStyle        bool   `json:"forcePathStyle,omitempty" url:"forcePathStyle,omitempty"`               //	description: Setting this flag will result in path-style access being used for all requests., example: , pattern:
	MaxConnectionPoolSize int64  `json:"maxConnectionPoolSize,omitempty" url:"maxConnectionPoolSize,omitempty"` //	description: Setting this value will override the default connection pool size of Nexus of the s3 client for this blobstore., example: , pattern:
	SignerType            string `json:"signerType,omitempty" url:"signerType,omitempty"`                       //	description: An API signature version which may be required for third party object stores using the S3 API., example: , pattern:
}

type S3BlobStoreApiBucket struct {
	Expiration int64  `json:"expiration,omitempty" url:"expiration,omitempty"` //	description: How many days until deleted blobs are finally removed from the S3 bucket (-1 to disable), example: 3, pattern:
	Name       string `json:"name,omitempty" url:"name,omitempty"`             //	description: The name of the S3 bucket, example: , pattern:
	Prefix     string `json:"prefix,omitempty" url:"prefix,omitempty"`         //	description: The S3 blob store (i.e S3 object) key prefix, example: , pattern:
	Region     string `json:"region,omitempty" url:"region,omitempty"`         //	description: The AWS region to create a new S3 bucket in or an existing S3 bucket's region, example: DEFAULT, pattern:
}

type S3BlobStoreApiBucketConfiguration struct {
	Bucket                   S3BlobStoreApiBucket                   `json:"bucket,omitempty" url:"bucket,omitempty"`                                     //	description: Details of the S3 bucket such as name and region, example: , pattern:
	BucketSecurity           S3BlobStoreApiBucketSecurity           `json:"bucketSecurity,omitempty" url:"bucketSecurity,omitempty"`                     //	description: Security details for granting access the S3 API, example: , pattern:
	Encryption               S3BlobStoreApiEncryption               `json:"encryption,omitempty" url:"encryption,omitempty"`                             //	description: The type of encryption to use if any, example: , pattern:
	AdvancedBucketConnection S3BlobStoreApiAdvancedBucketConnection `json:"advancedBucketConnection,omitempty" url:"advancedBucketConnection,omitempty"` //	description: A custom endpoint URL, signer type and whether path style access is enabled, example: , pattern:
}

type S3BlobStoreApiBucketSecurity struct {
	SessionToken    string `json:"sessionToken,omitempty" url:"sessionToken,omitempty"`       //	description: An AWS STS session token associated with temporary security credentials which grant access to the S3 bucket, example: , pattern:
	AccessKeyId     string `json:"accessKeyId,omitempty" url:"accessKeyId,omitempty"`         //	description: An IAM access key ID for granting access to the S3 bucket, example: , pattern:
	Role            string `json:"role,omitempty" url:"role,omitempty"`                       //	description: An IAM role to assume in order to access the S3 bucket, example: , pattern:
	SecretAccessKey string `json:"secretAccessKey,omitempty" url:"secretAccessKey,omitempty"` //	description: The secret access key associated with the specified IAM access key ID, example: , pattern:
}

type S3BlobStoreApiEncryption struct {
	EncryptionKey  string `json:"encryptionKey,omitempty" url:"encryptionKey,omitempty"`   //	description: The encryption key., example: , pattern:
	EncryptionType string `json:"encryptionType,omitempty" url:"encryptionType,omitempty"` //	description: The type of S3 server side encryption to use., example: , pattern:
}

type S3BlobStoreApiModel struct {
	BucketConfiguration S3BlobStoreApiBucketConfiguration `json:"bucketConfiguration,omitempty" url:"bucketConfiguration,omitempty"` //	description: The S3 specific configuration details for the S3 object that'll contain the blob store., example: , pattern:
	Name                string                            `json:"name,omitempty" url:"name,omitempty"`                               //	description: The name of the S3 blob store., example: s3, pattern:
	SoftQuota           BlobStoreApiSoftQuota             `json:"softQuota,omitempty" url:"softQuota,omitempty"`                     //	description: Settings to control the soft quota., example: , pattern:
	Type                string                            `json:"type,omitempty" url:"type,omitempty"`                               //	description: The blob store type., example: S3, pattern:
}

type ScriptResultXO struct {
	Name   string `json:"name,omitempty" url:"name,omitempty"`     //	description: , example: , pattern:
	Result string `json:"result,omitempty" url:"result,omitempty"` //	description: , example: , pattern:
}

type ScriptXO struct {
	Type    string `json:"type,omitempty" url:"type,omitempty"`       //	description: , example: , pattern:
	Content string `json:"content,omitempty" url:"content,omitempty"` //	description: , example: , pattern:
	Name    string `json:"name,omitempty" url:"name,omitempty"`       //	description: , example: , pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
}

type SimpleApiGroupDeployRepository struct {
	Group   GroupDeployAttributes `json:"group,omitempty" url:"group,omitempty"`     //	description: , example: , pattern:
	Name    string                `json:"name,omitempty" url:"name,omitempty"`       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online  bool                  `json:"online,omitempty" url:"online,omitempty"`   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage StorageAttributes     `json:"storage,omitempty" url:"storage,omitempty"` //	description: , example: , pattern:
}

type SimpleApiGroupRepository struct {
	Group   GroupAttributes   `json:"group,omitempty" url:"group,omitempty"`     //	description: , example: , pattern:
	Name    string            `json:"name,omitempty" url:"name,omitempty"`       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online  bool              `json:"online,omitempty" url:"online,omitempty"`   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage StorageAttributes `json:"storage,omitempty" url:"storage,omitempty"` //	description: , example: , pattern:
	Format  string            `json:"format,omitempty" url:"format,omitempty"`   // 补充 swagger.json 不存在的字段
	Url     string            `json:"url,omitempty" url:"url,omitempty"`         // 补充 swagger.json 不存在的字段
	Type    string            `json:"type,omitempty" url:"type,omitempty"`       // 补充 swagger.json 不存在的字段
}

type SimpleApiHostedRepository struct {
	Name      string                  `json:"name,omitempty" url:"name,omitempty"`           //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online    bool                    `json:"online,omitempty" url:"online,omitempty"`       //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage   HostedStorageAttributes `json:"storage,omitempty" url:"storage,omitempty"`     //	description: , example: , pattern:
	Cleanup   CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`     //	description: , example: , pattern:
	Component ComponentAttributes     `json:"component,omitempty" url:"component,omitempty"` //	description: , example: , pattern:
}

type SimpleApiProxyRepository struct {
	Cleanup         CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`                 //	description: , example: , pattern:
	NegativeCache   NegativeCacheAttributes `json:"negativeCache,omitempty" url:"negativeCache,omitempty"`     //	description: , example: , pattern:
	RoutingRuleName string                  `json:"routingRuleName,omitempty" url:"routingRuleName,omitempty"` //	description: The name of the routing rule assigned to this repository, example: , pattern:
	HttpClient      HttpClientAttributes    `json:"httpClient,omitempty" url:"httpClient,omitempty"`           //	description: , example: , pattern:
	Name            string                  `json:"name,omitempty" url:"name,omitempty"`                       //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online          bool                    `json:"online,omitempty" url:"online,omitempty"`                   //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Proxy           ProxyAttributes         `json:"proxy,omitempty" url:"proxy,omitempty"`                     //	description: , example: , pattern:
	Replication     ReplicationAttributes   `json:"replication,omitempty" url:"replication,omitempty"`         //	description: , example: , pattern:
	Storage         StorageAttributes       `json:"storage,omitempty" url:"storage,omitempty"`                 //	description: , example: , pattern:
}

type StackTraceElement struct {
	ClassName    string `json:"className,omitempty" url:"className,omitempty"`       //	description: , example: , pattern:
	FileName     string `json:"fileName,omitempty" url:"fileName,omitempty"`         //	description: , example: , pattern:
	LineNumber   int64  `json:"lineNumber,omitempty" url:"lineNumber,omitempty"`     //	description: , example: , pattern:
	MethodName   string `json:"methodName,omitempty" url:"methodName,omitempty"`     //	description: , example: , pattern:
	NativeMethod bool   `json:"nativeMethod,omitempty" url:"nativeMethod,omitempty"` //	description: , example: , pattern:
}

type StorageAttributes struct {
	BlobStoreName               string `json:"blobStoreName,omitempty" url:"blobStoreName,omitempty"`                             //	description: Blob store used to store repository contents, example: default, pattern:
	StrictContentTypeValidation bool   `json:"strictContentTypeValidation,omitempty" url:"strictContentTypeValidation,omitempty"` //	description: Whether to validate uploaded content's MIME type appropriate for the repository format, example: true, pattern:
}

type SupportZipGeneratorRequest struct {
	AuditLog          bool   `json:"auditLog,omitempty" url:"auditLog,omitempty"`                   //	description: , example: , pattern:
	LimitFileSizes    bool   `json:"limitFileSizes,omitempty" url:"limitFileSizes,omitempty"`       //	description: , example: , pattern:
	Security          bool   `json:"security,omitempty" url:"security,omitempty"`                   //	description: , example: , pattern:
	Jmx               bool   `json:"jmx,omitempty" url:"jmx,omitempty"`                             //	description: , example: , pattern:
	Metrics           bool   `json:"metrics,omitempty" url:"metrics,omitempty"`                     //	description: , example: , pattern:
	ThreadDump        bool   `json:"threadDump,omitempty" url:"threadDump,omitempty"`               //	description: , example: , pattern:
	Configuration     bool   `json:"configuration,omitempty" url:"configuration,omitempty"`         //	description: , example: , pattern:
	SystemInformation bool   `json:"systemInformation,omitempty" url:"systemInformation,omitempty"` //	description: , example: , pattern:
	Hostname          string `json:"hostname,omitempty" url:"hostname,omitempty"`                   //	description: , example: , pattern:
	LimitZipSize      bool   `json:"limitZipSize,omitempty" url:"limitZipSize,omitempty"`           //	description: , example: , pattern:
	Log               bool   `json:"log,omitempty" url:"log,omitempty"`                             //	description: , example: , pattern:
	Replication       bool   `json:"replication,omitempty" url:"replication,omitempty"`             //	description: , example: , pattern:
	TaskLog           bool   `json:"taskLog,omitempty" url:"taskLog,omitempty"`                     //	description: , example: , pattern:
}

type SupportZipXO struct {
	File      string `json:"file,omitempty" url:"file,omitempty"`           //	description: , example: , pattern:
	Name      string `json:"name,omitempty" url:"name,omitempty"`           //	description: , example: , pattern:
	Size      string `json:"size,omitempty" url:"size,omitempty"`           //	description: , example: , pattern:
	Truncated bool   `json:"truncated,omitempty" url:"truncated,omitempty"` //	description: , example: , pattern:
}

type TaskXO struct {
	LastRunResult string `json:"lastRunResult,omitempty" url:"lastRunResult,omitempty"` //	description: , example: , pattern:
	Message       string `json:"message,omitempty" url:"message,omitempty"`             //	description: , example: , pattern:
	Name          string `json:"name,omitempty" url:"name,omitempty"`                   //	description: , example: , pattern:
	NextRun       string `json:"nextRun,omitempty" url:"nextRun,omitempty"`             //	description: , example: , pattern:
	Type          string `json:"type,omitempty" url:"type,omitempty"`                   //	description: , example: , pattern:
	CurrentState  string `json:"currentState,omitempty" url:"currentState,omitempty"`   //	description: , example: , pattern:
	Id            string `json:"id,omitempty" url:"id,omitempty"`                       //	description: , example: , pattern:
	LastRun       string `json:"lastRun,omitempty" url:"lastRun,omitempty"`             //	description: , example: , pattern:
}

type Throwable struct {
	LocalizedMessage string              `json:"localizedMessage,omitempty" url:"localizedMessage,omitempty"` //	description: , example: , pattern:
	Message          string              `json:"message,omitempty" url:"message,omitempty"`                   //	description: , example: , pattern:
	StackTrace       []StackTraceElement `json:"stackTrace,omitempty" url:"stackTrace,omitempty"`             //	description: , example: , pattern:
	Suppressed       []Throwable         `json:"suppressed,omitempty" url:"suppressed,omitempty"`             //	description: , example: , pattern:
	Cause            *Throwable          `json:"cause,omitempty" url:"cause,omitempty"`                       //	description: , example: , pattern:
}

type UpdateLdapServerXo struct {
	SearchBase                  string `json:"searchBase,omitempty" url:"searchBase,omitempty"`                                   //	description: LDAP location to be added to the connection URL, example: dc=example,dc=com, pattern:
	UserRealNameAttribute       string `json:"userRealNameAttribute,omitempty" url:"userRealNameAttribute,omitempty"`             //	description: This is used to find a real name given the user ID, example: cn, pattern:
	AuthPassword                string `json:"authPassword,omitempty" url:"authPassword,omitempty"`                               //	description: The password to bind with. Required if authScheme other than none., example: , pattern:
	ConnectionRetryDelaySeconds int64  `json:"connectionRetryDelaySeconds,omitempty" url:"connectionRetryDelaySeconds,omitempty"` //	description: How long to wait before retrying, example: , pattern:
	LdapGroupsAsRoles           bool   `json:"ldapGroupsAsRoles,omitempty" url:"ldapGroupsAsRoles,omitempty"`                     //	description: Denotes whether LDAP assigned roles are used as Nexus Repository Manager roles, example: , pattern:
	UserEmailAddressAttribute   string `json:"userEmailAddressAttribute,omitempty" url:"userEmailAddressAttribute,omitempty"`     //	description: This is used to find an email address given the user ID, example: mail, pattern:
	UserIdAttribute             string `json:"userIdAttribute,omitempty" url:"userIdAttribute,omitempty"`                         //	description: This is used to find a user given its user ID, example: uid, pattern:
	UserMemberOfAttribute       string `json:"userMemberOfAttribute,omitempty" url:"userMemberOfAttribute,omitempty"`             //	description: Set this to the attribute used to store the attribute which holds groups DN in the user object. Required if groupType is dynamic, example: memberOf, pattern:
	UserPasswordAttribute       string `json:"userPasswordAttribute,omitempty" url:"userPasswordAttribute,omitempty"`             //	description: If this field is blank the user will be authenticated against a bind with the LDAP server, example: , pattern:
	AuthScheme                  string `json:"authScheme,omitempty" url:"authScheme,omitempty"`                                   //	description: Authentication scheme used for connecting to LDAP server, example: , pattern:
	AuthUsername                string `json:"authUsername,omitempty" url:"authUsername,omitempty"`                               //	description: This must be a fully qualified username if simple authentication is used. Required if authScheme other than none., example: , pattern:
	Name                        string `json:"name,omitempty" url:"name,omitempty"`                                               //	description: LDAP server name, example: , pattern:
	Id                          string `json:"id,omitempty" url:"id,omitempty"`                                                   //	description: LDAP server ID, example: , pattern:
	MaxIncidentsCount           int64  `json:"maxIncidentsCount,omitempty" url:"maxIncidentsCount,omitempty"`                     //	description: How many retry attempts, example: , pattern:
	GroupType                   string `json:"groupType,omitempty" url:"groupType,omitempty"`                                     //	description: Defines a type of groups used: static (a group contains a list of users) or dynamic (a user contains a list of groups). Required if ldapGroupsAsRoles is true., example: , pattern:
	UserLdapFilter              string `json:"userLdapFilter,omitempty" url:"userLdapFilter,omitempty"`                           //	description: LDAP search filter to limit user search, example: (|(mail=*@example.com)(uid=dom*)), pattern:
	AuthRealm                   string `json:"authRealm,omitempty" url:"authRealm,omitempty"`                                     //	description: The SASL realm to bind to. Required if authScheme is CRAM_MD5 or DIGEST_MD5, example: example.com, pattern:
	GroupMemberAttribute        string `json:"groupMemberAttribute,omitempty" url:"groupMemberAttribute,omitempty"`               //	description: LDAP attribute containing the usernames for the group. Required if groupType is static, example: memberUid, pattern:
	GroupSubtree                bool   `json:"groupSubtree,omitempty" url:"groupSubtree,omitempty"`                               //	description: Are groups located in structures below the group base DN, example: , pattern:
	Protocol                    string `json:"protocol,omitempty" url:"protocol,omitempty"`                                       //	description: LDAP server connection Protocol to use, example: , pattern:
	UseTrustStore               bool   `json:"useTrustStore,omitempty" url:"useTrustStore,omitempty"`                             //	description: Whether to use certificates stored in Nexus Repository Manager's truststore, example: , pattern:
	UserObjectClass             string `json:"userObjectClass,omitempty" url:"userObjectClass,omitempty"`                         //	description: LDAP class for user objects, example: inetOrgPerson, pattern:
	GroupMemberFormat           string `json:"groupMemberFormat,omitempty" url:"groupMemberFormat,omitempty"`                     //	description: The format of user ID stored in the group member attribute. Required if groupType is static, example: uid=${username},ou=people,dc=example,dc=com, pattern:
	Port                        int64  `json:"port,omitempty" url:"port,omitempty"`                                               //	description: LDAP server connection port to use, example: 636, pattern:
	UserSubtree                 bool   `json:"userSubtree,omitempty" url:"userSubtree,omitempty"`                                 //	description: Are users located in structures below the user base DN?, example: , pattern:
	ConnectionTimeoutSeconds    int64  `json:"connectionTimeoutSeconds,omitempty" url:"connectionTimeoutSeconds,omitempty"`       //	description: How long to wait before timeout, example: 1, pattern:
	GroupBaseDn                 string `json:"groupBaseDn,omitempty" url:"groupBaseDn,omitempty"`                                 //	description: The relative DN where group objects are found (e.g. ou=Group). This value will have the Search base DN value appended to form the full Group search base DN., example: ou=Group, pattern:
	UserBaseDn                  string `json:"userBaseDn,omitempty" url:"userBaseDn,omitempty"`                                   //	description: The relative DN where user objects are found (e.g. ou=people). This value will have the Search base DN value appended to form the full User search base DN., example: ou=people, pattern:
	GroupIdAttribute            string `json:"groupIdAttribute,omitempty" url:"groupIdAttribute,omitempty"`                       //	description: This field specifies the attribute of the Object class that defines the Group ID. Required if groupType is static, example: cn, pattern:
	GroupObjectClass            string `json:"groupObjectClass,omitempty" url:"groupObjectClass,omitempty"`                       //	description: LDAP class for group objects. Required if groupType is static, example: posixGroup, pattern:
	Host                        string `json:"host,omitempty" url:"host,omitempty"`                                               //	description: LDAP server connection hostname, example: , pattern:
}

type UploadDefinitionXO struct {
	MultipleUpload  bool                      `json:"multipleUpload,omitempty" url:"multipleUpload,omitempty"`   //	description: , example: , pattern:
	AssetFields     []UploadFieldDefinitionXO `json:"assetFields,omitempty" url:"assetFields,omitempty"`         //	description: , example: , pattern:
	ComponentFields []UploadFieldDefinitionXO `json:"componentFields,omitempty" url:"componentFields,omitempty"` //	description: , example: , pattern:
	Format          string                    `json:"format,omitempty" url:"format,omitempty"`                   //	description: , example: , pattern:
}

type UploadFieldDefinitionXO struct {
	Optional    bool   `json:"optional,omitempty" url:"optional,omitempty"`       //	description: , example: , pattern:
	Type        string `json:"type,omitempty" url:"type,omitempty"`               //	description: , example: , pattern:
	Description string `json:"description,omitempty" url:"description,omitempty"` //	description: , example: , pattern:
	Group       string `json:"group,omitempty" url:"group,omitempty"`             //	description: , example: , pattern:
	Name        string `json:"name,omitempty" url:"name,omitempty"`               //	description: , example: , pattern:
}

type YumAttributes struct {
	DeployPolicy  string `json:"deployPolicy,omitempty" url:"deployPolicy,omitempty"`   //	description: Validate that all paths are RPMs or yum metadata, example: STRICT, pattern:
	RepodataDepth int64  `json:"repodataDepth,omitempty" url:"repodataDepth,omitempty"` //	description: Specifies the repository depth where repodata folder(s) are created, example: 5, pattern:
}

type YumGroupRepositoryApiRequest struct {
	Name       string                           `json:"name,omitempty" url:"name,omitempty"`             //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online     bool                             `json:"online,omitempty" url:"online,omitempty"`         //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage    StorageAttributes                `json:"storage,omitempty" url:"storage,omitempty"`       //	description: , example: , pattern:
	YumSigning YumSigningRepositoriesAttributes `json:"yumSigning,omitempty" url:"yumSigning,omitempty"` //	description: , example: , pattern:
	Group      GroupAttributes                  `json:"group,omitempty" url:"group,omitempty"`           //	description: , example: , pattern:
}

type YumHostedApiRepository struct {
	Name      string                  `json:"name,omitempty" url:"name,omitempty"`           //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online    bool                    `json:"online,omitempty" url:"online,omitempty"`       //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage   HostedStorageAttributes `json:"storage,omitempty" url:"storage,omitempty"`     //	description: , example: , pattern:
	Yum       YumAttributes           `json:"yum,omitempty" url:"yum,omitempty"`             //	description: , example: , pattern:
	Cleanup   CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`     //	description: , example: , pattern:
	Component ComponentAttributes     `json:"component,omitempty" url:"component,omitempty"` //	description: , example: , pattern:
}

type YumHostedRepositoryApiRequest struct {
	Cleanup   CleanupPolicyAttributes `json:"cleanup,omitempty" url:"cleanup,omitempty"`     //	description: , example: , pattern:
	Component ComponentAttributes     `json:"component,omitempty" url:"component,omitempty"` //	description: , example: , pattern:
	Name      string                  `json:"name,omitempty" url:"name,omitempty"`           //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Online    bool                    `json:"online,omitempty" url:"online,omitempty"`       //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Storage   HostedStorageAttributes `json:"storage,omitempty" url:"storage,omitempty"`     //	description: , example: , pattern:
	Yum       YumAttributes           `json:"yum,omitempty" url:"yum,omitempty"`             //	description: , example: , pattern:
}

type YumProxyRepositoryApiRequest struct {
	Storage       StorageAttributes                `json:"storage,omitempty" url:"storage,omitempty"`             //	description: , example: , pattern:
	NegativeCache NegativeCacheAttributes          `json:"negativeCache,omitempty" url:"negativeCache,omitempty"` //	description: , example: , pattern:
	Online        bool                             `json:"online,omitempty" url:"online,omitempty"`               //	description: Whether this repository accepts incoming requests, example: true, pattern:
	Proxy         ProxyAttributes                  `json:"proxy,omitempty" url:"proxy,omitempty"`                 //	description: , example: , pattern:
	Replication   ReplicationAttributes            `json:"replication,omitempty" url:"replication,omitempty"`     //	description: , example: , pattern:
	RoutingRule   string                           `json:"routingRule,omitempty" url:"routingRule,omitempty"`     //	description: , example: , pattern:
	Cleanup       CleanupPolicyAttributes          `json:"cleanup,omitempty" url:"cleanup,omitempty"`             //	description: , example: , pattern:
	HttpClient    HttpClientAttributes             `json:"httpClient,omitempty" url:"httpClient,omitempty"`       //	description: , example: , pattern:
	Name          string                           `json:"name,omitempty" url:"name,omitempty"`                   //	description: A unique identifier for this repository, example: internal, pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	YumSigning    YumSigningRepositoriesAttributes `json:"yumSigning,omitempty" url:"yumSigning,omitempty"`       //	description: , example: , pattern:
}

type YumSigningRepositoriesAttributes struct {
	Passphrase string `json:"passphrase,omitempty" url:"passphrase,omitempty"` //	description: Passphrase to access PGP signing key, example: , pattern:
	Keypair    string `json:"keypair,omitempty" url:"keypair,omitempty"`       //	description: PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor), example: , pattern:
}
