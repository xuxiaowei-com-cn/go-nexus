package nexus

import "io"

// Asset 资产实体（Swagger: AssetXO）
// Asset entity (Swagger: AssetXO)
type Asset struct {
	BlobCreated    string            `json:"blobCreated"`
	BlobStoreName  string            `json:"blobStoreName"`
	Checksum       map[string]string `json:"checksum"`
	ContentType    string            `json:"contentType"`
	DownloadURL    string            `json:"downloadUrl"`
	FileSize       int64             `json:"fileSize"`
	Format         string            `json:"format"`
	ID             string            `json:"id"`
	LastDownloaded string            `json:"lastDownloaded"`
	LastModified   string            `json:"lastModified"`
	Path           string            `json:"path"`
	Repository     string            `json:"repository"`
	Uploader       string            `json:"uploader"`
	UploaderIP     string            `json:"uploaderIp"`
}

// AptHostedApiRepository APT 托管仓库返回体（Swagger: AptHostedApiRepository）
// APT hosted repository response (Swagger: AptHostedApiRepository)
// 包含基本信息、存储、APT 与签名属性 / includes basic info, storage, apt and signing attributes
type AptHostedApiRepository struct {
	Apt        AptHostedRepositoriesAttributes  `json:"apt"`
	AptSigning AptSigningRepositoriesAttributes `json:"aptSigning"`
	Format     string                           `json:"format"`
	Name       string                           `json:"name"`
	Online     bool                             `json:"online"`
	Storage    HostedStorageAttributes          `json:"storage"`
	Type       string                           `json:"type"`
	URL        string                           `json:"url"`
}

// AptHostedRepositoriesAttributes APT 托管仓库属性（Swagger: AptHostedRepositoriesAttributes）
// APT hosted repository attributes (Swagger: AptHostedRepositoriesAttributes)
// - distribution：发行版 / distribution（如 bionic、bookworm）
type AptHostedRepositoriesAttributes struct {
	Distribution string `json:"distribution,omitempty"`
}

// AptHostedRepositoryApiRequest 创建/更新 APT 托管仓库请求体（Swagger: AptHostedRepositoryApiRequest）
// Request payload to create/update APT hosted repository (Swagger: AptHostedRepositoryApiRequest)
// 字段：name/online/storage/apt/aptSigning
type AptHostedRepositoryApiRequest struct {
	Apt        AptHostedRepositoriesAttributes  `json:"apt"`
	AptSigning AptSigningRepositoriesAttributes `json:"aptSigning"`
	Name       string                           `json:"name"`
	Online     bool                             `json:"online"`
	Storage    HostedStorageAttributes          `json:"storage"`
}

// AptProxyApiRepository APT 代理仓库返回体（Swagger: AptProxyApiRepository）
// APT proxy repository response (Swagger: AptProxyApiRepository)
type AptProxyApiRepository struct {
	Apt           AptProxyRepositoriesAttributes `json:"apt"`
	Format        string                         `json:"format"`
	HttpClient    HttpClientAttributes           `json:"httpClient"`
	Name          string                         `json:"name"`
	NegativeCache NegativeCacheAttributes        `json:"negativeCache"`
	Online        bool                           `json:"online"`
	Proxy         ProxyAttributes                `json:"proxy"`
	Storage       StorageAttributes              `json:"storage"`
	Type          string                         `json:"type"`
	URL           string                         `json:"url"`
}

// AptProxyRepositoriesAttributes APT 代理仓库属性（Swagger: AptProxyRepositoriesAttributes）
// APT proxy repository attributes (Swagger: AptProxyRepositoriesAttributes)
type AptProxyRepositoriesAttributes struct {
	Flat         bool   `json:"flat"`
	Distribution string `json:"distribution,omitempty"`
}

// AptProxyRepositoryApiRequest 创建/更新 APT 代理仓库请求体（Swagger: AptProxyRepositoryApiRequest）
// Request payload to create/update APT proxy repository (Swagger: AptProxyRepositoryApiRequest)
type AptProxyRepositoryApiRequest struct {
	Apt           AptProxyRepositoriesAttributes `json:"apt"`
	HttpClient    HttpClientAttributes           `json:"httpClient"`
	Name          string                         `json:"name"`
	NegativeCache NegativeCacheAttributes        `json:"negativeCache"`
	Online        bool                           `json:"online"`
	Proxy         ProxyAttributes                `json:"proxy"`
	Storage       StorageAttributes              `json:"storage"`
}

// AptSigningRepositoriesAttributes APT 签名属性（Swagger: AptSigningRepositoriesAttributes）
// APT signing attributes (Swagger: AptSigningRepositoriesAttributes)
// - keypair：PGP 私钥（ASCII armored）/ PGP signing keypair
// - passphrase：私钥口令 / key passphrase
type AptSigningRepositoriesAttributes struct {
	Keypair    string `json:"keypair,omitempty"`
	Passphrase string `json:"passphrase,omitempty"`
}

// GroupAttributes 组仓库属性（Swagger: GroupAttributes）
// Group repository attributes (Swagger: GroupAttributes)
// - memberNames：成员仓库名称列表 / member repositories' names
type GroupAttributes struct {
	MemberNames []string `json:"memberNames"`
}

// HostedStorageAttributes 托管仓库存储属性（Swagger: HostedStorageAttributes）
// Hosted repository storage attributes (Swagger: HostedStorageAttributes)
// 字段映射/Field mapping：
// - blobStoreName：Blob 存储名称 / blob store name
// - strictContentTypeValidation：严格内容类型校验 / strict content type validation
// - writePolicy：写入策略（allow/allow_once/deny）/ write policy (allow/allow_once/deny)
type HostedStorageAttributes struct {
	BlobStoreName               string `json:"blobStoreName"`
	StrictContentTypeValidation bool   `json:"strictContentTypeValidation"`
	WritePolicy                 string `json:"writePolicy"`
}

// HttpClientAttributes HTTP 客户端属性（Swagger: HttpClientAttributes）
// HTTP client attributes (Swagger: HttpClientAttributes)
// 包含连接与认证属性（非预认证）/ Includes connection and authentication attributes (non-preemptive)
type HttpClientAttributes struct {
	Authentication *HttpClientConnectionAuthenticationAttributes `json:"authentication,omitempty"`
	AutoBlock      bool                                          `json:"autoBlock"`
	Blocked        bool                                          `json:"blocked"`
	Connection     *HttpClientConnectionAttributes               `json:"connection,omitempty"`
}

// HttpClientAttributesWithPreemptiveAuth HTTP 客户端属性（Swagger: HttpClientAttributesWithPreemptiveAuth）
// HTTP client attributes (Swagger: HttpClientAttributesWithPreemptiveAuth)
// 包含连接与认证子属性 / includes connection and authentication sub-attributes
type HttpClientAttributesWithPreemptiveAuth struct {
	Authentication *HttpClientConnectionAuthenticationAttributesWithPreemptive `json:"authentication,omitempty"`
	AutoBlock      bool                                                        `json:"autoBlock"`
	Blocked        bool                                                        `json:"blocked"`
	Connection     *HttpClientConnectionAttributes                             `json:"connection,omitempty"`
}

