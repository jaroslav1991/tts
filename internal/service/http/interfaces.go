//go:generate mockgen -source=$GOFILE -destination=interfaces_mocks.go -package=$GOPACKAGE
package http

type Service interface {
	SaveData(request any) error
}
