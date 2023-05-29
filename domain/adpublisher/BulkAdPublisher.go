package adpublisher

import "learning-go-challenges/domain/ad"

type BulkAdPublisher struct {
	adPublishers []AdPublisher
}

func NewBulkAdPublisher(adPublishers []AdPublisher) *BulkAdPublisher {
	return &BulkAdPublisher{adPublishers: adPublishers}
}
func (receiver BulkAdPublisher) Execute(ad ad.Ad) {
	done := make(chan bool)

	for _, publisher := range receiver.adPublishers {
		go func(currentPublisher AdPublisher) {
			currentPublisher.publish(ad)
			done <- true
		}(publisher)
	}

	for range receiver.adPublishers {
		<-done
	}
}
