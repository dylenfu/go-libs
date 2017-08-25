package datas

import (
	"github.com/dylenfu/go-libs/inject/basedatas"
	"log"
)

type RestApi struct {
	InternalData *basedatas.Base `inject:""`
}

func (s *RestApi) Ring() {
	log.Println("inject\t-", "restapi ring", s.InternalData.Id + " " + s.InternalData.Desc)
}

type JsonApi struct {
	InternalData *basedatas.Base `inject:""`
}

func (s *JsonApi) Loop() {
	log.Println("inject\t-", "jsonapi loop", s.InternalData.Id + " " + s.InternalData.Desc)
}