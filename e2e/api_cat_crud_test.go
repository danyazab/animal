package e2e

import (
	"context"
	"fmt"
	"github.com/danyazab/animal/internal/animal/model/util"
	"github.com/danyazab/animal/internal/api/http/request"
	"github.com/danyazab/animal/pkg/e2e"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// go test --tags=e2e ./e2e -run Test_Cat_Crud -v

func Test_Cat_Crud_GetList(t *testing.T) {
	t.Run(fmt.Sprintf("[e2e] test get list og cats"), e2e.Inject(func(t *testing.T, rc e2e.Transport) {
		route := "/pet/cat"
		var result []string

		fmt.Println()

		err := rc.Get(context.Background(), route, &result)
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
		assert.Equal(t, len(result), 1)
	}))
}

//go test --tags=e2e ./... -run Test_Cat_Crud_CreateCat -count=5
func Test_Cat_Crud_CreateCat(t *testing.T) {
	t.Run(fmt.Sprintf("[e2e] create cat"), e2e.Inject(func(t *testing.T, rc e2e.Transport) {
		route := "/pet/cat"
		var result struct {
			ID uint `json:"id"`
		}

		now := time.Now()

		requestBody := request.CreateCatReq{
			Name:        fmt.Sprintf("Test_%d", now.Nanosecond()),
			Description: fmt.Sprintf("test descr %s", now.Format(time.RFC1123)),
			Birthday:    now.Format("2006-01-02"),
			Breed:       "Spinx",
			Sex:         string(util.TypeSexMale),
			TailLength:  uint(now.Second()),
			Color:       "black",
			WoolType:    "short",
			IsChipped:   true,
			Weight:      2.2,
		}

		err := rc.Post(context.Background(), route, requestBody, &result)
		assert.NoError(t, err)
		assert.NotEmpty(t, result.ID)
	}))
}
