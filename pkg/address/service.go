package address

import (
	context "context"

	protoAddress "github.com/jcsw/address-grpc-api/pkg/proto/address"
	log "github.com/jcsw/address-grpc-service/pkg/system/log"
)

// Service difine the address service
type Service struct {
}

// SearchAddress search address
func (s *Service) SearchAddress(ctx context.Context, in *protoAddress.Input) (*protoAddress.Output, error) {
	log.Info("p=address f=SearchAddress input=%v", in)
	return &protoAddress.Output{Street: "Faria Lima"}, nil
}
