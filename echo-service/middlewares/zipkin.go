package middlewares

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kitopentracing "github.com/go-kit/kit/tracing/opentracing"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

type ServiceTracer interface {
	InjectOptions(handlerName string)
	Middlewares(handlerName string) []endpoint.Middleware
}

type BaseTracer struct {
	Reporter     reporter.Reporter
	ZipkinTracer *zipkin.Tracer     // zipkin tracer
	OpenTracer   opentracing.Tracer // opentracing tracer
}

type ClientTracer struct {
	*BaseTracer
	Options []grpctransport.ClientOption
}

type ServerTracer struct {
	*BaseTracer
	Options []grpctransport.ServerOption
}

func newBaseTracer(srvName, address, url string) (*BaseTracer, error) {
	// set up a span reporter
	reporter := zipkinhttp.NewReporter(url)

	// create our local service endpoint
	endpoint, err := zipkin.NewEndpoint(srvName, address)
	if err != nil {
		return nil, err
	}

	// initialize our openTracer
	zipkinTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		return nil, err
	}

	// use zipkin-go-opentracing to wrap our openTracer
	openTracer := zipkinot.Wrap(zipkinTracer)

	// optionally set as Global OpenTracing openTracer instance
	opentracing.SetGlobalTracer(openTracer)

	return &BaseTracer{
		Reporter:     reporter,
		ZipkinTracer: zipkinTracer,
		OpenTracer:   openTracer,
	}, nil
}

func NewServerTracer(srvName, address, url string, logger log.Logger) (*ServerTracer, error) {
	baseTracer, err := newBaseTracer(srvName, address, url)
	if err != nil {
		return nil, err
	}

	// set server serverOptions
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		// Zipkin GRPC BaseServer Trace can either be instantiated per gRPC method with a
		// provided operation name or a global tracing service can be instantiated
		// without an operation name and fed to each Go kit gRPC server as a GrpcServiceServerConfig.
		// In the latter case, the operation name will be the endpoint's grpc method
		// path if used in combination with the Go kit gRPC Interceptor.
		//
		kitzipkin.GRPCServerTrace(baseTracer.ZipkinTracer),
	}

	return &ServerTracer{
		BaseTracer: baseTracer,
		Options:    serverOptions,
	}, nil
}

func NewClientTracer(srvName, address, url string) (*ClientTracer, error) {
	baseTracer, err := newBaseTracer(srvName, address, url)
	if err != nil {
		return nil, err
	}

	clientOptions := []grpctransport.ClientOption{
		kitzipkin.GRPCClientTrace(baseTracer.ZipkinTracer),
	}

	return &ClientTracer{
		BaseTracer: baseTracer,
		Options:    clientOptions,
	}, nil
}

// server tracer options
// ServerOption???Serve???????????????????????????, ???????????????:
//1. ServerBefore: ?????????decode????????????????????????HTTP?????????????????????ServerBefore??????
//2. ServerAfter: ?????????endpoint??????, encode??????????????????
//3. ServerErrorHandler: ??????decode, endpoint, encode????????????????????????????????????????????????, ???????????????log
//4. ServerErrorEncoder: ??????decode, endpoint, encode????????????????????????????????????????????????, ??????????????????http.ResponseWriter???????????????
//5. ServerFinalizer: ?????????HTTP???????????????????????????encode??????ServerErrorEncoder?????????????????????
//
//?????????????????????: ServerBefore -> decode -> endpoint -> service -> ServerAfter -> encode -> ServerFinalizer
//???????????????????????????: ServerBefore -> ????????????(decode -> endpoint -> encode) -> ServerErrHandler -> ServerErrorEncoder(??????httpResponse) -> ServerFinalizer
func (t *ServerTracer) InjectOptions(handlerName string, logger log.Logger) {
	t.Options = append(t.Options,
		grpctransport.ServerBefore(
			kitopentracing.GRPCToContext(t.OpenTracer, handlerName, logger),
		),
	)
}

func (t *ServerTracer) Middlewares(handlerName string) []endpoint.Middleware {
	return []endpoint.Middleware{
		kitopentracing.TraceServer(t.OpenTracer, handlerName),
		kitzipkin.TraceEndpoint(t.ZipkinTracer, handlerName),
	}
}

// client tracer methods
func (t *ClientTracer) InjectOptions(logger log.Logger) {
	t.Options = append(t.Options,
		grpctransport.ClientBefore(
			kitopentracing.ContextToGRPC(t.OpenTracer, logger),
		),
	)
}

func (t *ClientTracer) Middlewares(handlerName string) []endpoint.Middleware {
	return []endpoint.Middleware{
		kitopentracing.TraceClient(t.OpenTracer, handlerName),
	}
}
