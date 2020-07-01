package ovnbindings

import (
	goovn "github.com/ebay/go-ovn"
)

type OVNInterface interface {
	// Get logical switch by name
	LSGet(ls string) ([]*goovn.LogicalSwitch, error)
	// Create ls named SWITCH
	LSAdd(ls string) (*goovn.OvnCommand, error)
	// Del ls and all its ports
	LSDel(ls string) (*goovn.OvnCommand, error)
	// Get all logical switches
	LSList() ([]*goovn.LogicalSwitch, error)
	// Add external_ids to logical switch
	LSExtIdsAdd(ls string, external_ids map[string]string) (*goovn.OvnCommand, error)
	// Del external_ids from logical_switch
	LSExtIdsDel(ls string, external_ids map[string]string) (*goovn.OvnCommand, error)
	// Link logical switch to router
	LinkSwitchToRouter(lsw, lsp, lr, lrp, lrpMac string, networks []string, externalIds map[string]string) (*goovn.OvnCommand, error)

	// Get logical switch port by name
	LSPGet(lsp string) (*goovn.LogicalSwitchPort, error)
	// Add logical port PORT on SWITCH
	LSPAdd(ls string, lsp string) (*goovn.OvnCommand, error)
	// Delete PORT from its attached switch
	LSPDel(lsp string) (*goovn.OvnCommand, error)
	// Set addressset per lport
	LSPSetAddress(lsp string, addresses ...string) (*goovn.OvnCommand, error)
	// Set port security per lport
	LSPSetPortSecurity(lsp string, security ...string) (*goovn.OvnCommand, error)
	// Get all lport by lswitch
	LSPList(ls string) ([]*goovn.LogicalSwitchPort, error)

	// Add LB to LSW
	LSLBAdd(ls string, lb string) (*goovn.OvnCommand, error)
	// Delete LB from LSW
	LSLBDel(ls string, lb string) (*goovn.OvnCommand, error)
	// List Load balancers for a LSW
	LSLBList(ls string) ([]*goovn.LoadBalancer, error)

	// Add ACL
	ACLAdd(ls, direct, match, action string, priority int, external_ids map[string]string, logflag bool, meter string, severity string) (*goovn.OvnCommand, error)
	// Delete acl
	ACLDel(ls, direct, match string, priority int, external_ids map[string]string) (*goovn.OvnCommand, error)
	// Get all acl by lswitch
	ACLList(ls string) ([]*goovn.ACL, error)

	// Get AS
	ASGet(name string) (*goovn.AddressSet, error)
	// Update address set
	ASUpdate(name string, addrs []string, external_ids map[string]string) (*goovn.OvnCommand, error)
	// Add addressset
	ASAdd(name string, addrs []string, external_ids map[string]string) (*goovn.OvnCommand, error)
	// Delete addressset
	ASDel(name string) (*goovn.OvnCommand, error)
	// Get all AS
	ASList() ([]*goovn.AddressSet, error)

	// Get LR with given name
	LRGet(name string) ([]*goovn.LogicalRouter, error)
	// Add LR with given name
	LRAdd(name string, external_ids map[string]string) (*goovn.OvnCommand, error)
	// Delete LR with given name
	LRDel(name string) (*goovn.OvnCommand, error)
	// Get LRs
	LRList() ([]*goovn.LogicalRouter, error)

	// Add LRP with given name on given lr
	LRPAdd(lr string, lrp string, mac string, network []string, peer string, external_ids map[string]string) (*goovn.OvnCommand, error)
	// Delete LRP with given name on given lr
	LRPDel(lr string, lrp string) (*goovn.OvnCommand, error)
	// Get all lrp by lr
	LRPList(lr string) ([]*goovn.LogicalRouterPort, error)

	// Add LRSR with given ip_prefix on given lr
	LRSRAdd(lr string, ip_prefix string, nexthop string, output_port []string, policy []string, external_ids map[string]string) (*goovn.OvnCommand, error)
	// Delete LRSR with given ip_prefix on given lr
	LRSRDel(lr string, ip_prefix string, nexthop, policy, outputPort *string) (*goovn.OvnCommand, error)
	// Get all LRSRs by lr
	LRSRList(lr string) ([]*goovn.LogicalRouterStaticRoute, error)

	// Add LB to LR
	LRLBAdd(lr string, lb string) (*goovn.OvnCommand, error)
	// Delete LB from LR
	LRLBDel(lr string, lb string) (*goovn.OvnCommand, error)
	// List Load balancers for a LR
	LRLBList(lr string) ([]*goovn.LoadBalancer, error)

	// Get LB with given name
	LBGet(name string) ([]*goovn.LoadBalancer, error)
	// Add LB
	LBAdd(name string, vipPort string, protocol string, addrs []string) (*goovn.OvnCommand, error)
	// Delete LB with given name
	LBDel(name string) (*goovn.OvnCommand, error)
	// Update existing LB
	LBUpdate(name string, vipPort string, protocol string, addrs []string) (*goovn.OvnCommand, error)

	// Set dhcp4_options uuid on lsp
	LSPSetDHCPv4Options(lsp string, options string) (*goovn.OvnCommand, error)
	// Get dhcp4_options from lsp
	LSPGetDHCPv4Options(lsp string) (*goovn.DHCPOptions, error)
	// Set dhcp6_options uuid on lsp
	LSPSetDHCPv6Options(lsp string, options string) (*goovn.OvnCommand, error)
	// Get dhcp6_options from lsp
	LSPGetDHCPv6Options(lsp string) (*goovn.DHCPOptions, error)
	// Set options in LSP
	LSPSetOptions(lsp string, options map[string]string) (*goovn.OvnCommand, error)
	// Get options from LSP
	LSPGetOptions(lsp string) (map[string]string, error)
	// Set dynamic addresses in LSP
	LSPSetDynamicAddresses(lsp string, address string) (*goovn.OvnCommand, error)
	// Get dynamic addresses from LSP
	LSPGetDynamicAddresses(lsp string) (string, error)
	// Set external_ids for LSP
	LSPSetExternalIds(lsp string, external_ids map[string]string) (*goovn.OvnCommand, error)
	// Get external_ids from LSP
	LSPGetExternalIds(lsp string) (map[string]string, error)
	// Add dhcp options for cidr and provided external_ids
	DHCPOptionsAdd(cidr string, options map[string]string, external_ids map[string]string) (*goovn.OvnCommand, error)
	// Set dhcp options and set external_ids for specific uuid
	DHCPOptionsSet(uuid string, options map[string]string, external_ids map[string]string) (*goovn.OvnCommand, error)
	// Del dhcp options via provided external_ids
	DHCPOptionsDel(uuid string) (*goovn.OvnCommand, error)
	// Get single dhcp via provided uuid
	DHCPOptionsGet(uuid string) (*goovn.DHCPOptions, error)
	// List dhcp options
	DHCPOptionsList() ([]*goovn.DHCPOptions, error)

	// Add qos rule
	QoSAdd(ls string, direction string, priority int, match string, action map[string]int, bandwidth map[string]int, external_ids map[string]string) (*goovn.OvnCommand, error)
	// Del qos rule, to delete wildcard specify priority -1 and string options as ""
	QoSDel(ls string, direction string, priority int, match string) (*goovn.OvnCommand, error)
	// Get qos rules by logical switch
	QoSList(ls string) ([]*goovn.QoS, error)

	//Add NAT to Logical Router
	LRNATAdd(lr string, ntype string, externalIp string, logicalIp string, external_ids map[string]string, logicalPortAndExternalMac ...string) (*goovn.OvnCommand, error)
	//Del NAT from Logical Router
	LRNATDel(lr string, ntype string, ip ...string) (*goovn.OvnCommand, error)
	// Get NAT List by Logical Router
	LRNATList(lr string) ([]*goovn.NAT, error)
	// Add Meter with a Meter Band
	MeterAdd(name, action string, rate int, unit string, external_ids map[string]string, burst int) (*goovn.OvnCommand, error)
	// Deletes meters
	MeterDel(name ...string) (*goovn.OvnCommand, error)
	// List Meters
	MeterList() ([]*goovn.Meter, error)
	// List Meter Bands
	MeterBandsList() ([]*goovn.MeterBand, error)
	// Exec command, support mul-commands in one transaction.
	Execute(cmds ...*goovn.OvnCommand) error

	// Add chassis with given name
	ChassisAdd(name string, hostname string, etype []string, ip string, external_ids map[string]string,
		transport_zones []string, vtep_lswitches []string) (*goovn.OvnCommand, error)
	// Delete chassis with given name
	ChassisDel(chName string) (*goovn.OvnCommand, error)
	// Get chassis by hostname or name
	ChassisGet(chname string) ([]*goovn.Chassis, error)

	// Get encaps by chassis name
	EncapList(chname string) ([]*goovn.Encap, error)

	// Set NB_Global table options
	NBGlobalSetOptions(options map[string]string) (*goovn.OvnCommand, error)

	// Get NB_Global table options
	NBGlobalGetOptions() (map[string]string, error)

	// Set SB_Global table options
	SBGlobalSetOptions(options map[string]string) (*goovn.OvnCommand, error)

	// Get SB_Global table options
	SBGlobalGetOptions() (map[string]string, error)

	// Close connection to OVN
	Close() error
}