// HttpClientConnectionAuthenticationAttributes HTTP 客户端认证属性（非预认证）（Swagger: HttpClientConnectionAuthenticationAttributes）
// HTTP client authentication attributes (non-preemptive) (Swagger: HttpClientConnectionAuthenticationAttributes)
// 支持类型：username、ntlm、bearerToken / Supported types: username, ntlm, bearerToken
type HttpClientConnectionAuthenticationAttributes struct {
	BearerToken string `json:"bearerToken,omitempty"`
	NTLMDomain  string `json:"ntlmDomain,omitempty"`
	NTLMHost    string `json:"ntlmHost,omitempty"`
	Password    string `json:"password,omitempty"`
	Type        string `json:"type,omitempty"`
	Username    string `json:"username,omitempty"`
}

// HttpClientConnectionAuthenticationAttributesWithPreemptive 预认证设置（Swagger: HttpClientConnectionAuthenticationAttributesWithPreemptive）
// Preemptive authentication settings (Swagger: HttpClientConnectionAuthenticationAttributesWithPreemptive)
// 支持类型：username/ntlm/bearerToken；可选 preemptive / supported types and preemptive flag
type HttpClientConnectionAuthenticationAttributesWithPreemptive struct {
	BearerToken string `json:"bearerToken,omitempty"`
	NTLMDomain  string `json:"ntlmDomain,omitempty"`
	NTLMHost    string `json:"ntlmHost,omitempty"`
	Password    string `json:"password,omitempty"`
	Preemptive  bool   `json:"preemptive,omitempty"`
	Type        string `json:"type,omitempty"`
	Username    string `json:"username,omitempty"`
}

// HttpClientConnectionAttributes HTTP 客户端连接属性（Swagger: HttpClientConnectionAttributes）
// HTTP client connection attributes (Swagger: HttpClientConnectionAttributes)
// 示例：retries/timeout/enableCookies/useTrustStore 等 / e.g., retries, timeout, cookies, trust store
type HttpClientConnectionAttributes struct {
	EnableCircularRedirects bool   `json:"enableCircularRedirects,omitempty"`
	EnableCookies           bool   `json:"enableCookies,omitempty"`
	Retries                 int    `json:"retries,omitempty"`
	Timeout                 int    `json:"timeout,omitempty"`
	UseTrustStore           bool   `json:"useTrustStore,omitempty"`
	UserAgentSuffix         string `json:"userAgentSuffix,omitempty"`
}

// MavenAttributes Maven 属性（Swagger: MavenAttributes）
// Maven attributes (Swagger: MavenAttributes)
// 字段映射/Field mapping：
// - versionPolicy：版本策略（RELEASE/SNAPSHOT/MIXED）/ version policy (RELEASE/SNAPSHOT/MIXED)
// - layoutPolicy：布局策略（STRICT/PERMISSIVE）/ layout policy (STRICT/PERMISSIVE)
// - contentDisposition：内容处置（INLINE/ATTACHMENT）/ content disposition (INLINE/ATTACHMENT)
type MavenAttributes struct {
	ContentDisposition string `json:"contentDisposition"`
	LayoutPolicy       string `json:"layoutPolicy"`
	VersionPolicy      string `json:"versionPolicy"`
}

// MavenGroupRepositoryApiRequest 创建/更新 Maven 组仓库的请求体（Swagger: MavenGroupRepositoryApiRequest）
// Request payload to create/update Maven group repository (Swagger: MavenGroupRepositoryApiRequest)
// 字段：name/online/storage/group（含 memberNames）/ fields: name/online/storage/group (with memberNames)
type MavenGroupRepositoryApiRequest struct {
	Group   GroupAttributes   `json:"group"`
	Name    string            `json:"name"`
	Online  bool              `json:"online"`
	Storage StorageAttributes `json:"storage"`
}

// MavenHostedApiRepository Maven 托管仓库返回体（Swagger: MavenHostedApiRepository）
// Maven hosted repository response (Swagger: MavenHostedApiRepository)
// 字段映射/Field mapping：
// - name/format/type/url/online：仓库基本信息 / repository basic info
// - storage：托管存储属性 / hosted storage attributes
// - maven：Maven 属性 / maven attributes
type MavenHostedApiRepository struct {
	Format  string                  `json:"format"`
	Maven   MavenAttributes         `json:"maven"`
	Name    string                  `json:"name"`
	Online  bool                    `json:"online"`
	Storage HostedStorageAttributes `json:"storage"`
	Type    string                  `json:"type"`
	URL     string                  `json:"url"`
}

// MavenHostedRepositoryApiRequest 创建/更新 Maven 托管仓库的请求体（Swagger: MavenHostedRepositoryApiRequest）
// Request payload to create/update Maven hosted repository (Swagger: MavenHostedRepositoryApiRequest)
// 字段映射/Field mapping：
// - name：仓库名 / repository name
// - online：是否在线 / whether repository accepts requests
// - storage：托管存储属性 / hosted storage attributes
// - maven：Maven 属性 / maven attributes
type MavenHostedRepositoryApiRequest struct {
	Maven   MavenAttributes         `json:"maven"`
	Name    string                  `json:"name"`
	Online  bool                    `json:"online"`
	Storage HostedStorageAttributes `json:"storage"`
}

// MavenProxyApiRepository Maven 代理仓库返回体（Swagger: MavenProxyApiRepository）
// Maven proxy repository response (Swagger: MavenProxyApiRepository)
// 包含基本信息与存储/代理/负缓存/Maven 属性 / includes basic info and attributes
type MavenProxyApiRepository struct {
	Format        string                  `json:"format"`
	Maven         MavenAttributes         `json:"maven"`
	Name          string                  `json:"name"`
	NegativeCache NegativeCacheAttributes `json:"negativeCache"`
	Online        bool                    `json:"online"`
	Proxy         ProxyAttributes         `json:"proxy"`
	Storage       StorageAttributes       `json:"storage"`
	Type          string                  `json:"type"`
	URL           string                  `json:"url"`
}

// MavenProxyRepositoryApiRequest 创建/更新 Maven 代理仓库的请求体（Swagger: MavenProxyRepositoryApiRequest）
// Request payload to create/update Maven proxy repository (Swagger: MavenProxyRepositoryApiRequest)
// 关键字段：storage/proxy/negativeCache/httpClient/maven / key fields
type MavenProxyRepositoryApiRequest struct {
	HttpClient    HttpClientAttributesWithPreemptiveAuth `json:"httpClient"`
	Maven         MavenAttributes                        `json:"maven"`
	Name          string                                 `json:"name"`
	NegativeCache NegativeCacheAttributes                `json:"negativeCache"`
	Online        bool                                   `json:"online"`
	Proxy         ProxyAttributes                        `json:"proxy"`
	Storage       StorageAttributes                      `json:"storage"`
}

