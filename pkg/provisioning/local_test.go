package provisioning

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
	"github.com/intelsdi-x/swan/pkg/isolation"
)

func TestShell(t *testing.T) {
	Convey("Creating a new shell with `sleep 3`", t, func() {
		l := NewLocal("root", []isolation.Isolation{})

		Convey("Should take more than three seconds to execute", func() {
			start := time.Now()

			task := l.Run("sleep 3")

			task.Wait(0)

			duration := time.Since(start)
			durationsMs := duration.Nanoseconds() / 1e6

			Convey("The Duration should last longer than 3s", func() {
				So(durationsMs, ShouldBeGreaterThan, 3000)
			})

			Convey("And the exit status should be zero", func() {
				So(task.Status().code, ShouldEqual, 0)
			})
		})

		Convey("Should NOT take more than three seconds to execute, since timeout is 1s",
				func() {
			start := time.Now()

			task := l.Run("sleep 3")

			task.Wait(1)

			duration := time.Since(start)
			durationsMs := duration.Nanoseconds() / 1e6

			Convey("The Duration should last less than 3s", func() {
				So(durationsMs, ShouldBeLessThan, 3000)
			})

			Convey("And the exit status should be zero", func() {
				So(task.Status().code, ShouldEqual, 0)
			})
		})
	})
}
