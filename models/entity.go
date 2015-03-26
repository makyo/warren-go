// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
		ContentType:   contentType,
		Owner:         owner,
		OriginalOwner: originalOwner,
		IsShare:       isShare,
		Title:         title,
		Content:       content,
	}
}

func (e *Entity) Save(db *mgo.Database) error {
	e.updateRenderedContent()
	e.updateIndexedContent()
	if e.Id.Hex() == "" {
		e.Id = bson.NewObjectId()
	}
	_, err := db.C("entities").UpsertId(e.Id, e)
	return err
}

func (e *Entity) updateRenderedContent() {
	e.RenderedContent = e.Content
}

func (e *Entity) updateIndexedContent() {
	e.IndexedContent = e.Content
}

func (e *Entity) Delete(db *mgo.Database) error {
	return db.C("entities").RemoveId(e.Id)
}