// NegativeCacheAttributes 负缓存属性（Swagger: NegativeCacheAttributes）
// Negative cache attributes (Swagger: NegativeCacheAttributes)
// - enabled：启用负缓存 / whether enabled
// - timeToLive：负缓存生存时间（分钟）/ TTL in minutes
type NegativeCacheAttributes struct {
	Enabled    bool `json:"enabled"`
	TimeToLive int  `json:"timeToLive"`
}

// PageAsset 资产分页（Swagger: PageAssetXO）
// Asset page (Swagger: PageAssetXO)
type PageAsset struct {
	ContinuationToken string  `json:"continuationToken"`
	Items             []Asset `json:"items"`
}

// ProxyAttributes 代理属性（Swagger: ProxyAttributes）
// Proxy attributes (Swagger: ProxyAttributes)
// - remoteUrl：远端仓库地址 / remote repository URL
// - contentMaxAge：内容最大缓存分钟数 / content cache max age (minutes)
// - metadataMaxAge：元数据最大缓存分钟数 / metadata cache max age (minutes)
type ProxyAttributes struct {
	ContentMaxAge  int    `json:"contentMaxAge"`
	MetadataMaxAge int    `json:"metadataMaxAge"`
	RemoteURL      string `json:"remoteUrl"`
}

// RawAttributes Raw 仓库属性（Swagger: RawAttributes）
// Raw repository attributes (Swagger: RawAttributes)
// - contentDisposition：内容处置（INLINE/ATTACHMENT）/ content disposition (INLINE/ATTACHMENT)
type RawAttributes struct {
	ContentDisposition string `json:"contentDisposition"`
}

// RawHostedRepositoryApiRequest 创建/更新 Raw 托管仓库请求体（Swagger: RawHostedRepositoryApiRequest）
// Request payload to create/update raw hosted repository (Swagger: RawHostedRepositoryApiRequest)
// 字段：name/online/storage/raw
type RawHostedRepositoryApiRequest struct {
	Name    string                  `json:"name"`
	Online  bool                    `json:"online"`
	Raw     RawAttributes           `json:"raw"`
	Storage HostedStorageAttributes `json:"storage"`
}

// CleanupPolicyAttributes 清理策略属性（Swagger: CleanupPolicyAttributes）
// Cleanup policy attributes (Swagger: CleanupPolicyAttributes)
type CleanupPolicyAttributes struct {
	PolicyNames []string `json:"policyNames,omitempty"`
}

// ComponentAttributes 组件属性（Swagger: ComponentAttributes）
// Component attributes (Swagger: ComponentAttributes)
type ComponentAttributes struct {
	ProprietaryComponents bool `json:"proprietaryComponents,omitempty"`
}

// CargoAttributes Cargo 属性（Swagger: CargoAttributes）
// Cargo attributes (Swagger: CargoAttributes)
type CargoAttributes struct {
	RequireAuthentication bool `json:"requireAuthentication"`
}

// CargoGroupRepositoryApiRequest 创建/更新 Cargo 组仓库请求体（Swagger: CargoGroupRepositoryApiRequest）
// Request payload to create/update Cargo group repository (Swagger: CargoGroupRepositoryApiRequest)
type CargoGroupRepositoryApiRequest struct {
	Cargo   CargoAttributes   `json:"cargo"`
	Group   GroupAttributes   `json:"group"`
	Name    string            `json:"name"`
	Online  bool              `json:"online"`
	Storage StorageAttributes `json:"storage"`
}

// CargoHostedRepositoryApiRequest 创建/更新 Cargo 托管仓库请求体（Swagger: CargoHostedRepositoryApiRequest）
// Request payload to create/update Cargo hosted repository (Swagger: CargoHostedRepositoryApiRequest)
type CargoHostedRepositoryApiRequest struct {
	Cleanup   *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes     `json:"component,omitempty"`
	Name      string                   `json:"name"`
	Online    bool                     `json:"online"`
	Storage   HostedStorageAttributes  `json:"storage"`
}

// CargoProxyRepositoryApiRequest 创建/更新 Cargo 代理仓库请求体（Swagger: CargoProxyRepositoryApiRequest）
// Request payload to create/update Cargo proxy repository (Swagger: CargoProxyRepositoryApiRequest)
type CargoProxyRepositoryApiRequest struct {
	Cargo         CargoAttributes          `json:"cargo"`
	Cleanup       *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	HttpClient    HttpClientAttributes     `json:"httpClient"`
	Name          string                   `json:"name"`
	NegativeCache NegativeCacheAttributes  `json:"negativeCache"`
	Online        bool                     `json:"online"`
	Proxy         ProxyAttributes          `json:"proxy"`
	RoutingRule   string                   `json:"routingRule,omitempty"`
	Storage       StorageAttributes        `json:"storage"`
}

// CargoGroupApiRepository Cargo 组仓库返回体（Swagger: CargoGroupApiRepository）
// Cargo group repository response (Swagger: CargoGroupApiRepository)
type CargoGroupApiRepository struct {
	Cargo   CargoAttributes   `json:"cargo"`
	Format  string            `json:"format"`
	Group   GroupAttributes   `json:"group"`
	Name    string            `json:"name"`
	Online  bool              `json:"online"`
	Storage StorageAttributes `json:"storage"`
	Type    string            `json:"type"`
	URL     string            `json:"url"`
}

// CargoProxyApiRepository Cargo 代理仓库返回体（Swagger: CargoProxyApiRepository）
// Cargo proxy repository response (Swagger: CargoProxyApiRepository)
type CargoProxyApiRepository struct {
	Cargo           CargoAttributes          `json:"cargo"`
	Cleanup         *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Format          string                   `json:"format"`
	HttpClient      HttpClientAttributes     `json:"httpClient"`
	Name            string                   `json:"name"`
	NegativeCache   NegativeCacheAttributes  `json:"negativeCache"`
	Online          bool                     `json:"online"`
	Proxy           ProxyAttributes          `json:"proxy"`
	RoutingRuleName string                   `json:"routingRuleName,omitempty"`
	Storage         StorageAttributes        `json:"storage"`
	Type            string                   `json:"type"`
	URL             string                   `json:"url"`
}

// SimpleApiProxyRepository 通用代理仓库返回体（Swagger: SimpleApiProxyRepository）
// Generic proxy repository response (Swagger: SimpleApiProxyRepository)
type SimpleApiProxyRepository struct {
	Cleanup         *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Format          string                   `json:"format"`
	HttpClient      HttpClientAttributes     `json:"httpClient"`
	Name            string                   `json:"name"`
	NegativeCache   NegativeCacheAttributes  `json:"negativeCache"`
	Online          bool                     `json:"online"`
	Proxy           ProxyAttributes          `json:"proxy"`
	RoutingRuleName string                   `json:"routingRuleName,omitempty"`
	Storage         StorageAttributes        `json:"storage"`
	Type            string                   `json:"type"`
	URL             string                   `json:"url"`
}

