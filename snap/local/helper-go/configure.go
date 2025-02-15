/*
 * Copyright (C) 2021 Canonical Ltd
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *
 * SPDX-License-Identifier: Apache-2.0'
 */

package main

import (
	hooks "github.com/canonical/edgex-snap-hooks/v2"
	"github.com/canonical/edgex-snap-hooks/v2/log"
	"github.com/canonical/edgex-snap-hooks/v2/options"
)

// ConfToEnv defines mappings from snap config keys to EdgeX environment variable
// names that are used to override individual app-rfid-llrp-inventory's [AppCustom]  configuration
// values via a .env file read by the snap service wrapper.
//
// The syntax to set a configuration key is:
//
// env.<section>.<keyname>
//
var ConfToEnv = map[string]string{

	//  [AppCustom.Aliases]
	"appcustom.appsettings.device-service-name":             "APPCUSTOM_APPSETTINGS_DEVICESERVICENAME",
	"appcustom.appsettings.adjust-last-read-on-by-origin":   "APPCUSTOM_APPSETTINGS_ADJUSTLASTREADONBYORIGIN",
	"appcustom.appsettings.departed-threshold-seconds":      "APPCUSTOM_APPSETTINGS_DEPARTEDTHRESHOLDSECONDS",
	"appcustom.appsettings.departed-check-interval-seconds": "APPCUSTOM_APPSETTINGS_DEPARTEDCHECKINTERVALSECONDS",
	"appcustom.appsettings.age-out-hours":                   "APPCUSTOM_APPSETTINGS_AGEOUTHOURS",
	"appcustom.appsettings.mobility-profile-threshold":      "APPCUSTOM_APPSETTINGS_MOBILITYPROFILETHRESHOLD",
	"appcustom.appsettings.mobility-profile-holdoff-millis": "APPCUSTOM_APPSETTINGS_MOBILITYPROFILEHOLDOFFMILLIS",
	"appcustom.appsettings.mobility-profile-slope":          "APPCUSTOM_APPSETTINGS_MOBILITYPROFILESLOPE",
}

// configure is called by the main function
func configure() {

	const service = "app-rfid-llrp-inventory"

	log.SetComponentName("configure")

	log.Info("Processing legacy env options")
	envJSON, err := hooks.NewSnapCtl().Config(hooks.EnvConfig)
	if err != nil {
		log.Fatalf("Reading config 'env' failed: %v", err)
	}
	if envJSON != "" {
		log.Debugf("envJSON: %s", envJSON)
		err = hooks.HandleEdgeXConfig(service, envJSON, ConfToEnv)
		if err != nil {
			log.Fatalf("HandleEdgeXConfig failed: %v", err)
		}
	}

	log.Info("Processing config options")
	err = options.ProcessConfig(service)
	if err != nil {
		log.Fatalf("could not process config options: %v", err)
	}

	log.Info("Processing autostart options")
	err = options.ProcessAutostart(service)
	if err != nil {
		log.Fatalf("could not process autostart options: %v", err)
	}
}
