package server

import (
	"crypto/subtle"
	"net"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dylanlott/pangolin/pkg/database"
	"github.com/dylanlott/pangolin/pkg/pb"
)

var (
	// Error is the main error class for the server package
	Error = errs.Class("server error")
)

// Server holds the db and logger
type Server struct {
	GRPC   *grpc.Server
	DB     *db.DB
	logger *zap.Logger
	apiKey string
}

// NewServer instantiates a server
func NewServer(database *db.DB, apiKey string, logger *zap.Logger, lis net.Listener) *Server {
	return &Server{
		GRPC:   grpc.NewServer(),
		DB:     database,
		logger: logger,
		apiKey: apiKey,
	}
}

func (s *Server) validateAuth(key string) error {
	var expected = []byte(s.apiKey)
	var actual = []byte(key)

	if 1 != subtle.ConstantTimeCompare(expected, actual) {
		s.logger.Error("unauthorized request: ", zap.Error(status.Errorf(codes.Unauthenticated, "invalid api key")))
		return status.Errorf(codes.Unauthenticated, "invalid api key")
	}
	return nil
}

// Put inserts an item into a collection
func (s *Server) Put(req *pb.PutRequest) (res *pb.PutResponse, err error) {
	err = s.validateAuth(req.ApiKey)
	if err != nil {
		return nil, err
	}

	coll := s.DB.Collections[req.CollectionName]

	data, err := coll.Put(req.Data)
	if err != nil {
		return nil, err
	}

	d, ok := data.(*any.Any)
	if !ok {
		return nil, db.Error.New("error putting data")
	}

	res = &pb.PutResponse{
		Data: d,
	}
	return res, nil
}

// Insert inserts a data interface (blob?) into a collection
func (s *Server) Insert(req *pb.InsertRequest) (res *pb.InsertResponse, err error) {
	err = s.validateAuth(req.ApiKey)
	if err != nil {
		return nil, err
	}

	data, err := db.Insert(req.Data, req.CollectionName)
	if err != nil {
		return nil, err
	}

	d, ok := data.(*any.Any)
	if !ok {
		return nil, db.Error.New("error inserting data")
	}

	res = &pb.InsertResponse{
		Data: d,
	}
	return res, nil
}

// FindOne finds exactly one collection
func (s *Server) FindOne(req *pb.FindOneRequest) (res *pb.FindOneResponse, err error) {
	err = s.validateAuth(req.ApiKey)
	if err != nil {
		return nil, err
	}

	coll := s.DB.Collections[req.CollectionName]

	data, err := coll.FindOne(req.Query)
	if err != nil {
		return nil, err
	}

	d, ok := data.(*any.Any)
	if !ok {
		return nil, db.Error.New("error with FindOne data")
	}

	res = &pb.FindOneResponse{
		Data: d,
	}
	return res, nil
}

// Find queries collections
func (s *Server) Find(req *pb.FindRequest) (res *pb.FindResponse, err error) {
	err = s.validateAuth(req.ApiKey)
	if err != nil {
		return nil, err
	}

	coll := s.DB.Collections[req.CollectionName]

	// TODO: make Find return more stuff
	err = coll.Find(req.Query)
	if err != nil {
		return nil, err
	}

	res = &pb.FindResponse{}
	return res, nil
}

// Delete deletes a key from a collection
func (s *Server) Delete(req *pb.DeleteRequest) (res *pb.DeleteResponse, err error) {
	err = s.validateAuth(req.ApiKey)
	if err != nil {
		return nil, err
	}

	coll := s.DB.Collections[req.CollectionName]

	data, err := coll.Delete(req.Key)
	if err != nil {
		return nil, err
	}

	d, ok := data.(*any.Any)
	if !ok {
		return nil, db.Error.New("error with FindOne data")
	}

	res = &pb.DeleteResponse{
		Data: d,
	}
	return res, nil
}
