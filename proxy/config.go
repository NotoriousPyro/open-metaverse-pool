package proxy

import (
    "github.com/NotoriousPyro/open-metaverse-pool/api"
    "github.com/NotoriousPyro/open-metaverse-pool/payouts"
    "github.com/NotoriousPyro/open-metaverse-pool/policy"
    "github.com/NotoriousPyro/open-metaverse-pool/storage"
)

type Config struct {
    Account                   string           `json:"account"`
    Password                  string           `json:"password"`
    Name                      string           `json:"name"`
    Proxy                     Proxy            `json:"proxy"`
    Api                       api.ApiConfig    `json:"api"`
    Upstream                  []Upstream       `json:"upstream"`
    UpstreamCheckInterval     string           `json:"upstreamCheckInterval"`

    Threads                   int              `json:"threads"`

    Coin                      string              `json:"coin"`
    Redis                     storage.Config      `json:"redis"`
    
    BlockUnlocker             payouts.UnlockerConfig       `json:"unlocker"`
    Payouts                   payouts.PayoutsConfig        `json:"payouts"`

    NewrelicName              string    `json:"newrelicName"`
    NewrelicKey               string    `json:"newrelicKey"`
    NewrelicVerbose           bool      `json:"newrelicVerbose"`
    NewrelicEnabled           bool      `json:"newrelicEnabled"`
}

type Proxy struct {
    Enabled                 bool        `json:"enabled"`
    Name                    string      `json:"name"`
    Address                 string      `json:"address"`
    Listen                  string      `json:"listen"`
    LimitHeadersSize        int         `json:"limitHeadersSize"`
    LimitBodySize           int64       `json:"limitBodySize"`
    BehindReverseProxy      bool        `json:"behindReverseProxy"`
    BlockRefreshInterval    string      `json:"blockRefreshInterval"`
    StateUpdateInterval     string      `json:"stateUpdateInterval"`
    HashrateExpiration      string      `json:"hashrateExpiration"`

    Policy                  policy.Config   `json:"policy"`

    MaxFails                int64           `json:"maxFails"`
    HealthCheck             bool            `json:"healthCheck"`

    Stratum                 []Stratum       `json:"stratum"`
}

type Stratum struct {
    Name           string      `json:"name"`
    Enabled        bool        `json:"enabled"`
    Listen         string      `json:"listen"`
    Timeout        string      `json:"timeout"`
    MaxConn        int         `json:"maxConn"`
    Difficulty     int64       `json:"difficulty"`
}

type Upstream struct {
    Name           string      `json:"name"`
    Url            string      `json:"url"`
    Timeout        string      `json:"timeout"`
}