// CocoapodsProxyRepositoryApiRequest 创建/更新 CocoaPods 代理仓库请求体（Swagger: CocoapodsProxyRepositoryApiRequest）
// Request payload to create/update CocoaPods proxy repository (Swagger: CocoapodsProxyRepositoryApiRequest)
type CocoapodsProxyRepositoryApiRequest struct {
	Cleanup       *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	HttpClient    HttpClientAttributes     `json:"httpClient"`
	Name          string                   `json:"name"`
	NegativeCache NegativeCacheAttributes  `json:"negativeCache"`
	Online        bool                     `json:"online"`
	Proxy         ProxyAttributes          `json:"proxy"`
	RoutingRule   string                   `json:"routingRule,omitempty"`
	Storage       StorageAttributes        `json:"storage"`
}

// ComposerProxyRepositoryApiRequest 创建/更新 Composer 代理仓库请求体（Swagger: ComposerProxyRepositoryApiRequest）
// Request payload to create/update Composer proxy repository (Swagger: ComposerProxyRepositoryApiRequest)
type ComposerProxyRepositoryApiRequest struct {
	Cleanup       *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	HttpClient    HttpClientAttributes     `json:"httpClient"`
	Name          string                   `json:"name"`
	NegativeCache NegativeCacheAttributes  `json:"negativeCache"`
	Online        bool                     `json:"online"`
	Proxy         ProxyAttributes          `json:"proxy"`
	RoutingRule   string                   `json:"routingRule,omitempty"`
	Storage       StorageAttributes        `json:"storage"`
}

// GolangGroupRepositoryApiRequest 创建/更新 Go 组仓库请求体（Swagger: GolangGroupRepositoryApiRequest）
// Request payload to create/update Go group repository (Swagger: GolangGroupRepositoryApiRequest)
type GolangGroupRepositoryApiRequest struct {
	Group   GroupAttributes   `json:"group"`
	Name    string            `json:"name"`
	Online  bool              `json:"online"`
	Storage StorageAttributes `json:"storage"`
}

// GolangProxyRepositoryApiRequest 创建/更新 Go 代理仓库请求体（Swagger: GolangProxyRepositoryApiRequest）
// Request payload to create/update Go proxy repository (Swagger: GolangProxyRepositoryApiRequest)
type GolangProxyRepositoryApiRequest struct {
	Cleanup       *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	HttpClient    HttpClientAttributes     `json:"httpClient"`
	Name          string                   `json:"name"`
	NegativeCache NegativeCacheAttributes  `json:"negativeCache"`
	Online        bool                     `json:"online"`
	Proxy         ProxyAttributes          `json:"proxy"`
	RoutingRule   string                   `json:"routingRule,omitempty"`
	Storage       StorageAttributes        `json:"storage"`
}

// GitLfsHostedRepositoryApiRequest 创建/更新 Git LFS 托管仓库请求体（Swagger: GitLfsHostedRepositoryApiRequest）
// Request payload to create/update Git LFS hosted repository (Swagger: GitLfsHostedRepositoryApiRequest)
type GitLfsHostedRepositoryApiRequest struct {
	Cleanup   *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes     `json:"component,omitempty"`
	Name      string                   `json:"name"`
	Online    bool                     `json:"online"`
	Storage   HostedStorageAttributes  `json:"storage"`
}

// HelmHostedRepositoryApiRequest 创建/更新 Helm 托管仓库请求体（Swagger: HelmHostedRepositoryApiRequest）
// Request payload to create/update Helm hosted repository (Swagger: HelmHostedRepositoryApiRequest)
type HelmHostedRepositoryApiRequest struct {
	Cleanup   *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes     `json:"component,omitempty"`
	Name      string                   `json:"name"`
	Online    bool                     `json:"online"`
	Storage   HostedStorageAttributes  `json:"storage"`
}

// HelmProxyRepositoryApiRequest 创建/更新 Helm 代理仓库请求体（Swagger: HelmProxyRepositoryApiRequest）
// Request payload to create/update Helm proxy repository (Swagger: HelmProxyRepositoryApiRequest)
type HelmProxyRepositoryApiRequest struct {
	Cleanup       *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	HttpClient    HttpClientAttributes     `json:"httpClient"`
	Name          string                   `json:"name"`
	NegativeCache NegativeCacheAttributes  `json:"negativeCache"`
	Online        bool                     `json:"online"`
	Proxy         ProxyAttributes          `json:"proxy"`
	RoutingRule   string                   `json:"routingRule,omitempty"`
	Storage       StorageAttributes        `json:"storage"`
}

// HuggingFaceProxyRepositoryApiRequest 创建/更新 HuggingFace 代理仓库请求体（Swagger: HuggingFaceProxyRepositoryApiRequest）
// Request payload to create/update HuggingFace proxy repository (Swagger: HuggingFaceProxyRepositoryApiRequest)
type HuggingFaceProxyRepositoryApiRequest struct {
	Cleanup       *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	HttpClient    HttpClientAttributes     `json:"httpClient"`
	Name          string                   `json:"name"`
	NegativeCache NegativeCacheAttributes  `json:"negativeCache"`
	Online        bool                     `json:"online"`
	Proxy         ProxyAttributes          `json:"proxy"`
	RoutingRule   string                   `json:"routingRule,omitempty"`
	Storage       StorageAttributes        `json:"storage"`
}

// NpmAttributes NPM 属性（Swagger: NpmAttributes）
// NPM attributes (Swagger: NpmAttributes)
type NpmAttributes struct {
	RemoveQuarantined bool `json:"removeQuarantined"`
}

// NpmGroupRepositoryApiRequest 创建/更新 NPM 组仓库请求体（Swagger: NpmGroupRepositoryApiRequest）
// Request payload to create/update NPM group repository (Swagger: NpmGroupRepositoryApiRequest)
type NpmGroupRepositoryApiRequest struct {
	Group   GroupDeployAttributes `json:"group"`
	Name    string                `json:"name"`
	Online  bool                  `json:"online"`
	Storage StorageAttributes     `json:"storage"`
}

// NpmHostedRepositoryApiRequest 创建/更新 NPM 托管仓库请求体（Swagger: NpmHostedRepositoryApiRequest）
// Request payload to create/update NPM hosted repository (Swagger: NpmHostedRepositoryApiRequest)
type NpmHostedRepositoryApiRequest struct {
	Cleanup   *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes     `json:"component,omitempty"`
	Name      string                   `json:"name"`
	Online    bool                     `json:"online"`
	Storage   HostedStorageAttributes  `json:"storage"`
}

