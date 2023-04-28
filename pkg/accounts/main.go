package accounts

var fakeAccounts = map[string]Account{
	"1": {1, "first", "first"},
	"2": {2, "second", "second"},
}

func Authenticate(token *string) Account {
	return fakeAccounts[*token];
}
