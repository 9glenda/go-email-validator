package ev

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev/evcache"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	mock_evcache "github.com/go-email-validator/go-email-validator/test/mock/ev/evcache"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_cacheDecorator_Validate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	results := make([]ValidationResult, 0)
	key := validEmail.String()

	type fields struct {
		validator Validator
		cache     func() evcache.Interface
		getKey    CacheKeyGetter
	}
	type args struct {
		email   evmail.Address
		results []ValidationResult
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult ValidationResult
	}{
		{
			name: "without cache, error and with set error",
			fields: fields{
				validator: inValidMockValidator,
				cache: func() evcache.Interface {
					cacheMock := mock_evcache.NewMockInterface(ctrl)
					cacheMock.EXPECT().Get(key).Return(nil, nil).Times(1)
					cacheMock.EXPECT().Set(key, invalidResult).Return(simpleError).Times(1)

					return cacheMock
				},
				getKey: EmailCacheKeyGetter,
			},
			args: args{
				email:   validEmail,
				results: results,
			},
			wantResult: invalidResult,
		},
		{
			name: "without cache and with get error",
			fields: fields{
				validator: validMockValidator,
				cache: func() evcache.Interface {
					cacheMock := mock_evcache.NewMockInterface(ctrl)
					cacheMock.EXPECT().Get(key).Return(nil, simpleError).Times(1)
					cacheMock.EXPECT().Set(key, validResult).Return(nil).Times(1)

					return cacheMock
				},
				getKey: nil,
			},
			args: args{
				email:   validEmail,
				results: results,
			},
			wantResult: validResult,
		},
		{
			name: "with cache",
			fields: fields{
				validator: validMockValidator,
				cache: func() evcache.Interface {
					cacheMock := mock_evcache.NewMockInterface(ctrl)
					cacheMock.EXPECT().Get(key).Return(validResult, nil).Times(1)

					return cacheMock
				},
				getKey: EmailCacheKeyGetter,
			},
			args: args{
				email:   validEmail,
				results: nil,
			},
			wantResult: validResult,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCacheDecorator(tt.fields.validator, tt.fields.cache(), tt.fields.getKey)
			if gotResult := c.Validate(tt.args.email, tt.args.results...); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Validate() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_cacheDecorator_GetDeps(t *testing.T) {

	deps := []ValidatorName{OtherValidator}

	type fields struct {
		validator Validator
		cache     evcache.Interface
		getKey    CacheKeyGetter
	}
	tests := []struct {
		name   string
		fields fields
		want   []ValidatorName
	}{
		{
			name: "return deps",
			fields: fields{
				validator: mockValidator{deps: deps},
			},
			want: deps,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cacheDecorator{
				validator: tt.fields.validator,
				cache:     tt.fields.cache,
				getKey:    tt.fields.getKey,
			}
			if got := c.GetDeps(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDeps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmailCacheKeyGetter(t *testing.T) {
	type args struct {
		email   evmail.Address
		results []ValidationResult
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "success",
			args: args{
				email:   validEmail,
				results: nil,
			},
			want: validEmail.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EmailCacheKeyGetter(tt.args.email, tt.args.results...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmailCacheKeyGetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomainCacheKeyGetter(t *testing.T) {
	type args struct {
		email evmail.Address
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "success",
			args: args{
				email: validEmail,
			},
			want: validEmail.Domain(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DomainCacheKeyGetter(tt.args.email); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DomainCacheKeyGetter() = %v, want %v", got, tt.want)
			}
		})
	}
}