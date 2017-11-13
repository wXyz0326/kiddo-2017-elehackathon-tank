// Auto-generated by "nex regen" at "2017-11-13 21:16:49.016348 +0800 CST"
// ** DO NOT EDIT **

package handler

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/eleme/nex/endpoint"
	"github.com/eleme/nex/log"
	json "github.com/json-iterator/go"

	"github.com/eleme/purchaseMeiTuan/services/player"
)

// needed to ensure import safety.
var _ = player.GoUnusedProtection__

// makeAssignTanksEndpoint create a wrapped handler of PlayerService.
func makeAssignTanksEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(assigntanksRequest)
		err := s.AssignTanks(ctx, req.Tanks)
		return assigntanksResponse{}, err
	}
}

// makeGetNewOrdersEndpoint create a wrapped handler of PlayerService.
func makeGetNewOrdersEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		v, err := s.GetNewOrders(ctx)
		return getnewordersResponse{V: v}, err
	}
}

// makeLatestStateEndpoint create a wrapped handler of PlayerService.
func makeLatestStateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(lateststateRequest)
		err := s.LatestState(ctx, req.State)
		return lateststateResponse{}, err
	}
}

// makePingEndpoint create a wrapped handler of PlayerService.
func makePingEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		v, err := s.Ping(ctx)
		return pingResponse{V: v}, err
	}
}

// makeUploadMapEndpoint create a wrapped handler of PlayerService.
func makeUploadMapEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uploadmapRequest)
		err := s.UploadMap(ctx, req.Gamemap)
		return uploadmapResponse{}, err
	}
}

// makeUploadParamtersEndpoint create a wrapped handler of PlayerService.
func makeUploadParamtersEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uploadparamtersRequest)
		err := s.UploadParamters(ctx, req.Arguments)
		return uploadparamtersResponse{}, err
	}
}

// The following structs is for internal use.

func marshalThriftArg(v interface{}) string {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		return fmt.Sprintf("%#v", v)
	}
	return strings.Trim(string(buf.Bytes()), " \n")
}

type assigntanksRequest struct {
	Tanks []int32
}

// String defines the "native" format for assigntanksRequest.
func (req assigntanksRequest) String() string {
	var ireq interface{} = req
	if imp, ok := ireq.(log.ALogStringer); ok {
		return imp.ALogString()
	}
	return fmt.Sprintf("%s", marshalThriftArg(req.Tanks))
}

type assigntanksResponse struct {
}

type getnewordersRequest struct {
}

// String defines the "native" format for getnewordersRequest.
func (req getnewordersRequest) String() string {
	var ireq interface{} = req
	if imp, ok := ireq.(log.ALogStringer); ok {
		return imp.ALogString()
	}
	return ""
}

type getnewordersResponse struct {
	V []*player.Order
}

type lateststateRequest struct {
	State *player.GameState
}

// String defines the "native" format for lateststateRequest.
func (req lateststateRequest) String() string {
	var ireq interface{} = req
	if imp, ok := ireq.(log.ALogStringer); ok {
		return imp.ALogString()
	}
	return fmt.Sprintf("%s", marshalThriftArg(req.State))
}

type lateststateResponse struct {
}

type pingRequest struct {
}

// String defines the "native" format for pingRequest.
func (req pingRequest) String() string {
	var ireq interface{} = req
	if imp, ok := ireq.(log.ALogStringer); ok {
		return imp.ALogString()
	}
	return ""
}

type pingResponse struct {
	V bool
}

type uploadmapRequest struct {
	Gamemap [][]int32
}

// String defines the "native" format for uploadmapRequest.
func (req uploadmapRequest) String() string {
	var ireq interface{} = req
	if imp, ok := ireq.(log.ALogStringer); ok {
		return imp.ALogString()
	}
	return fmt.Sprintf("%s", marshalThriftArg(req.Gamemap))
}

type uploadmapResponse struct {
}

type uploadparamtersRequest struct {
	Arguments *player.Args_
}

// String defines the "native" format for uploadparamtersRequest.
func (req uploadparamtersRequest) String() string {
	var ireq interface{} = req
	if imp, ok := ireq.(log.ALogStringer); ok {
		return imp.ALogString()
	}
	return fmt.Sprintf("%s", marshalThriftArg(req.Arguments))
}

type uploadparamtersResponse struct {
}
