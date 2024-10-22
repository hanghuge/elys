// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	math "cosmossdk.io/math"
	ammtypes "github.com/elys-network/elys/x/amm/types"

	mock "github.com/stretchr/testify/mock"

	perpetualtypes "github.com/elys-network/elys/x/perpetual/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// CloseEstimationChecker is an autogenerated mock type for the CloseEstimationChecker type
type CloseEstimationChecker struct {
	mock.Mock
}

type CloseEstimationChecker_Expecter struct {
	mock *mock.Mock
}

func (_m *CloseEstimationChecker) EXPECT() *CloseEstimationChecker_Expecter {
	return &CloseEstimationChecker_Expecter{mock: &_m.Mock}
}

// EstimateSwap provides a mock function with given fields: ctx, leveragedAmtTokenIn, borrowAsset, ammPool
func (_m *CloseEstimationChecker) EstimateSwap(ctx types.Context, leveragedAmtTokenIn types.Coin, borrowAsset string, ammPool ammtypes.Pool) (math.Int, math.LegacyDec, error) {
	ret := _m.Called(ctx, leveragedAmtTokenIn, borrowAsset, ammPool)

	if len(ret) == 0 {
		panic("no return value specified for EstimateSwap")
	}

	var r0 math.Int
	var r1 math.LegacyDec
	var r2 error
	if rf, ok := ret.Get(0).(func(types.Context, types.Coin, string, ammtypes.Pool) (math.Int, math.LegacyDec, error)); ok {
		return rf(ctx, leveragedAmtTokenIn, borrowAsset, ammPool)
	}
	if rf, ok := ret.Get(0).(func(types.Context, types.Coin, string, ammtypes.Pool) math.Int); ok {
		r0 = rf(ctx, leveragedAmtTokenIn, borrowAsset, ammPool)
	} else {
		r0 = ret.Get(0).(math.Int)
	}

	if rf, ok := ret.Get(1).(func(types.Context, types.Coin, string, ammtypes.Pool) math.LegacyDec); ok {
		r1 = rf(ctx, leveragedAmtTokenIn, borrowAsset, ammPool)
	} else {
		r1 = ret.Get(1).(math.LegacyDec)
	}

	if rf, ok := ret.Get(2).(func(types.Context, types.Coin, string, ammtypes.Pool) error); ok {
		r2 = rf(ctx, leveragedAmtTokenIn, borrowAsset, ammPool)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CloseEstimationChecker_EstimateSwap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EstimateSwap'
type CloseEstimationChecker_EstimateSwap_Call struct {
	*mock.Call
}

// EstimateSwap is a helper method to define mock.On call
//   - ctx types.Context
//   - leveragedAmtTokenIn types.Coin
//   - borrowAsset string
//   - ammPool ammtypes.Pool
func (_e *CloseEstimationChecker_Expecter) EstimateSwap(ctx interface{}, leveragedAmtTokenIn interface{}, borrowAsset interface{}, ammPool interface{}) *CloseEstimationChecker_EstimateSwap_Call {
	return &CloseEstimationChecker_EstimateSwap_Call{Call: _e.mock.On("EstimateSwap", ctx, leveragedAmtTokenIn, borrowAsset, ammPool)}
}

func (_c *CloseEstimationChecker_EstimateSwap_Call) Run(run func(ctx types.Context, leveragedAmtTokenIn types.Coin, borrowAsset string, ammPool ammtypes.Pool)) *CloseEstimationChecker_EstimateSwap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(types.Coin), args[2].(string), args[3].(ammtypes.Pool))
	})
	return _c
}

func (_c *CloseEstimationChecker_EstimateSwap_Call) Return(_a0 math.Int, _a1 math.LegacyDec, _a2 error) *CloseEstimationChecker_EstimateSwap_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *CloseEstimationChecker_EstimateSwap_Call) RunAndReturn(run func(types.Context, types.Coin, string, ammtypes.Pool) (math.Int, math.LegacyDec, error)) *CloseEstimationChecker_EstimateSwap_Call {
	_c.Call.Return(run)
	return _c
}

