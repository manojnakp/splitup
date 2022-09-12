package schema

import "time"

type ErrorString struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	ErrorString `json:"error"`
}

type BalanceUnit struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int    `json:"amount"`
}

func DefaultBalance() (BalanceUnit, error) {
	return BalanceUnit{
		From:   "Alex",
		To:     "Julia",
		Amount: 250,
	}, nil
}

type ExpenseUnit struct {
	Name   string    `json:"name"`
	Amount int       `json:"amount"`
	By     string    `json:"by"`
	When   time.Time `json:"when"`
	Among  []int     `json:"among"`
}

func DefaultExpense() (ExpenseUnit, error) {
	datetime, err := time.Parse("2006-01-02", "2016-02-26")
	if err != nil {
		return ExpenseUnit{}, err
	}
	return ExpenseUnit{
		Name:   "Lunch",
		Amount: 300,
		By:     "Julia",
		When:   datetime,
		Among:  []int{120, 60, 60, 60},
	}, nil
}

type SplitCount struct {
	Title        string    `json:"title"`
	Desc         string    `json:"desc"`
	By           string    `json:"by"`
	CreatedOn    time.Time `json:"on"`
	Participants []string  `json:"participants"`
}

func DefaultSplitCount() (SplitCount, error) {
	datetime, err := time.Parse("2006-01-02", "2016-02-26")
	if err != nil {
		return SplitCount{}, err
	}
	return SplitCount{
		Title:     "City Trip",
		Desc:      "Sample splitcount sharing expenses of a city trip among friends",
		By:        "Alex",
		CreatedOn: datetime,
		Participants: []string{
			"Alex",
			"Brian",
			"Julia",
			"Thomas",
		},
	}, nil
}

type MiniSplitCount struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func DefaultMiniSplitup() (MiniSplitCount, error) {
	return MiniSplitCount{
		Title: "City Trip",
		Desc:  "Sample splitcount sharing expenses of a city trip among friends",
	}, nil
}
