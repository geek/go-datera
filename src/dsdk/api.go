package dsdk

import (
	"encoding/json"
	//"fmt"
	"errors"
	"strings"
)

// Using a global here because screw having to pass this around to everything
// even via an autogeneration script.  We may hit some limitations later with
// concurrency, but I need this working now.
var conn *ApiConnection

type RootEp struct {
	Path                 string
	AppInstances         AppInstancesEndpoint
	Api                  ApiEndpoint
	AppTemplates         AppTemplatesEndpoint
	Initiators           InitiatorsEndpoint
	InitiatorGroups      InitiatorGroupsEndpoint
	AccessNetworkIpPools AccessNetworkIpPoolsEndpoint
	StorageNodes         StorageNodesEndpoint
	System               SystemEndpoint
	EventLogs            EventLogsEndpoint
	AuditLogs            AuditLogsEndpoint
	FaultLogs            FaultLogsEndpoint
	Roles                RolesEndpoint
	Users                UsersEndpoint
	Upgrade              UpgradeEndpoint
	Time                 TimeEndpoint
	Tenants              TenantsEndpoint
}

func NewRootEp(hostname, port, username, password, apiVersion, tenant, timeout string, headers map[string]string, secure bool) (*RootEp, error) {
	var err error
	//Initialize global connection object
	conn, err = NewApiConnection(hostname, port, username, password, apiVersion, tenant, timeout, headers, secure)
	if err != nil {
		return nil, err
	}
	err = conn.Login()
	if err != nil {
		return nil, err
	}
	return &RootEp{
		Path:                 "",
		AppInstances:         NewAppInstancesEndpoint(""),
		Api:                  NewApiEndpoint(""),
		AppTemplates:         NewAppTemplatesEndpoint(""),
		Initiators:           NewInitiatorsEndpoint(""),
		InitiatorGroups:      NewInitiatorGroupsEndpoint(""),
		AccessNetworkIpPools: NewAccessNetworkIpPoolsEndpoint(""),
		StorageNodes:         NewStorageNodesEndpoint(""),
		System:               NewSystemEndpoint(""),
		EventLogs:            NewEventLogsEndpoint(""),
		AuditLogs:            NewAuditLogsEndpoint(""),
		FaultLogs:            NewFaultLogsEndpoint(""),
		Roles:                NewRolesEndpoint(""),
		Users:                NewUsersEndpoint(""),
		Upgrade:              NewUpgradeEndpoint(""),
		Time:                 NewTimeEndpoint(""),
		Tenants:              NewTenantsEndpoint(""),
	}, nil
}

type AccessNetworkIpPoolEntity struct {
	Descr        string      `json:"descr,omitempty"`
	Name         string      `json:"name,omitempty"`
	NetworkPaths interface{} `json:"network_paths,omitempty"`
	Path         string      `json:"path,omitempty"`
}

