package logic

import (
	"hash/crc32"
)

var _defaultBucketGroup *BucketGroup

const bucketSize = 100

type BucketGroup struct {
	bucketArray [bucketSize]*Hub
	//lockArray   [bucketSize]sync.RWMutex
}

func InitDefaultBucketGroup() {
	_defaultBucketGroup = NewBucketGroup()
}

func GetDefaultBucketGroup() *BucketGroup {
	return _defaultBucketGroup
}

func NewBucketGroup() *BucketGroup {
	group := &BucketGroup{}
	for i := 0; i < bucketSize; i++ {
		hub := NewHub()
		group.bucketArray[i] = hub
		//list.lockArray[i] = sync.RWMutex{}
		go hub.Run()
	}
	return group
}

func (b *BucketGroup) GetHub(identifyCode string) *Hub {
	num := crc32.ChecksumIEEE([]byte(identifyCode))
	return b.bucketArray[num%bucketSize]
}
