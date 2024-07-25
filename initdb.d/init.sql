-- Попытка создания пользователя с игнорированием ошибки, если пользователь уже существует
DO $$
    BEGIN
        -- Если роль уже существует, она не будет повторно создана
        BEGIN
            EXECUTE 'CREATE ROLE postgres WITH LOGIN SUPERUSER PASSWORD ''12345''';
        EXCEPTION
            WHEN duplicate_object THEN
                -- Игнорируем ошибку, если роль уже существует
                RAISE NOTICE 'Role postgres already exists.';
        END;
    END $$;

-- Создание таблицы message, если она не существует
CREATE TABLE IF NOT EXISTS message
(
    id   SERIAL PRIMARY KEY,
    data TEXT DEFAULT ''::text NOT NULL
);

-- Установка владельца таблицы
ALTER TABLE message
    OWNER TO postgres;
