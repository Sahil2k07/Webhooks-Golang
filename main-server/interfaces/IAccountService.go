package interfaces

type IAccountService interface {
	UpgradeUserAccount() error
	MakeUserPayment() error
}