// EstimateSwapGivenOut provides a mock function with given fields: ctx, tokenOutAmount, tokenInDenom, ammPool
func (_m *CloseEstimationChecker) EstimateSwapGivenOut(ctx types.Context, tokenOutAmount types.Coin, tokenInDenom string, ammPool ammtypes.Pool) (math.Int, math.LegacyDec, error) {
	ret := _m.Called(ctx, tokenOutAmount, tokenInDenom, ammPool)

	if len(ret) == 0 {
		panic("no return value specified for EstimateSwapGivenOut")
	}

	var r0 math.Int
	var r1 math.LegacyDec
	var r2 error
	if rf, ok := ret.Get(0).(func(types.Context, types.Coin, string, ammtypes.Pool) (math.Int, math.LegacyDec, error)); ok {
		return rf(ctx, tokenOutAmount, tokenInDenom, ammPool)
	}
	if rf, ok := ret.Get(0).(func(types.Context, types.Coin, string, ammtypes.Pool) math.Int); ok {
		r0 = rf(ctx, tokenOutAmount, tokenInDenom, ammPool)
	} else {
		r0 = ret.Get(0).(math.Int)
	}

	if rf, ok := ret.Get(1).(func(types.Context, types.Coin, string, ammtypes.Pool) math.LegacyDec); ok {
		r1 = rf(ctx, tokenOutAmount, tokenInDenom, ammPool)
	} else {
		r1 = ret.Get(1).(math.LegacyDec)
	}

	if rf, ok := ret.Get(2).(func(types.Context, types.Coin, string, ammtypes.Pool) error); ok {
		r2 = rf(ctx, tokenOutAmount, tokenInDenom, ammPool)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CloseEstimationChecker_EstimateSwapGivenOut_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EstimateSwapGivenOut'
type CloseEstimationChecker_EstimateSwapGivenOut_Call struct {
	*mock.Call
}

// EstimateSwapGivenOut is a helper method to define mock.On call
//   - ctx types.Context
//   - tokenOutAmount types.Coin
//   - tokenInDenom string
//   - ammPool ammtypes.Pool
func (_e *CloseEstimationChecker_Expecter) EstimateSwapGivenOut(ctx interface{}, tokenOutAmount interface{}, tokenInDenom interface{}, ammPool interface{}) *CloseEstimationChecker_EstimateSwapGivenOut_Call {
	return &CloseEstimationChecker_EstimateSwapGivenOut_Call{Call: _e.mock.On("EstimateSwapGivenOut", ctx, tokenOutAmount, tokenInDenom, ammPool)}
}

func (_c *CloseEstimationChecker_EstimateSwapGivenOut_Call) Run(run func(ctx types.Context, tokenOutAmount types.Coin, tokenInDenom string, ammPool ammtypes.Pool)) *CloseEstimationChecker_EstimateSwapGivenOut_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(types.Coin), args[2].(string), args[3].(ammtypes.Pool))
	})
	return _c
}

func (_c *CloseEstimationChecker_EstimateSwapGivenOut_Call) Return(_a0 math.Int, _a1 math.LegacyDec, _a2 error) *CloseEstimationChecker_EstimateSwapGivenOut_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *CloseEstimationChecker_EstimateSwapGivenOut_Call) RunAndReturn(run func(types.Context, types.Coin, string, ammtypes.Pool) (math.Int, math.LegacyDec, error)) *CloseEstimationChecker_EstimateSwapGivenOut_Call {
	_c.Call.Return(run)
	return _c
}

// GetAmmPool provides a mock function with given fields: ctx, poolId
func (_m *CloseEstimationChecker) GetAmmPool(ctx types.Context, poolId uint64) (ammtypes.Pool, error) {
	ret := _m.Called(ctx, poolId)

	if len(ret) == 0 {
		panic("no return value specified for GetAmmPool")
	}

	var r0 ammtypes.Pool
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, uint64) (ammtypes.Pool, error)); ok {
		return rf(ctx, poolId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint64) ammtypes.Pool); ok {
		r0 = rf(ctx, poolId)
	} else {
		r0 = ret.Get(0).(ammtypes.Pool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint64) error); ok {
		r1 = rf(ctx, poolId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloseEstimationChecker_GetAmmPool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAmmPool'
type CloseEstimationChecker_GetAmmPool_Call struct {
	*mock.Call
}

// GetAmmPool is a helper method to define mock.On call
//   - ctx types.Context
//   - poolId uint64
func (_e *CloseEstimationChecker_Expecter) GetAmmPool(ctx interface{}, poolId interface{}) *CloseEstimationChecker_GetAmmPool_Call {
	return &CloseEstimationChecker_GetAmmPool_Call{Call: _e.mock.On("GetAmmPool", ctx, poolId)}
}

func (_c *CloseEstimationChecker_GetAmmPool_Call) Run(run func(ctx types.Context, poolId uint64)) *CloseEstimationChecker_GetAmmPool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(uint64))
	})
	return _c
}

