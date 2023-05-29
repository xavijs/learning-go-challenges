package adpublisher

import (
	"fmt"
	"learning-go-challenges/domain/ad"
	"math/rand"
	"time"
)

type AdPublisher interface {
	publish(ad ad.Ad) bool
}

type WallapopAdPublisher struct{}

func NewWallapopAdPublisher() *WallapopAdPublisher {
	return &WallapopAdPublisher{}
}

func (receiver WallapopAdPublisher) publish(ad ad.Ad) bool {
	return randomDelayedBooleanResponse(fmt.Sprintf("%T", receiver), ad)
}

type LeboncoinAdPublisher struct{}

func NewLeboncoinAdPublisher() *LeboncoinAdPublisher {
	return &LeboncoinAdPublisher{}
}

func (receiver LeboncoinAdPublisher) publish(ad ad.Ad) bool {
	return randomDelayedBooleanResponse(fmt.Sprintf("%T", receiver), ad)
}

type MilanunciosAdPublisher struct{}

func NewMilanunciosAdPublisher() *MilanunciosAdPublisher {
	return &MilanunciosAdPublisher{}
}

func (receiver MilanunciosAdPublisher) publish(ad ad.Ad) bool {
	return randomDelayedBooleanResponse(fmt.Sprintf("%T", receiver), ad)
}

func randomDelayedBooleanResponse(publisherName string, adToPublish ad.Ad) bool {
	fmt.Printf("Publishing ad %v to publisher %v \n", adToPublish, publisherName)
	time.Sleep(time.Second)
	rand.NewSource(time.Now().UnixNano())
	isSucceeded := rand.Float64() < 0.5
	fmt.Printf("Publishing ad to publisher succeeded? %v \n", isSucceeded)
	return isSucceeded
}
