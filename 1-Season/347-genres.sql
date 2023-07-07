-- Спасибо Кутьин Захар Сергеевич
-- Студент группы БББО-05-20

-- 347. Музыкальные жанры
-- USE KutinMusic;
WITH track_q(track_id, genre_id, parent_id, track_name, genre_name) as (
    SELECT tg.track_id, tg.genre_id, gen.parent_genre_id as parent_id, t.name as track_name, gen.name as genre_name
    FROM track_genre as tg, genre as gen, track as t
    WHERE tg.genre_id = gen.id and t.id = tg.track_id
    UNION ALL
    SELECT tq.track_id, tq.parent_id, gen.parent_genre_id, t.name as track_name, gen.name as genre_name
    FROM track_q as tq, genre as gen, track as t
    WHERE tq.parent_id = gen.id AND t.id = tq.track_id
)
SELECT DISTINCT track_id, genre_id, track_name, genre_name FROM track_q
ORDER BY track_id, genre_id ASC