CREATE KEYSPACE demo_db WITH replication = {'class': 'org.apache.cassandra.locator.SimpleStrategy', 'replication_factor': '1'}
                         AND durable_writes = true AND tablets = {'enabled': false};

CREATE TABLE demo_db.events_by_user
(
    user_id    uuid,
    created_at timestamp,
    id         uuid,
    payload    text,
    type       text,
    PRIMARY KEY (user_id, created_at, id)
) WITH CLUSTERING ORDER BY (created_at DESC, id ASC)
   AND bloom_filter_fp_chance = 0.01
   AND caching = {'keys': 'ALL', 'rows_per_partition': 'ALL'}
   AND comment = ''
   AND compaction = {'class': 'IncrementalCompactionStrategy'}
   AND compression = {'sstable_compression': 'LZ4WithDictsCompressor'}
   AND crc_check_chance = 1
   AND default_time_to_live = 0
   AND gc_grace_seconds = 864000
   AND max_index_interval = 2048
   AND memtable_flush_period_in_ms = 0
   AND min_index_interval = 128
   AND speculative_retry = '99.0PERCENTILE'
   AND tombstone_gc = {'mode': 'timeout', 'propagation_delay_in_seconds': '3600'};


CREATE TABLE demo_db.posts_by_user
(
    user_id    uuid,
    created_at timestamp,
    id         uuid,
    body       text,
    published  boolean,
    title      text,
    PRIMARY KEY (user_id, created_at, id)
) WITH CLUSTERING ORDER BY (created_at DESC, id ASC)
   AND bloom_filter_fp_chance = 0.01
   AND caching = {'keys': 'ALL', 'rows_per_partition': 'ALL'}
   AND comment = ''
   AND compaction = {'class': 'IncrementalCompactionStrategy'}
   AND compression = {'sstable_compression': 'LZ4WithDictsCompressor'}
   AND crc_check_chance = 1
   AND default_time_to_live = 0
   AND gc_grace_seconds = 864000
   AND max_index_interval = 2048
   AND memtable_flush_period_in_ms = 0
   AND min_index_interval = 128
   AND speculative_retry = '99.0PERCENTILE'
   AND tombstone_gc = {'mode': 'timeout', 'propagation_delay_in_seconds': '3600'};


CREATE TABLE demo_db.users_by_email
(
    email      text,
    active     boolean,
    created_at timestamp,
    id         uuid,
    name       text,
    PRIMARY KEY (email)
) WITH bloom_filter_fp_chance = 0.01
   AND caching = {'keys': 'ALL', 'rows_per_partition': 'ALL'}
   AND comment = ''
   AND compaction = {'class': 'IncrementalCompactionStrategy'}
   AND compression = {'sstable_compression': 'LZ4WithDictsCompressor'}
   AND crc_check_chance = 1
   AND default_time_to_live = 0
   AND gc_grace_seconds = 864000
   AND max_index_interval = 2048
   AND memtable_flush_period_in_ms = 0
   AND min_index_interval = 128
   AND speculative_retry = '99.0PERCENTILE'
   AND tombstone_gc = {'mode': 'timeout', 'propagation_delay_in_seconds': '3600'};



USE demo_db;

-- Users
INSERT INTO users_by_email (email, active, created_at, id, name)
VALUES ('anna@example.com',
        true,
        '2026-06-01 09:00:00',
        11111111-1111-1111-1111-111111111111,
        'Anna Novak');

INSERT INTO users_by_email (email, active, created_at, id, name)
VALUES ('peter@example.com',
        true,
        '2026-06-02 10:30:00',
        22222222-2222-2222-2222-222222222222,
        'Peter Smith');

INSERT INTO users_by_email (email, active, created_at, id, name)
VALUES ('lucia@example.com',
        false,
        '2026-06-03 14:15:00',
        33333333-3333-3333-3333-333333333333,
        'Lucy Brown');


-- Posts by user: Anna
INSERT INTO posts_by_user (user_id, created_at, id, body, published, title)
VALUES (11111111-1111-1111-1111-111111111111,
        '2026-06-10 08:30:00',
        aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa1,
        'This is Anna''s first demo post.',
        true,
        'First Post');