func (en AccessNetworkIpPoolEntity) Reload() (AccessNetworkIpPoolEntity, error) {
	var n AccessNetworkIpPoolEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en AccessNetworkIpPoolEntity) Set(bodyp ...string) (AccessNetworkIpPoolEntity, error) {
	var n AccessNetworkIpPoolEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en AccessNetworkIpPoolEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type AclPolicyEntity struct {
	InitiatorGroups []InitiatorGroupEntity `json:"initiator_groups,omitempty"`
	Initiators      []InitiatorEntity      `json:"initiators,omitempty"`
	Path            string                 `json:"path,omitempty"`
}

func (en AclPolicyEntity) Reload() (AclPolicyEntity, error) {
	var n AclPolicyEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en AclPolicyEntity) Set(bodyp ...string) (AclPolicyEntity, error) {
	var n AclPolicyEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en AclPolicyEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type AppInstanceEntity struct {
	AccessControlMode string                  `json:"access_control_mode,omitempty"`
	AdminState        string                  `json:"admin_state,omitempty"`
	AppTemplate       AppTemplateEntity       `json:"app_template,omitempty"`
	Causes            []interface{}           `json:"causes,omitempty"`
	CloneSrc          map[string]string       `json:"clone_src,omitempty"`
	CreateMode        string                  `json:"create_mode,omitempty"`
	Descr             string                  `json:"descr,omitempty"`
	Health            string                  `json:"health,omitempty"`
	Id                string                  `json:"id,omitempty"`
	Name              string                  `json:"name,omitempty"`
	OpStatus          string                  `json:"op_status,omitempty"`
	Path              string                  `json:"path,omitempty"`
	RestorePoint      string                  `json:"restore_point,omitempty"`
	SnapshotPolicies  []SnapshotPolicyEntity  `json:"snapshot_policies,omitempty"`
	Snapshots         []SnapshotEntity        `json:"snapshots,omitempty"`
	StorageInstances  []StorageInstanceEntity `json:"storage_instances,omitempty"`
	Uuid              string                  `json:"uuid,omitempty"`
}

func (en AppInstanceEntity) Reload() (AppInstanceEntity, error) {
	var n AppInstanceEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en AppInstanceEntity) Set(bodyp ...string) (AppInstanceEntity, error) {
	var n AppInstanceEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en AppInstanceEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type AppTemplateEntity struct {
	AppInstances     []AppInstanceEntity     `json:"app_instances,omitempty"`
	Descr            string                  `json:"descr,omitempty"`
	Name             string                  `json:"name,omitempty"`
	Path             string                  `json:"path,omitempty"`
	SnapshotPolicies []SnapshotPolicyEntity  `json:"snapshot_policies,omitempty"`
	StorageTemplates []StorageTemplateEntity `json:"storage_templates,omitempty"`
}

func (en AppTemplateEntity) Reload() (AppTemplateEntity, error) {
	var n AppTemplateEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en AppTemplateEntity) Set(bodyp ...string) (AppTemplateEntity, error) {
	var n AppTemplateEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en AppTemplateEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type AuditLogEntity struct {
	Description string     `json:"description,omitempty"`
	Id          string     `json:"id,omitempty"`
	ObjectName  string     `json:"object_name,omitempty"`
	ObjectType  string     `json:"object_type,omitempty"`
	ObjectUrl   string     `json:"object_url,omitempty"`
	Operation   string     `json:"operation,omitempty"`
	ParamInfo   string     `json:"param_info,omitempty"`
	Path        string     `json:"path,omitempty"`
	SessionInfo string     `json:"session_info,omitempty"`
	Timestamp   string     `json:"timestamp,omitempty"`
	User        UserEntity `json:"user,omitempty"`
	Version     string     `json:"version,omitempty"`
}

func (en AuditLogEntity) Reload() (AuditLogEntity, error) {
	var n AuditLogEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en AuditLogEntity) Set(bodyp ...string) (AuditLogEntity, error) {
	var n AuditLogEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en AuditLogEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type AuthEntity struct {
	InitiatorPswd     string `json:"initiator_pswd,omitempty"`
	InitiatorUserName string `json:"initiator_user_name,omitempty"`
	Path              string `json:"path,omitempty"`
	TargetPswd        string `json:"target_pswd,omitempty"`
	TargetUserName    string `json:"target_user_name,omitempty"`
	Type              string `json:"type,omitempty"`
}

func (en AuthEntity) Reload() (AuthEntity, error) {
	var n AuthEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en AuthEntity) Set(bodyp ...string) (AuthEntity, error) {
	var n AuthEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en AuthEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type BootDriveEntity struct {
	Causes    []interface{} `json:"causes,omitempty"`
	Health    string        `json:"health,omitempty"`
	Id        string        `json:"id,omitempty"`
	OpState   string        `json:"op_state,omitempty"`
	Path      string        `json:"path,omitempty"`
	Size      int           `json:"size,omitempty"`
	SlotLabel string        `json:"slot_label,omitempty"`
}

func (en BootDriveEntity) Reload() (BootDriveEntity, error) {
	var n BootDriveEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en BootDriveEntity) Set(bodyp ...string) (BootDriveEntity, error) {
	var n BootDriveEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en BootDriveEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type DnsEntity struct {
	Domain string `json:"domain,omitempty"`
	Path   string `json:"path,omitempty"`
}

func (en DnsEntity) Reload() (DnsEntity, error) {
	var n DnsEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en DnsEntity) Set(bodyp ...string) (DnsEntity, error) {
	var n DnsEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en DnsEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type DnsSearchDomainEntity struct {
	Domain string `json:"domain,omitempty"`
	Order  int    `json:"order,omitempty"`
	Path   string `json:"path,omitempty"`
}

func (en DnsSearchDomainEntity) Reload() (DnsSearchDomainEntity, error) {
	var n DnsSearchDomainEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en DnsSearchDomainEntity) Set(bodyp ...string) (DnsSearchDomainEntity, error) {
	var n DnsSearchDomainEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en DnsSearchDomainEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type DnsServerEntity struct {
	Ip    string `json:"ip,omitempty"`
	Order int    `json:"order,omitempty"`
	Path  string `json:"path,omitempty"`
}

func (en DnsServerEntity) Reload() (DnsServerEntity, error) {
	var n DnsServerEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en DnsServerEntity) Set(bodyp ...string) (DnsServerEntity, error) {
	var n DnsServerEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en DnsServerEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type EventLogEntity struct {
	Cause        string `json:"cause,omitempty"`
	Code         string `json:"code,omitempty"`
	Id           string `json:"id,omitempty"`
	ObjectId     string `json:"object_id,omitempty"`
	ObjectLbl    string `json:"object_lbl,omitempty"`
	ObjectPath   string `json:"object_path,omitempty"`
	ObjectTenant string `json:"object_tenant,omitempty"`
	ObjectType   string `json:"object_type,omitempty"`
	ObjectUrl    string `json:"object_url,omitempty"`
	Path         string `json:"path,omitempty"`
	Severity     string `json:"severity,omitempty"`
	Timestamp    string `json:"timestamp,omitempty"`
	Type         string `json:"type,omitempty"`
}

func (en EventLogEntity) Reload() (EventLogEntity, error) {
	var n EventLogEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en EventLogEntity) Set(bodyp ...string) (EventLogEntity, error) {
	var n EventLogEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en EventLogEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type FaultLogEntity struct {
	Acknowledged    bool   `json:"acknowledged,omitempty"`
	CallhomeEnabled bool   `json:"callhome_enabled,omitempty"`
	Cause           string `json:"cause,omitempty"`
	Cleared         bool   `json:"cleared,omitempty"`
	Code            string `json:"code,omitempty"`
	Count           int    `json:"count,omitempty"`
	Id              string `json:"id,omitempty"`
	ObjectId        string `json:"object_id,omitempty"`
	ObjectLbl       string `json:"object_lbl,omitempty"`
	ObjectPath      string `json:"object_path,omitempty"`
	ObjectTenant    string `json:"object_tenant,omitempty"`
	ObjectType      string `json:"object_type,omitempty"`
	ObjectUrl       string `json:"object_url,omitempty"`
	Path            string `json:"path,omitempty"`
	Repeat          string `json:"repeat,omitempty"`
	Severity        string `json:"severity,omitempty"`
	Timestamp       int    `json:"timestamp,omitempty"`
	Type            string `json:"type,omitempty"`
}

func (en FaultLogEntity) Reload() (FaultLogEntity, error) {
	var n FaultLogEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en FaultLogEntity) Set(bodyp ...string) (FaultLogEntity, error) {
	var n FaultLogEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en FaultLogEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type HddEntity struct {
	Causes    []interface{} `json:"causes,omitempty"`
	Health    string        `json:"health,omitempty"`
	Id        string        `json:"id,omitempty"`
	OpState   string        `json:"op_state,omitempty"`
	Path      string        `json:"path,omitempty"`
	Size      int           `json:"size,omitempty"`
	SlotLabel string        `json:"slot_label,omitempty"`
}

func (en HddEntity) Reload() (HddEntity, error) {
	var n HddEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en HddEntity) Set(bodyp ...string) (HddEntity, error) {
	var n HddEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en HddEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type HttpProxyEntity struct {
	Enabled  bool       `json:"enabled,omitempty"`
	Host     string     `json:"host,omitempty"`
	Password string     `json:"password,omitempty"`
	Path     string     `json:"path,omitempty"`
	Port     int        `json:"port,omitempty"`
	User     UserEntity `json:"user,omitempty"`
}

func (en HttpProxyEntity) Reload() (HttpProxyEntity, error) {
	var n HttpProxyEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en HttpProxyEntity) Set(bodyp ...string) (HttpProxyEntity, error) {
	var n HttpProxyEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en HttpProxyEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type InitiatorEntity struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

func (en InitiatorEntity) Reload() (InitiatorEntity, error) {
	var n InitiatorEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en InitiatorEntity) Set(bodyp ...string) (InitiatorEntity, error) {
	var n InitiatorEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en InitiatorEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type InitiatorGroupEntity struct {
	Members []interface{} `json:"members,omitempty"`
	Name    string        `json:"name,omitempty"`
	Path    string        `json:"path,omitempty"`
}

func (en InitiatorGroupEntity) Reload() (InitiatorGroupEntity, error) {
	var n InitiatorGroupEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en InitiatorGroupEntity) Set(bodyp ...string) (InitiatorGroupEntity, error) {
	var n InitiatorGroupEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en InitiatorGroupEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type InternalIpBlockEntity struct {
	Gateway string `json:"gateway,omitempty"`
	Mtu     int    `json:"mtu,omitempty"`
	Name    string `json:"name,omitempty"`
	Netmask int    `json:"netmask,omitempty"`
	Path    string `json:"path,omitempty"`
	Range   int    `json:"range,omitempty"`
	StartIp string `json:"start_ip,omitempty"`
	Vlan    int    `json:"vlan,omitempty"`
}

func (en InternalIpBlockEntity) Reload() (InternalIpBlockEntity, error) {
	var n InternalIpBlockEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en InternalIpBlockEntity) Set(bodyp ...string) (InternalIpBlockEntity, error) {
	var n InternalIpBlockEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en InternalIpBlockEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type IpAddressEntity struct {
	Gateway string `json:"gateway,omitempty"`
	Ip      string `json:"ip,omitempty"`
	Mtu     int    `json:"mtu,omitempty"`
	Name    string `json:"name,omitempty"`
	Netmask int    `json:"netmask,omitempty"`
	Path    string `json:"path,omitempty"`
	Vlan    int    `json:"vlan,omitempty"`
}

func (en IpAddressEntity) Reload() (IpAddressEntity, error) {
	var n IpAddressEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en IpAddressEntity) Set(bodyp ...string) (IpAddressEntity, error) {
	var n IpAddressEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en IpAddressEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type IpBlockEntity struct {
	Gateway string `json:"gateway,omitempty"`
	Mtu     int    `json:"mtu,omitempty"`
	Name    string `json:"name,omitempty"`
	Netmask int    `json:"netmask,omitempty"`
	Path    string `json:"path,omitempty"`
	Range   int    `json:"range,omitempty"`
	StartIp string `json:"start_ip,omitempty"`
	Vlan    int    `json:"vlan,omitempty"`
}

func (en IpBlockEntity) Reload() (IpBlockEntity, error) {
	var n IpBlockEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en IpBlockEntity) Set(bodyp ...string) (IpBlockEntity, error) {
	var n IpBlockEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en IpBlockEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type IpPoolEntity struct {
	Descr        string      `json:"descr,omitempty"`
	Name         string      `json:"name,omitempty"`
	NetworkPaths interface{} `json:"network_paths,omitempty"`
	Path         string      `json:"path,omitempty"`
}

func (en IpPoolEntity) Reload() (IpPoolEntity, error) {
	var n IpPoolEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en IpPoolEntity) Set(bodyp ...string) (IpPoolEntity, error) {
	var n IpPoolEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en IpPoolEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type MonitoringDestinationEntity struct {
	Facility  string `json:"facility,omitempty"`
	Host      string `json:"host,omitempty"`
	LastMsgTs string `json:"last_msg_ts,omitempty"`
	Name      string `json:"name,omitempty"`
	OpState   string `json:"op_state,omitempty"`
	Path      string `json:"path,omitempty"`
	Port      int    `json:"port,omitempty"`
	Type      string `json:"type,omitempty"`
}

func (en MonitoringDestinationEntity) Reload() (MonitoringDestinationEntity, error) {
	var n MonitoringDestinationEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en MonitoringDestinationEntity) Set(bodyp ...string) (MonitoringDestinationEntity, error) {
	var n MonitoringDestinationEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en MonitoringDestinationEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type MonitoringPolicyEntity struct {
	Destinations []interface{} `json:"destinations,omitempty"`
	Enabled      bool          `json:"enabled,omitempty"`
	Name         string        `json:"name,omitempty"`
	Path         string        `json:"path,omitempty"`
}

func (en MonitoringPolicyEntity) Reload() (MonitoringPolicyEntity, error) {
	var n MonitoringPolicyEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en MonitoringPolicyEntity) Set(bodyp ...string) (MonitoringPolicyEntity, error) {
	var n MonitoringPolicyEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en MonitoringPolicyEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type NetworkEntity struct {
	AccessNetworks  []interface{} `json:"access_networks,omitempty"`
	AccessVip       interface{}   `json:"access_vip,omitempty"`
	InternalNetwork interface{}   `json:"internal_network,omitempty"`
	Mapping         interface{}   `json:"mapping,omitempty"`
	MgmtVip         interface{}   `json:"mgmt_vip,omitempty"`
	Path            string        `json:"path,omitempty"`
}

func (en NetworkEntity) Reload() (NetworkEntity, error) {
	var n NetworkEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en NetworkEntity) Set(bodyp ...string) (NetworkEntity, error) {
	var n NetworkEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en NetworkEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type NetworkMappingEntity struct {
	Access1   string `json:"access_1,omitempty"`
	Access2   string `json:"access_2,omitempty"`
	Internal1 string `json:"internal_1,omitempty"`
	Internal2 string `json:"internal_2,omitempty"`
	Mgmt1     string `json:"mgmt_1,omitempty"`
	Mgmt2     string `json:"mgmt_2,omitempty"`
	Path      string `json:"path,omitempty"`
}

func (en NetworkMappingEntity) Reload() (NetworkMappingEntity, error) {
	var n NetworkMappingEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en NetworkMappingEntity) Set(bodyp ...string) (NetworkMappingEntity, error) {
	var n NetworkMappingEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en NetworkMappingEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type NicEntity struct {
	Causes    []interface{} `json:"causes,omitempty"`
	Health    string        `json:"health,omitempty"`
	Id        string        `json:"id,omitempty"`
	OpState   string        `json:"op_state,omitempty"`
	Path      string        `json:"path,omitempty"`
	SlotLabel string        `json:"slot_label,omitempty"`
}

func (en NicEntity) Reload() (NicEntity, error) {
	var n NicEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en NicEntity) Set(bodyp ...string) (NicEntity, error) {
	var n NicEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en NicEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type NtpServerEntity struct {
	Ip    string `json:"ip,omitempty"`
	Order int    `json:"order,omitempty"`
	Path  string `json:"path,omitempty"`
}

func (en NtpServerEntity) Reload() (NtpServerEntity, error) {
	var n NtpServerEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en NtpServerEntity) Set(bodyp ...string) (NtpServerEntity, error) {
	var n NtpServerEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en NtpServerEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type NvmFlashDeviceEntity struct {
	Causes    []interface{} `json:"causes,omitempty"`
	Health    string        `json:"health,omitempty"`
	Id        string        `json:"id,omitempty"`
	OpState   string        `json:"op_state,omitempty"`
	Path      string        `json:"path,omitempty"`
	Size      int           `json:"size,omitempty"`
	SlotLabel string        `json:"slot_label,omitempty"`
}

func (en NvmFlashDeviceEntity) Reload() (NvmFlashDeviceEntity, error) {
	var n NvmFlashDeviceEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en NvmFlashDeviceEntity) Set(bodyp ...string) (NvmFlashDeviceEntity, error) {
	var n NvmFlashDeviceEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en NvmFlashDeviceEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type PerformancePolicyEntity struct {
	Path              string `json:"path,omitempty"`
	ReadBandwidthMax  int    `json:"read_bandwidth_max,omitempty"`
	ReadIopsMax       int    `json:"read_iops_max,omitempty"`
	TotalBandwidthMax int    `json:"total_bandwidth_max,omitempty"`
	TotalIopsMax      int    `json:"total_iops_max,omitempty"`
	WriteBandwidthMax int    `json:"write_bandwidth_max,omitempty"`
	WriteIopsMax      int    `json:"write_iops_max,omitempty"`
}

func (en PerformancePolicyEntity) Reload() (PerformancePolicyEntity, error) {
	var n PerformancePolicyEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en PerformancePolicyEntity) Set(bodyp ...string) (PerformancePolicyEntity, error) {
	var n PerformancePolicyEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en PerformancePolicyEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type RoleEntity struct {
	Path       string        `json:"path,omitempty"`
	Privileges []interface{} `json:"privileges,omitempty"`
	RoleId     string        `json:"role_id,omitempty"`
}

func (en RoleEntity) Reload() (RoleEntity, error) {
	var n RoleEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en RoleEntity) Set(bodyp ...string) (RoleEntity, error) {
	var n RoleEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en RoleEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type SnapshotEntity struct {
	OpState   string `json:"op_state,omitempty"`
	Path      string `json:"path,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	UtcTs     int    `json:"utc_ts,omitempty"`
	Uuid      string `json:"uuid,omitempty"`
}

func (en SnapshotEntity) Reload() (SnapshotEntity, error) {
	var n SnapshotEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en SnapshotEntity) Set(bodyp ...string) (SnapshotEntity, error) {
	var n SnapshotEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en SnapshotEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type SnapshotPolicyEntity struct {
	Interval       string `json:"interval,omitempty"`
	Name           string `json:"name,omitempty"`
	Path           string `json:"path,omitempty"`
	RetentionCount int    `json:"retention_count,omitempty"`
	StartTime      string `json:"start_time,omitempty"`
}

func (en SnapshotPolicyEntity) Reload() (SnapshotPolicyEntity, error) {
	var n SnapshotPolicyEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en SnapshotPolicyEntity) Set(bodyp ...string) (SnapshotPolicyEntity, error) {
	var n SnapshotPolicyEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en SnapshotPolicyEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type SnmpPolicyEntity struct {
	Contact  string       `json:"contact,omitempty"`
	Enabled  bool         `json:"enabled,omitempty"`
	Location string       `json:"location,omitempty"`
	Path     string       `json:"path,omitempty"`
	Users    []UserEntity `json:"users,omitempty"`
}

func (en SnmpPolicyEntity) Reload() (SnmpPolicyEntity, error) {
	var n SnmpPolicyEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en SnmpPolicyEntity) Set(bodyp ...string) (SnmpPolicyEntity, error) {
	var n SnmpPolicyEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en SnmpPolicyEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type SnmpUserEntity struct {
	AuthPass      string `json:"auth_pass,omitempty"`
	AuthProtocol  string `json:"auth_protocol,omitempty"`
	EncrPass      string `json:"encr_pass,omitempty"`
	EncrProtocol  string `json:"encr_protocol,omitempty"`
	Path          string `json:"path,omitempty"`
	SecurityLevel string `json:"security_level,omitempty"`
	UserId        string `json:"user_id,omitempty"`
	Version       string `json:"version,omitempty"`
	View          string `json:"view,omitempty"`
}

func (en SnmpUserEntity) Reload() (SnmpUserEntity, error) {
	var n SnmpUserEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en SnmpUserEntity) Set(bodyp ...string) (SnmpUserEntity, error) {
	var n SnmpUserEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en SnmpUserEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type StorageInstanceEntity struct {
	Access             interface{}     `json:"access,omitempty"`
	AccessControlMode  string          `json:"access_control_mode,omitempty"`
	AclPolicy          AclPolicyEntity `json:"acl_policy,omitempty"`
	ActiveInitiators   []interface{}   `json:"active_initiators,omitempty"`
	ActiveStorageNodes []interface{}   `json:"active_storage_nodes,omitempty"`
	AdminState         string          `json:"admin_state,omitempty"`
	Auth               AuthEntity      `json:"auth,omitempty"`
	Causes             []interface{}   `json:"causes,omitempty"`
	Health             string          `json:"health,omitempty"`
	IpPool             IpPoolEntity    `json:"ip_pool,omitempty"`
	Name               string          `json:"name,omitempty"`
	OpState            string          `json:"op_state,omitempty"`
	Path               string          `json:"path,omitempty"`
	Uuid               string          `json:"uuid,omitempty"`
	Volumes            []VolumeEntity  `json:"volumes,omitempty"`
}

func (en StorageInstanceEntity) Reload() (StorageInstanceEntity, error) {
	var n StorageInstanceEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en StorageInstanceEntity) Set(bodyp ...string) (StorageInstanceEntity, error) {
	var n StorageInstanceEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en StorageInstanceEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type StorageNodeEntity struct {
	AdminState          string                  `json:"admin_state,omitempty"`
	AvailableCapacity   int                     `json:"available_capacity,omitempty"`
	BiosVersion         string                  `json:"bios_version,omitempty"`
	BootDrives          interface{}             `json:"boot_drives,omitempty"`
	BuildVersion        string                  `json:"build_version,omitempty"`
	Causes              []interface{}           `json:"causes,omitempty"`
	Disconnected        bool                    `json:"disconnected,omitempty"`
	FlashDevices        interface{}             `json:"flash_devices,omitempty"`
	Hdds                []HddEntity             `json:"hdds,omitempty"`
	Health              string                  `json:"health,omitempty"`
	HwHealth            string                  `json:"hw_health,omitempty"`
	HwState             string                  `json:"hw_state,omitempty"`
	InternalIp1         string                  `json:"internal_ip_1,omitempty"`
	InternalIp2         string                  `json:"internal_ip_2,omitempty"`
	LastRebootTimestamp int                     `json:"last_reboot_timestamp,omitempty"`
	MgmtIp1             string                  `json:"mgmt_ip_1,omitempty"`
	MgmtIp2             string                  `json:"mgmt_ip_2,omitempty"`
	Model               string                  `json:"model,omitempty"`
	Name                string                  `json:"name,omitempty"`
	Nics                []NicEntity             `json:"nics,omitempty"`
	NvmFlashDevices     []NvmFlashDeviceEntity  `json:"nvm_flash_devices,omitempty"`
	OpProgress          interface{}             `json:"op_progress,omitempty"`
	OpState             string                  `json:"op_state,omitempty"`
	OpStatus            string                  `json:"op_status,omitempty"`
	OsVersion           string                  `json:"os_version,omitempty"`
	Path                string                  `json:"path,omitempty"`
	Psus                interface{}             `json:"psus,omitempty"`
	SerialNo            string                  `json:"serial_no,omitempty"`
	StorageInstances    []StorageInstanceEntity `json:"storage_instances,omitempty"`
	SubsystemHealth     interface{}             `json:"subsystem_health,omitempty"`
	SubsystemStates     interface{}             `json:"subsystem_states,omitempty"`
	SwHealth            string                  `json:"sw_health,omitempty"`
	SwState             string                  `json:"sw_state,omitempty"`
	SwVersion           string                  `json:"sw_version,omitempty"`
	TotalCapacity       int                     `json:"total_capacity,omitempty"`
	TotalRawCapacity    int                     `json:"total_raw_capacity,omitempty"`
	Type                string                  `json:"type,omitempty"`
	Upgrade             interface{}             `json:"upgrade,omitempty"`
	Uuid                string                  `json:"uuid,omitempty"`
	Vendor              string                  `json:"vendor,omitempty"`
	Volumes             []VolumeEntity          `json:"volumes,omitempty"`
}

func (en StorageNodeEntity) Reload() (StorageNodeEntity, error) {
	var n StorageNodeEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en StorageNodeEntity) Set(bodyp ...string) (StorageNodeEntity, error) {
	var n StorageNodeEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en StorageNodeEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type StorageTemplateEntity struct {
	Auth            AuthEntity             `json:"auth,omitempty"`
	IpPool          IpPoolEntity           `json:"ip_pool,omitempty"`
	Name            string                 `json:"name,omitempty"`
	Path            string                 `json:"path,omitempty"`
	VolumeTemplates []VolumeTemplateEntity `json:"volume_templates,omitempty"`
}

func (en StorageTemplateEntity) Reload() (StorageTemplateEntity, error) {
	var n StorageTemplateEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en StorageTemplateEntity) Set(bodyp ...string) (StorageTemplateEntity, error) {
	var n StorageTemplateEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en StorageTemplateEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type SubsystemEntity struct {
	Causes      []interface{} `json:"causes,omitempty"`
	Fan         string        `json:"fan,omitempty"`
	Health      string        `json:"health,omitempty"`
	Network     NetworkEntity `json:"network,omitempty"`
	Path        string        `json:"path,omitempty"`
	Power       string        `json:"power,omitempty"`
	Temperature string        `json:"temperature,omitempty"`
	Voltage     string        `json:"voltage,omitempty"`
}

func (en SubsystemEntity) Reload() (SubsystemEntity, error) {
	var n SubsystemEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en SubsystemEntity) Set(bodyp ...string) (SubsystemEntity, error) {
	var n SubsystemEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en SubsystemEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type SystemEntity struct {
	AllFlashAvailableCapacity   int               `json:"all_flash_available_capacity,omitempty"`
	AllFlashProvisionedCapacity int               `json:"all_flash_provisioned_capacity,omitempty"`
	AllFlashTotalCapacity       int               `json:"all_flash_total_capacity,omitempty"`
	AvailableCapacity           int               `json:"available_capacity,omitempty"`
	BuildVersion                string            `json:"build_version,omitempty"`
	CallhomeEnabled             bool              `json:"callhome_enabled,omitempty"`
	Causes                      []interface{}     `json:"causes,omitempty"`
	Dns                         DnsEntity         `json:"dns,omitempty"`
	Health                      string            `json:"health,omitempty"`
	HttpProxy                   HttpProxyEntity   `json:"http_proxy,omitempty"`
	HybridAvailableCapacity     int               `json:"hybrid_available_capacity,omitempty"`
	HybridProvisionedCapacity   int               `json:"hybrid_provisioned_capacity,omitempty"`
	HybridTotalCapacity         int               `json:"hybrid_total_capacity,omitempty"`
	LastRebootTimestamp         string            `json:"last_reboot_timestamp,omitempty"`
	Name                        string            `json:"name,omitempty"`
	Network                     NetworkEntity     `json:"network,omitempty"`
	NtpServers                  []NtpServerEntity `json:"ntp_servers,omitempty"`
	OpState                     string            `json:"op_state,omitempty"`
	Path                        string            `json:"path,omitempty"`
	SwVersion                   string            `json:"sw_version,omitempty"`
	TotalCapacity               int               `json:"total_capacity,omitempty"`
	TotalProvisionedCapacity    int               `json:"total_provisioned_capacity,omitempty"`
	Upgrade                     interface{}       `json:"upgrade,omitempty"`
	Uptime                      int               `json:"uptime,omitempty"`
	Uuid                        string            `json:"uuid,omitempty"`
}

func (en SystemEntity) Reload() (SystemEntity, error) {
	var n SystemEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en SystemEntity) Set(bodyp ...string) (SystemEntity, error) {
	var n SystemEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en SystemEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type TenantEntity struct {
	Descr      string      `json:"descr,omitempty"`
	Name       string      `json:"name,omitempty"`
	ParentPath string      `json:"parent_path,omitempty"`
	Path       string      `json:"path,omitempty"`
	Subtenants interface{} `json:"subtenants,omitempty"`
}

func (en TenantEntity) Reload() (TenantEntity, error) {
	var n TenantEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en TenantEntity) Set(bodyp ...string) (TenantEntity, error) {
	var n TenantEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en TenantEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type UserEntity struct {
	Email    string         `json:"email,omitempty"`
	Enabled  bool           `json:"enabled,omitempty"`
	FullName string         `json:"full_name,omitempty"`
	Password string         `json:"password,omitempty"`
	Path     string         `json:"path,omitempty"`
	Roles    []RoleEntity   `json:"roles,omitempty"`
	Tenants  []TenantEntity `json:"tenants,omitempty"`
	UserId   string         `json:"user_id,omitempty"`
	Version  string         `json:"version,omitempty"`
}

func (en UserEntity) Reload() (UserEntity, error) {
	var n UserEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en UserEntity) Set(bodyp ...string) (UserEntity, error) {
	var n UserEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en UserEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type VipEntity struct {
	Name         string      `json:"name,omitempty"`
	NetworkPaths interface{} `json:"network_paths,omitempty"`
	Path         string      `json:"path,omitempty"`
}

func (en VipEntity) Reload() (VipEntity, error) {
	var n VipEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en VipEntity) Set(bodyp ...string) (VipEntity, error) {
	var n VipEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en VipEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type VolumeEntity struct {
	ActiveStorageNodes []interface{}    `json:"active_storage_nodes,omitempty"`
	CapacityInUse      int              `json:"capacity_in_use,omitempty"`
	Causes             []interface{}    `json:"causes,omitempty"`
	Health             string           `json:"health,omitempty"`
	Name               string           `json:"name,omitempty"`
	OpState            string           `json:"op_state,omitempty"`
	OpStatus           string           `json:"op_status,omitempty"`
	Path               string           `json:"path,omitempty"`
	PlacementMode      string           `json:"placement_mode,omitempty"`
	ReplicaCount       int              `json:"replica_count,omitempty"`
	RestorePoint       string           `json:"restore_point,omitempty"`
	Size               int              `json:"size,omitempty"`
	Snapshots          []SnapshotEntity `json:"snapshots,omitempty"`
	Uuid               string           `json:"uuid,omitempty"`
}

func (en VolumeEntity) Reload() (VolumeEntity, error) {
	var n VolumeEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en VolumeEntity) Set(bodyp ...string) (VolumeEntity, error) {
	var n VolumeEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en VolumeEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type VolumeTemplateEntity struct {
	Name          string `json:"name,omitempty"`
	Path          string `json:"path,omitempty"`
	PlacementMode string `json:"placement_mode,omitempty"`
	ReplicaCount  int    `json:"replica_count,omitempty"`
	Size          int    `json:"size,omitempty"`
}

func (en VolumeTemplateEntity) Reload() (VolumeTemplateEntity, error) {
	var n VolumeTemplateEntity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en VolumeTemplateEntity) Set(bodyp ...string) (VolumeTemplateEntity, error) {
	var n VolumeTemplateEntity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &n)
	return n, nil
}
func (en VolumeTemplateEntity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}

type MetadataEndpoint struct {
	Path string
}

func NewMetadataEndpoint(parent string) MetadataEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "metadata"}, "/"), "/")
	return MetadataEndpoint{
		Path: path,
	}
}

type LoginEndpoint struct {
	Path string
}

func NewLoginEndpoint(parent string) LoginEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "login"}, "/"), "/")
	return LoginEndpoint{
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameVolumesEndpoint struct {
	Path       string
	VolumeName AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameEndpoint
}

func (ep AppInstancesIdStorageInstancesStorageInstanceNameVolumesEndpoint) Create(bodyp ...string) (VolumeEntity, error) {
	var en VolumeEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppInstancesIdStorageInstancesStorageInstanceNameVolumesEndpoint) List(queryp ...string) ([]VolumeEntity, error) {
	var ens []VolumeEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesEndpoint(parent string) AppInstancesIdStorageInstancesStorageInstanceNameVolumesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/volumes"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameVolumesEndpoint{
		Path:       path,
		VolumeName: NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameEndpoint(path),
	}
}

type MetricsHwMetricEndpoint struct {
	Path   string
	Latest MetricsHwMetricLatestEndpoint
}

func NewMetricsHwMetricEndpoint(parent string) MetricsHwMetricEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "metrics/hw/:metric"}, "/"), "/")
	return MetricsHwMetricEndpoint{
		Path:   path,
		Latest: NewMetricsHwMetricLatestEndpoint(path),
	}
}

type EventLogsEndpoint struct {
	Path string
	Id   EventLogsIdEndpoint
}

func (ep EventLogsEndpoint) Create(bodyp ...string) (EventLogEntity, error) {
	var en EventLogEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep EventLogsEndpoint) List(queryp ...string) ([]EventLogEntity, error) {
	var ens []EventLogEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewEventLogsEndpoint(parent string) EventLogsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "event_logs"}, "/"), "/")
	return EventLogsEndpoint{
		Path: path,
		Id:   NewEventLogsIdEndpoint(path),
	}
}

type StorageNodesUuidFlashDevicesEndpoint struct {
	Path string
	Id   StorageNodesUuidFlashDevicesIdEndpoint
}

func NewStorageNodesUuidFlashDevicesEndpoint(parent string) StorageNodesUuidFlashDevicesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/flash_devices"}, "/"), "/")
	return StorageNodesUuidFlashDevicesEndpoint{
		Path: path,
		Id:   NewStorageNodesUuidFlashDevicesIdEndpoint(path),
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsTimestampEndpoint struct {
	Path string
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsTimestampEndpoint(parent string) AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsTimestampEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/volumes/:volume_name/snapshots/:timestamp"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsTimestampEndpoint{
		Path: path,
	}
}

type LogoutEndpoint struct {
	Path string
}

func NewLogoutEndpoint(parent string) LogoutEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "logout"}, "/"), "/")
	return LogoutEndpoint{
		Path: path,
	}
}

type SystemSnmpPolicyUsersEndpoint struct {
	Path   string
	UserId SystemSnmpPolicyUsersUserIdEndpoint
}

func (ep SystemSnmpPolicyUsersEndpoint) Create(bodyp ...string) (UserEntity, error) {
	var en UserEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep SystemSnmpPolicyUsersEndpoint) List(queryp ...string) ([]UserEntity, error) {
	var ens []UserEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewSystemSnmpPolicyUsersEndpoint(parent string) SystemSnmpPolicyUsersEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/snmp_policy/users"}, "/"), "/")
	return SystemSnmpPolicyUsersEndpoint{
		Path:   path,
		UserId: NewSystemSnmpPolicyUsersUserIdEndpoint(path),
	}
}

type SystemSnmpPolicyUsersUserIdEndpoint struct {
	Path string
}

func NewSystemSnmpPolicyUsersUserIdEndpoint(parent string) SystemSnmpPolicyUsersUserIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/snmp_policy/users/:user_id"}, "/"), "/")
	return SystemSnmpPolicyUsersUserIdEndpoint{
		Path: path,
	}
}

type UsersUserIdEndpoint struct {
	Path  string
	Roles UsersUserIdRolesEndpoint
}

func NewUsersUserIdEndpoint(parent string) UsersUserIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "users/:user_id"}, "/"), "/")
	return UsersUserIdEndpoint{
		Path:  path,
		Roles: NewUsersUserIdRolesEndpoint(path),
	}
}

type AccessNetworkIpPoolsPoolNameEndpoint struct {
	Path         string
	NetworkPaths AccessNetworkIpPoolsPoolNameNetworkPathsEndpoint
}

func NewAccessNetworkIpPoolsPoolNameEndpoint(parent string) AccessNetworkIpPoolsPoolNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "access_network_ip_pools/:pool_name"}, "/"), "/")
	return AccessNetworkIpPoolsPoolNameEndpoint{
		Path:         path,
		NetworkPaths: NewAccessNetworkIpPoolsPoolNameNetworkPathsEndpoint(path),
	}
}

type ApiEndpoint struct {
	Path string
}

func NewApiEndpoint(parent string) ApiEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "api"}, "/"), "/")
	return ApiEndpoint{
		Path: path,
	}
}

type SystemDnsSearchDomainsEndpoint struct {
	Path string
}

func NewSystemDnsSearchDomainsEndpoint(parent string) SystemDnsSearchDomainsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/dns/search_domains"}, "/"), "/")
	return SystemDnsSearchDomainsEndpoint{
		Path: path,
	}
}

type AppInstancesIdEndpoint struct {
	Path             string
	Metadata         AppInstancesIdMetadataEndpoint
	SnapshotPolicies AppInstancesIdSnapshotPoliciesEndpoint
	Snapshots        AppInstancesIdSnapshotsEndpoint
	StorageInstances AppInstancesIdStorageInstancesEndpoint
}

func NewAppInstancesIdEndpoint(parent string) AppInstancesIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id"}, "/"), "/")
	return AppInstancesIdEndpoint{
		Path:             path,
		Metadata:         NewAppInstancesIdMetadataEndpoint(path),
		SnapshotPolicies: NewAppInstancesIdSnapshotPoliciesEndpoint(path),
		Snapshots:        NewAppInstancesIdSnapshotsEndpoint(path),
		StorageInstances: NewAppInstancesIdStorageInstancesEndpoint(path),
	}
}

type MetricsIoMetricLatestEndpoint struct {
	Path string
}

func NewMetricsIoMetricLatestEndpoint(parent string) MetricsIoMetricLatestEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "metrics/io/:metric/latest"}, "/"), "/")
	return MetricsIoMetricLatestEndpoint{
		Path: path,
	}
}

type AccessNetworkIpPoolsEndpoint struct {
	Path     string
	PoolName AccessNetworkIpPoolsPoolNameEndpoint
}

func (ep AccessNetworkIpPoolsEndpoint) Create(bodyp ...string) (AccessNetworkIpPoolEntity, error) {
	var en AccessNetworkIpPoolEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AccessNetworkIpPoolsEndpoint) List(queryp ...string) ([]AccessNetworkIpPoolEntity, error) {
	var ens []AccessNetworkIpPoolEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAccessNetworkIpPoolsEndpoint(parent string) AccessNetworkIpPoolsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "access_network_ip_pools"}, "/"), "/")
	return AccessNetworkIpPoolsEndpoint{
		Path:     path,
		PoolName: NewAccessNetworkIpPoolsPoolNameEndpoint(path),
	}
}

type AccessNetworkIpPoolsPoolNameNetworkPathsEndpoint struct {
	Path     string
	PathName AccessNetworkIpPoolsPoolNameNetworkPathsPathNameEndpoint
}

func NewAccessNetworkIpPoolsPoolNameNetworkPathsEndpoint(parent string) AccessNetworkIpPoolsPoolNameNetworkPathsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "access_network_ip_pools/:pool_name/network_paths"}, "/"), "/")
	return AccessNetworkIpPoolsPoolNameNetworkPathsEndpoint{
		Path:     path,
		PathName: NewAccessNetworkIpPoolsPoolNameNetworkPathsPathNameEndpoint(path),
	}
}

type InitConfigEndpoint struct {
	Path string
}

func NewInitConfigEndpoint(parent string) InitConfigEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "init/config"}, "/"), "/")
	return InitConfigEndpoint{
		Path: path,
	}
}

type AppTemplatesNameSnapshotPoliciesEndpoint struct {
	Path string
}

func (ep AppTemplatesNameSnapshotPoliciesEndpoint) Create(bodyp ...string) (SnapshotPolicyEntity, error) {
	var en SnapshotPolicyEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppTemplatesNameSnapshotPoliciesEndpoint) List(queryp ...string) ([]SnapshotPolicyEntity, error) {
	var ens []SnapshotPolicyEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppTemplatesNameSnapshotPoliciesEndpoint(parent string) AppTemplatesNameSnapshotPoliciesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:name/snapshot_policies"}, "/"), "/")
	return AppTemplatesNameSnapshotPoliciesEndpoint{
		Path: path,
	}
}

type PrivilegesEndpoint struct {
	Path string
}

func NewPrivilegesEndpoint(parent string) PrivilegesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "privileges"}, "/"), "/")
	return PrivilegesEndpoint{
		Path: path,
	}
}

type SystemNetworkMgmtVipNetworkPathsEndpoint struct {
	Path        string
	NetworkPath SystemNetworkMgmtVipNetworkPathsNetworkPathEndpoint
}

func NewSystemNetworkMgmtVipNetworkPathsEndpoint(parent string) SystemNetworkMgmtVipNetworkPathsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/mgmt_vip/network_paths"}, "/"), "/")
	return SystemNetworkMgmtVipNetworkPathsEndpoint{
		Path:        path,
		NetworkPath: NewSystemNetworkMgmtVipNetworkPathsNetworkPathEndpoint(path),
	}
}

type SystemNetworkMgmtVipNetworkPathsNetworkPathEndpoint struct {
	Path string
}

func NewSystemNetworkMgmtVipNetworkPathsNetworkPathEndpoint(parent string) SystemNetworkMgmtVipNetworkPathsNetworkPathEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/mgmt_vip/network_paths/:network_path"}, "/"), "/")
	return SystemNetworkMgmtVipNetworkPathsNetworkPathEndpoint{
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameEndpoint struct {
	Path              string
	PerformancePolicy AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNamePerformancePolicyEndpoint
	SnapshotPolicies  AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesEndpoint
	Snapshots         AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsEndpoint
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameEndpoint(parent string) AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/volumes/:volume_name"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameEndpoint{
		Path:              path,
		PerformancePolicy: NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNamePerformancePolicyEndpoint(path),
		SnapshotPolicies:  NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesEndpoint(path),
		Snapshots:         NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsEndpoint(path),
	}
}

type AuditLogsIdEndpoint struct {
	Path string
}

func NewAuditLogsIdEndpoint(parent string) AuditLogsIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "audit_logs/:id"}, "/"), "/")
	return AuditLogsIdEndpoint{
		Path: path,
	}
}

type MonitoringDestinationsDefaultEndpoint struct {
	Path string
}

func NewMonitoringDestinationsDefaultEndpoint(parent string) MonitoringDestinationsDefaultEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "monitoring/destinations/default"}, "/"), "/")
	return MonitoringDestinationsDefaultEndpoint{
		Path: path,
	}
}

type EventsSystemEndpoint struct {
	Path string
}

func (ep EventsSystemEndpoint) Create(bodyp ...string) (SystemEntity, error) {
	var en SystemEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep EventsSystemEndpoint) List(queryp ...string) ([]SystemEntity, error) {
	var ens []SystemEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewEventsSystemEndpoint(parent string) EventsSystemEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "events/system"}, "/"), "/")
	return EventsSystemEndpoint{
		Path: path,
	}
}

type SystemDnsServersEndpoint struct {
	Path string
}

func NewSystemDnsServersEndpoint(parent string) SystemDnsServersEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/dns/servers"}, "/"), "/")
	return SystemDnsServersEndpoint{
		Path: path,
	}
}

type SystemDnsEndpoint struct {
	Path          string
	SearchDomains SystemDnsSearchDomainsEndpoint
	Servers       SystemDnsServersEndpoint
}

func (ep SystemDnsEndpoint) Create(bodyp ...string) (DnsEntity, error) {
	var en DnsEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep SystemDnsEndpoint) List(queryp ...string) ([]DnsEntity, error) {
	var ens []DnsEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewSystemDnsEndpoint(parent string) SystemDnsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/dns"}, "/"), "/")
	return SystemDnsEndpoint{
		Path:          path,
		SearchDomains: NewSystemDnsSearchDomainsEndpoint(path),
		Servers:       NewSystemDnsServersEndpoint(path),
	}
}

type SystemNetworkInternalNetworkNetworkPathsNetworkPathEndpoint struct {
	Path string
}

func NewSystemNetworkInternalNetworkNetworkPathsNetworkPathEndpoint(parent string) SystemNetworkInternalNetworkNetworkPathsNetworkPathEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/internal_network/network_paths/:network_path"}, "/"), "/")
	return SystemNetworkInternalNetworkNetworkPathsNetworkPathEndpoint{
		Path: path,
	}
}

type SystemNetworkMgmtVipEndpoint struct {
	Path         string
	NetworkPaths SystemNetworkMgmtVipNetworkPathsEndpoint
}

func NewSystemNetworkMgmtVipEndpoint(parent string) SystemNetworkMgmtVipEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/mgmt_vip"}, "/"), "/")
	return SystemNetworkMgmtVipEndpoint{
		Path:         path,
		NetworkPaths: NewSystemNetworkMgmtVipNetworkPathsEndpoint(path),
	}
}

type AppTemplatesNameStorageTemplatesEndpoint struct {
	Path string
}

func (ep AppTemplatesNameStorageTemplatesEndpoint) Create(bodyp ...string) (StorageTemplateEntity, error) {
	var en StorageTemplateEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppTemplatesNameStorageTemplatesEndpoint) List(queryp ...string) ([]StorageTemplateEntity, error) {
	var ens []StorageTemplateEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppTemplatesNameStorageTemplatesEndpoint(parent string) AppTemplatesNameStorageTemplatesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:name/storage_templates"}, "/"), "/")
	return AppTemplatesNameStorageTemplatesEndpoint{
		Path: path,
	}
}

type TenantsTenantPathEndpoint struct {
	Path string
}

func NewTenantsTenantPathEndpoint(parent string) TenantsTenantPathEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "tenants/:tenant_path"}, "/"), "/")
	return TenantsTenantPathEndpoint{
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyEndpoint struct {
	Path            string
	InitiatorGroups AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorGroupsEndpoint
	Initiators      AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorsEndpoint
}

func (ep AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyEndpoint) Create(bodyp ...string) (AclPolicyEntity, error) {
	var en AclPolicyEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyEndpoint) List(queryp ...string) ([]AclPolicyEntity, error) {
	var ens []AclPolicyEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameAclPolicyEndpoint(parent string) AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/acl_policy"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyEndpoint{
		Path:            path,
		InitiatorGroups: NewAppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorGroupsEndpoint(path),
		Initiators:      NewAppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorsEndpoint(path),
	}
}

type UpgradeAvailableEndpoint struct {
	Path string
}

func NewUpgradeAvailableEndpoint(parent string) UpgradeAvailableEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "upgrade/available"}, "/"), "/")
	return UpgradeAvailableEndpoint{
		Path: path,
	}
}

type StorageNodesUuidNicsIdEndpoint struct {
	Path string
}

func NewStorageNodesUuidNicsIdEndpoint(parent string) StorageNodesUuidNicsIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/nics/:id"}, "/"), "/")
	return StorageNodesUuidNicsIdEndpoint{
		Path: path,
	}
}

type AppInstancesIdSnapshotsTimestampEndpoint struct {
	Path string
}

func NewAppInstancesIdSnapshotsTimestampEndpoint(parent string) AppInstancesIdSnapshotsTimestampEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/snapshots/:timestamp"}, "/"), "/")
	return AppInstancesIdSnapshotsTimestampEndpoint{
		Path: path,
	}
}

type SystemVersionConfigEndpoint struct {
	Path string
}

func NewSystemVersionConfigEndpoint(parent string) SystemVersionConfigEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/version_config"}, "/"), "/")
	return SystemVersionConfigEndpoint{
		Path: path,
	}
}

type AppInstancesIdSnapshotPoliciesSnapshotPolicyNameEndpoint struct {
	Path string
}

func NewAppInstancesIdSnapshotPoliciesSnapshotPolicyNameEndpoint(parent string) AppInstancesIdSnapshotPoliciesSnapshotPolicyNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/snapshot_policies/:snapshot_policy_name"}, "/"), "/")
	return AppInstancesIdSnapshotPoliciesSnapshotPolicyNameEndpoint{
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorsEndpoint struct {
	Path string
}

func (ep AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorsEndpoint) Create(bodyp ...string) (InitiatorEntity, error) {
	var en InitiatorEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorsEndpoint) List(queryp ...string) ([]InitiatorEntity, error) {
	var ens []InitiatorEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorsEndpoint(parent string) AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/acl_policy/initiators"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorsEndpoint{
		Path: path,
	}
}

type EventsDebugEndpoint struct {
	Path string
}

func NewEventsDebugEndpoint(parent string) EventsDebugEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "events/debug"}, "/"), "/")
	return EventsDebugEndpoint{
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNamePerformancePolicyEndpoint struct {
	Path string
}

func (ep AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNamePerformancePolicyEndpoint) Create(bodyp ...string) (PerformancePolicyEntity, error) {
	var en PerformancePolicyEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNamePerformancePolicyEndpoint) List(queryp ...string) ([]PerformancePolicyEntity, error) {
	var ens []PerformancePolicyEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNamePerformancePolicyEndpoint(parent string) AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNamePerformancePolicyEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/volumes/:volume_name/performance_policy"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNamePerformancePolicyEndpoint{
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameAuthEndpoint struct {
	Path string
}

func (ep AppInstancesIdStorageInstancesStorageInstanceNameAuthEndpoint) Create(bodyp ...string) (AuthEntity, error) {
	var en AuthEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppInstancesIdStorageInstancesStorageInstanceNameAuthEndpoint) List(queryp ...string) ([]AuthEntity, error) {
	var ens []AuthEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameAuthEndpoint(parent string) AppInstancesIdStorageInstancesStorageInstanceNameAuthEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/auth"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameAuthEndpoint{
		Path: path,
	}
}

type UsersUserIdRolesEndpoint struct {
	Path string
}

func (ep UsersUserIdRolesEndpoint) Create(bodyp ...string) (RoleEntity, error) {
	var en RoleEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep UsersUserIdRolesEndpoint) List(queryp ...string) ([]RoleEntity, error) {
	var ens []RoleEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewUsersUserIdRolesEndpoint(parent string) UsersUserIdRolesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "users/:user_id/roles"}, "/"), "/")
	return UsersUserIdRolesEndpoint{
		Path: path,
	}
}

type TenantsEndpoint struct {
	Path       string
	TenantPath TenantsTenantPathEndpoint
}

func (ep TenantsEndpoint) Create(bodyp ...string) (TenantEntity, error) {
	var en TenantEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep TenantsEndpoint) List(queryp ...string) ([]TenantEntity, error) {
	var ens []TenantEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewTenantsEndpoint(parent string) TenantsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "tenants"}, "/"), "/")
	return TenantsEndpoint{
		Path:       path,
		TenantPath: NewTenantsTenantPathEndpoint(path),
	}
}

type InitiatorsEndpoint struct {
	Path string
	Id   InitiatorsIdEndpoint
}

func (ep InitiatorsEndpoint) Create(bodyp ...string) (InitiatorEntity, error) {
	var en InitiatorEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep InitiatorsEndpoint) List(queryp ...string) ([]InitiatorEntity, error) {
	var ens []InitiatorEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewInitiatorsEndpoint(parent string) InitiatorsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "initiators"}, "/"), "/")
	return InitiatorsEndpoint{
		Path: path,
		Id:   NewInitiatorsIdEndpoint(path),
	}
}

type StorageNodesUuidSubsystemStatesEndpoint struct {
	Path string
}

func NewStorageNodesUuidSubsystemStatesEndpoint(parent string) StorageNodesUuidSubsystemStatesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/subsystem_states"}, "/"), "/")
	return StorageNodesUuidSubsystemStatesEndpoint{
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesEndpoint struct {
	Path               string
	SnapshotPolicyName AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesSnapshotPolicyNameEndpoint
}

func (ep AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesEndpoint) Create(bodyp ...string) (SnapshotPolicyEntity, error) {
	var en SnapshotPolicyEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesEndpoint) List(queryp ...string) ([]SnapshotPolicyEntity, error) {
	var ens []SnapshotPolicyEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesEndpoint(parent string) AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/volumes/:volume_name/snapshot_policies"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesEndpoint{
		Path:               path,
		SnapshotPolicyName: NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesSnapshotPolicyNameEndpoint(path),
	}
}

type MetricsIoMetricEndpoint struct {
	Path   string
	Latest MetricsIoMetricLatestEndpoint
}

func NewMetricsIoMetricEndpoint(parent string) MetricsIoMetricEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "metrics/io/:metric"}, "/"), "/")
	return MetricsIoMetricEndpoint{
		Path:   path,
		Latest: NewMetricsIoMetricLatestEndpoint(path),
	}
}

type TimeEndpoint struct {
	Path string
}

func NewTimeEndpoint(parent string) TimeEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "time"}, "/"), "/")
	return TimeEndpoint{
		Path: path,
	}
}

type AppTemplatesNameEndpoint struct {
	Path             string
	SnapshotPolicies AppTemplatesNameSnapshotPoliciesEndpoint
	StorageTemplates AppTemplatesNameStorageTemplatesEndpoint
}

func NewAppTemplatesNameEndpoint(parent string) AppTemplatesNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:name"}, "/"), "/")
	return AppTemplatesNameEndpoint{
		Path:             path,
		SnapshotPolicies: NewAppTemplatesNameSnapshotPoliciesEndpoint(path),
		StorageTemplates: NewAppTemplatesNameStorageTemplatesEndpoint(path),
	}
}

type AuditLogsEndpoint struct {
	Path string
	Id   AuditLogsIdEndpoint
}

func (ep AuditLogsEndpoint) Create(bodyp ...string) (AuditLogEntity, error) {
	var en AuditLogEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AuditLogsEndpoint) List(queryp ...string) ([]AuditLogEntity, error) {
	var ens []AuditLogEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAuditLogsEndpoint(parent string) AuditLogsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "audit_logs"}, "/"), "/")
	return AuditLogsEndpoint{
		Path: path,
		Id:   NewAuditLogsIdEndpoint(path),
	}
}

type RolesEndpoint struct {
	Path   string
	RoleId RolesRoleIdEndpoint
}

func (ep RolesEndpoint) Create(bodyp ...string) (RoleEntity, error) {
	var en RoleEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep RolesEndpoint) List(queryp ...string) ([]RoleEntity, error) {
	var ens []RoleEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewRolesEndpoint(parent string) RolesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "roles"}, "/"), "/")
	return RolesEndpoint{
		Path:   path,
		RoleId: NewRolesRoleIdEndpoint(path),
	}
}

type AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameSnapshotPoliciesEndpoint struct {
	Path string
}

func (ep AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameSnapshotPoliciesEndpoint) Create(bodyp ...string) (SnapshotPolicyEntity, error) {
	var en SnapshotPolicyEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameSnapshotPoliciesEndpoint) List(queryp ...string) ([]SnapshotPolicyEntity, error) {
	var ens []SnapshotPolicyEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameSnapshotPoliciesEndpoint(parent string) AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameSnapshotPoliciesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:app_template_name/storage_templates/:storage_template_name/volume_templates/:volume_template_name/snapshot_policies"}, "/"), "/")
	return AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameSnapshotPoliciesEndpoint{
		Path: path,
	}
}

type SystemNetworkInternalNetworkEndpoint struct {
	Path         string
	NetworkPaths SystemNetworkInternalNetworkNetworkPathsEndpoint
}

func NewSystemNetworkInternalNetworkEndpoint(parent string) SystemNetworkInternalNetworkEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/internal_network"}, "/"), "/")
	return SystemNetworkInternalNetworkEndpoint{
		Path:         path,
		NetworkPaths: NewSystemNetworkInternalNetworkNetworkPathsEndpoint(path),
	}
}

type EventsUserEndpoint struct {
	Path string
}

func (ep EventsUserEndpoint) Create(bodyp ...string) (UserEntity, error) {
	var en UserEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep EventsUserEndpoint) List(queryp ...string) ([]UserEntity, error) {
	var ens []UserEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewEventsUserEndpoint(parent string) EventsUserEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "events/user"}, "/"), "/")
	return EventsUserEndpoint{
		Path: path,
	}
}

type InitiatorGroupsNameEndpoint struct {
	Path    string
	Members InitiatorGroupsNameMembersEndpoint
}

func NewInitiatorGroupsNameEndpoint(parent string) InitiatorGroupsNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "initiator_groups/:name"}, "/"), "/")
	return InitiatorGroupsNameEndpoint{
		Path:    path,
		Members: NewInitiatorGroupsNameMembersEndpoint(path),
	}
}

type FaultLogsEndpoint struct {
	Path string
	Id   FaultLogsIdEndpoint
}

func (ep FaultLogsEndpoint) Create(bodyp ...string) (FaultLogEntity, error) {
	var en FaultLogEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep FaultLogsEndpoint) List(queryp ...string) ([]FaultLogEntity, error) {
	var ens []FaultLogEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewFaultLogsEndpoint(parent string) FaultLogsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "fault_logs"}, "/"), "/")
	return FaultLogsEndpoint{
		Path: path,
		Id:   NewFaultLogsIdEndpoint(path),
	}
}

type StorageNodesUuidHddsIdEndpoint struct {
	Path string
}

func NewStorageNodesUuidHddsIdEndpoint(parent string) StorageNodesUuidHddsIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/hdds/:id"}, "/"), "/")
	return StorageNodesUuidHddsIdEndpoint{
		Path: path,
	}
}

type SystemHttpProxyEndpoint struct {
	Path string
}

func (ep SystemHttpProxyEndpoint) Create(bodyp ...string) (HttpProxyEntity, error) {
	var en HttpProxyEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep SystemHttpProxyEndpoint) List(queryp ...string) ([]HttpProxyEntity, error) {
	var ens []HttpProxyEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewSystemHttpProxyEndpoint(parent string) SystemHttpProxyEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/http_proxy"}, "/"), "/")
	return SystemHttpProxyEndpoint{
		Path: path,
	}
}

type StorageNodesUuidEndpoint struct {
	Path            string
	BootDrives      StorageNodesUuidBootDrivesEndpoint
	FlashDevices    StorageNodesUuidFlashDevicesEndpoint
	Hdds            StorageNodesUuidHddsEndpoint
	Nics            StorageNodesUuidNicsEndpoint
	SubsystemStates StorageNodesUuidSubsystemStatesEndpoint
}

func NewStorageNodesUuidEndpoint(parent string) StorageNodesUuidEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid"}, "/"), "/")
	return StorageNodesUuidEndpoint{
		Path:            path,
		BootDrives:      NewStorageNodesUuidBootDrivesEndpoint(path),
		FlashDevices:    NewStorageNodesUuidFlashDevicesEndpoint(path),
		Hdds:            NewStorageNodesUuidHddsEndpoint(path),
		Nics:            NewStorageNodesUuidNicsEndpoint(path),
		SubsystemStates: NewStorageNodesUuidSubsystemStatesEndpoint(path),
	}
}

type EventsUuidEndpoint struct {
	Path string
}

func NewEventsUuidEndpoint(parent string) EventsUuidEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "events/:uuid"}, "/"), "/")
	return EventsUuidEndpoint{
		Path: path,
	}
}

type SystemEndpoint struct {
	Path          string
	Dns           SystemDnsEndpoint
	HttpProxy     SystemHttpProxyEndpoint
	Network       SystemNetworkEndpoint
	NtpServers    SystemNtpServersEndpoint
	SnmpPolicy    SystemSnmpPolicyEndpoint
	VersionConfig SystemVersionConfigEndpoint
}

func (ep SystemEndpoint) Create(bodyp ...string) (SystemEntity, error) {
	var en SystemEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep SystemEndpoint) List(queryp ...string) ([]SystemEntity, error) {
	var ens []SystemEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewSystemEndpoint(parent string) SystemEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system"}, "/"), "/")
	return SystemEndpoint{
		Path:          path,
		Dns:           NewSystemDnsEndpoint(path),
		HttpProxy:     NewSystemHttpProxyEndpoint(path),
		Network:       NewSystemNetworkEndpoint(path),
		NtpServers:    NewSystemNtpServersEndpoint(path),
		SnmpPolicy:    NewSystemSnmpPolicyEndpoint(path),
		VersionConfig: NewSystemVersionConfigEndpoint(path),
	}
}

type StorageNodesEndpoint struct {
	Path string
	Uuid StorageNodesUuidEndpoint
}

func (ep StorageNodesEndpoint) Create(bodyp ...string) (StorageNodeEntity, error) {
	var en StorageNodeEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep StorageNodesEndpoint) List(queryp ...string) ([]StorageNodeEntity, error) {
	var ens []StorageNodeEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewStorageNodesEndpoint(parent string) StorageNodesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes"}, "/"), "/")
	return StorageNodesEndpoint{
		Path: path,
		Uuid: NewStorageNodesUuidEndpoint(path),
	}
}

type AppTemplatesEndpoint struct {
	Path string
	Name AppTemplatesNameEndpoint
}

func (ep AppTemplatesEndpoint) Create(bodyp ...string) (AppTemplateEntity, error) {
	var en AppTemplateEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppTemplatesEndpoint) List(queryp ...string) ([]AppTemplateEntity, error) {
	var ens []AppTemplateEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppTemplatesEndpoint(parent string) AppTemplatesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates"}, "/"), "/")
	return AppTemplatesEndpoint{
		Path: path,
		Name: NewAppTemplatesNameEndpoint(path),
	}
}

type UpgradeEndpoint struct {
	Path      string
	Available UpgradeAvailableEndpoint
}

func NewUpgradeEndpoint(parent string) UpgradeEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "upgrade"}, "/"), "/")
	return UpgradeEndpoint{
		Path:      path,
		Available: NewUpgradeAvailableEndpoint(path),
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsEndpoint struct {
	Path      string
	Timestamp AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsTimestampEndpoint
}

func (ep AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsEndpoint) Create(bodyp ...string) (SnapshotEntity, error) {
	var en SnapshotEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsEndpoint) List(queryp ...string) ([]SnapshotEntity, error) {
	var ens []SnapshotEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsEndpoint(parent string) AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/volumes/:volume_name/snapshots"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsEndpoint{
		Path:      path,
		Timestamp: NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsTimestampEndpoint(path),
	}
}

type InitiatorsIdEndpoint struct {
	Path string
}

func NewInitiatorsIdEndpoint(parent string) InitiatorsIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "initiators/:id"}, "/"), "/")
	return InitiatorsIdEndpoint{
		Path: path,
	}
}

type SystemNetworkAccessVipNetworkPathsEndpoint struct {
	Path        string
	NetworkPath SystemNetworkAccessVipNetworkPathsNetworkPathEndpoint
}

func NewSystemNetworkAccessVipNetworkPathsEndpoint(parent string) SystemNetworkAccessVipNetworkPathsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/access_vip/network_paths"}, "/"), "/")
	return SystemNetworkAccessVipNetworkPathsEndpoint{
		Path:        path,
		NetworkPath: NewSystemNetworkAccessVipNetworkPathsNetworkPathEndpoint(path),
	}
}

type PolicyadmEndpoint struct {
	Path string
}

func NewPolicyadmEndpoint(parent string) PolicyadmEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "policyadm"}, "/"), "/")
	return PolicyadmEndpoint{
		Path: path,
	}
}

type EventLogsIdEndpoint struct {
	Path string
}

func NewEventLogsIdEndpoint(parent string) EventLogsIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "event_logs/:id"}, "/"), "/")
	return EventLogsIdEndpoint{
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorGroupsEndpoint struct {
	Path string
}

func (ep AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorGroupsEndpoint) Create(bodyp ...string) (InitiatorGroupEntity, error) {
	var en InitiatorGroupEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorGroupsEndpoint) List(queryp ...string) ([]InitiatorGroupEntity, error) {
	var ens []InitiatorGroupEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorGroupsEndpoint(parent string) AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorGroupsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/acl_policy/initiator_groups"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorGroupsEndpoint{
		Path: path,
	}
}

type MonitoringAlertsEndpoint struct {
	Path string
}

func NewMonitoringAlertsEndpoint(parent string) MonitoringAlertsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "monitoring/alerts"}, "/"), "/")
	return MonitoringAlertsEndpoint{
		Path: path,
	}
}

type StorageNodesUuidFlashDevicesIdEndpoint struct {
	Path string
}

func NewStorageNodesUuidFlashDevicesIdEndpoint(parent string) StorageNodesUuidFlashDevicesIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/flash_devices/:id"}, "/"), "/")
	return StorageNodesUuidFlashDevicesIdEndpoint{
		Path: path,
	}
}

type MetricsHwMetricLatestEndpoint struct {
	Path string
}

func NewMetricsHwMetricLatestEndpoint(parent string) MetricsHwMetricLatestEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "metrics/hw/:metric/latest"}, "/"), "/")
	return MetricsHwMetricLatestEndpoint{
		Path: path,
	}
}

type FaultLogsIdEndpoint struct {
	Path string
}

func NewFaultLogsIdEndpoint(parent string) FaultLogsIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "fault_logs/:id"}, "/"), "/")
	return FaultLogsIdEndpoint{
		Path: path,
	}
}

type InitiatorGroupsEndpoint struct {
	Path string
	Name InitiatorGroupsNameEndpoint
}

func (ep InitiatorGroupsEndpoint) Create(bodyp ...string) (InitiatorGroupEntity, error) {
	var en InitiatorGroupEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep InitiatorGroupsEndpoint) List(queryp ...string) ([]InitiatorGroupEntity, error) {
	var ens []InitiatorGroupEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewInitiatorGroupsEndpoint(parent string) InitiatorGroupsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "initiator_groups"}, "/"), "/")
	return InitiatorGroupsEndpoint{
		Path: path,
		Name: NewInitiatorGroupsNameEndpoint(path),
	}
}

type SystemNetworkMappingEndpoint struct {
	Path string
}

func NewSystemNetworkMappingEndpoint(parent string) SystemNetworkMappingEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/mapping"}, "/"), "/")
	return SystemNetworkMappingEndpoint{
		Path: path,
	}
}

type AppTemplatesAtNameStorageTemplatesStNameVolumeTemplatesVtNameSnapshotPoliciesSnapshotPolicyNameEndpoint struct {
	Path string
}

func NewAppTemplatesAtNameStorageTemplatesStNameVolumeTemplatesVtNameSnapshotPoliciesSnapshotPolicyNameEndpoint(parent string) AppTemplatesAtNameStorageTemplatesStNameVolumeTemplatesVtNameSnapshotPoliciesSnapshotPolicyNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:at_name/storage_templates/:st_name/volume_templates/:vt_name/snapshot_policies/:snapshot_policy_name"}, "/"), "/")
	return AppTemplatesAtNameStorageTemplatesStNameVolumeTemplatesVtNameSnapshotPoliciesSnapshotPolicyNameEndpoint{
		Path: path,
	}
}

type AppInstancesIdSnapshotsEndpoint struct {
	Path      string
	Timestamp AppInstancesIdSnapshotsTimestampEndpoint
}

func (ep AppInstancesIdSnapshotsEndpoint) Create(bodyp ...string) (SnapshotEntity, error) {
	var en SnapshotEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppInstancesIdSnapshotsEndpoint) List(queryp ...string) ([]SnapshotEntity, error) {
	var ens []SnapshotEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppInstancesIdSnapshotsEndpoint(parent string) AppInstancesIdSnapshotsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/snapshots"}, "/"), "/")
	return AppInstancesIdSnapshotsEndpoint{
		Path:      path,
		Timestamp: NewAppInstancesIdSnapshotsTimestampEndpoint(path),
	}
}

type SystemSnmpPolicyEndpoint struct {
	Path  string
	Users SystemSnmpPolicyUsersEndpoint
}

func (ep SystemSnmpPolicyEndpoint) Create(bodyp ...string) (SnmpPolicyEntity, error) {
	var en SnmpPolicyEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep SystemSnmpPolicyEndpoint) List(queryp ...string) ([]SnmpPolicyEntity, error) {
	var ens []SnmpPolicyEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewSystemSnmpPolicyEndpoint(parent string) SystemSnmpPolicyEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/snmp_policy"}, "/"), "/")
	return SystemSnmpPolicyEndpoint{
		Path:  path,
		Users: NewSystemSnmpPolicyUsersEndpoint(path),
	}
}

type UserinfoEndpoint struct {
	Path string
}

func NewUserinfoEndpoint(parent string) UserinfoEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "userinfo"}, "/"), "/")
	return UserinfoEndpoint{
		Path: path,
	}
}

type AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesEndpoint struct {
	Path               string
	VolumeTemplateName AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameEndpoint
}

func (ep AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesEndpoint) Create(bodyp ...string) (VolumeTemplateEntity, error) {
	var en VolumeTemplateEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesEndpoint) List(queryp ...string) ([]VolumeTemplateEntity, error) {
	var ens []VolumeTemplateEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesEndpoint(parent string) AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:app_template_name/storage_templates/:storage_template_name/volume_templates"}, "/"), "/")
	return AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesEndpoint{
		Path:               path,
		VolumeTemplateName: NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameEndpoint(path),
	}
}

type HealthAttrsEndpoint struct {
	Path string
}

func NewHealthAttrsEndpoint(parent string) HealthAttrsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "health_attrs"}, "/"), "/")
	return HealthAttrsEndpoint{
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameEndpoint struct {
	Path      string
	AclPolicy AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyEndpoint
	Auth      AppInstancesIdStorageInstancesStorageInstanceNameAuthEndpoint
	Volumes   AppInstancesIdStorageInstancesStorageInstanceNameVolumesEndpoint
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameEndpoint(parent string) AppInstancesIdStorageInstancesStorageInstanceNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameEndpoint{
		Path:      path,
		AclPolicy: NewAppInstancesIdStorageInstancesStorageInstanceNameAclPolicyEndpoint(path),
		Auth:      NewAppInstancesIdStorageInstancesStorageInstanceNameAuthEndpoint(path),
		Volumes:   NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesEndpoint(path),
	}
}

type AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameAuthEndpoint struct {
	Path string
}

func (ep AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameAuthEndpoint) Create(bodyp ...string) (AuthEntity, error) {
	var en AuthEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameAuthEndpoint) List(queryp ...string) ([]AuthEntity, error) {
	var ens []AuthEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameAuthEndpoint(parent string) AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameAuthEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:app_template_name/storage_templates/:storage_template_name/auth"}, "/"), "/")
	return AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameAuthEndpoint{
		Path: path,
	}
}

type StorageNodesUuidHddsEndpoint struct {
	Path string
	Id   StorageNodesUuidHddsIdEndpoint
}

func (ep StorageNodesUuidHddsEndpoint) Create(bodyp ...string) (HddEntity, error) {
	var en HddEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep StorageNodesUuidHddsEndpoint) List(queryp ...string) ([]HddEntity, error) {
	var ens []HddEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewStorageNodesUuidHddsEndpoint(parent string) StorageNodesUuidHddsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/hdds"}, "/"), "/")
	return StorageNodesUuidHddsEndpoint{
		Path: path,
		Id:   NewStorageNodesUuidHddsIdEndpoint(path),
	}
}

type AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameEndpoint struct {
	Path            string
	Auth            AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameAuthEndpoint
	VolumeTemplates AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesEndpoint
}

func NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameEndpoint(parent string) AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:app_template_name/storage_templates/:storage_template_name"}, "/"), "/")
	return AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameEndpoint{
		Path:            path,
		Auth:            NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameAuthEndpoint(path),
		VolumeTemplates: NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesEndpoint(path),
	}
}

type MonitoringPoliciesDefaultEndpoint struct {
	Path string
}

func NewMonitoringPoliciesDefaultEndpoint(parent string) MonitoringPoliciesDefaultEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "monitoring/policies/default"}, "/"), "/")
	return MonitoringPoliciesDefaultEndpoint{
		Path: path,
	}
}

type SystemNetworkAccessVipEndpoint struct {
	Path         string
	NetworkPaths SystemNetworkAccessVipNetworkPathsEndpoint
}

func NewSystemNetworkAccessVipEndpoint(parent string) SystemNetworkAccessVipEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/access_vip"}, "/"), "/")
	return SystemNetworkAccessVipEndpoint{
		Path:         path,
		NetworkPaths: NewSystemNetworkAccessVipNetworkPathsEndpoint(path),
	}
}

type AppInstancesIdSnapshotPoliciesEndpoint struct {
	Path               string
	SnapshotPolicyName AppInstancesIdSnapshotPoliciesSnapshotPolicyNameEndpoint
}

func (ep AppInstancesIdSnapshotPoliciesEndpoint) Create(bodyp ...string) (SnapshotPolicyEntity, error) {
	var en SnapshotPolicyEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppInstancesIdSnapshotPoliciesEndpoint) List(queryp ...string) ([]SnapshotPolicyEntity, error) {
	var ens []SnapshotPolicyEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppInstancesIdSnapshotPoliciesEndpoint(parent string) AppInstancesIdSnapshotPoliciesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/snapshot_policies"}, "/"), "/")
	return AppInstancesIdSnapshotPoliciesEndpoint{
		Path:               path,
		SnapshotPolicyName: NewAppInstancesIdSnapshotPoliciesSnapshotPolicyNameEndpoint(path),
	}
}

type AppInstancesAiIdStorageInstancesSiIdMetadataEndpoint struct {
	Path string
}

func NewAppInstancesAiIdStorageInstancesSiIdMetadataEndpoint(parent string) AppInstancesAiIdStorageInstancesSiIdMetadataEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:ai_id/storage_instances/:si_id/metadata"}, "/"), "/")
	return AppInstancesAiIdStorageInstancesSiIdMetadataEndpoint{
		Path: path,
	}
}

type AppInstancesIdStorageInstancesEndpoint struct {
	Path                string
	StorageInstanceName AppInstancesIdStorageInstancesStorageInstanceNameEndpoint
}

func (ep AppInstancesIdStorageInstancesEndpoint) Create(bodyp ...string) (StorageInstanceEntity, error) {
	var en StorageInstanceEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppInstancesIdStorageInstancesEndpoint) List(queryp ...string) ([]StorageInstanceEntity, error) {
	var ens []StorageInstanceEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesEndpoint(parent string) AppInstancesIdStorageInstancesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances"}, "/"), "/")
	return AppInstancesIdStorageInstancesEndpoint{
		Path:                path,
		StorageInstanceName: NewAppInstancesIdStorageInstancesStorageInstanceNameEndpoint(path),
	}
}

type StorageNodesUuidBootDrivesIdEndpoint struct {
	Path string
}

func NewStorageNodesUuidBootDrivesIdEndpoint(parent string) StorageNodesUuidBootDrivesIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/boot_drives/:id"}, "/"), "/")
	return StorageNodesUuidBootDrivesIdEndpoint{
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesSnapshotPolicyNameEndpoint struct {
	Path string
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesSnapshotPolicyNameEndpoint(parent string) AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesSnapshotPolicyNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/volumes/:volume_name/snapshot_policies/:snapshot_policy_name"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesSnapshotPolicyNameEndpoint{
		Path: path,
	}
}

type SystemNetworkEndpoint struct {
	Path            string
	AccessVip       SystemNetworkAccessVipEndpoint
	InternalNetwork SystemNetworkInternalNetworkEndpoint
	Mapping         SystemNetworkMappingEndpoint
	MgmtVip         SystemNetworkMgmtVipEndpoint
}

func (ep SystemNetworkEndpoint) Create(bodyp ...string) (NetworkEntity, error) {
	var en NetworkEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep SystemNetworkEndpoint) List(queryp ...string) ([]NetworkEntity, error) {
	var ens []NetworkEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewSystemNetworkEndpoint(parent string) SystemNetworkEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network"}, "/"), "/")
	return SystemNetworkEndpoint{
		Path:            path,
		AccessVip:       NewSystemNetworkAccessVipEndpoint(path),
		InternalNetwork: NewSystemNetworkInternalNetworkEndpoint(path),
		Mapping:         NewSystemNetworkMappingEndpoint(path),
		MgmtVip:         NewSystemNetworkMgmtVipEndpoint(path),
	}
}

type UsersEndpoint struct {
	Path   string
	UserId UsersUserIdEndpoint
}

func (ep UsersEndpoint) Create(bodyp ...string) (UserEntity, error) {
	var en UserEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep UsersEndpoint) List(queryp ...string) ([]UserEntity, error) {
	var ens []UserEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewUsersEndpoint(parent string) UsersEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "users"}, "/"), "/")
	return UsersEndpoint{
		Path:   path,
		UserId: NewUsersUserIdEndpoint(path),
	}
}

type InitiatorGroupsNameMembersEndpoint struct {
	Path string
}

func NewInitiatorGroupsNameMembersEndpoint(parent string) InitiatorGroupsNameMembersEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "initiator_groups/:name/members"}, "/"), "/")
	return InitiatorGroupsNameMembersEndpoint{
		Path: path,
	}
}

type StorageNodesUuidBootDrivesEndpoint struct {
	Path string
	Id   StorageNodesUuidBootDrivesIdEndpoint
}

func NewStorageNodesUuidBootDrivesEndpoint(parent string) StorageNodesUuidBootDrivesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/boot_drives"}, "/"), "/")
	return StorageNodesUuidBootDrivesEndpoint{
		Path: path,
		Id:   NewStorageNodesUuidBootDrivesIdEndpoint(path),
	}
}

type SystemNtpServersEndpoint struct {
	Path string
}

func (ep SystemNtpServersEndpoint) Create(bodyp ...string) (NtpServerEntity, error) {
	var en NtpServerEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep SystemNtpServersEndpoint) List(queryp ...string) ([]NtpServerEntity, error) {
	var ens []NtpServerEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewSystemNtpServersEndpoint(parent string) SystemNtpServersEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/ntp_servers"}, "/"), "/")
	return SystemNtpServersEndpoint{
		Path: path,
	}
}

type SystemNetworkInternalNetworkNetworkPathsEndpoint struct {
	Path        string
	NetworkPath SystemNetworkInternalNetworkNetworkPathsNetworkPathEndpoint
}

func NewSystemNetworkInternalNetworkNetworkPathsEndpoint(parent string) SystemNetworkInternalNetworkNetworkPathsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/internal_network/network_paths"}, "/"), "/")
	return SystemNetworkInternalNetworkNetworkPathsEndpoint{
		Path:        path,
		NetworkPath: NewSystemNetworkInternalNetworkNetworkPathsNetworkPathEndpoint(path),
	}
}

type AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameEndpoint struct {
	Path              string
	PerformancePolicy AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNamePerformancePolicyEndpoint
	SnapshotPolicies  AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameSnapshotPoliciesEndpoint
}

func NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameEndpoint(parent string) AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:app_template_name/storage_templates/:storage_template_name/volume_templates/:volume_template_name"}, "/"), "/")
	return AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameEndpoint{
		Path:              path,
		PerformancePolicy: NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNamePerformancePolicyEndpoint(path),
		SnapshotPolicies:  NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameSnapshotPoliciesEndpoint(path),
	}
}

type AppInstancesEndpoint struct {
	Path string
	Id   AppInstancesIdEndpoint
}

func (ep AppInstancesEndpoint) Create(bodyp ...string) (AppInstanceEntity, error) {
	var en AppInstanceEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppInstancesEndpoint) List(queryp ...string) ([]AppInstanceEntity, error) {
	var ens []AppInstanceEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppInstancesEndpoint(parent string) AppInstancesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances"}, "/"), "/")
	return AppInstancesEndpoint{
		Path: path,
		Id:   NewAppInstancesIdEndpoint(path),
	}
}

type AppInstancesIdMetadataEndpoint struct {
	Path string
}

func NewAppInstancesIdMetadataEndpoint(parent string) AppInstancesIdMetadataEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/metadata"}, "/"), "/")
	return AppInstancesIdMetadataEndpoint{
		Path: path,
	}
}

type StorageNodesUuidNicsEndpoint struct {
	Path string
	Id   StorageNodesUuidNicsIdEndpoint
}

func (ep StorageNodesUuidNicsEndpoint) Create(bodyp ...string) (NicEntity, error) {
	var en NicEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep StorageNodesUuidNicsEndpoint) List(queryp ...string) ([]NicEntity, error) {
	var ens []NicEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewStorageNodesUuidNicsEndpoint(parent string) StorageNodesUuidNicsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/nics"}, "/"), "/")
	return StorageNodesUuidNicsEndpoint{
		Path: path,
		Id:   NewStorageNodesUuidNicsIdEndpoint(path),
	}
}

type SystemNetworkAccessVipNetworkPathsNetworkPathEndpoint struct {
	Path string
}

func NewSystemNetworkAccessVipNetworkPathsNetworkPathEndpoint(parent string) SystemNetworkAccessVipNetworkPathsNetworkPathEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/access_vip/network_paths/:network_path"}, "/"), "/")
	return SystemNetworkAccessVipNetworkPathsNetworkPathEndpoint{
		Path: path,
	}
}

type RolesRoleIdEndpoint struct {
	Path string
}

func NewRolesRoleIdEndpoint(parent string) RolesRoleIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "roles/:role_id"}, "/"), "/")
	return RolesRoleIdEndpoint{
		Path: path,
	}
}

type AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNamePerformancePolicyEndpoint struct {
	Path string
}

func (ep AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNamePerformancePolicyEndpoint) Create(bodyp ...string) (PerformancePolicyEntity, error) {
	var en PerformancePolicyEntity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &en)
	if err != nil {
		panic(err)
	}
	return en, nil
}
func (ep AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNamePerformancePolicyEndpoint) List(queryp ...string) ([]PerformancePolicyEntity, error) {
	var ens []PerformancePolicyEntity
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	return ens, nil
}

func NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNamePerformancePolicyEndpoint(parent string) AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNamePerformancePolicyEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:app_template_name/storage_templates/:storage_template_name/volume_templates/:volume_template_name/performance_policy"}, "/"), "/")
	return AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNamePerformancePolicyEndpoint{
		Path: path,
	}
}

type AccessNetworkIpPoolsPoolNameNetworkPathsPathNameEndpoint struct {
	Path string
}

func NewAccessNetworkIpPoolsPoolNameNetworkPathsPathNameEndpoint(parent string) AccessNetworkIpPoolsPoolNameNetworkPathsPathNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "access_network_ip_pools/:pool_name/network_paths/:path_name"}, "/"), "/")
	return AccessNetworkIpPoolsPoolNameNetworkPathsPathNameEndpoint{
		Path: path,
	}
}

type AppTemplatesAppTemplateNameSnapshotPoliciesSnapshotPolicyNameEndpoint struct {
	Path string
}

func NewAppTemplatesAppTemplateNameSnapshotPoliciesSnapshotPolicyNameEndpoint(parent string) AppTemplatesAppTemplateNameSnapshotPoliciesSnapshotPolicyNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:app_template_name/snapshot_policies/:snapshot_policy_name"}, "/"), "/")
	return AppTemplatesAppTemplateNameSnapshotPoliciesSnapshotPolicyNameEndpoint{
		Path: path,
	}
}