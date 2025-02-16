package mongo

import (
	"context"

	"github.com/rigdev/rig-go-api/api/v1/database"
	"github.com/rigdev/rig/pkg/auth"
	"github.com/rigdev/rig/pkg/uuid"
	"github.com/rigdev/rig/internal/repository/database/mongo/schema"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoRepository) Get(ctx context.Context, databaseID uuid.UUID) (*database.Database, error) {
	var d schema.Database
	projectId, err := auth.GetProjectID(ctx)
	if err != nil {
		return nil, err
	}
	if err := m.DatabaseCollection.FindOne(ctx, bson.M{"database_id": databaseID, "project_id": projectId}).Decode(&d); err != nil {
		return nil, err
	}
	return d.ToProto()
}
