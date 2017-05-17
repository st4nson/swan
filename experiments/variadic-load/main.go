// Copyright (c) 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"math"
	"os"
	"time"

	"github.com/intelsdi-x/swan/experiments/memcached-sensitivity-profile/common"
	"github.com/intelsdi-x/swan/pkg/executor"
	"github.com/intelsdi-x/swan/pkg/experiment"
	"github.com/intelsdi-x/swan/pkg/experiment/logger"
	"github.com/intelsdi-x/swan/pkg/experiment/sensitivity"
	"github.com/intelsdi-x/swan/pkg/experiment/sensitivity/topology"
	"github.com/intelsdi-x/swan/pkg/experiment/sensitivity/validate"
	"github.com/intelsdi-x/swan/pkg/utils/errutil"
	_ "github.com/intelsdi-x/swan/pkg/utils/unshare"
	"github.com/intelsdi-x/swan/pkg/workloads/memcached"

	"github.com/Sirupsen/logrus"
	"github.com/intelsdi-x/swan/pkg/utils/uuid"
)

var (
	appName = os.Args[0]
)

func main() {
	// Preparing application - setting name, help, aprsing flags etc.
	experimentStart := time.Now()
	experiment.Configure()

	// Generate an experiment ID and start the metadata session.
	uid := uuid.New()

	// Initialize logger.
	logger.Initialize(appName, uid)

	// Validate preconditions.
	validate.OS()

	// Create isolations.
	hpIsolation, l1Isolation, llcIsolation := topology.NewIsolations()

	// Create executors with cleanup function.
	hpExecutor, beExecutorFactory, cleanup, err := sensitivity.PrepareExecutors(hpIsolation)
	errutil.Check(err)
	defer cleanup()

	// Create BE workloads.
	beLaunchers, err := sensitivity.PrepareAggressors(l1Isolation, llcIsolation, beExecutorFactory)
	errutil.Check(err)

	// Create HP workload.
	memcachedConfig := memcached.DefaultMemcachedConfig()
	hpLauncher := executor.ServiceLauncher{Launcher: memcached.New(hpExecutor, memcachedConfig)}

	// Load generator.
	loadGenerator, err := common.PrepareMutilateGenerator(memcachedConfig.IP, memcachedConfig.Port)
	errutil.Check(err)

	// Peak load
	peakload := sensitivity.PeakLoadFlag.Value()
	logrus.Infof("using peakload %d", peakload)

	loadPoints := sensitivity.LoadPointsCountFlag.Value()
	loadDuration := sensitivity.LoadDurationFlag.Value()
	repetitions := sensitivity.RepetitionsFlag.Value()

	for _, beLauncher := range beLaunchers {

		// hp
		hpHandle, err := hpLauncher.Launch()
		errutil.Check(err)

		// populate
		errutil.Check(loadGenerator.Populate())

		// be
		beHandle, err := beLauncher.Launcher.Launch()
		errutil.Check(err)

		for repetition := 0; repetition < repetitions; repetition++ {

			for loadPoint := 0; loadPoint < loadPoints; loadPoint++ {
				qps := int(math.Sin(float64(loadPoint)/float64(loadPoints)*math.Pi) * float64(peakload))
				if qps < 10000 {
					qps = 10000
				}
				logrus.Info("qps:", qps)
				// load gen
				loadGeneratorHandle, err := loadGenerator.Load(qps, loadDuration)
				errutil.Check(err)
				loadGeneratorHandle.Wait(0)
			}
		}
		hpHandle.Stop()
		beHandle.Stop()
	}
	logrus.Infof("Ended experiment %s with uid %s in %s", appName, uid, time.Since(experimentStart).String())
}
