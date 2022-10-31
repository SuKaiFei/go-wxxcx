package service

import (
	"context"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"os"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
)

var (
	id, _ = os.Hostname()
	ctx   = context.Background()
	tSVC  *unitTestSvc
	_, _  = ctx, tSVC
)

func TestMain(m *testing.M) {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", "unit test",
		"service.version", "local",
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)

	var bc conf.Bootstrap
	c := config.New(
		config.WithSource(
			file.NewSource("../../configs/config-prod.yaml"),
		))
	if err := c.Load(); err != nil {
		panic(err)
	}
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	defer c.Close()

	var (
		closeFunc func()
		err       error
	)

	tSVC, closeFunc, err = NewTestUnitTestSvcService(bc.GetServer(), logger, &bc, bc.GetData())
	if err != nil {
		panic(err)
	}
	defer closeFunc()

	os.Exit(m.Run())
}
