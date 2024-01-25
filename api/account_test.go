package api

//func TestGetAccount(t *testing.T) {
//
//	account := randomAccount()
//
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	store := mockdb.NewMockStore(ctrl)
//	store.EXPECT().
//		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
//		Times(1).
//		Return(account, nil)
//
//	// start test server and send request
//	server := NewTestServer(t, store)
//	recorder := httptest.NewRecorder()
//
//	url := fmt.Sprintf("/accounts/%d", account.ID)
//	request, err := http.NewRequest(http.MethodGet, url, nil)
//	require.NoError(t, err)
//
//	server.router.ServeHTTP(recorder, request)
//	// check response
//	require.Equal(t, http.StatusOK, recorder.Code)
//}
//
//func randomAccount() db.Account {
//	return db.Account{
//		ID:       util.RandomInt(1, 1000),
//		Owner:    util.RandomOwner(),
//		Balance:  util.RandomMoney(),
//		Currency: util.RandomCurrency(),
//	}
//}