// NpmProxyRepositoryApiRequest 创建/更新 NPM 代理仓库请求体（Swagger: NpmProxyRepositoryApiRequest）
// Request payload to create/update NPM proxy repository (Swagger: NpmProxyRepositoryApiRequest)
type NpmProxyRepositoryApiRequest struct {
	Cleanup       *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	HttpClient    HttpClientAttributes     `json:"httpClient"`
	Name          string                   `json:"name"`
	NegativeCache NegativeCacheAttributes  `json:"negativeCache"`
	Npm           *NpmAttributes           `json:"npm,omitempty"`
	Online        bool                     `json:"online"`
	Proxy         ProxyAttributes          `json:"proxy"`
	RoutingRule   string                   `json:"routingRule,omitempty"`
	Storage       StorageAttributes        `json:"storage"`
}

// NpmProxyApiRepository NPM 代理仓库返回体（Swagger: NpmProxyApiRepository）
// NPM proxy repository response (Swagger: NpmProxyApiRepository)
type NpmProxyApiRepository struct {
	Cleanup         *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Format          string                   `json:"format"`
	HttpClient      HttpClientAttributes     `json:"httpClient"`
	Name            string                   `json:"name"`
	NegativeCache   NegativeCacheAttributes  `json:"negativeCache"`
	Npm             NpmAttributes            `json:"npm"`
	Online          bool                     `json:"online"`
	Proxy           ProxyAttributes          `json:"proxy"`
	RoutingRuleName string                   `json:"routingRuleName,omitempty"`
	Storage         StorageAttributes        `json:"storage"`
	Type            string                   `json:"type"`
	URL             string                   `json:"url"`
}

// NugetAttributes NuGet 属性（Swagger: NugetAttributes）
// NuGet attributes (Swagger: NugetAttributes)
type NugetAttributes struct {
	QueryCacheItemMaxAge int    `json:"queryCacheItemMaxAge,omitempty"`
	NugetVersion         string `json:"nugetVersion,omitempty"`
}

// NugetGroupRepositoryApiRequest 创建/更新 NuGet 组仓库请求体（Swagger: NugetGroupRepositoryApiRequest）
// Request payload to create/update NuGet group repository (Swagger: NugetGroupRepositoryApiRequest)
type NugetGroupRepositoryApiRequest struct {
	Group   GroupAttributes   `json:"group"`
	Name    string            `json:"name"`
	Online  bool              `json:"online"`
	Storage StorageAttributes `json:"storage"`
}

// NugetHostedRepositoryApiRequest 创建/更新 NuGet 托管仓库请求体（Swagger: NugetHostedRepositoryApiRequest）
// Request payload to create/update NuGet hosted repository (Swagger: NugetHostedRepositoryApiRequest)
type NugetHostedRepositoryApiRequest struct {
	Cleanup   *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes     `json:"component,omitempty"`
	Name      string                   `json:"name"`
	Online    bool                     `json:"online"`
	Storage   HostedStorageAttributes  `json:"storage"`
}

// NugetProxyRepositoryApiRequest 创建/更新 NuGet 代理仓库请求体（Swagger: NugetProxyRepositoryApiRequest）
// Request payload to create/update NuGet proxy repository (Swagger: NugetProxyRepositoryApiRequest)
type NugetProxyRepositoryApiRequest struct {
	Cleanup       *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	HttpClient    HttpClientAttributes     `json:"httpClient"`
	Name          string                   `json:"name"`
	NegativeCache NegativeCacheAttributes  `json:"negativeCache"`
	NugetProxy    NugetAttributes          `json:"nugetProxy"`
	Online        bool                     `json:"online"`
	Proxy         ProxyAttributes          `json:"proxy"`
	RoutingRule   string                   `json:"routingRule,omitempty"`
	Storage       StorageAttributes        `json:"storage"`
}

// NugetProxyApiRepository NuGet 代理仓库返回体（Swagger: NugetProxyApiRepository）
// NuGet proxy repository response (Swagger: NugetProxyApiRepository)
type NugetProxyApiRepository struct {
	Cleanup         *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Format          string                   `json:"format"`
	HttpClient      HttpClientAttributes     `json:"httpClient"`
	Name            string                   `json:"name"`
	NegativeCache   NegativeCacheAttributes  `json:"negativeCache"`
	NugetProxy      NugetAttributes          `json:"nugetProxy"`
	Online          bool                     `json:"online"`
	Proxy           ProxyAttributes          `json:"proxy"`
	RoutingRuleName string                   `json:"routingRuleName,omitempty"`
	Storage         StorageAttributes        `json:"storage"`
	Type            string                   `json:"type"`
	URL             string                   `json:"url"`
}

// PyPiProxyAttributes PyPI 代理属性（Swagger: PyPiProxyAttributes）
// PyPI proxy attributes (Swagger: PyPiProxyAttributes)
type PyPiProxyAttributes struct {
	RemoveQuarantined bool `json:"removeQuarantined"`
}

// PypiHostedRepositoryApiRequest 创建/更新 PyPI 托管仓库请求体（Swagger: PypiHostedRepositoryApiRequest）
// Request payload to create/update PyPI hosted repository (Swagger: PypiHostedRepositoryApiRequest)
type PypiHostedRepositoryApiRequest struct {
	Cleanup   *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes     `json:"component,omitempty"`
	Name      string                   `json:"name"`
	Online    bool                     `json:"online"`
	Storage   HostedStorageAttributes  `json:"storage"`
}

// PypiGroupRepositoryApiRequest 创建/更新 PyPI 组仓库请求体（Swagger: PypiGroupRepositoryApiRequest）
// Request payload to create/update PyPI group repository (Swagger: PypiGroupRepositoryApiRequest)
type PypiGroupRepositoryApiRequest struct {
	Group   GroupDeployAttributes   `json:"group"`
	Name    string                  `json:"name"`
	Online  bool                    `json:"online"`
	Storage HostedStorageAttributes `json:"storage"`
}

// PypiProxyRepositoryApiRequest 创建/更新 PyPI 代理仓库请求体（Swagger: PypiProxyRepositoryApiRequest）
// Request payload to create/update PyPI proxy repository (Swagger: PypiProxyRepositoryApiRequest)
type PypiProxyRepositoryApiRequest struct {
	Cleanup       *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	HttpClient    HttpClientAttributes     `json:"httpClient"`
	Name          string                   `json:"name"`
	NegativeCache NegativeCacheAttributes  `json:"negativeCache"`
	Online        bool                     `json:"online"`
	Proxy         ProxyAttributes          `json:"proxy"`
	Pypi          PyPiProxyAttributes      `json:"pypi"`
	RoutingRule   string                   `json:"routingRule,omitempty"`
	Storage       StorageAttributes        `json:"storage"`
}

// RGroupRepositoryApiRequest 创建/更新 R 组仓库请求体（Swagger: RGroupRepositoryApiRequest）
// Request payload to create/update R group repository (Swagger: RGroupRepositoryApiRequest)
type RGroupRepositoryApiRequest struct {
	Group   GroupAttributes   `json:"group"`
	Name    string            `json:"name"`
	Online  bool              `json:"online"`
	Storage StorageAttributes `json:"storage"`
}