func (_c *CloseEstimationChecker_GetAmmPool_Call) Return(_a0 ammtypes.Pool, _a1 error) *CloseEstimationChecker_GetAmmPool_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CloseEstimationChecker_GetAmmPool_Call) RunAndReturn(run func(types.Context, uint64) (ammtypes.Pool, error)) *CloseEstimationChecker_GetAmmPool_Call {
	_c.Call.Return(run)
	return _c
}

// GetMTP provides a mock function with given fields: ctx, mtpAddress, id
func (_m *CloseEstimationChecker) GetMTP(ctx types.Context, mtpAddress types.AccAddress, id uint64) (perpetualtypes.MTP, error) {
	ret := _m.Called(ctx, mtpAddress, id)

	if len(ret) == 0 {
		panic("no return value specified for GetMTP")
	}

	var r0 perpetualtypes.MTP
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, uint64) (perpetualtypes.MTP, error)); ok {
		return rf(ctx, mtpAddress, id)
	}
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, uint64) perpetualtypes.MTP); ok {
		r0 = rf(ctx, mtpAddress, id)
	} else {
		r0 = ret.Get(0).(perpetualtypes.MTP)
	}

	if rf, ok := ret.Get(1).(func(types.Context, types.AccAddress, uint64) error); ok {
		r1 = rf(ctx, mtpAddress, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloseEstimationChecker_GetMTP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMTP'
type CloseEstimationChecker_GetMTP_Call struct {
	*mock.Call
}

// GetMTP is a helper method to define mock.On call
//   - ctx types.Context
//   - mtpAddress types.AccAddress
//   - id uint64
func (_e *CloseEstimationChecker_Expecter) GetMTP(ctx interface{}, mtpAddress interface{}, id interface{}) *CloseEstimationChecker_GetMTP_Call {
	return &CloseEstimationChecker_GetMTP_Call{Call: _e.mock.On("GetMTP", ctx, mtpAddress, id)}
}

func (_c *CloseEstimationChecker_GetMTP_Call) Run(run func(ctx types.Context, mtpAddress types.AccAddress, id uint64)) *CloseEstimationChecker_GetMTP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(types.AccAddress), args[2].(uint64))
	})
	return _c
}

func (_c *CloseEstimationChecker_GetMTP_Call) Return(_a0 perpetualtypes.MTP, _a1 error) *CloseEstimationChecker_GetMTP_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CloseEstimationChecker_GetMTP_Call) RunAndReturn(run func(types.Context, types.AccAddress, uint64) (perpetualtypes.MTP, error)) *CloseEstimationChecker_GetMTP_Call {
	_c.Call.Return(run)
	return _c
}

// GetPool provides a mock function with given fields: ctx, poolId
func (_m *CloseEstimationChecker) GetPool(ctx types.Context, poolId uint64) (perpetualtypes.Pool, bool) {
	ret := _m.Called(ctx, poolId)

	if len(ret) == 0 {
		panic("no return value specified for GetPool")
	}

	var r0 perpetualtypes.Pool
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, uint64) (perpetualtypes.Pool, bool)); ok {
		return rf(ctx, poolId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint64) perpetualtypes.Pool); ok {
		r0 = rf(ctx, poolId)
	} else {
		r0 = ret.Get(0).(perpetualtypes.Pool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint64) bool); ok {
		r1 = rf(ctx, poolId)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// CloseEstimationChecker_GetPool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPool'
type CloseEstimationChecker_GetPool_Call struct {
	*mock.Call
}

// GetPool is a helper method to define mock.On call
//   - ctx types.Context
//   - poolId uint64
func (_e *CloseEstimationChecker_Expecter) GetPool(ctx interface{}, poolId interface{}) *CloseEstimationChecker_GetPool_Call {
	return &CloseEstimationChecker_GetPool_Call{Call: _e.mock.On("GetPool", ctx, poolId)}
}

func (_c *CloseEstimationChecker_GetPool_Call) Run(run func(ctx types.Context, poolId uint64)) *CloseEstimationChecker_GetPool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(uint64))
	})
	return _c
}

func (_c *CloseEstimationChecker_GetPool_Call) Return(val perpetualtypes.Pool, found bool) *CloseEstimationChecker_GetPool_Call {
	_c.Call.Return(val, found)
	return _c
}

func (_c *CloseEstimationChecker_GetPool_Call) RunAndReturn(run func(types.Context, uint64) (perpetualtypes.Pool, bool)) *CloseEstimationChecker_GetPool_Call {
	_c.Call.Return(run)
	return _c
}

// NewCloseEstimationChecker creates a new instance of CloseEstimationChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCloseEstimationChecker(t interface {
	mock.TestingT
	Cleanup(func())
}) *CloseEstimationChecker {
	mock := &CloseEstimationChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
