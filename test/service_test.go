package test

import (
	"log"
	"testing"

	"github.com/google/uuid"
)

func TestService_Ping(t *testing.T) {
	log.Println(uuid.NewString())
}