// RHostedRepositoryApiRequest 创建/更新 R 托管仓库请求体（Swagger: RHostedRepositoryApiRequest）
// Request payload to create/update R hosted repository (Swagger: RHostedRepositoryApiRequest)
type RHostedRepositoryApiRequest struct {
	Cleanup   *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes     `json:"component,omitempty"`
	Name      string                   `json:"name"`
	Online    bool                     `json:"online"`
	Storage   HostedStorageAttributes  `json:"storage"`
}

// RProxyRepositoryApiRequest 创建/更新 R 代理仓库请求体（Swagger: RProxyRepositoryApiRequest）
// Request payload to create/update R proxy repository (Swagger: RProxyRepositoryApiRequest)
type RProxyRepositoryApiRequest struct {
	Cleanup       *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	HttpClient    HttpClientAttributes     `json:"httpClient"`
	Name          string                   `json:"name"`
	NegativeCache NegativeCacheAttributes  `json:"negativeCache"`
	Online        bool                     `json:"online"`
	Proxy         ProxyAttributes          `json:"proxy"`
	RoutingRule   string                   `json:"routingRule,omitempty"`
	Storage       StorageAttributes        `json:"storage"`
}

// RubyGemsGroupRepositoryApiRequest 创建/更新 RubyGems 组仓库请求体（Swagger: RubyGemsGroupRepositoryApiRequest）
// Request payload to create/update RubyGems group repository (Swagger: RubyGemsGroupRepositoryApiRequest)
type RubyGemsGroupRepositoryApiRequest struct {
	Group   GroupAttributes   `json:"group"`
	Name    string            `json:"name"`
	Online  bool              `json:"online"`
	Storage StorageAttributes `json:"storage"`
}

// RubyGemsHostedRepositoryApiRequest 创建/更新 RubyGems 托管仓库请求体（Swagger: RubyGemsHostedRepositoryApiRequest）
// Request payload to create/update RubyGems hosted repository (Swagger: RubyGemsHostedRepositoryApiRequest)
type RubyGemsHostedRepositoryApiRequest struct {
	Cleanup   *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes     `json:"component,omitempty"`
	Name      string                   `json:"name"`
	Online    bool                     `json:"online"`
	Storage   HostedStorageAttributes  `json:"storage"`
}

// RubyGemsProxyRepositoryApiRequest 创建/更新 RubyGems 代理仓库请求体（Swagger: RubyGemsProxyRepositoryApiRequest）
// Request payload to create/update RubyGems proxy repository (Swagger: RubyGemsProxyRepositoryApiRequest)
type RubyGemsProxyRepositoryApiRequest struct {
	Cleanup       *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	HttpClient    HttpClientAttributes     `json:"httpClient"`
	Name          string                   `json:"name"`
	NegativeCache NegativeCacheAttributes  `json:"negativeCache"`
	Online        bool                     `json:"online"`
	Proxy         ProxyAttributes          `json:"proxy"`
	RoutingRule   string                   `json:"routingRule,omitempty"`
	Storage       StorageAttributes        `json:"storage"`
}

// YumSigningRepositoriesAttributes Yum 签名属性（Swagger: YumSigningRepositoriesAttributes）
// Yum signing attributes (Swagger: YumSigningRepositoriesAttributes)
type YumSigningRepositoriesAttributes struct {
	Keypair    string `json:"keypair,omitempty"`
	Passphrase string `json:"passphrase,omitempty"`
}

// YumAttributes Yum 属性（Swagger: YumAttributes）
// Yum attributes (Swagger: YumAttributes)
type YumAttributes struct {
	RepodataDepth int    `json:"repodataDepth"`
	DeployPolicy  string `json:"deployPolicy,omitempty"`
}

// YumHostedApiRepository Yum 托管仓库返回体（Swagger: YumHostedApiRepository）
// Yum hosted repository response (Swagger: YumHostedApiRepository)
type YumHostedApiRepository struct {
	Cleanup   *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes     `json:"component,omitempty"`
	Format    string                   `json:"format"`
	Name      string                   `json:"name"`
	Online    bool                     `json:"online"`
	Storage   HostedStorageAttributes  `json:"storage"`
	Type      string                   `json:"type"`
	URL       string                   `json:"url"`
	Yum       YumAttributes            `json:"yum"`
}

// YumHostedRepositoryApiRequest 创建/更新 Yum 托管仓库请求体（Swagger: YumHostedRepositoryApiRequest）
// Request payload to create/update Yum hosted repository (Swagger: YumHostedRepositoryApiRequest)
type YumHostedRepositoryApiRequest struct {
	Cleanup   *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes     `json:"component,omitempty"`
	Name      string                   `json:"name"`
	Online    bool                     `json:"online"`
	Storage   HostedStorageAttributes  `json:"storage"`
	Yum       YumAttributes            `json:"yum"`
}

// YumGroupRepositoryApiRequest 创建/更新 Yum 组仓库请求体（Swagger: YumGroupRepositoryApiRequest）
// Request payload to create/update Yum group repository (Swagger: YumGroupRepositoryApiRequest)
type YumGroupRepositoryApiRequest struct {
	Group      GroupAttributes                   `json:"group"`
	Name       string                            `json:"name"`
	Online     bool                              `json:"online"`
	Storage    StorageAttributes                 `json:"storage"`
	YumSigning *YumSigningRepositoriesAttributes `json:"yumSigning,omitempty"`
}

// YumProxyRepositoryApiRequest 创建/更新 Yum 代理仓库请求体（Swagger: YumProxyRepositoryApiRequest）
// Request payload to create/update Yum proxy repository (Swagger: YumProxyRepositoryApiRequest)
type YumProxyRepositoryApiRequest struct {
	Cleanup       *CleanupPolicyAttributes          `json:"cleanup,omitempty"`
	HttpClient    HttpClientAttributes              `json:"httpClient"`
	Name          string                            `json:"name"`
	NegativeCache NegativeCacheAttributes           `json:"negativeCache"`
	Online        bool                              `json:"online"`
	Proxy         ProxyAttributes                   `json:"proxy"`
	RoutingRule   string                            `json:"routingRule,omitempty"`
	Storage       StorageAttributes                 `json:"storage"`
	YumSigning    *YumSigningRepositoriesAttributes `json:"yumSigning,omitempty"`
}

// P2ProxyRepositoryApiRequest 创建/更新 P2 代理仓库请求体（Swagger: P2ProxyRepositoryApiRequest）
// Request payload to create/update P2 proxy repository (Swagger: P2ProxyRepositoryApiRequest)
type P2ProxyRepositoryApiRequest struct {
	Cleanup       *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	HttpClient    HttpClientAttributes     `json:"httpClient"`
	Name          string                   `json:"name"`
	NegativeCache NegativeCacheAttributes  `json:"negativeCache"`
	Online        bool                     `json:"online"`
	Proxy         ProxyAttributes          `json:"proxy"`
	RoutingRule   string                   `json:"routingRule,omitempty"`
	Storage       StorageAttributes        `json:"storage"`
}

