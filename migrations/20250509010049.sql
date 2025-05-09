-- Create enum type "action_type"
CREATE TYPE "public.action_type" AS ENUM ('create', 'read', 'update', 'delete');
-- Create enum type "resource_type"
CREATE TYPE "public.resource_type" AS ENUM ('project', 'task', 'user', 'user2');
-- Create enum type "condition_operator"
CREATE TYPE "public.condition_operator" AS ENUM ('eq', 'neq', 'contains', 'startsWith');
