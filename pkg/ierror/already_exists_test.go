package ierror_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ruslanSorokin/lock-manager/pkg/ierror"
)

func TestInstantiateAlreadyExists(t *testing.T) {
	assert := assert.New(t)

	userStatic := ierror.NewAlreadyExists(
		"user with given login already exists",
		"USER_ALREADY_EXISTS",
	)
	userDynamic := ierror.InstantiateAlreadyExists(userStatic, "73546234")

	assert.True(errors.Is(userStatic, userDynamic))

	productStatic := ierror.NewAlreadyExists(
		"product with given name already exists",
		"PRODUCT_ALREADY_EXISTS",
	)
	productDynamic := ierror.InstantiateAlreadyExists(productStatic, "6456235")

	assert.True(errors.Is(productStatic, productDynamic))

	assert.False(errors.Is(productStatic, userDynamic))
	assert.False(errors.Is(userStatic, productDynamic))

	assert.False(errors.Is(userStatic, errors.New("random error here")))
}
