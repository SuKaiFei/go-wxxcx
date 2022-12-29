package service

import (
	"context"
	"testing"

	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServiceGetCosCredential(t *testing.T) {
	Convey("GetCosCredential", t, func() {
		var (
			ctx = context.Background()
			in  = &v1.GetCosCredentialRequest{}
		)
		Convey("When everything goes positive", func() {
			p1, err := tSVC.communitySvc.GetCosCredential(ctx, in)
			t.Logf("%+v", p1)
			Convey("Then err should be nil.p1 should not be nil.", func() {
				So(err, ShouldBeNil)
				So(p1, ShouldNotBeNil)
			})
		})
	})
}
