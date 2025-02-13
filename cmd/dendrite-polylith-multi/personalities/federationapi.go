// Copyright 2020 The Matrix.org Foundation C.I.C.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package personalities

import (
	"github.com/matrix-org/dendrite/federationapi"
	"github.com/matrix-org/dendrite/setup"
	"github.com/matrix-org/dendrite/setup/config"
)

func FederationAPI(base *setup.BaseDendrite, cfg *config.Dendrite) {
	userAPI := base.UserAPIClient()
	federation := base.CreateFederationClient()
	serverKeyAPI := base.SigningKeyServerHTTPClient()
	keyRing := serverKeyAPI.KeyRing()
	fsAPI := base.FederationSenderHTTPClient()
	rsAPI := base.RoomserverHTTPClient()
	keyAPI := base.KeyServerHTTPClient()

	federationapi.AddPublicRoutes(
		base.PublicFederationAPIMux, base.PublicKeyAPIMux, base.PublicWellKnownAPIMux,
		&base.Cfg.FederationAPI, userAPI, federation, keyRing,
		rsAPI, fsAPI, base.EDUServerClient(), keyAPI,
		&base.Cfg.MSCs, nil,
	)

	base.SetupAndServeHTTP(
		base.Cfg.FederationAPI.InternalAPI.Listen,
		base.Cfg.FederationAPI.ExternalAPI.Listen,
		nil, nil,
	)
}
