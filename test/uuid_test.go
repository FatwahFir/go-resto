package test

import (
	"fmt"
	database "go-resto/db"
	"testing"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestUuid(t *testing.T) {
	uuid := uuid.NewString()
	fmt.Println(uuid)
}

func TestDb(t *testing.T) {
	var env *viper.Viper = viper.New()

	err := database.NewDB(env)
	assert.Nil(t, err)
}
