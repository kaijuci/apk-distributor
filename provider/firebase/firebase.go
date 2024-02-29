package firebase

import (
	"context"
	"errors"

	"github.com/kaijuci/apk-distributor/distributor"
	"google.golang.org/api/firebaseappdistribution/v1"
	"google.golang.org/api/option"
)

func NewFirebaseDistributor(credentialsFile string) (distributor.Distributor, error) {
	ctx := context.Background()
	svc, err := firebaseappdistribution.NewService(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, err
	}

	return &impl{svc}, nil
}

type impl struct {
	svc *firebaseappdistribution.Service
}

func (i *impl) Distribute(apkRelease *distributor.APKRelease) (string, error) {
	return "", errors.New("not implemented")
}
