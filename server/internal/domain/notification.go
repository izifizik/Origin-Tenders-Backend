package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Notification struct {
	ID            int                `json:"id,omitempty"`
	Type          string             `json:"type,omitempty"`
	Description   string             `json:"description,omitempty"`
	TenderID      primitive.ObjectID `json:"tender_id,omitempty"`
	TenderName    string             `json:"tender_name,omitempty"`
	Price         float64            `json:"price,omitempty"`
	IsNeedApprove bool               `json:"is_need_approve,omitempty"`
}
