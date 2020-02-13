package futures

import (
	"context"
	"encoding/json"
)

// ChangeLeverageService change user's initial leverage of specific symbol market
type ChangeLeverageService struct {
	c        *Client
	symbol   string
	leverage int
}

// Symbol set symbol
func (s *ChangeLeverageService) Symbol(symbol string) *ChangeLeverageService {
	s.symbol = symbol
	return s
}

// Leverage set leverage
func (s *ChangeLeverageService) Leverage(leverage int) *ChangeLeverageService {
	s.leverage = leverage
	return s
}

// Do send request
func (s *ChangeLeverageService) Do(ctx context.Context, opts ...RequestOption) (res *SymbolLeverage, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/fapi/v1/leverage",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"symbol":   s.symbol,
		"leverage": s.leverage,
	})
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SymbolLeverage)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SymbolLeverage define leverage info of symbol
type SymbolLeverage struct {
	Leverage         int    `json:"leverage"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	Symbol           string `json:"symbol"`
}

// ChangeMarginTypeService change user's margin type of specific symbol market
type ChangeMarginTypeService struct {
	c          *Client
	symbol     string
	marginType MarginType
}

// Symbol set symbol
func (s *ChangeMarginTypeService) Symbol(symbol string) *ChangeMarginTypeService {
	s.symbol = symbol
	return s
}

// MarginType set margin type
func (s *ChangeMarginTypeService) MarginType(marginType MarginType) *ChangeMarginTypeService {
	s.marginType = marginType
	return s
}

// Do send request
func (s *ChangeMarginTypeService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   "POST",
		endpoint: "/fapi/v1/marginType",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"symbol":     s.symbol,
		"marginType": s.marginType,
	})
	_, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePositionMarginService update isolated position margin
type UpdatePositionMarginService struct {
	c          *Client
	symbol     string
	amount     string
	actionType int
}

// Symbol set symbol
func (s *UpdatePositionMarginService) Symbol(symbol string) *UpdatePositionMarginService {
	s.symbol = symbol
	return s
}

// Amount set position margin amount
func (s *UpdatePositionMarginService) Amount(amount string) *UpdatePositionMarginService {
	s.amount = amount
	return s
}

// Type set action type: 1: Add postion margin，2: Reduce postion margin
func (s *UpdatePositionMarginService) Type(actionType int) *UpdatePositionMarginService {
	s.actionType = actionType
	return s
}

// Do send request
func (s *UpdatePositionMarginService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   "POST",
		endpoint: "/fapi/v1/positionMargin",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"symbol": s.symbol,
		"amount": s.amount,
		"type":   s.actionType,
	})
	_, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

// PositionInformation define deep position info
type PositionInformation struct {
	EntryPrice       string `json:"entryPrice"`
	MarginType       string `json:"marginType"`
	IsAutoAddMargin  string `json:"isAutoAddMargin"`
	IsolatedMargin   string `json:"isolatedMargin"`
	Leverage         string `json:"leverage"`
	LiquidationPrice string `json:"liquidationPrice"`
	MarkPrice        string `json:"markPrice"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	PositionAmt      string `json:"positionAmt"`
	Symbol           string `json:"symbol"`
	UnRealizedProfit string `json:"unRealizedProfit"`
}

// GetPositionInformationService get full position info
type GetPositionInformationService struct {
	c *Client
}

// Do send request
func (s *GetPositionInformationService) Do(ctx context.Context, opts ...RequestOption) (res []*PositionInformation, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/positionRisk",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = make([]*PositionInformation, 0)
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}
	return res, nil
}
