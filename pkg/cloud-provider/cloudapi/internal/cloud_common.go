// Copyright 2022 Antrea Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/multierr"
	"k8s.io/apimachinery/pkg/types"

	cloudv1alpha1 "antrea.io/antreacloud/apis/crd/v1alpha1"
	"antrea.io/antreacloud/pkg/logging"
)

// CloudCommonHelperInterface interface needs to be implemented by each cloud-plugin. It provides a way to inject
// cloud dependent functionality into plugin-common-framework. Cloud dependent functionality can include cloud
// service operations, credentials management etc.
type CloudCommonHelperInterface interface {
	GetCloudServicesCreateFunc() CloudServiceConfigCreatorFunc
	GetCloudCredentialsValidatorFunc() CloudCredentialValidatorFunc
	GetCloudCredentialsComparatorFunc() CloudCredentialComparatorFunc
}

// CloudCommonInterface implements functionality common across all supported cloud-plugins. Each cloud plugin uses
// this interface by composition.
type CloudCommonInterface interface {
	GetCloudAccountByName(namespacedName *types.NamespacedName) (CloudAccountInterface, bool)
	GetCloudAccounts() map[types.NamespacedName]CloudAccountInterface

	GetCloudAccountComputeResourceCRDs(namespacedName *types.NamespacedName) ([]*cloudv1alpha1.VirtualMachine,
		[]*cloudv1alpha1.NetworkInterface, error)
	GetAllCloudAccountsComputeResourceCRDs() ([]*cloudv1alpha1.VirtualMachine, []*cloudv1alpha1.NetworkInterface, error)

	AddCloudAccount(account *cloudv1alpha1.CloudProviderAccount, credentials interface{}) error
	RemoveCloudAccount(namespacedName *types.NamespacedName)

	AddSelector(namespacedName *types.NamespacedName, selector *cloudv1alpha1.CloudEntitySelector) error
	RemoveSelector(accNamespacedName *types.NamespacedName, selectorName string)

	GetStatus(accNamespacedName *types.NamespacedName) (*cloudv1alpha1.CloudProviderAccountStatus, error)
}

type cloudCommon struct {
	mutex               sync.Mutex
	commonHelper        CloudCommonHelperInterface
	logger              func() logging.Logger
	accountConfigs      map[types.NamespacedName]CloudAccountInterface
	cloudSpecificHelper interface{}
	Status              string
}

func NewCloudCommon(logger func() logging.Logger, commonHelper CloudCommonHelperInterface,
	cloudSpecificHelper interface{}) CloudCommonInterface {
	return &cloudCommon{
		logger:              logger,
		commonHelper:        commonHelper,
		cloudSpecificHelper: cloudSpecificHelper,
		accountConfigs:      make(map[types.NamespacedName]CloudAccountInterface),
	}
}

func (c *cloudCommon) AddCloudAccount(account *cloudv1alpha1.CloudProviderAccount, credentials interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	namespacedName := &types.NamespacedName{
		Namespace: account.GetNamespace(),
		Name:      account.GetName(),
	}

	existingConfig, found := c.accountConfigs[*namespacedName]
	if found {
		return c.updateCloudAccountConfig(credentials, existingConfig)
	}

	config, err := c.newCloudAccountConfig(namespacedName, credentials,
		time.Duration(*account.Spec.PollIntervalInSeconds)*time.Second, c.logger)
	if err != nil {
		return err
	}

	c.accountConfigs[*config.GetNamespacedName()] = config

	return nil
}

func (c *cloudCommon) deleteCloudAccount(namespacedName *types.NamespacedName) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.accountConfigs, *namespacedName)
}

func (c *cloudCommon) RemoveCloudAccount(namespacedName *types.NamespacedName) {
	_, found := c.GetCloudAccountByName(namespacedName)
	if !found {
		c.logger().V(0).Info("unable to find cloud account", "account", *namespacedName)
		return
	}
	c.deleteCloudAccount(namespacedName)
}

