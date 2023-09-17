package conf

import "os"

var ColName string = "links"
var DbName string = "ZipLink"
var ConnectionString string = "mongodb://localhost:27017"

var ENV string = os.Getenv("STAGE")
var SentryDSN string = os.Getenv("SENTRY_DSN")

// Server ENV constants
const ENV_PROD = "prod"
