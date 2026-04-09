CREATE SCHEMA IF NOT EXISTS rategate;

CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- gen_random_uuid

CREATE TABLE rategate.tenants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_tenants_status 
ON rategate.tenants(status);

CREATE TABLE rategate.api_keys (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    key_hash TEXT NOT NULL,
    name TEXT,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    expires_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_api_keys_tenant
        FOREIGN KEY (tenant_id)
        REFERENCES rategate.tenants(id)
        ON DELETE CASCADE
);

CREATE UNIQUE INDEX idx_api_keys_hash 
ON rategate.api_keys(key_hash);

CREATE INDEX idx_api_keys_tenant 
ON rategate.api_keys(tenant_id);

CREATE TABLE rategate.services (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    base_path TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE rategate.endpoints (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_id UUID NOT NULL,
    path TEXT NOT NULL,
    method VARCHAR(10) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_endpoints_service
        FOREIGN KEY (service_id)
        REFERENCES rategate.services(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_endpoints_service 
ON rategate.endpoints(service_id);

CREATE TABLE rategate.rate_limit_policies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    tenant_id UUID NOT NULL,
    api_key_id UUID,
    endpoint_id UUID,

    limit_count INT NOT NULL CHECK (limit_count > 0),
    window_seconds INT NOT NULL CHECK (window_seconds > 0),

    strategy VARCHAR(20) NOT NULL, -- fixed, sliding, token_bucket
    burst_limit INT,

    priority INT NOT NULL DEFAULT 0,

    version INT NOT NULL DEFAULT 1,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,

    CONSTRAINT fk_policy_tenant
        FOREIGN KEY (tenant_id)
        REFERENCES rategate.tenants(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_policy_api_key
        FOREIGN KEY (api_key_id)
        REFERENCES rategate.api_keys(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_policy_endpoint
        FOREIGN KEY (endpoint_id)
        REFERENCES rategate.endpoints(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_policy_lookup 
ON rategate.rate_limit_policies(
    tenant_id,
    api_key_id,
    endpoint_id,
    priority DESC
)
WHERE deleted_at IS NULL AND is_active = TRUE;

CREATE TABLE rategate.policy_audit_logs (
    id BIGSERIAL PRIMARY KEY,
    policy_id UUID NOT NULL,
    change_type VARCHAR(20) NOT NULL, -- create, update, delete
    old_value JSONB,
    new_value JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_audit_policy
        FOREIGN KEY (policy_id)
        REFERENCES rategate.rate_limit_policies(id)
        ON DELETE CASCADE
);

CREATE TABLE rategate.request_logs (
    id BIGSERIAL,
    tenant_id UUID,
    api_key_id UUID,
    endpoint_id UUID,
    status_code INT,
    latency_ms INT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id, created_at)
) PARTITION BY RANGE (created_at);

CREATE OR REPLACE FUNCTION rategate.set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_tenants_updated_at
BEFORE UPDATE ON rategate.tenants
FOR EACH ROW
EXECUTE FUNCTION rategate.set_updated_at();

CREATE TRIGGER trg_policies_updated_at
BEFORE UPDATE ON rategate.rate_limit_policies
FOR EACH ROW
EXECUTE FUNCTION rategate.set_updated_at();