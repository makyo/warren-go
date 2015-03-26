// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package models

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/warren-community/warren/contenttype"
)

type Entity struct {
	Id              bson.ObjectId `bson:"_id"`
	ContentType     string
	Owner           string
	OriginalOwner   string
	IsShare         bool
	Title           string
	Content         string
	RenderedContent string
	IndexedContent  string
	Assets          []bson.ObjectId
}

func GetEntity(id string, db *mgo.Database) (Entity, error) {
	var entity Entity
	q := db.C("entities").FindId(bson.ObjectIdHex(id))
	if c, err := q.Count(); c == 0 {
		return entity, err
	}
	err := q.One(&entity)
	return entity, err
}

func NewEntity(contentType string, owner string, originalOwner string, isShare bool, title string, content string) Entity {
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

func (e *Entity) Save(db *mgo.Database) error {
	err := e.updateRenderedContent(false)
	if err != nil {
		return err
	}
	err = e.updateIndexedContent(false)
	if err != nil {
		return err
	}
	_, err = db.C("entities").UpsertId(e.Id, e)
	return err
}

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

func (e *Entity) Delete(db *mgo.Database) error {
	return db.C("entities").RemoveId(e.Id)
}

func (e *Entity) BelongsToUser(user User) bool {
	return e.Owner == user.Username
}
