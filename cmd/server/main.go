package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/nartvt/smart-service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	jwtlib "github.com/golang-jwt/jwt/v5"
	coreService "github.com/nartvt/go-core/service"
	proto "github.com/nartvt/smart-service/api/smart/v1"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func initService(logger log.Logger,
	gs *grpc.Server,
	hs *http.Server,
	userService *service.UserServiceService) coreService.Service {
	proto.RegisterUserServiceServer(gs, userService)
	proto.RegisterUserServiceHTTPServer(hs, userService)
	return coreService.NewService(logger, hs, gs, kratos.ID(id), kratos.Name(Name), kratos.Version(Version))
}

// func main() {
// 	flag.Parse()
// 	logger := log.With(log.NewStdLogger(os.Stdout),
// 		"ts", log.DefaultTimestamp,
// 		"caller", log.DefaultCaller,
// 		"service.id", id,
// 		"service.name", Name,
// 		"service.version", Version,
// 		"trace.id", tracing.TraceID(),
// 		"span.id", tracing.SpanID(),
// 	)
// 	c := config.New(
// 		config.WithSource(
// 			file.NewSource(flagconf),
// 		),
// 	)
// 	defer c.Close()
//
// 	if err := c.Load(); err != nil {
// 		panic(err)
// 	}
//
// 	var bc conf.Bootstrap
// 	if err := c.Scan(&bc); err != nil {
// 		panic(err)
// 	}
//
// 	app, cleanup, err := initApp(bc.Server, bc.Data, logger)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer cleanup()
//
// 	// start and wait for stop signal
// 	if err := app.Run(); err != nil {
// 		panic(err)
// 	}
// }

type options struct {
	required      bool
	signingMethod jwtlib.SigningMethod
	excludes      []string
	autoParse     bool
	claims        func() jwtlib.Claims
	tokenHeader   map[string]interface{}
	key           string
	keyFunc       jwtlib.Keyfunc
}

func GeneratorJwtToken(ctx context.Context, userId string) (context.Context, error) {
	claims := jwtlib.RegisteredClaims{
		Subject: userId,
		ExpiresAt: jwtlib.NewNumericDate(
			time.Now().Add(24 * time.Hour),
		),
	}
	secretKey := os.Getenv("ENV_JWT_SECRET")
	o := &options{
		signingMethod: jwtlib.SigningMethodHS256,
		key:           secretKey,
		required:      false,
		claims:        func() jwtlib.Claims { return claims },
		tokenHeader: map[string]interface{}{
			"typ": "JWT",
			"alg": "HS256",
		},
		keyFunc: func(token *jwtlib.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	}

	token := jwtlib.NewWithClaims(o.signingMethod, o.claims())
	if o.tokenHeader != nil {
		for k, v := range o.tokenHeader {
			token.Header[k] = v
		}
	}
	key, err := o.keyFunc(token)
	if err != nil {
		return ctx, errors.New("keyFunc error")
	}
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return ctx, errors.New("SignedString error")
	}

	return context.WithValue(ctx, "accessToken", tokenStr), nil
}

func main() {
	ctx, err := GeneratorJwtToken(context.Background(), "tranvantai0011")
	fmt.Println("Token: ", ctx.Value("accessToken"), err)
}
