ALTER TABLE public.articles ALTER COLUMN deleted_at DROP NOT NULL;
ALTER TABLE public.articles DROP COLUMN tag_id;