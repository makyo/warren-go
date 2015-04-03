// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package models

import (
	"fmt"

	elastigo "github.com/mattbaird/elastigo/lib"
	"gopkg.in/errgo.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/warren-community/warren/contenttype"
)

// An entity represents a post to be displayed on the site.  It has an
// associated content type, is associated with an owner, and may be a share
// of a previous post.
type Entity struct {
	Id              bson.ObjectId `bson:"_id"`
	ContentType     string
	Owner           string
	OriginalOwner   string
	IsShare         bool
	Title           string
	Content         interface{}
	RenderedContent string
	IndexedContent  string
	Assets          []bson.ObjectId
}

// Retrive an entity given an ID
func GetEntity(id string, db *mgo.Database) (Entity, error) {
	var entity Entity
	q := db.C("entities").FindId(bson.ObjectIdHex(id))
	if c, err := q.Count(); c == 0 {
		return entity, err
	}
	err := q.One(&entity)
	return entity, err
}

// Create a new entity with a new ID.
func NewEntity(contentType string, owner string, originalOwner string, isShare bool, title string, content interface{}) Entity {
	return Entity{
		Id:            bson.NewObjectId(),
		ContentType:   contentType,
		Owner:         owner,
		OriginalOwner: originalOwner,
		IsShare:       isShare,
		Title:         title,
		Content:       content,
	}
}

// Save the current entity in the database, generating display and index
// content in the process
func (e *Entity) Save(db *mgo.Database, es *elastigo.Conn) error {
	err := e.updateRenderedContent(false)
	if err != nil {
		return errgo.Mask(err)
	}
	err = e.updateIndexedContent(false)
	if err != nil {
		return errgo.Mask(err)
	}
	_, err = es.Index("warren", "entity", e.Id.Hex(), nil, map[string]string{"indexedContent": e.IndexedContent})
	if err != nil {
		return errgo.Mask(err)
	}
	_, err = db.C("entities").UpsertId(e.Id, e)
	return errgo.Mask(err)
}

// Render the content using the content type renderer for display.
func (e *Entity) updateRenderedContent(allowUnsafe bool) error {
	ct, ok := contenttype.Registry[e.ContentType]
	if !ok {
		ct = contenttype.DefaultContentType
	}
	if !allowUnsafe && !ct.Safe() {
		return fmt.Errorf("Attempted unsafe content-type usage: %s", e.ContentType)
	}
	rendered, err := ct.RenderDisplayContent(e.Content)
	if err != nil {
		return err
	}
	e.RenderedContent = rendered
	return nil
}

// Render the content using the content type renderer for indexing.
func (e *Entity) updateIndexedContent(allowUnsafe bool) error {
	ct, ok := contenttype.Registry[e.ContentType]
	if !ok {
		ct = contenttype.DefaultContentType
	}
	if !allowUnsafe && !ct.Safe() {
		return fmt.Errorf("Attempted unsafe content-type usage: %s", e.ContentType)
	}
	rendered, err := ct.RenderIndexContent(e.Content)
	if err != nil {
		return err
	}
	e.IndexedContent = rendered
	return nil
}

// Delete an entity from the database
func (e *Entity) Delete(db *mgo.Database) error {
	return db.C("entities").RemoveId(e.Id)
}

// Determine whether or not the entity belongs to the user.
func (e *Entity) BelongsToUser(user User) bool {
	return e.Owner == user.Username
}
