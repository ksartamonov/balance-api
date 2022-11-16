CREATE TABLE IF NOT EXISTS users
(
    id           SERIAL,
    name         VARCHAR NOT NULL,
    cardNumber   VARCHAR NOT NULL,
    balance      INT     DEFAULT 0,
    reserved     INT     DEFAULT 0
);

CREATE TABLE IF NOT EXISTS operations
(
    OperationId         SERIAL,
    SenderId            INT         DEFAULT 0,
    ReceiverId          INT         DEFAULT 0,
    Money               INT         DEFAULT 0,
    TransactionTime     TIMESTAMP
);

INSERT INTO users (name, cardNumber)
VALUES ('Ivan', '1111 1111 1111 1111'),
       ('Artem', '2222 2222 2222 2222'),
       ('Anna', '3333 3333 3333 3333'),
       ('Kate', '4444 4444 4444 4444'),
       ('Anastasia', '5555 5555 5555 5555'),
       ('Vadim', '6666 6666 6666 6666');
