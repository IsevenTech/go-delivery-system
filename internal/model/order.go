package model

type Order struct {
    ID       int    `json:"id"`
    Customer string `json:"customer"`
    Address  string `json:"address"`
    Status   string `json:"status"`
}
