package ad

import (
	"github.com/stretchr/testify/assert"
	"learning-go-challenges/domain/ad/exception"
	"testing"
	"time"
)

func TestNewAd(t *testing.T) {
	currentTimestamp := time.Now()

	ad, err := NewAd(Id{Value: "random"}, "a title", "a description", 99, currentTimestamp)

	assert.NoError(t, err)
	assert.Equal(
		t,
		Ad{
			Id:          Id{Value: "random"},
			Title:       "a title",
			Description: "a description",
			Price:       99,
			PublishedAt: currentTimestamp,
		},
		*ad)
}

func TestDoNotFailWhenAdDescriptionIsLongerThan50Characters(t *testing.T) {
	tests := []struct {
		description string
	}{
		{description: "This is a valid title"},
		{description: "Valid_Title"},
		{description: ""},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			_, err := NewAd(Id{Value: "random"}, "a title", test.description, 99, time.Now())
			assert.NoError(t, err)
		})
	}
}

func TestFailWhenAdDescriptionIsLongerThan50Characters(t *testing.T) {
	tests := []struct {
		description string
	}{
		{description: "          iiiiiii         oooooo      m      m  m              k          k          k      "},
		{description: "                                                                                            "},
		{description: "iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii"},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			ad, err := NewAd(Id{Value: "random"}, "a title", test.description, 99, time.Now())

			assert.Nil(t, ad)
			assert.Error(t, err, exception.AdDescriptionTooLongException{})
		})
	}
}
