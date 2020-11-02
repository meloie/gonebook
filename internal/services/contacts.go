package services

import (
	"context"
	"github.com/meloie/gonebook/internal/ent"
	"github.com/meloie/gonebook/internal/ent/contact"
	"github.com/meloie/gonebook/internal/ent/user"
	"github.com/meloie/gonebook/pkg/models"

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

func (svc *Service) GetContact(ctx context.Context, contactID int) (*ent.Contact, error) {
	cnt, err := svc.Database.Contact.Query().Where(contact.IDEQ(contactID)).WithOwner().Only(ctx)
	if err != nil {
		return nil, err
	}
	return cnt, nil
}

func (svc *Service) UpdateContact(
	ctx context.Context, contct *ent.Contact, data models.ContactPayload) (*ent.Contact, error) {
	cnt, err := contct.Update().SetAddress(data.Address).
		SetName(data.Name).SetPhone(data.Phone).Save(ctx)
	return cnt, err
}

func (svc *Service) DeleteContact(ctx context.Context, cnc *ent.Contact) error {
	return svc.Database.Contact.DeleteOne(cnc).Exec(ctx)
}
