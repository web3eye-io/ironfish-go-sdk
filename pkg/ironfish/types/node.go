package types

const (
	NodeStarted NodeStatus = "started"
	NodeStopped NodeStatus = "stopped"
	NodeError   NodeStatus = "error"
)

type NodeStatus string

const (
	BlockSyncerStopped  BlockSyncerStatus = "stopped"
	BlockSyncerIdle     BlockSyncerStatus = "idle"
	BlockSyncerStopping BlockSyncerStatus = "stopping"
	BlockSyncerSyncing  BlockSyncerStatus = "syncing"
)

type BlockSyncerStatus string

const GetNodeStatusPath = "node/getStatus"

type GetNodeStatusRequest struct {
	Stream bool `json:"stream"`
}

type GetNodeStatusResponse struct {
	Node struct {
		Status   NodeStatus `json:"status"`
		Version  string     `json:"version"`
		Git      string     `json:"git"`
		NodeName string     `json:"nodeName"`
	} `json:"node"`
	Cpu struct {
		Cores             int     `json:"cores"`
		PercentRollingAvg float32 `json:"percentRollingAvg"`
		PercentCurrent    float32 `json:"percentCurrent"`
	} `json:"cpu"`
	Memory struct {
		HeapMax   int `json:"heapMax"`
		HeapTotal int `json:"heapTotal"`
		HeapUsed  int `json:"heapUsed"`
		Rss       int `json:"rss"`
		MemFree   int `json:"memFree"`
		MemTotal  int `json:"memTotal"`
	} `json:"memory"`
	MiningDirector struct {
		Status                    string  `json:"status"`
		Miners                    int     `json:"miners"`
		Blocks                    int     `json:"blocks"`
		BlockGraffiti             string  `json:"blockGraffiti"`
		NewBlockTemplateSpeed     float32 `json:"newBlockTemplateSpeed"`
		NewBlockTransactionsSpeed float32 `json:"newBlockTransactionsSpeed"`
	} `json:"miningDirector"`
	MemPool struct {
		Size      int `json:"size"`
		SizeBytes int `json:"sizeBytes"`
	} `json:"memPool"`
	Blockchain struct {
		Synced bool `json:"synced"`
		Head   struct {
			Hash     string `json:"hash"`
			Sequence int    `json:"sequence"`
		} `json:"head"`
		HeadTimestamp int     `json:"headTimestamp"`
		NewBlockSpeed float32 `json:"newBlockSpeed"`
	} `json:"blockchain"`
	BlockSyncer struct {
		Status  BlockSyncerStatus `json:"status"`
		Syncing struct {
			BlockSpeed    float32 `json:"blockSpeed"`
			Speed         float32 `json:"speed"`
			DownloadSpeed float32 `json:"downloadSpeed"`
			Progress      float32 `json:"progress"`
		} `json:"syncing"`
	} `json:"blockSyncer"`
	PeerNetwork struct {
		Peers           int     `json:"peers"`
		IsReady         bool    `json:"isReady"`
		InboundTraffic  float32 `json:"inboundTraffic"`
		OutboundTraffic float32 `json:"outboundTraffic"`
	} `json:"peerNetwork"`
	Telemetry struct {
		Status    string `json:"status"`
		Pending   int    `json:"pending"`
		Submitted int    `json:"submitted"`
	} `json:"telemetry"`
	Workers struct {
		Started   bool    `json:"started"`
		Workers   int     `json:"workers"`
		Queued    int     `json:"queued"`
		Capacity  int     `json:"capacity"`
		Executing int     `json:"executing"`
		Change    float32 `json:"change"`
		Speed     float32 `json:"speed"`
	} `json:"workers"`
	Accounts struct {
		Scanning struct {
			Sequence    int `json:"sequence"`
			EndSequence int `json:"endSequence"`
			StartedAt   int `json:"startedAt"`
		} `json:"scanning"`
		Head struct {
			Hash     string `json:"hash"`
			Sequence int    `json:"sequence"`
		} `json:"head"`
	} `json:"accounts"`
}
