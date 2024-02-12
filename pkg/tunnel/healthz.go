package tunnel

import (
	"context"
	"net/url"

	"github.com/paralus/paralus/pkg/grpc"
	rpcv3 "github.com/paralus/paralus/proto/rpc/scheduler"
	commonv3 "github.com/paralus/paralus/proto/types/commonpb/v3"
	infrav3 "github.com/paralus/paralus/proto/types/infrapb/v3"
	"github.com/paralus/relay/pkg/utils"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func syncClusterHeath(sni string, status commonv3.ParalusConditionStatus, reason string) {
	id, err := getClusterID(sni)
	if err != nil {
		_log.Error("unable to get clusterID", "error", err)
	}

	u, err := url.Parse(utils.PeerServiceURI)
	if err != nil {
		_log.Error("unable to parse peer service url", "error", err)
	}
	//Load certificates
	tlsConfig, err := ClientTLSConfigFromBytes(utils.PeerCertificate, utils.PeerPrivateKey, utils.PeerCACertificate, u.Host)
	if err != nil {
		_log.Error("unable to build tls config for peer service", "error", err)
	}
	transportCreds := credentials.NewTLS(tlsConfig)
	peerSeviceHost := u.Host

	ctx := context.Background()
	conn, err := grpc.NewSecureClientConn(ctx, peerSeviceHost, transportCreds)
	if err != nil {
		_log.Error("unable to connect to core", "error", err)
	}
	defer conn.Close()

	client := rpcv3.NewClusterServiceClient(conn)
	_, err = client.UpdateClusterStatus(ctx, &rpcv3.UpdateClusterStatusRequest{
		Metadata: &commonv3.Metadata{
			Id:      id,
			Project: id,
		},
		ClusterStatus: &infrav3.ClusterStatus{
			Conditions: []*infrav3.ClusterCondition{
				{
					Type:        infrav3.ClusterConditionType_ClusterHealth,
					Status:      status,
					LastUpdated: timestamppb.Now(),
					Reason:      reason,
				},
			},
		},
	})
	if err != nil {
		_log.Error("failed to update cluster status", "error", err)
	}
	_log.Debug("successfully update cluster ", sni, " status ", status)
}
