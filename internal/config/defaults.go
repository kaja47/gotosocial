/*
   GoToSocial
   Copyright (C) 2021-2022 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package config

import "github.com/coreos/go-oidc/v3/oidc"

// Defaults contains a populated Configuration with reasonable defaults. Note that
// if you use this, you will still need to set Host, and, if desired, ConfigPath.
var Defaults = Configuration{
	LogLevel:        "info",
	LogDbQueries:    false,
	ApplicationName: "gotosocial",
	LandingPageUser: "",
	ConfigPath:      "",
	Host:            "",
	AccountDomain:   "",
	Protocol:        "https",
	BindAddress:     "0.0.0.0",
	Port:            8080,
	TrustedProxies:  []string{"127.0.0.1/32", "::1"}, // localhost

	DbType:      "postgres",
	DbAddress:   "",
	DbPort:      5432,
	DbUser:      "",
	DbPassword:  "",
	DbDatabase:  "gotosocial",
	DbTLSMode:   "disable",
	DbTLSCACert: "",

	WebTemplateBaseDir: "./web/template/",
	WebAssetBaseDir:    "./web/assets/",

	InstanceExposePeers:            false,
	InstanceExposeSuspended:        false,
	InstanceDeliverToSharedInboxes: true,

	AccountsRegistrationOpen: true,
	AccountsApprovalRequired: true,
	AccountsReasonRequired:   true,
	AccountsAllowCustomCSS:   false,

	MediaImageMaxSize:        10485760, // 10mb
	MediaVideoMaxSize:        41943040, // 40mb
	MediaDescriptionMinChars: 0,
	MediaDescriptionMaxChars: 500,
	MediaRemoteCacheDays:     30,
	MediaEmojiLocalMaxSize:   51200,  // 50kb
	MediaEmojiRemoteMaxSize:  102400, // 100kb

	StorageBackend:       "local",
	StorageLocalBasePath: "/gotosocial/storage",
	StorageS3UseSSL:      true,

	StatusesMaxChars:           5000,
	StatusesCWMaxChars:         100,
	StatusesPollMaxOptions:     6,
	StatusesPollOptionMaxChars: 50,
	StatusesMediaMaxFiles:      6,

	LetsEncryptEnabled:      false,
	LetsEncryptPort:         80,
	LetsEncryptCertDir:      "/gotosocial/storage/certs",
	LetsEncryptEmailAddress: "",

	OIDCEnabled:          false,
	OIDCIdpName:          "",
	OIDCSkipVerification: false,
	OIDCIssuer:           "",
	OIDCClientID:         "",
	OIDCClientSecret:     "",
	OIDCScopes:           []string{oidc.ScopeOpenID, "profile", "email", "groups"},

	SMTPHost:     "",
	SMTPPort:     0,
	SMTPUsername: "",
	SMTPPassword: "",
	SMTPFrom:     "GoToSocial",

	SyslogEnabled:  false,
	SyslogProtocol: "udp",
	SyslogAddress:  "localhost:514",

	AdvancedCookiesSamesite:   "lax",
	AdvancedRateLimitRequests: 1000, // per 5 minutes
}
