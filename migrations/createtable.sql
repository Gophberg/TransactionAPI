-- DROP TABLE IF EXISTS transactions;
-- DROP TYPE IF EXISTS status;
-- DROP TYPE IF EXISTS currency;

-- CREATE TYPE status AS ENUM ('New', 'Success', 'Failure', 'Error', 'Canceled');

CREATE TABLE IF NOT EXISTS transactions (
    id bigserial not null primary key,
    userid bigserial,
    useremail text,
    amount money,
    currency text,
    creationdate timestamp,
    updatedate timestamp,
    status text
);
