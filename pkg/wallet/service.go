package wallet

import (
	"github.com/google/uuid"
	"errors"
	
	"github.com/AzizRahimov/wallet/pkg/types"
)


type Error string
func (e Error) Error() string  {
	
	return string(e)
	
}


var ErrPhoneNumberRegistred = errors.New("phone already registred")
var ErrAmountMustBePositive = errors.New("amount must be greater that zero")
var ErrAccountNotFound = errors.New("account not found")
var ErrNotEnoughBalance = errors.New("not enough balance")


//Service -dasdasdsa
type Service struct {
	nextAccountID int64
	accounts      []*types.Account
	payments      []*types.Payment
}

// RegisterAccount создаем тут ак
func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error) {

	for _, account := range s.accounts {
		if account.Phone == phone{
			return nil,  ErrPhoneNumberRegistred
		}
	}

	s.nextAccountID++
	account := &types.Account{
		ID: s.nextAccountID,
		Phone: phone,
		Balance: 0,

	}
		
		s.accounts = append (s.accounts,account)

		return account,  nil
		
		}
	





func (s *Service) Deposit(accountID int64, amount types.Money) error  {
	if amount <= 0{
		return ErrAmountMustBePositive
		
	}

	var account *types.Account

	// Смотри, твоя структура, пока что пустая, одним словом Nil
	for _, acc := range s.accounts {
		if acc.ID == accountID{
			account = acc
		}
		
	}
	if account == nil{
		return ErrAccountNotFound
	}
	account.Balance += amount


	
	return nil
}



func (s *Service) Pay(accountID int64, amount types.Money, category types.PaymentCategory ) (*types.Payment, error)  {
	if amount <= 0{
		return nil, ErrAmountMustBePositive
	}

	var account *types.Account

	for _, acc := range s.accounts {
		if acc.ID == accountID{
			// в прошлый раз мы делали 
			// account := &acc
			// но как оказалось, он с оригиналом почему то не работал
			account = acc
			break
		}
		
	}
	// и делаем проверк на сущ ли акк
	if account == nil{
		return nil, ErrAccountNotFound
	}

	if account.Balance < amount{
		return nil, ErrNotEnoughBalance

	}
	account.Balance -= amount
	paymentID := uuid.New().String()
	payment := &types.Payment{
		ID: paymentID,
		AccountID: accountID,
		Amount: amount,
		Category: category,
		Status: types.PaymentStatusInProgress,
	}
	s.payments =append(s.payments, payment)
	return payment, nil


}


func (s *Service) FindAccountByID(accountID int64) (*types.Account, error) {
	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID{
			account = acc
		}
	}
	if account == nil{
		return nil, ErrAccountNotFound
	}
	return account, nil
	
}