// GroupDeployAttributes 组仓库部署属性（Swagger: GroupDeployAttributes）
// Group deployment attributes (Swagger: GroupDeployAttributes)
type GroupDeployAttributes struct {
	MemberNames    []string `json:"memberNames"`
	WritableMember string   `json:"writableMember,omitempty"`
}

// ConanProxyAttributes Conan 代理属性（Swagger: ConanProxyAttributes）
// Conan proxy attributes (Swagger: ConanProxyAttributes)
type ConanProxyAttributes struct {
	ConanVersion string `json:"conanVersion,omitempty"`
}

// SimpleApiGroupDeployRepository 通用组仓库（含可写成员）返回体（Swagger: SimpleApiGroupDeployRepository）
// Generic group repository (with writable member) response (Swagger: SimpleApiGroupDeployRepository)
type SimpleApiGroupDeployRepository struct {
	Format  string                `json:"format"`
	Group   GroupDeployAttributes `json:"group"`
	Name    string                `json:"name"`
	Online  bool                  `json:"online"`
	Storage StorageAttributes     `json:"storage"`
	Type    string                `json:"type"`
	URL     string                `json:"url"`
}

// ConanHostedRepositoryApiRequest 创建/更新 Conan 托管仓库请求体（Swagger: ConanHostedRepositoryApiRequest）
// Request payload to create/update Conan hosted repository (Swagger: ConanHostedRepositoryApiRequest)
type ConanHostedRepositoryApiRequest struct {
	Cleanup   *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes     `json:"component,omitempty"`
	Name      string                   `json:"name"`
	Online    bool                     `json:"online"`
	Storage   HostedStorageAttributes  `json:"storage"`
}

// ConanGroupRepositoryApiRequest 创建/更新 Conan 组仓库请求体（Swagger: ConanGroupRepositoryApiRequest）
// Request payload to create/update Conan group repository (Swagger: ConanGroupRepositoryApiRequest)
type ConanGroupRepositoryApiRequest struct {
	Group   GroupDeployAttributes `json:"group"`
	Name    string                `json:"name"`
	Online  bool                  `json:"online"`
	Storage StorageAttributes     `json:"storage"`
}

// ConanProxyRepositoryApiRequest 创建/更新 Conan 代理仓库请求体（Swagger: ConanProxyRepositoryApiRequest）
// Request payload to create/update Conan proxy repository (Swagger: ConanProxyRepositoryApiRequest)
type ConanProxyRepositoryApiRequest struct {
	Cleanup       *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	HttpClient    HttpClientAttributes     `json:"httpClient"`
	Name          string                   `json:"name"`
	NegativeCache NegativeCacheAttributes  `json:"negativeCache"`
	Online        bool                     `json:"online"`
	Proxy         ProxyAttributes          `json:"proxy"`
	RoutingRule   string                   `json:"routingRule,omitempty"`
	Storage       StorageAttributes        `json:"storage"`
	ConanProxy    *ConanProxyAttributes    `json:"conanProxy,omitempty"`
}

// Repository 仓库实体（Swagger: RepositoryXO）
// Repository entity (Swagger: RepositoryXO)
// 字段映射/Field mapping：
// - name：仓库名称 / repository name
// - format：仓库格式（如 maven2、npm 等）/ repository format (e.g., maven2, npm)
// - type：仓库类型（hosted、proxy、group）/ repository type (hosted, proxy, group)
// - url：仓库浏览地址 / repository browse URL
// - attributes：仓库属性的扩展字典 / extended attributes map
type Repository struct {
	Attributes map[string]map[string]interface{} `json:"attributes"`
	Format     string                            `json:"format"`
	Name       string                            `json:"name"`
	Type       string                            `json:"type"`
	URL        string                            `json:"url"`
}

// SimpleApiGroupRepository Maven 组仓库返回体（Swagger: SimpleApiGroupRepository）
// Maven group repository response (Swagger: SimpleApiGroupRepository)
// 包含基本信息与组属性（memberNames）/ includes basic info and group attributes (memberNames)
type SimpleApiGroupRepository struct {
	Format  string            `json:"format"`
	Group   GroupAttributes   `json:"group"`
	Name    string            `json:"name"`
	Online  bool              `json:"online"`
	Storage StorageAttributes `json:"storage"`
	Type    string            `json:"type"`
	URL     string            `json:"url"`
}

// SimpleApiHostedRepository Raw 托管仓库返回体（Swagger: SimpleApiHostedRepository）
// Raw hosted repository response (Swagger: SimpleApiHostedRepository)
// 包含基本信息与存储属性 / includes basic info and storage attributes
type SimpleApiHostedRepository struct {
	Format  string                  `json:"format"`
	Name    string                  `json:"name"`
	Online  bool                    `json:"online"`
	Raw     RawAttributes           `json:"raw"`
	Storage HostedStorageAttributes `json:"storage"`
	Type    string                  `json:"type"`
	URL     string                  `json:"url"`
}

// DockerAttributes Docker 属性（Swagger: DockerAttributes）
// Docker attributes (Swagger: DockerAttributes)
type DockerAttributes struct {
	ForceBasicAuth bool   `json:"forceBasicAuth"`
	HttpPort       int    `json:"httpPort,omitempty"`
	HttpsPort      int    `json:"httpsPort,omitempty"`
	PathEnabled    bool   `json:"pathEnabled,omitempty"`
	Subdomain      string `json:"subdomain,omitempty"`
	V1Enabled      bool   `json:"v1Enabled"`
}

// DockerHostedStorageAttributes Docker 托管存储属性（Swagger: DockerHostedStorageAttributes）
// Docker hosted storage attributes (Swagger: DockerHostedStorageAttributes)
type DockerHostedStorageAttributes struct {
	BlobStoreName               string `json:"blobStoreName"`
	StrictContentTypeValidation bool   `json:"strictContentTypeValidation"`
	WritePolicy                 string `json:"writePolicy"`
	LatestPolicy                bool   `json:"latestPolicy"`
}

// DockerProxyAttributes Docker 代理属性（Swagger: DockerProxyAttributes）
// Docker proxy attributes (Swagger: DockerProxyAttributes)
type DockerProxyAttributes struct {
	CacheForeignLayers       bool     `json:"cacheForeignLayers,omitempty"`
	ForeignLayerUrlWhitelist []string `json:"foreignLayerUrlWhitelist,omitempty"`
	IndexType                string   `json:"indexType,omitempty"`
	IndexUrl                 string   `json:"indexUrl,omitempty"`
}

// DockerGroupRepositoryApiRequest 创建/更新 Docker 组仓库请求体（Swagger: DockerGroupRepositoryApiRequest）
// Request payload to create/update Docker group repository (Swagger: DockerGroupRepositoryApiRequest)
type DockerGroupRepositoryApiRequest struct {
	Docker  DockerAttributes      `json:"docker"`
	Group   GroupDeployAttributes `json:"group"`
	Name    string                `json:"name"`
	Online  bool                  `json:"online"`
	Storage StorageAttributes     `json:"storage"`
}

