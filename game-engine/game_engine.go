package gameengine

import (
	"fmt"
	"github.com/h-abranches-dev/connect-4/pkg/colors"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	version = new(versions.Version)
)

func SetVersion(v string) error {
	pv, err := versions.GetVersion(versions.Get(), v)
	if err != nil {
		return err
	}
	*version = *pv
	return nil
}

func DisplayVersion() {
	gameTitle := colors.FgRed("CONNECT-4")
	v := colors.FgRed(version.Tag)
	fmt.Printf("\n%s game engine\n\n", gameTitle)
	fmt.Printf("version: %s\n\n", v)
}

func StartGRPCServer(port int) error {
	listSrvAddr := utils.ListSrvAddr(port)
	lis, err := net.Listen("tcp", listSrvAddr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	RegisterRouteServer(grpcServer, NewGameEngine())

	log.Printf(">>> game engine is listening on the address %s\n\n", colors.FgRed(lis.Addr().String()))

	if err = grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