func (c *cloudCommon) GetCloudAccountByName(namespacedName *types.NamespacedName) (CloudAccountInterface, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	accCfg, found := c.accountConfigs[*namespacedName]
	return accCfg, found
}

func (c *cloudCommon) GetCloudAccounts() map[types.NamespacedName]CloudAccountInterface {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.accountConfigs
}

func (c *cloudCommon) GetCloudAccountComputeResourceCRDs(accountNamespacedName *types.NamespacedName) ([]*cloudv1alpha1.VirtualMachine,
	[]*cloudv1alpha1.NetworkInterface, error) {
	accCfg, found := c.GetCloudAccountByName(accountNamespacedName)
	if !found {
		return nil, nil, fmt.Errorf("unable to find cloud account:%v", *accountNamespacedName)
	}
	namespace := accCfg.GetNamespacedName().Namespace

	var computeCRDs []*cloudv1alpha1.VirtualMachine
	var computeInterfaceCRDs []*cloudv1alpha1.NetworkInterface
	serviceConfigs := accCfg.GetServiceConfigs()
	for _, serviceConfig := range serviceConfigs {
		if serviceConfig.getType() == CloudServiceTypeCompute {
			resourceCRDs := serviceConfig.getResourceCRDs(namespace)
			computeCRDs = append(computeCRDs, resourceCRDs.virtualMachines...)
			computeInterfaceCRDs = append(computeInterfaceCRDs, resourceCRDs.virtualMachineNetworkInterfaces...)
		}
	}

	c.logger().Info("account CRDs", "account", accountNamespacedName, "service-type", CloudServiceTypeCompute,
		"compute", len(computeCRDs), "network-interfaces", len(computeInterfaceCRDs))

	return computeCRDs, computeInterfaceCRDs, nil
}

func (c *cloudCommon) GetAllCloudAccountsComputeResourceCRDs() ([]*cloudv1alpha1.VirtualMachine,
	[]*cloudv1alpha1.NetworkInterface, error) {
	var err error
	var computeCRDs []*cloudv1alpha1.VirtualMachine
	var computeInterfaceCRDs []*cloudv1alpha1.NetworkInterface
	for namespacedName := range c.GetCloudAccounts() {
		accountComputeCRDs, accountComputeIntfCRDs, e := c.GetCloudAccountComputeResourceCRDs(&namespacedName)
		if e == nil {
			computeCRDs = append(computeCRDs, accountComputeCRDs...)
			computeInterfaceCRDs = append(computeInterfaceCRDs, accountComputeIntfCRDs...)
		} else {
			err = multierr.Append(err, e)
		}
	}

	return computeCRDs, computeInterfaceCRDs, err
}

func (c *cloudCommon) AddSelector(accountNamespacedName *types.NamespacedName, selector *cloudv1alpha1.CloudEntitySelector) error {
	accCfg, found := c.GetCloudAccountByName(accountNamespacedName)
	if !found {
		return fmt.Errorf("account not found %v", *accountNamespacedName)
	}

	for _, serviceCfg := range accCfg.GetServiceConfigs() {
		serviceCfg.setResourceFilters(selector)
	}

	err := accCfg.startPeriodicInventorySync()
	if err != nil {
		return fmt.Errorf("inventory sync failed [ account : %v, err: %v ]", *accountNamespacedName, err)
	}

	return nil
}

func (c *cloudCommon) RemoveSelector(accNamespacedName *types.NamespacedName, selectorName string) {
	accCfg, found := c.GetCloudAccountByName(accNamespacedName)
	if !found {
		c.logger().Info("account not found", "account", *accNamespacedName)
		return
	}

	accCfg.stopPeriodicInventorySync()
}

func (c *cloudCommon) GetStatus(accountNamespacedName *types.NamespacedName) (*cloudv1alpha1.CloudProviderAccountStatus, error) {
	accCfg, found := c.GetCloudAccountByName(accountNamespacedName)
	if !found {
		return nil, fmt.Errorf("account not found %v", *accountNamespacedName)
	}

	return accCfg.GetStatus(), nil
}