INSERT INTO posts_by_user (user_id, created_at, id, body, published, title)
VALUES (11111111-1111-1111-1111-111111111111,
        '2026-06-09 17:45:00',
        aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa2,
        'A short draft post that has not been published yet.',
        false,
        'Draft Post');


-- Posts by user: Peter
INSERT INTO posts_by_user (user_id, created_at, id, body, published, title)
VALUES (22222222-2222-2222-2222-222222222222,
        '2026-06-08 12:00:00',
        bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb1,
        'Peter writes about testing a Cassandra data model.',
        true,
        'Testing the Data Model');

INSERT INTO posts_by_user (user_id, created_at, id, body, published, title)
VALUES (22222222-2222-2222-2222-222222222222,
        '2026-06-07 18:20:00',
        bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb2,
        'Another sample text for the demo application.',
        true,
        'Demo Application');


-- Posts by user: Lucy
INSERT INTO posts_by_user (user_id, created_at, id, body, published, title)
VALUES (33333333-3333-3333-3333-333333333333,
        '2026-06-06 11:10:00',
        cccccccc-cccc-cccc-cccc-ccccccccccc1,
        'Lucy''s unpublished private notebook.',
        false,
        'Private Notebook');


-- Events by user: Anna
INSERT INTO events_by_user (user_id, created_at, id, payload, type)
VALUES (11111111-1111-1111-1111-111111111111,
        '2026-06-10 08:35:00',
        dddddddd-dddd-dddd-dddd-ddddddddddd1,
        '{"email":"anna@example.com","action":"login"}',
        'USER_LOGIN');

INSERT INTO events_by_user (user_id, created_at, id, payload, type)
VALUES (11111111-1111-1111-1111-111111111111,
        '2026-06-10 08:40:00',
        dddddddd-dddd-dddd-dddd-ddddddddddd2,
        '{"post_id":"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa1","title":"First Post"}',
        'POST_CREATED');

INSERT INTO events_by_user (user_id, created_at, id, payload, type)
VALUES (11111111-1111-1111-1111-111111111111,
        '2026-06-10 09:00:00',
        dddddddd-dddd-dddd-dddd-ddddddddddd3,
        '{"post_id":"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa1","published":true}',
        'POST_PUBLISHED');


-- Events by user: Peter
INSERT INTO events_by_user (user_id, created_at, id, payload, type)
VALUES (22222222-2222-2222-2222-222222222222,
        '2026-06-08 12:05:00',
        eeeeeeee-eeee-eeee-eeee-eeeeeeeeeee1,
        '{"email":"peter@example.com","action":"login"}',
        'USER_LOGIN');

INSERT INTO events_by_user (user_id, created_at, id, payload, type)
VALUES (22222222-2222-2222-2222-222222222222,
        '2026-06-08 12:10:00',
        eeeeeeee-eeee-eeee-eeee-eeeeeeeeeee2,
        '{"post_id":"bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb1","title":"Testing the Data Model"}',
        'POST_CREATED');

INSERT INTO events_by_user (user_id, created_at, id, payload, type)
VALUES (22222222-2222-2222-2222-222222222222,
        '2026-06-08 12:20:00',
        eeeeeeee-eeee-eeee-eeee-eeeeeeeeeee3,
        '{"post_id":"bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb1","published":true}',
        'POST_PUBLISHED');


-- Events by user: Lucy
INSERT INTO events_by_user (user_id, created_at, id, payload, type)
VALUES (33333333-3333-3333-3333-333333333333,
        '2026-06-06 11:00:00',
        ffffffff-ffff-ffff-ffff-fffffffffff1,
        '{"email":"lucia@example.com","action":"signup"}',
        'USER_CREATED');

INSERT INTO events_by_user (user_id, created_at, id, payload, type)
VALUES (33333333-3333-3333-3333-333333333333,
        '2026-06-06 11:15:00',
        ffffffff-ffff-ffff-ffff-fffffffffff2,
        '{"post_id":"cccccccc-cccc-cccc-cccc-ccccccccccc1","title":"Private Notebook"}',
        'POST_CREATED');

