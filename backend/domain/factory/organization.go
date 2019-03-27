package factory

import (
	"go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/entity"
	"go-ddd/backend/domain/error"
	vo "go-ddd/backend/domain/valueobject"
)

// OrganizationFactory ... 「組織」集約のあらゆる生成方法を担う
type OrganizationFactory struct {
	uid vo.UniqueID
}

// CreateRootOrganization ...
func (f *OrganizationFactory) CreateRootOrganization(organizationName vo.OrganizationName) (aggregate.Organization, error.ApplicationError) {
	attribute, err := entity.NewOrganizationAttribute(f.uid, organizationName)
	if err != nil {
		return nil, err
	}

	return aggregate.NewOrganization(attribute, nil, nil), nil
}

// CreateRootOrganizationWithChildren ...
func (f *OrganizationFactory) CreateRootOrganizationWithChildren(organizationName vo.OrganizationName, children []entity.OrganizationAttribute) (aggregate.Organization, error.ApplicationError) {
	attribute, err := entity.NewOrganizationAttribute(f.uid, organizationName)
	if err != nil {
		return nil, err
	}

	return aggregate.NewOrganization(attribute, nil, children), nil
}

// CreateOrganization ...
func (f *OrganizationFactory) CreateOrganization(organizationName vo.OrganizationName, parent entity.OrganizationAttribute) (aggregate.Organization, error.ApplicationError) {
	attribute, err := entity.NewOrganizationAttribute(f.uid, organizationName)
	if err != nil {
		return nil, err
	}

	return aggregate.NewOrganization(attribute, parent, nil), nil
}

// CreateOrganizationWithChildren ...
func (f *OrganizationFactory) CreateOrganizationWithChildren(organizationName vo.OrganizationName, parent entity.OrganizationAttribute, children []entity.OrganizationAttribute) (aggregate.Organization, error.ApplicationError) {
	attribute, err := entity.NewOrganizationAttribute(f.uid, organizationName)
	if err != nil {
		return nil, err
	}

	return aggregate.NewOrganization(attribute, parent, children), nil
}
