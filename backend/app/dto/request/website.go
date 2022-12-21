package request

import "github.com/1Panel-dev/1Panel/backend/app/dto"

type WebsiteSearch struct {
	dto.PageInfo
}

type WebsiteCreate struct {
	PrimaryDomain  string `json:"primaryDomain" validate:"required"`
	Type           string `json:"type" validate:"required"`
	Alias          string `json:"alias" validate:"required"`
	Remark         string `json:"remark"`
	OtherDomains   string `json:"otherDomains"`
	Proxy          string `json:"proxy"`
	WebsiteGroupID uint   `json:"webSiteGroupID" validate:"required"`

	AppType      string        `json:"appType" validate:"oneof=new installed"`
	AppInstall   NewAppInstall `json:"appInstall"`
	AppID        uint          `json:"appID"`
	AppInstallID uint          `json:"appInstallID"`
}

type NewAppInstall struct {
	Name        string                 `json:"name"`
	AppDetailId uint                   `json:"appDetailID"`
	Params      map[string]interface{} `json:"params"`
}

type WebsiteInstallCheckReq struct {
	InstallIds []uint `json:"InstallIds" validate:"required"`
}

type WebsiteUpdate struct {
	ID             uint   `json:"id" validate:"required"`
	PrimaryDomain  string `json:"primaryDomain" validate:"required"`
	Remark         string `json:"remark"`
	WebsiteGroupID uint   `json:"webSiteGroupID" validate:"required"`
}

type WebsiteDelete struct {
	ID           uint `json:"id" validate:"required"`
	DeleteApp    bool `json:"deleteApp"`
	DeleteBackup bool `json:"deleteBackup"`
	ForceDelete  bool `json:"forceDelete"`
}

type WebsiteWafReq struct {
	WebsiteID uint   `json:"websiteId" validate:"required"`
	Key       string `json:"key" validate:"required"`
	Rule      string `json:"rule" validate:"required"`
}

type WebsiteWafUpdate struct {
	WebsiteID uint   `json:"websiteId" validate:"required"`
	Key       string `json:"key" validate:"required"`
	Enable    bool   `json:"enable" validate:"required"`
}

type WebsiteRecover struct {
	WebsiteName string `json:"websiteName" validate:"required"`
	Type        string `json:"type" validate:"required"`
	BackupName  string `json:"backupName" validate:"required"`
}

type WebsiteRecoverByFile struct {
	WebsiteName string `json:"websiteName" validate:"required"`
	Type        string `json:"type" validate:"required"`
	FileDir     string `json:"fileDir" validate:"required"`
	FileName    string `json:"fileName" validate:"required"`
}

type WebsiteGroupCreate struct {
	Name string `json:"name" validate:"required"`
}

type WebsiteGroupUpdate struct {
	ID      uint   `json:"id" validate:"required"`
	Name    string `json:"name"`
	Default bool   `json:"default"`
}

type WebsiteDomainCreate struct {
	WebsiteID uint   `json:"websiteId" validate:"required"`
	Port      int    `json:"port" validate:"required"`
	Domain    string `json:"domain" validate:"required"`
}

type WebsiteDomainDelete struct {
	ID uint `json:"id" validate:"required"`
}

type WebsiteHTTPSOp struct {
	WebsiteID    uint   `json:"websiteId" validate:"required"`
	Enable       bool   `json:"enable" validate:"required"`
	WebsiteSSLID uint   `json:"websiteSSLId"`
	Type         string `json:"type"  validate:"oneof=existed auto manual"`
	PrivateKey   string `json:"privateKey"`
	Certificate  string `json:"certificate"`
}