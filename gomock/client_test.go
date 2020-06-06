package gomock

import (
	"fmt"
	"testing"

	"github.com/budougumi0617/til/go/gomock/mock"
	"github.com/golang/mock/gomock"
)

func TestReturn(t *testing.T) {
	in := "any value"
	want := "hoge"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mc := mock.NewMockClient(ctrl)

	mc.EXPECT().Method(in).Return("hoge", nil)

	if got, _ := mc.Method(in); got != want {
		t.Errorf("want %#v, but got %#v\n", want, got)
	}

}

func TestDoAndReturn(t *testing.T) {
	in := "value"
	want := "value modified in mock"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mc := mock.NewMockClient(ctrl)

	mc.EXPECT().Method(in).DoAndReturn(
		func(in string) (string, error) {
			return fmt.Sprint(in, " modified in mock"), nil
		})
	if got, _ := mc.Method(in); got != want {
		t.Errorf("want %#v, but got %#v\n", want, got)
	}

}

func TestDoAndReturn2(t *testing.T) {

	in := "any value"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mc := mock.NewMockClient(ctrl)

	mc.EXPECT().Method(in).DoAndReturn(
		func(in string) (string, error) {
			return "", fmt.Errorf("%s", in)
		})

	if _, err := mc.Method(in); err == nil {
		t.Error("cannot get error")
	}

}

func TestDoAndReturn3(t *testing.T) {

	errin := "raise error"
	specify := "specify"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mc := mock.NewMockClient(ctrl)

	mc.EXPECT().Method(gomock.Any()).DoAndReturn(
		func(in string) (string, error) {
			switch in {
			case errin:
				return "", fmt.Errorf("%s", in)
			case specify:
				return "!!!!!", nil
			}
			return in, nil
		}).AnyTimes() // By default, method is called only once.

	if _, err := mc.Method(errin); err == nil {
		t.Error("cannot get error")
	}

	in := "hogehoge"
	if got, _ := mc.Method(in); got != in {
		t.Errorf("want %s, but got %s\n", in, got)
	}

	if got, _ := mc.Method(specify); got != "!!!!!" {
		t.Errorf("want %s, but got %s\n", "!!!!!", got)
	}
}

func TestBySubTest(t *testing.T) {
	tests := []struct {
		name string
		// Inject rule to mock.
		setClient func(*mock.MockClient, string)
		in        string
		want      string
		wantErr   bool
	}{
		{
			name: "Return",
			setClient: func(mc *mock.MockClient, in string) {
				mc.EXPECT().Method(in).Return("hoge", nil)
			},
			in:      "any value",
			want:    "hoge",
			wantErr: false,
		},
		{
			name: "DoAndReturn",
			setClient: func(mc *mock.MockClient, in string) {
				mc.EXPECT().Method(in).DoAndReturn(
					func(in string) (string, error) {
						return in, nil
					})
			},
			in:      "input value",
			want:    "input value",
			wantErr: false,
		},
		{
			name: "DoAndReturn2",
			setClient: func(mc *mock.MockClient, in string) {
				mc.EXPECT().Method(in).DoAndReturn(
					func(in string) (string, error) {
						return "", fmt.Errorf("%s", in)
					})
			},
			in:      "any  value",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mc := mock.NewMockClient(ctrl)
			tt.setClient(mc, tt.in)
			got, err := mc.Method(tt.in)
			if !tt.wantErr && err != nil {
				t.Fatal(err)
			}

			if got != tt.want {
				t.Errorf("want %#v, but got %#v\n", tt.want, got)
			}
		})
	}

}