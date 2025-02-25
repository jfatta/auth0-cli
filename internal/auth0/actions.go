//go:generate mockgen -source=actions.go -destination=actions_mock.go -package=auth0
package auth0

import "gopkg.in/auth0.v5/management"

type ActionAPI interface {
	Create(a *management.Action) error
	Read(id string) (*management.Action, error)
	Update(id string, a *management.Action) error
	Delete(id string, opts ...management.RequestOption) error
	List(opts ...management.RequestOption) (c *management.ActionList, err error)
}

type ActionVersionAPI interface {
	Create(actionID string, v *management.ActionVersion) error
	Read(actionID string, id string) (*management.ActionVersion, error)
	UpsertDraft(actionID string, v *management.ActionVersion) error
	ReadDraft(actionID string) (*management.ActionVersion, error)
	Delete(actionID string, id string, opts ...management.RequestOption) error
	List(actionID string, opts ...management.RequestOption) (c *management.ActionVersionList, err error)
	Test(actionID string, id string, payload management.Object) (management.Object, error)
	Deploy(actionID string, id string) (*management.ActionVersion, error)
}

type ActionBindingAPI interface {
	Create(triggerID management.TriggerID, action *management.Action) (ab *management.ActionBinding, err error)
	List(triggerID management.TriggerID, opts ...management.RequestOption) (c *management.ActionBindingList, err error)
	Update(triggerID management.TriggerID, v []*management.ActionBinding) (list *management.ActionBindingList, err error)
}

type ActionExecutionAPI interface {
	Read(id string) (*management.ActionExecution, error)
}
