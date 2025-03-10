// Copyright 2019-2025 The Liqo Authors
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

package utils

import (
	networkingv1beta1 "github.com/liqotech/liqo/apis/networking/v1beta1"
	cidrutils "github.com/liqotech/liqo/pkg/utils/cidr"
)

// IsConfigurationStatusSet check if a Configuration is ready by checking if its status is correctly set.
func IsConfigurationStatusSet(confStatus networkingv1beta1.ConfigurationStatus) bool {
	return confStatus.Remote != nil &&
		!cidrutils.IsVoid(cidrutils.GetPrimary(confStatus.Remote.CIDR.Pod)) &&
		!cidrutils.IsVoid(cidrutils.GetPrimary(confStatus.Remote.CIDR.External))
}
