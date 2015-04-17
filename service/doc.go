// Copyright 2015 JoongSeob Vito Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package service - bussiness logics are contained in separated packages.

HttpStatusCode

Goyangi uses 11 http status code for bussiness logics. It's from net/http/status.go
StatusOK                  = 200
StatusCreated             = 201
StatusSeeOther            = 303
StatusNotModified         = 304
StatusBadRequest          = 400
StatusUnauthorized        = 401
StatusPaymentRequired     = 402
StatusForbidden           = 403
StatusNotFound            = 404
StatusMethodNotAllowed    = 405
StatusInternalServerError = 500
