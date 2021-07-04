// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package accounts

import (
	"context"
	"github.com/fernandodr19/library/pkg/domain/vos"
	"sync"
)

// AccountsMockUsecase is a mock implementation of Usecase.
//
// 	func TestSomethingThatUsesUsecase(t *testing.T) {
//
// 		// make and configure a mocked Usecase
// 		mockedUsecase := &AccountsMockUsecase{
// 			CreateAccountFunc: func(contextMoqParam context.Context, email vos.Email, password vos.Password) error {
// 				panic("mock out the CreateAccount method")
// 			},
// 			LoginFunc: func(contextMoqParam context.Context, email vos.Email, password vos.Password) (vos.Tokens, error) {
// 				panic("mock out the Login method")
// 			},
// 		}
//
// 		// use mockedUsecase in code that requires Usecase
// 		// and then make assertions.
//
// 	}
type AccountsMockUsecase struct {
	// CreateAccountFunc mocks the CreateAccount method.
	CreateAccountFunc func(contextMoqParam context.Context, email vos.Email, password vos.Password) error

	// LoginFunc mocks the Login method.
	LoginFunc func(contextMoqParam context.Context, email vos.Email, password vos.Password) (vos.Tokens, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateAccount holds details about calls to the CreateAccount method.
		CreateAccount []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Email is the email argument value.
			Email vos.Email
			// Password is the password argument value.
			Password vos.Password
		}
		// Login holds details about calls to the Login method.
		Login []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Email is the email argument value.
			Email vos.Email
			// Password is the password argument value.
			Password vos.Password
		}
	}
	lockCreateAccount sync.RWMutex
	lockLogin         sync.RWMutex
}

// CreateAccount calls CreateAccountFunc.
func (mock *AccountsMockUsecase) CreateAccount(contextMoqParam context.Context, email vos.Email, password vos.Password) error {
	callInfo := struct {
		ContextMoqParam context.Context
		Email           vos.Email
		Password        vos.Password
	}{
		ContextMoqParam: contextMoqParam,
		Email:           email,
		Password:        password,
	}
	mock.lockCreateAccount.Lock()
	mock.calls.CreateAccount = append(mock.calls.CreateAccount, callInfo)
	mock.lockCreateAccount.Unlock()
	if mock.CreateAccountFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.CreateAccountFunc(contextMoqParam, email, password)
}

// CreateAccountCalls gets all the calls that were made to CreateAccount.
// Check the length with:
//     len(mockedUsecase.CreateAccountCalls())
func (mock *AccountsMockUsecase) CreateAccountCalls() []struct {
	ContextMoqParam context.Context
	Email           vos.Email
	Password        vos.Password
} {
	var calls []struct {
		ContextMoqParam context.Context
		Email           vos.Email
		Password        vos.Password
	}
	mock.lockCreateAccount.RLock()
	calls = mock.calls.CreateAccount
	mock.lockCreateAccount.RUnlock()
	return calls
}

// Login calls LoginFunc.
func (mock *AccountsMockUsecase) Login(contextMoqParam context.Context, email vos.Email, password vos.Password) (vos.Tokens, error) {
	callInfo := struct {
		ContextMoqParam context.Context
		Email           vos.Email
		Password        vos.Password
	}{
		ContextMoqParam: contextMoqParam,
		Email:           email,
		Password:        password,
	}
	mock.lockLogin.Lock()
	mock.calls.Login = append(mock.calls.Login, callInfo)
	mock.lockLogin.Unlock()
	if mock.LoginFunc == nil {
		var (
			tokensOut vos.Tokens
			errOut    error
		)
		return tokensOut, errOut
	}
	return mock.LoginFunc(contextMoqParam, email, password)
}

// LoginCalls gets all the calls that were made to Login.
// Check the length with:
//     len(mockedUsecase.LoginCalls())
func (mock *AccountsMockUsecase) LoginCalls() []struct {
	ContextMoqParam context.Context
	Email           vos.Email
	Password        vos.Password
} {
	var calls []struct {
		ContextMoqParam context.Context
		Email           vos.Email
		Password        vos.Password
	}
	mock.lockLogin.RLock()
	calls = mock.calls.Login
	mock.lockLogin.RUnlock()
	return calls
}