// DockerHostedRepositoryApiRequest 创建/更新 Docker 托管仓库请求体（Swagger: DockerHostedRepositoryApiRequest）
// Request payload to create/update Docker hosted repository (Swagger: DockerHostedRepositoryApiRequest)
type DockerHostedRepositoryApiRequest struct {
	Cleanup   *CleanupPolicyAttributes      `json:"cleanup,omitempty"`
	Component *ComponentAttributes          `json:"component,omitempty"`
	Docker    DockerAttributes              `json:"docker"`
	Name      string                        `json:"name"`
	Online    bool                          `json:"online"`
	Storage   DockerHostedStorageAttributes `json:"storage"`
}

// DockerProxyRepositoryApiRequest 创建/更新 Docker 代理仓库请求体（Swagger: DockerProxyRepositoryApiRequest）
// Request payload to create/update Docker proxy repository (Swagger: DockerProxyRepositoryApiRequest)
type DockerProxyRepositoryApiRequest struct {
	Cleanup       *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Docker        DockerAttributes         `json:"docker"`
	DockerProxy   DockerProxyAttributes    `json:"dockerProxy"`
	HttpClient    HttpClientAttributes     `json:"httpClient"`
	Name          string                   `json:"name"`
	NegativeCache NegativeCacheAttributes  `json:"negativeCache"`
	Online        bool                     `json:"online"`
	Proxy         ProxyAttributes          `json:"proxy"`
	RoutingRule   string                   `json:"routingRule,omitempty"`
	Storage       StorageAttributes        `json:"storage"`
}

// DockerGroupApiRepository Docker 组仓库返回体（Swagger: DockerGroupApiRepository）
// Docker group repository response (Swagger: DockerGroupApiRepository)
type DockerGroupApiRepository struct {
	Docker  DockerAttributes      `json:"docker"`
	Format  string                `json:"format"`
	Group   GroupDeployAttributes `json:"group"`
	Name    string                `json:"name"`
	Online  bool                  `json:"online"`
	Storage StorageAttributes     `json:"storage"`
	Type    string                `json:"type"`
	URL     string                `json:"url"`
}

// DockerHostedApiRepository Docker 托管仓库返回体（Swagger: DockerHostedApiRepository）
// Docker hosted repository response (Swagger: DockerHostedApiRepository)
type DockerHostedApiRepository struct {
	Cleanup   *CleanupPolicyAttributes      `json:"cleanup,omitempty"`
	Component *ComponentAttributes          `json:"component,omitempty"`
	Docker    DockerAttributes              `json:"docker"`
	Format    string                        `json:"format"`
	Name      string                        `json:"name"`
	Online    bool                          `json:"online"`
	Storage   DockerHostedStorageAttributes `json:"storage"`
	Type      string                        `json:"type"`
	URL       string                        `json:"url"`
}

// DockerProxyApiRepository Docker 代理仓库返回体（Swagger: DockerProxyApiRepository）
// Docker proxy repository response (Swagger: DockerProxyApiRepository)
type DockerProxyApiRepository struct {
	Cleanup         *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Docker          DockerAttributes         `json:"docker"`
	DockerProxy     DockerProxyAttributes    `json:"dockerProxy"`
	Format          string                   `json:"format"`
	HttpClient      HttpClientAttributes     `json:"httpClient"`
	Name            string                   `json:"name"`
	NegativeCache   NegativeCacheAttributes  `json:"negativeCache"`
	Online          bool                     `json:"online"`
	Proxy           ProxyAttributes          `json:"proxy"`
	RoutingRuleName string                   `json:"routingRuleName,omitempty"`
	Storage         StorageAttributes        `json:"storage"`
	Type            string                   `json:"type"`
	URL             string                   `json:"url"`
}

// StorageAttributes 通用存储属性（Swagger: StorageAttributes）
// Generic storage attributes (Swagger: StorageAttributes)
// - blobStoreName：Blob 存储名称 / blob store name
// - strictContentTypeValidation：严格内容类型校验 / strict content type validation
type StorageAttributes struct {
	BlobStoreName               string `json:"blobStoreName"`
	StrictContentTypeValidation bool   `json:"strictContentTypeValidation"`
}

// Component 组件实体（Swagger: ComponentXO）
// Component entity (Swagger: ComponentXO)
type Component struct {
	Assets     []Asset `json:"assets"`
	Format     string  `json:"format"`
	Group      string  `json:"group"`
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Repository string  `json:"repository"`
	Version    string  `json:"version"`
}

// PageComponent 组件分页（Swagger: PageComponentXO）
// Component page (Swagger: PageComponentXO)
type PageComponent struct {
	ContinuationToken string      `json:"continuationToken"`
	Items             []Component `json:"items"`
}

// UploadAssets 通用构件上传载体（Swagger: POST /v1/components, operationId: uploadComponent）
// Generic upload carrier for components (Swagger: POST /v1/components, operationId: uploadComponent)
//
// 说明 / Notes:
// - swagger.json 同时支持 apt/pypi/raw/npm/nuget/rubygems/helm/docker 等格式的表单字段；
// - 当前实现聚焦 Maven2（maven2.* 字段），其他格式可按需扩展。
type UploadAssets struct {
	yum    *UploadAssetYum
	maven2 *UploadAssetMaven2
}

type UploadAssetYum struct {
	directory     string
	asset         io.Reader
	assetFilename string
}

// UploadAssetMaven2 Maven2 构件上传字段映射（multipart/form-data）
// Maven2 upload form fields mapping (multipart/form-data)
//
// 字段 / Fields：
// - maven2.groupId：Group ID（如 org.springframework.boot）
// - maven2.artifactId：Artifact ID（如 spring-boot-starter-web）
// - maven2.version：版本号（如 2.7.14）
// - maven2.generate-pom：按坐标生成 POM（boolean，可选）/ generate POM by coordinates (optional)
// - maven2.packaging：打包类型（如 jar/pom，可选）/ packaging type (optional)
// - maven2.asset1：主构件文件 / primary artifact file
// - maven2.asset1.classifier：主构件分类（如 sources/javadoc，可选）/ classifier for primary artifact (optional)
// - maven2.asset1.extension：主构件扩展名（如 jar/pom）/ extension for primary artifact
// - maven2.asset2/maven2.asset3：副构件文件 / secondary artifact files
// - maven2.assetN.classifier：副构件分类（可选）/ classifier for secondary artifacts (optional)
// - maven2.assetN.extension：副构件扩展名 / extension for secondary artifacts
type UploadAssetMaven2 struct {
	groupId    string
	artifactId string
	version    string

	generatePom *bool
	packaging   string

	asset1Extension  string
	asset1Classifier string
	asset1           io.Reader

	asset2Extension  string
	asset2Classifier string
	asset2           io.Reader

	asset3Extension  string
	asset3Classifier string
	asset3           io.Reader
}
