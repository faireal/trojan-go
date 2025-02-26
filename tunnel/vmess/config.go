package vmess

// Request Options
const (
	OptBasicFormat byte = 0 // 不加密传输
	OptChunkStream byte = 1 // 分块传输，每个分块使用如下Security方法加密
	// OptReuseTCPConnection byte = 2
	// OptMetadataObfuscate  byte = 4
)

// Security types
const (
	SecurityAES128GCM        byte = 3
	SecurityChacha20Poly1305 byte = 4
	SecurityNone             byte = 5
)

// CMD types
const (
	CmdTCP byte = 1
	CmdUDP byte = 2
)

// Atyp
const (
	AtypIP4    byte = 1
	AtypDomain byte = 2
	AtypIP6    byte = 3
)

type Config struct {
	LocalHost  string `json:"local_addr" yaml:"local-addr"`
	LocalPort  int    `json:"local_port" yaml:"local-port"`
	RemoteHost string `json:"remote_addr" yaml:"remote-addr"`
	RemotePort int    `json:"remote_port" yaml:"remote-port"`
	UUID       string `json:"uuid" yaml:"uuid"`
	AlterID    int    `json:"alterId" yaml:"alterId"`
	Security   string `json:"security" yaml:"security"`
	Port       string `json:"port" yaml:"port"`
	HostName   string `json:"hostName" yaml:"hostName"`
	IsAead     bool   `json:"isAead" yaml:"isAead"`
}
