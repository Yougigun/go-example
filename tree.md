.
├── MapNotConcurrentWrite
│   └── main.go
├── MarshallNotThreadSave
│   └── main.go
├── atomic.value
│   └── main.go
├── base64
│   └── main.go
├── bigint
│   └── maing.go
├── browser-cache
│   └── main.go
├── bufio
│   ├── ScannerToReplace
│   └── main.go
│
├── conditions
│   ├── Example.flow
│   ├── Example.png
│   ├── LICENSE
│   ├── README.md
│   ├── ast.go
│   ├── doc.go
│   ├── evaluator.go
│   ├── parser.go
│   ├── parser_test.go
│   └── token.go
├── fsnotify
│   └── main.go
├── generatetoken
│   └── main.go
├── generic
│   ├── go.mod
│   └── main.go
├── gin-abort-next
│   └── main.go
├── go-test
│   ├── Summary.md
│   ├── math.go
│   ├── mathadd.go
│   ├── mathadd_test.go
│   └── othermath
│       ├── c.out
│       └── math.go
├── go.mod
├── go.sum
├── goroutine-pattern
│   └── generator.go
├── govaluate
│   └── main.go
├── gracefulshutdown
│   └── main.go
├── json
│   └── main.go
├── jsonCustomUnmarshal
│   └── main.go
├── main.go
├── mapCurrentOps
│   └── main.go
├── mq
├── multiservice
│   └── main.go
├── mysql
│   └── main.go
├── okta
│   ├── CONTRIBUTING.md
│   ├── LICENSE.md
│   ├── README.md
│   ├── custom-login
│   │   ├── README.md
│   │   ├── main.go
│   │   ├── templates
│   │   │   ├── _footer.gohtml
│   │   │   ├── _header.gohtml
│   │   │   ├── home.gohtml
│   │   │   ├── login.gohtml
│   │   │   └── profile.gohtml
│   │   └── utils
│   │       ├── nonce.go
│   │       └── parseEnv.go
│   ├── identity-engine
│   │   ├── embedded-auth-with-sdk
│   │   │   ├── README.md
│   │   │   ├── config
│   │   │   │   └── config.go
│   │   │   ├── css
│   │   │   │   └── samples.css
│   │   │   ├── feature_runner_test.go
│   │   │   ├── features
│   │   │   │   ├── 00_root_page.feature
│   │   │   │   ├── 01_basic_login.feature
│   │   │   │   ├── 03_password_recovery.feature
│   │   │   │   ├── 04_self_service_registration.feature
│   │   │   │   ├── 05_1_1_social_idp.feature
│   │   │   │   ├── 05_1_2_social_idp.feature
│   │   │   │   ├── 06_1_2fa_auth.feature
│   │   │   │   └── 06_2_2fa_auth.feature
│   │   │   ├── go.mod
│   │   │   ├── go.sum
│   │   │   ├── harness
│   │   │   │   ├── sdk.go
│   │   │   │   ├── testing_harness.go
│   │   │   │   └── testing_harness_util.go
│   │   │   ├── main.go
│   │   │   ├── server
│   │   │   │   ├── login.go
│   │   │   │   └── server.go
│   │   │   └── views
│   │   │       ├── _error.gohtml
│   │   │       ├── _flows.gohtml
│   │   │       ├── _footer.gohtml
│   │   │       ├── _header.gohtml
│   │   │       ├── _serverConfig.gohtml
│   │   │       ├── enroll.gohtml
│   │   │       ├── enrollEmail.gohtml
│   │   │       ├── enrollPassword.gohtml
│   │   │       ├── enrollPhone.gohtml
│   │   │       ├── enrollPhoneCode.gohtml
│   │   │       ├── enrollPhoneMethod.gohtml
│   │   │       ├── home.gohtml
│   │   │       ├── login.gohtml
│   │   │       ├── loginFactorEmail.gohtml
│   │   │       ├── loginFactorPhone.gohtml
│   │   │       ├── loginFactorPhoneMethod.gohtml
│   │   │       ├── loginSecondaryFactors.gohtml
│   │   │       ├── register.gohtml
│   │   │       ├── resetPassword.gohtml
│   │   │       ├── resetPasswordCode.gohtml
│   │   │       ├── resetPasswordNewPassword.gohtml
│   │   │       ├── verify.gohtml
│   │   │       ├── verifyEmailCode.gohtml
│   │   │       └── views.go
│   │   └── embedded-sign-in-widget
│   │       ├── README.md
│   │       ├── config
│   │       │   └── config.go
│   │       ├── feature_runner_test.go
│   │       ├── features
│   │       │   ├── 07_1_1_social_login.feature
│   │       │   ├── 07_1_2_social_login.feature
│   │       │   └── 08_basic_login.feature
│   │       ├── go.mod
│   │       ├── go.sum
│   │       ├── harness
│   │       │   ├── harness.go
│   │       │   ├── sdk.go
│   │       │   └── util.go
│   │       ├── main.go
│   │       ├── server
│   │       │   └── server.go
│   │       └── templates
│   │           ├── _footer.gohtml
│   │           ├── _header.gohtml
│   │           ├── home.gohtml
│   │           ├── login.gohtml
│   │           └── profile.gohtml
│   ├── namespace-themis.yaml
│   ├── okta-hosted-login
│   │   ├── README.md
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── main.go
│   │   ├── main_test.go
│   │   ├── templates
│   │   │   ├── _footer.gohtml
│   │   │   ├── _header.gohtml
│   │   │   ├── home.gohtml
│   │   │   └── profile.gohtml
│   │   └── utils
│   │       ├── nonce.go
│   │       └── parseEnv.go
│   ├── package-lock.json
│   ├── package.json
│   ├── resource-server
│   │   ├── README.md
│   │   ├── main.go
│   │   └── utils
│   │       ├── nonce.go
│   │       └── parseEnv.go
│   ├── scripts
│   │   ├── setup-env.js
│   │   └── update-se-drivers.js
│   └── yarn.lock
├── opa
│   ├── example.rego
│   ├── go.mod
│   ├── go.sum
│   ├── iam
│   ├── iam.rego
│   ├── input.json
│   ├── input_example.json
│   ├── main.go
│   ├── sample copy.rego
│   └── sample.rego
├── pktest
│   └── pktest.go
├── priorityQueue
│   └── main.go
├── proxy
│   └── main.go
├── reflection
│   ├── main.go
│   ├── reflect.Method
│   │   └── main.go
│   ├── reflect.StructField
│   │   └── main.go
│   ├── reflect.Type
│   │   └── main.go
│   └── reflect.Value
│       └── main.go
├── serverMethod
│   └── main.go
├── tickerBuffer
│   └── main.go
├── tickleak
│   └── main.go
├── tree.md
└── unitTest
    ├── c.out
    ├── main.go
    └── main_test.go

65 directories, 165 files
