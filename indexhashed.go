package carbs

import (
	"io"

	"github.com/ipfs/go-cid"
	cbor "github.com/whyrusleeping/cbor/go"
)

type mapIndex map[cid.Cid]uint64

func (m *mapIndex) Get(c cid.Cid) uint64 {
	return (*m)[c]
}

func (m *mapIndex) Marshal(w io.Writer) error {
	return cbor.Encode(w, m)
}

func (m *mapIndex) Unmarshal(r io.Reader) error {
	d := cbor.NewDecoder(r)
	return d.Decode(&m)
}

func (m *mapIndex) Codec() IndexCodec {
	return IndexHashed
}

func (m *mapIndex) Load(rs []Record) error {
	for _, r := range rs {
		(*m)[r.Cid] = r.idx
	}
	return nil
}

func mkHashed() Index {
	mi := make(mapIndex)
	return &mi
}
