package mocks

//go:generate mockery --name "AdRepository" --dir ../domain/ad --output ./ --filename "adRepositoryMock.go" --with-expecter
//go:generate mockery --name "Clock" --dir ../domain/clock --output ./ --filename "clockMock.go" --with-expecter
//go:generate mockery --name "UUIDGenerator" --dir ../domain/uuid --output ./ --filename "uuidGeneratorMock.go" --with-expecter
