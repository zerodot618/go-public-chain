package blc

type Version struct {
	Version    int    // 版本
	BestHeight int64  // 当前节点区块的高度
	AddrFrom   string //当前节点的地址
}
