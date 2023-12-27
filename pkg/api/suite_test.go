package api

import (
	"github.com/tonkeeper/opentonapi/pkg/addressbook"
	"github.com/tonkeeper/tongo"
)

type mockAddressBook struct {
}

func (m mockAddressBook) GetAddressInfoByAddress(a tongo.AccountID) (addressbook.KnownAddress, bool) {
	return addressbook.KnownAddress{}, false
}

func (m mockAddressBook) GetCollectionInfoByAddress(a tongo.AccountID) (addressbook.KnownCollection, bool) {
	return addressbook.KnownCollection{}, false
}

func (m mockAddressBook) GetJettonInfoByAddress(a tongo.AccountID) (addressbook.KnownJetton, bool) {
	return addressbook.KnownJetton{}, false
}

func (m mockAddressBook) GetTFPoolInfo(a tongo.AccountID) (addressbook.TFPoolInfo, bool) {
	return addressbook.TFPoolInfo{}, false
}

func (m mockAddressBook) GetKnownJettons() map[tongo.AccountID]addressbook.KnownJetton {
	return map[tongo.AccountID]addressbook.KnownJetton{}
}

func (m mockAddressBook) GetKnownCollections() map[tongo.AccountID]addressbook.KnownCollection {
	return map[tongo.AccountID]addressbook.KnownCollection{}
}

func (m mockAddressBook) SearchAttachedAccountsByPrefix(prefix string) []addressbook.AttachedAccount {
	return []addressbook.AttachedAccount{}
}

var _ addressBook = &mockAddressBook{}
