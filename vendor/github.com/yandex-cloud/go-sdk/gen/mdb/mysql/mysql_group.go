// Code generated by sdkgen. DO NOT EDIT.

package mysql

import (
	"context"

	"google.golang.org/grpc"
)

// MySQL provides access to "mysql" component of Yandex.Cloud
type MySQL struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewMySQL creates instance of MySQL
func NewMySQL(g func(ctx context.Context) (*grpc.ClientConn, error)) *MySQL {
	return &MySQL{g}
}

// Backup gets BackupService client
func (m *MySQL) Backup() *BackupServiceClient {
	return &BackupServiceClient{getConn: m.getConn}
}

// Cluster gets ClusterService client
func (m *MySQL) Cluster() *ClusterServiceClient {
	return &ClusterServiceClient{getConn: m.getConn}
}

// Database gets DatabaseService client
func (m *MySQL) Database() *DatabaseServiceClient {
	return &DatabaseServiceClient{getConn: m.getConn}
}

// ResourcePreset gets ResourcePresetService client
func (m *MySQL) ResourcePreset() *ResourcePresetServiceClient {
	return &ResourcePresetServiceClient{getConn: m.getConn}
}

// User gets UserService client
func (m *MySQL) User() *UserServiceClient {
	return &UserServiceClient{getConn: m.getConn}
}