INSERT INTO events_by_user (user_id, created_at, id, payload, type)
VALUES (33333333-3333-3333-3333-333333333333,
        '2026-06-09 16:45:00',
        ffffffff-ffff-ffff-ffff-fffffffffff3,
        '{"active":false,"reason":"demo inactive user"}',
        'USER_DEACTIVATED');


-- User: Jane Doe
INSERT INTO users_by_email (email, active, created_at, id, name)
VALUES (
           'jane@example.com',
           true,
           '2026-06-01 09:00:00',
           11111111-1111-1111-1111-111111111111,
           'Jane Doe'
       );


-- Posts by user: Jane
INSERT INTO posts_by_user (user_id, created_at, id, body, published, title)
VALUES (
           11111111-1111-1111-1111-111111111111,
           '2026-06-10 08:30:00',
           aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa1,
           'This is Jane''s first demo post.',
           true,
           'First Post'
       );

INSERT INTO posts_by_user (user_id, created_at, id, body, published, title)
VALUES (
           11111111-1111-1111-1111-111111111111,
           '2026-06-09 17:45:00',
           aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa2,
           'A short draft post that has not been published yet.',
           false,
           'Draft Post'
       );


-- Events by user: Jane
INSERT INTO events_by_user (user_id, created_at, id, payload, type)
VALUES (
           11111111-1111-1111-1111-111111111111,
           '2026-06-10 08:35:00',
           dddddddd-dddd-dddd-dddd-ddddddddddd1,
           '{"email":"jane@example.com","action":"login"}',
           'USER_LOGIN'
       );

INSERT INTO events_by_user (user_id, created_at, id, payload, type)
VALUES (
           11111111-1111-1111-1111-111111111111,
           '2026-06-10 08:40:00',
           dddddddd-dddd-dddd-dddd-ddddddddddd2,
           '{"post_id":"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa1","title":"First Post"}',
           'POST_CREATED'
       );

INSERT INTO events_by_user (user_id, created_at, id, payload, type)
VALUES (
           11111111-1111-1111-1111-111111111111,
           '2026-06-10 09:00:00',
           dddddddd-dddd-dddd-dddd-ddddddddddd3,
           '{"post_id":"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa1","published":true}',
           'POST_PUBLISHED'
       );

-- User: John Doe
INSERT INTO users_by_email (email, active, created_at, id, name)
VALUES (
           'john@example.com',
           true,
           '2026-06-02 10:30:00',
           22222222-2222-2222-2222-222222222222,
           'John Doe'
       );


-- Posts by user: John
INSERT INTO posts_by_user (user_id, created_at, id, body, published, title)
VALUES (
           22222222-2222-2222-2222-222222222222,
           '2026-06-08 12:00:00',
           bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb1,
           'John writes about testing a Cassandra data model.',
           true,
           'Testing the Data Model'
       );

INSERT INTO posts_by_user (user_id, created_at, id, body, published, title)
VALUES (
           22222222-2222-2222-2222-222222222222,
           '2026-06-07 18:20:00',
           bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb2,
           'Another sample text for the demo application.',
           true,
           'Demo Application'
       );


-- Events by user: John
INSERT INTO events_by_user (user_id, created_at, id, payload, type)
VALUES (
           22222222-2222-2222-2222-222222222222,
           '2026-06-08 12:05:00',
           eeeeeeee-eeee-eeee-eeee-eeeeeeeeeee1,
           '{"email":"john@example.com","action":"login"}',
           'USER_LOGIN'
       );

INSERT INTO events_by_user (user_id, created_at, id, payload, type)
VALUES (
           22222222-2222-2222-2222-222222222222,
           '2026-06-08 12:10:00',
           eeeeeeee-eeee-eeee-eeee-eeeeeeeeeee2,
           '{"post_id":"bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb1","title":"Testing the Data Model"}',
           'POST_CREATED'
       );

INSERT INTO events_by_user (user_id, created_at, id, payload, type)
VALUES (
           22222222-2222-2222-2222-222222222222,
           '2026-06-08 12:20:00',
           eeeeeeee-eeee-eeee-eeee-eeeeeeeeeee3,
           '{"post_id":"bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb1","published":true}',
           'POST_PUBLISHED'
       );