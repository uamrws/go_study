package trip

import (
	"context"
	trippb "go_study/coolcar_study/server/proto/gen/go"
)

// Service implements trip service
type Service struct {
	trippb.UnimplementedTripServerServer
}

func (s Service) GetTrip(ctx context.Context, request *trippb.GetTripRequest) (*trippb.GetTripRespone, error) {
	return &trippb.GetTripRespone{
		Id: request.Id,
		Trip: &trippb.Trip{
			Start:       "start",
			End:         "end",
			DurationSec: 3600,
			FeeCent:     10000,
			StartPos: &trippb.Location{
				Longitude: 30,
				Latitude:  120,
			},
			EndPos: &trippb.Location{
				Longitude: 35,
				Latitude:  115,
			},
			PathLocations: []*trippb.Location{
				{Longitude: 31, Latitude: 119},
				{Longitude: 32, Latitude: 118},
				{Longitude: 33, Latitude: 117},
				{Longitude: 34, Latitude: 116},
			},
			Status: trippb.TripStatus_NOT_STARTED,
		},
	}, nil
}
