//
// DISCLAIMER
//
// Copyright 2024 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

package gateway

import (
	"fmt"
	"sort"

	bootstrapAPI "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	clusterAPI "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	coreAPI "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	listenerAPI "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	routeAPI "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	routerAPI "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
	httpConnectionManagerAPI "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/anypb"
	"sigs.k8s.io/yaml"

	shared "github.com/arangodb/kube-arangodb/pkg/apis/shared"
	"github.com/arangodb/kube-arangodb/pkg/util"
	"github.com/arangodb/kube-arangodb/pkg/util/errors"
)

type Config struct {
	DefaultDestination ConfigDestination `json:"defaultDestination,omitempty"`

	Destinations ConfigDestinations `json:"destinations,omitempty"`

	DefaultTLS *ConfigTLS `json:"defaultTLS,omitempty"`
}

func (c Config) Validate() error {
	return errors.Errors(
		shared.PrefixResourceErrors("defaultDestination", c.DefaultDestination.Validate()),
		shared.PrefixResourceErrors("destinations", c.Destinations.Validate()),
	)
}

func (c Config) RenderYAML() ([]byte, string, *bootstrapAPI.Bootstrap, error) {
	cfg, err := c.Render()
	if err != nil {
		return nil, "", nil, err
	}

	data, err := protojson.MarshalOptions{
		UseProtoNames: true,
	}.Marshal(cfg)
	if err != nil {
		return nil, "", nil, err
	}

	data, err = yaml.JSONToYAML(data)
	return data, util.SHA256(data), cfg, err
}

func (c Config) Render() (*bootstrapAPI.Bootstrap, error) {
	if err := c.Validate(); err != nil {
		return nil, errors.Wrapf(err, "Validation failed")
	}

	clusters, err := c.RenderClusters()
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to render clusters")
	}

	listener, err := c.RenderListener()
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to render listener")
	}

	return &bootstrapAPI.Bootstrap{
		Admin: &bootstrapAPI.Admin{
			Address: &coreAPI.Address{
				Address: &coreAPI.Address_SocketAddress{
					SocketAddress: &coreAPI.SocketAddress{
						Address:       "127.0.0.1",
						PortSpecifier: &coreAPI.SocketAddress_PortValue{PortValue: 9901},
					},
				},
			},
		},
		StaticResources: &bootstrapAPI.Bootstrap_StaticResources{
			Listeners: []*listenerAPI.Listener{
				listener,
			},
			Clusters: clusters,
		},
	}, nil
}

func (c Config) RenderClusters() ([]*clusterAPI.Cluster, error) {
	def, err := c.DefaultDestination.RenderCluster("default")
	if err != nil {
		return nil, err
	}
	clusters := []*clusterAPI.Cluster{
		def,
	}

	for k, v := range c.Destinations {
		name := fmt.Sprintf("cluster_%s", util.SHA256FromString(k))
		c, err := v.RenderCluster(name)
		if err != nil {
			return nil, err
		}

		clusters = append(clusters, c)
	}

	sort.Slice(clusters, func(i, j int) bool {
		return clusters[i].Name < clusters[j].Name
	})

	return clusters, nil
}

func (c Config) RenderRoutes() ([]*routeAPI.Route, error) {
	def, err := c.DefaultDestination.RenderRoute("default", "/")
	if err != nil {
		return nil, err
	}
	routes := []*routeAPI.Route{
		def,
	}

	for k, v := range c.Destinations {
		name := fmt.Sprintf("cluster_%s", util.SHA256FromString(k))
		c, err := v.RenderRoute(name, k)
		if err != nil {
			return nil, err
		}

		routes = append(routes, c)
	}

	sort.Slice(routes, func(i, j int) bool {
		return routes[i].GetMatch().GetPrefix() > routes[j].GetMatch().GetPrefix()
	})

	return routes, nil
}

func (c Config) RenderFilters() ([]*listenerAPI.Filter, error) {
	httpFilterConfigType, err := anypb.New(&routerAPI.Router{})
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to render route config")
	}

	routes, err := c.RenderRoutes()
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to render routes")
	}

	filterConfigType, err := anypb.New(&httpConnectionManagerAPI.HttpConnectionManager{
		StatPrefix: "ingress_http",
		CodecType:  httpConnectionManagerAPI.HttpConnectionManager_AUTO,
		RouteSpecifier: &httpConnectionManagerAPI.HttpConnectionManager_RouteConfig{
			RouteConfig: &routeAPI.RouteConfiguration{
				Name: "default",
				VirtualHosts: []*routeAPI.VirtualHost{
					{
						Name:    "default",
						Domains: []string{"*"},
						Routes:  routes,
					},
				},
			},
		},
		HttpFilters: []*httpConnectionManagerAPI.HttpFilter{
			{
				Name: "envoy.filters.http.routerAPI",
				ConfigType: &httpConnectionManagerAPI.HttpFilter_TypedConfig{
					TypedConfig: httpFilterConfigType,
				},
			},
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to render http connection manager")
	}

	return []*listenerAPI.Filter{
		{
			Name: "envoy.filters.network.httpConnectionManagerAPI",
			ConfigType: &listenerAPI.Filter_TypedConfig{
				TypedConfig: filterConfigType,
			},
		},
	}, nil
}

func (c Config) RenderDefaultFilterChain() (*listenerAPI.FilterChain, error) {
	filters, err := c.RenderFilters()
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to render filters")
	}

	ret := &listenerAPI.FilterChain{
		Filters: filters,
	}

	if tls, err := c.DefaultTLS.RenderListenerTransportSocket(); err != nil {
		return nil, err
	} else {
		ret.TransportSocket = tls
	}

	return ret, nil
}

func (c Config) RenderSecondaryFilterChains() ([]*listenerAPI.FilterChain, error) {
	return nil, nil
}

func (c Config) RenderListener() (*listenerAPI.Listener, error) {
	filterChains, err := c.RenderSecondaryFilterChains()
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to render secondary filter chains")
	}

	defaultFilterChain, err := c.RenderDefaultFilterChain()
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to render default filter")
	}

	return &listenerAPI.Listener{
		Name: "default",
		Address: &coreAPI.Address{
			Address: &coreAPI.Address_SocketAddress{
				SocketAddress: &coreAPI.SocketAddress{
					Address:       "0.0.0.0",
					PortSpecifier: &coreAPI.SocketAddress_PortValue{PortValue: shared.ArangoPort},
				},
			},
		},
		FilterChains: filterChains,

		DefaultFilterChain: defaultFilterChain,
	}, nil
}
