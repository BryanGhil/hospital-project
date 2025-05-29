create table roles (
    id bigserial primary key,
    role varchar not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp
);

create table users (
    id bigserial primary key,
    email varchar not null unique,
    password varchar not null,
    role_id bigint not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp,
    foreign key (role_id) references roles(id)
);

insert into roles (role, created_at, updated_at) 
values 
    ('admin', NOW(), NOW()),
    ('user', NOW(), NOW());


CREATE TABLE patients (
    id BIGSERIAL PRIMARY KEY,
    full_name varchar NOT NULL,
    dob DATE NOT NULL,
    gender varchar,
    address varchar,
    phone varchar,
    created_by INTEGER REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp not null,
    deleted_at timestamp
);

CREATE TABLE medicines (
    id bigserial PRIMARY KEY,
    name varchar NOT NULL,
    stock INTEGER DEFAULT 0,
    price NUMERIC(12, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp not null,
    deleted_at timestamp
);

CREATE TABLE medical_histories (
    id bigserial PRIMARY KEY,
    patient_id bigint NOT NULL REFERENCES patients(id),
    doctor_id bigint NOT NULL REFERENCES users(id),
    diagnosis TEXT,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp not null,
    deleted_at timestamp
);

CREATE TABLE prescriptions (
    id bigserial PRIMARY KEY,
    history_id bigint NOT NULL REFERENCES medical_histories(id),
    medicine_id bigint NOT NULL REFERENCES medicines(id),
    medicine_name TEXT NOT NULL,
    dosage TEXT,
    quantity INTEGER NOT NULL,
    unit_price NUMERIC(12, 2) NOT NULL,
    total_price NUMERIC(12, 2) NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp not null,
    deleted_at timestamp
);

CREATE TABLE transactions (
    id bigserial PRIMARY KEY,
    patient_id bigint REFERENCES patients(id),
    doctor_id bigint REFERENCES users(id),
    total NUMERIC(12, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp not null,
    deleted_at timestamp
);



