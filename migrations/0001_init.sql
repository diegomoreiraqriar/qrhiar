create table if not exists companies (
  id uuid primary key default gen_random_uuid(),
  tenant_id varchar(64) not null,
  cnpj varchar(20) not null,
  razao_social text not null,
  contrato text,
  inicio_vigencia date,
  fim_vigencia date,
  status varchar(20) not null default 'active',
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);
create index if not exists idx_companies_tenant on companies(tenant_id);

create table if not exists third_party_users (
  id uuid primary key default gen_random_uuid(),
  tenant_id varchar(64) not null,
  company_id uuid not null references companies(id) on delete cascade,
  nome text not null,
  cpf varchar(20),
  email text,
  cargo text,
  funcao text,
  manager_email text,
  status varchar(20) not null default 'active', -- active|inactive|pending
  data_inicio date,
  data_fim date,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);
create index if not exists idx_tpu_tenant on third_party_users(tenant_id);

create table if not exists audit_logs (
  id bigserial primary key,
  tenant_id varchar(64) not null,
  entity varchar(64) not null,
  entity_id uuid,
  action varchar(32) not null,
  actor varchar(128),
  diff jsonb,
  created_at timestamptz not null default now()
);
create index if not exists idx_audit_tenant on audit_logs(tenant_id);
