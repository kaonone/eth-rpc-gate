package transformer

import (
	"context"

	"github.com/dcb9/go-ethereum/common/hexutil"
	"github.com/kaonone/eth-rpc-gate/pkg/eth"
	"github.com/kaonone/eth-rpc-gate/pkg/kaon"
	"github.com/labstack/echo"
)

// ProxyNetPeerCount implements ETHProxy
type ProxyNetPeerCount struct {
	*kaon.Kaon
}

func (p *ProxyNetPeerCount) Method() string {
	return "net_peerCount"
}

func (p *ProxyNetPeerCount) Request(rawreq *eth.JSONRPCRequest, c echo.Context) (interface{}, *eth.JSONRPCError) {
	return p.request(c.Request().Context())
}

func (p *ProxyNetPeerCount) request(ctx context.Context) (*eth.NetPeerCountResponse, *eth.JSONRPCError) {
	peerInfos, err := p.GetPeerInfo(ctx)
	if err != nil {
		return nil, eth.NewCallbackError(err.Error())
	}

	resp := eth.NetPeerCountResponse(hexutil.EncodeUint64(uint64(len(peerInfos))))
	return &resp, nil
}
