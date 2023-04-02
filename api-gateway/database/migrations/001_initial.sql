-- +goose Up
CREATE TABLE IF NOT EXISTS details
(
    id         serial primary key NOT NULL,
    long       real               NOT NULL,
    width      real               NOT NULL,
    height     real               NOT NULL,
    color      text               NOT NULL,
    event_date timestamp          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool               NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS detail_stages
(
    id   serial primary key NOT NULL,
    name text               NOT NULL
);

CREATE TABLE IF NOT EXISTS detail_stage_versions
(
    id         serial primary key NOT NULL,
    detail_id  int                NOT NULL,
    stage_id   int                NOT NULL,
    comment    text,
    event_date timestamp          NOT NULL,

    CONSTRAINT fk_detail_stage_versions_details
    FOREIGN KEY (detail_id)
    REFERENCES details (id)
    ON DELETE SET NULL,

    CONSTRAINT fk_detail_stage_versions_detail_stages
    FOREIGN KEY (stage_id)
    REFERENCES detail_stages (id)
    ON DELETE SET NULL
    );

create index detail_stage_versions_detail_id on detail_stage_versions (detail_id);
create index detail_stage_versions_stage_id on detail_stage_versions (stage_id);

insert into detail_stages (name)
values ('Storage');
insert into detail_stages (name)
values ('Verification');
insert into detail_stages (name)
values ('Processing');
insert into detail_stages (name)
values ('Done');

-- +goose Down
DROP TABLE IF EXISTS detail_stage_versions;
DROP TABLE IF EXISTS detail_stages;
DROP TABLE IF EXISTS details;
