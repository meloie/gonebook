package services

import (
	"context"
	"gonebook/internal/ent"
	"gonebook/internal/ent/contact"
	"gonebook/internal/ent/user"
	"gonebook/pkg/models"

	"github.com/pkg/errors"
)

func (svc *Service) GetUserContacts(ctx context.Context, userId int) ([]*ent.Contact, error) {
	contactsList, err := svc.Database.Contact.Query().
		Where(contact.HasOwnerWith(user.ID(userId))).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "falied to fetch user's contacts")
	}
	return contactsList, nil
}

func (svc *Service) CreateContact(ctx context.Context, data *models.ContactPayload, userID int) (*ent.Contact, error) {
	cnt, err := svc.Database.Contact.Create().SetOwnerID(userID).SetName(data.Name).
		SetPhone(data.Phone).SetAddress(data.Address).Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save contact in database")
	}
	return cnt, err

}
