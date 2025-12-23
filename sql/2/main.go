select sender_code
  , receiver_code
  , content
  , documentschemaname
  , document_type
  , documents.docuuid as docuuid
  , permanentobjectuuid
  , create_date
  , version
  , difference
  , previous_container
  , sentbybus_at
  , invalid_signatures
  , documents.payload as payload
  , events.event as event
from documents
  left join (
    select distinct on (docuuid) docuuid, event
    from doc_event -- 2млн
    order by docuuid, createdat desc
  ) as events on documents.docuuid = events.docuuid
where true
  and lower(documentschemaname) like 'aosr%'
  and permanentobjectuuid = 'd4d43f21-ba73-4e51-8dc2-edba2a6c9667'
  and events.event = ANY('{"ACCEPTEDBYRECEIVER"}')
order by create_date desc
limit 100
offset 0;

-- написать запрос что бы ответить на вопрос РП почему пользователь с логином user1@mail.ru не видит заявки со статусом на подтверждении.
-- доступные Роли на проекте:
-- 1. USER - подача заявок
-- 2. APPROVER - подтверждение заявок
-- 3. EXECUTOR - исполнение заявок

create table "user"
(
    id         bigint primary key generated always as identity,
    login      varchar not null,
    password   varchar not null,
    active     bool      default true,
    updatedat timestamp default currenttimestamp(0),
    org_id     bigint  not null,
    attributes jsonb     default '{}'::jsonb,
    createdat timestamp default currenttimestamp(0)
);

create table organization
(
    id            bigint primary key generated always as identity,
    parentorgid bigint,
    role_codes    jsonb     default '{}'::jsonb,
    name          varchar not null,
    createdat    timestamp default currenttimestamp(0)
);
