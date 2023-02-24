package bound

import (
	"context"
	"fmt"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/errno"
	"github.com/shirou/gopsutil/cpu"
)

var _ remote.InboundHandler = &cpuLimitHandler{}

type cpuLimitHandler struct{}

func NewCpuLimitHandler() remote.InboundHandler {
	return &cpuLimitHandler{}
}

func (c *cpuLimitHandler) OnActive(ctx context.Context, conn net.Conn) (context.Context, error) {
	return ctx, nil
}

func (c *cpuLimitHandler) OnRead(ctx context.Context, conn net.Conn) (context.Context, error) {
	p := cpuPercent()
	klog.CtxInfof(ctx, "current cpu is %.2g", p)
	if constants.NeedCPURateLimit && p > constants.CPURateLimit {
		return ctx, errno.ServiceErr.WithMessage(fmt.Sprintf("cpu = %.2g", c))
	}
	return ctx, nil
}

func (c *cpuLimitHandler) OnInactive(ctx context.Context, conn net.Conn) context.Context {
	return ctx
}

func (c *cpuLimitHandler) OnMessage(ctx context.Context, args, result remote.Message) (context.Context, error) {
	return ctx, nil
}

func cpuPercent() float64 {
	percent, _ := cpu.Percent(0, false)
	return percent[0]
}
