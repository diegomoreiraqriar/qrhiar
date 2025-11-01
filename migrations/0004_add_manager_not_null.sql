
ALTER TABLE third_party_users
ADD COLUMN IF NOT EXISTS manager_id UUID NOT NULL;

ALTER TABLE third_party_users
ADD CONSTRAINT fk_manager
FOREIGN KEY (manager_id)
REFERENCES third_party_users (id)
ON DELETE RESTRICT;
