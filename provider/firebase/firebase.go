package firebase

import (
	"errors"

	"github.com/kaijuci/apk-distributor/distributor"
)

func NewFirebaseDistributor() (distributor.Distributor, error) {
	return &impl{}, nil
}

type impl struct {
}

func (i *impl) Distribute(apkRelease *distributor.APKRelease) (string, error) {
	return "", errors.New("not implemented")
}
