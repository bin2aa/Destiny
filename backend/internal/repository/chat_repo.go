package repository

import (
	"context"
	"fmt"

	"destiny-backend/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ChatRepo struct {
	pool *pgxpool.Pool
}

func NewChatRepo(pool *pgxpool.Pool) *ChatRepo {
	return &ChatRepo{pool: pool}
}

func (r *ChatRepo) CreateChat(ctx context.Context, chat *model.AIChat) error {
	query := `INSERT INTO ai_chat (user_id, birth_profile_id, title) VALUES ($1,$2,$3) RETURNING id, created_at, updated_at`
	return r.pool.QueryRow(ctx, query, chat.UserID, chat.BirthProfileID, chat.Title).Scan(&chat.ID, &chat.CreatedAt, &chat.UpdatedAt)
}

func (r *ChatRepo) GetChatByID(ctx context.Context, id string) (*model.AIChat, error) {
	query := `SELECT id, user_id, birth_profile_id, title, created_at, updated_at FROM ai_chat WHERE id = $1`
	c := &model.AIChat{}
	err := r.pool.QueryRow(ctx, query, id).Scan(&c.ID, &c.UserID, &c.BirthProfileID, &c.Title, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("get chat by id: %w", err)
	}
	return c, nil
}

func (r *ChatRepo) ListByUserID(ctx context.Context, userID string) ([]*model.AIChat, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, user_id, birth_profile_id, title, created_at, updated_at FROM ai_chat WHERE user_id = $1 ORDER BY updated_at DESC", userID)
	if err != nil {
		return nil, fmt.Errorf("list chats: %w", err)
	}
	defer rows.Close()

	var chats []*model.AIChat
	for rows.Next() {
		c := &model.AIChat{}
		if err := rows.Scan(&c.ID, &c.UserID, &c.BirthProfileID, &c.Title, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan chat: %w", err)
		}
		chats = append(chats, c)
	}
	return chats, nil
}

func (r *ChatRepo) CreateMessage(ctx context.Context, msg *model.AIMessage) error {
	query := `INSERT INTO ai_messages (chat_id, role, content, token) VALUES ($1,$2,$3,$4) RETURNING id, created_at`
	return r.pool.QueryRow(ctx, query, msg.ChatID, msg.Role, msg.Content, msg.Token).Scan(&msg.ID, &msg.CreatedAt)
}

func (r *ChatRepo) GetMessagesByChatID(ctx context.Context, chatID string) ([]*model.AIMessage, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, chat_id, role, content, token, created_at FROM ai_messages WHERE chat_id = $1 ORDER BY created_at ASC", chatID)
	if err != nil {
		return nil, fmt.Errorf("get messages: %w", err)
	}
	defer rows.Close()

	var msgs []*model.AIMessage
	for rows.Next() {
		m := &model.AIMessage{}
		if err := rows.Scan(&m.ID, &m.ChatID, &m.Role, &m.Content, &m.Token, &m.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan message: %w", err)
		}
		msgs = append(msgs, m)
	}
	return msgs, nil
}
