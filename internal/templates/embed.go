package templates

import "embed"

// Templates hold all the template files. Line6 tells go to include all the files inside that folder
//
//go:embed express-basic/* express-ts/*
var Templates embed.FS
