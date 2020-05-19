// Copyright (C) MongoDB, Inc. 2019-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

// Code generated by operationgen. DO NOT EDIT.

package operation

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"go.mongodb.org/mongo-driver/x/mongo/driver"
	"go.mongodb.org/mongo-driver/x/mongo/driver/description"
	"go.mongodb.org/mongo-driver/x/mongo/driver/session"
)

// ListCollections performs a listCollections operation.
type ListCollections struct {
	filter         bsoncore.Document
	nameOnly       *bool
	session        *session.Client
	clock          *session.ClusterClock
	monitor        *event.CommandMonitor
	crypt          *driver.Crypt
	database       string
	deployment     driver.Deployment
	readPreference *readpref.ReadPref
	selector       description.ServerSelector
	retry          *driver.RetryMode
	result         driver.CursorResponse
}

// NewListCollections constructs and returns a new ListCollections.
func NewListCollections(filter bsoncore.Document) *ListCollections {
	return &ListCollections{
		filter: filter,
	}
}

// Result returns the result of executing this operation.
func (lc *ListCollections) Result(opts driver.CursorOptions) (*driver.ListCollectionsBatchCursor, error) {
	bc, err := driver.NewBatchCursor(lc.result, lc.session, lc.clock, opts)
	if err != nil {
		return nil, err
	}
	desc := lc.result.Desc
	if desc.WireVersion == nil || desc.WireVersion.Max < 3 {
		return driver.NewLegacyListCollectionsBatchCursor(bc)
	}
	return driver.NewListCollectionsBatchCursor(bc)
}

func (lc *ListCollections) processResponse(response bsoncore.Document, srvr driver.Server, desc description.Server) error {
	var err error
	lc.result, err = driver.NewCursorResponse(response, srvr, desc)
	return err
}

// Execute runs this operations and returns an error if the operaiton did not execute successfully.
func (lc *ListCollections) Execute(ctx context.Context) error {
	if lc.deployment == nil {
		return errors.New("the ListCollections operation must have a Deployment set before Execute can be called")
	}

	return driver.Operation{
		CommandFn:         lc.command,
		ProcessResponseFn: lc.processResponse,
		RetryMode:         lc.retry,
		Type:              driver.Read,
		Client:            lc.session,
		Clock:             lc.clock,
		CommandMonitor:    lc.monitor,
		Crypt:             lc.crypt,
		Database:          lc.database,
		Deployment:        lc.deployment,
		ReadPreference:    lc.readPreference,
		Selector:          lc.selector,
		Legacy:            driver.LegacyListCollections,
	}.Execute(ctx, nil)

}

func (lc *ListCollections) command(dst []byte, desc description.SelectedServer) ([]byte, error) {

	dst = bsoncore.AppendInt32Element(dst, "listCollections", 1)
	if lc.filter != nil {
		dst = bsoncore.AppendDocumentElement(dst, "filter", lc.filter)
	}
	if lc.nameOnly != nil {
		dst = bsoncore.AppendBooleanElement(dst, "nameOnly", *lc.nameOnly)
	}
	return dst, nil
}

// Filter determines what results are returned from listCollections.
func (lc *ListCollections) Filter(filter bsoncore.Document) *ListCollections {
	if lc == nil {
		lc = new(ListCollections)
	}

	lc.filter = filter
	return lc
}

// NameOnly specifies whether to only return collection names.
func (lc *ListCollections) NameOnly(nameOnly bool) *ListCollections {
	if lc == nil {
		lc = new(ListCollections)
	}

	lc.nameOnly = &nameOnly
	return lc
}

// Session sets the session for this operation.
func (lc *ListCollections) Session(session *session.Client) *ListCollections {
	if lc == nil {
		lc = new(ListCollections)
	}

	lc.session = session
	return lc
}

// ClusterClock sets the cluster clock for this operation.
func (lc *ListCollections) ClusterClock(clock *session.ClusterClock) *ListCollections {
	if lc == nil {
		lc = new(ListCollections)
	}

	lc.clock = clock
	return lc
}

// CommandMonitor sets the monitor to use for APM events.
func (lc *ListCollections) CommandMonitor(monitor *event.CommandMonitor) *ListCollections {
	if lc == nil {
		lc = new(ListCollections)
	}

	lc.monitor = monitor
	return lc
}

// Crypt sets the Crypt object to use for automatic encryption and decryption.
func (lc *ListCollections) Crypt(crypt *driver.Crypt) *ListCollections {
	if lc == nil {
		lc = new(ListCollections)
	}

	lc.crypt = crypt
	return lc
}

// Database sets the database to run this operation against.
func (lc *ListCollections) Database(database string) *ListCollections {
	if lc == nil {
		lc = new(ListCollections)
	}

	lc.database = database
	return lc
}

// Deployment sets the deployment to use for this operation.
func (lc *ListCollections) Deployment(deployment driver.Deployment) *ListCollections {
	if lc == nil {
		lc = new(ListCollections)
	}

	lc.deployment = deployment
	return lc
}

// ReadPreference set the read prefernce used with this operation.
func (lc *ListCollections) ReadPreference(readPreference *readpref.ReadPref) *ListCollections {
	if lc == nil {
		lc = new(ListCollections)
	}

	lc.readPreference = readPreference
	return lc
}

// ServerSelector sets the selector used to retrieve a server.
func (lc *ListCollections) ServerSelector(selector description.ServerSelector) *ListCollections {
	if lc == nil {
		lc = new(ListCollections)
	}

	lc.selector = selector
	return lc
}

// Retry enables retryable mode for this operation. Retries are handled automatically in driver.Operation.Execute based
// on how the operation is set.
func (lc *ListCollections) Retry(retry driver.RetryMode) *ListCollections {
	if lc == nil {
		lc = new(ListCollections)
	}

	lc.retry = &retry
	return lc
}
