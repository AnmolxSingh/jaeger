// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package cassandra

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/configtls"
	"go.uber.org/zap"

	viperize "github.com/jaegertracing/jaeger/internal/config"
	"github.com/jaegertracing/jaeger/internal/metrics"
	"github.com/jaegertracing/jaeger/internal/storage/cassandra/config"
	"github.com/jaegertracing/jaeger/internal/storage/cassandra/mocks"
	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestCassandraFactory(t *testing.T) {
	logger, _ := testutils.NewLogger()

	tests := []struct {
		name      string
		factoryFn func() *Factory
		namespace string
	}{
		{
			name:      "CassandraFactory",
			factoryFn: NewFactory,
			namespace: primaryStorageNamespace,
		},
		{
			name:      "CassandraArchiveFactory",
			factoryFn: NewArchiveFactory,
			namespace: archiveStorageNamespace,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := test.factoryFn()
			require.Equal(t, test.namespace, f.Options.namespace)
			v, command := viperize.Viperize(f.AddFlags)
			command.ParseFlags([]string{})
			f.InitFromViper(v, zap.NewNop())

			MockSession(f, nil, errors.New("made-up primary error"))
			require.EqualError(t, f.Initialize(metrics.NullFactory, zap.NewNop()), "made-up primary error")

			var (
				session = &mocks.Session{}
				query   = &mocks.Query{}
			)
			session.On("Query", mock.AnythingOfType("string"), mock.Anything).Return(query)
			session.On("Close").Return()
			query.On("Exec").Return(nil)

			MockSession(f, session, nil)
			require.NoError(t, f.Initialize(metrics.NullFactory, logger))

			_, err := f.CreateSpanReader()
			require.NoError(t, err)

			_, err = f.CreateSpanWriter()
			require.NoError(t, err)

			_, err = f.CreateDependencyReader()
			require.NoError(t, err)

			MockSession(f, session, nil)
			require.NoError(t, f.Initialize(metrics.NullFactory, zap.NewNop()))

			_, err = f.CreateLock()
			require.NoError(t, err)

			_, err = f.CreateSamplingStore(0)
			require.NoError(t, err)

			require.NoError(t, f.Close())
		})
	}
}

func TestCreateSpanReaderError(t *testing.T) {
	session := &mocks.Session{}
	query := &mocks.Query{}
	session.On("Query",
		mock.AnythingOfType("string"),
		mock.Anything).Return(query)
	session.On("Query",
		mock.AnythingOfType("string"),
		mock.Anything).Return(query)
	query.On("Exec").Return(errors.New("table does not exist"))
	f := NewFactory()
	MockSession(f, session, nil)
	require.NoError(t, f.Initialize(metrics.NullFactory, zap.NewNop()))
	r, err := f.CreateSpanReader()
	require.Error(t, err)
	require.Nil(t, r)
}

func TestExclusiveWhitelistBlacklist(t *testing.T) {
	f := NewFactory()
	v, command := viperize.Viperize(f.AddFlags)
	command.ParseFlags([]string{
		"--cassandra.index.tag-whitelist=a,b,c",
		"--cassandra.index.tag-blacklist=a,b,c",
	})
	f.InitFromViper(v, zap.NewNop())

	var (
		session = &mocks.Session{}
		query   = &mocks.Query{}
	)
	session.On("Query", mock.AnythingOfType("string"), mock.Anything).Return(query)
	query.On("Exec").Return(nil)
	MockSession(f, session, nil)

	_, err := f.CreateSpanWriter()
	require.EqualError(t, err, "only one of TagIndexBlacklist and TagIndexWhitelist can be specified")

	MockSession(f, session, nil)
	require.NoError(t, f.Initialize(metrics.NullFactory, zap.NewNop()))
}

func TestWriterOptions(t *testing.T) {
	opts := NewOptions("cassandra")
	v, command := viperize.Viperize(opts.AddFlags)
	command.ParseFlags([]string{"--cassandra.index.tag-whitelist=a,b,c"})
	opts.InitFromViper(v)

	options, _ := writerOptions(opts)
	assert.Len(t, options, 1)

	opts = NewOptions("cassandra")
	v, command = viperize.Viperize(opts.AddFlags)
	command.ParseFlags([]string{"--cassandra.index.tag-blacklist=a,b,c"})
	opts.InitFromViper(v)

	options, _ = writerOptions(opts)
	assert.Len(t, options, 1)

	opts = NewOptions("cassandra")
	v, command = viperize.Viperize(opts.AddFlags)
	command.ParseFlags([]string{"--cassandra.index.tags=false"})
	opts.InitFromViper(v)

	options, _ = writerOptions(opts)
	assert.Len(t, options, 1)

	opts = NewOptions("cassandra")
	v, command = viperize.Viperize(opts.AddFlags)
	command.ParseFlags([]string{"--cassandra.index.tags=false", "--cassandra.index.tag-blacklist=a,b,c"})
	opts.InitFromViper(v)

	options, _ = writerOptions(opts)
	assert.Len(t, options, 1)

	opts = NewOptions("cassandra")
	v, command = viperize.Viperize(opts.AddFlags)
	command.ParseFlags([]string{""})
	opts.InitFromViper(v)

	options, _ = writerOptions(opts)
	assert.Empty(t, options)
}

func TestConfigureFromOptions(t *testing.T) {
	f := NewFactory()
	o := NewOptions("foo")
	f.ConfigureFromOptions(o)
	assert.Equal(t, o, f.Options)
	assert.Equal(t, o.GetConfig(), f.config)
}

func TestFactory_Purge(t *testing.T) {
	f := NewFactory()
	var (
		session = &mocks.Session{}
		query   = &mocks.Query{}
	)
	session.On("Query", mock.AnythingOfType("string"), mock.Anything).Return(query)
	query.On("Exec").Return(nil)
	f.session = session

	err := f.Purge(context.Background())
	require.NoError(t, err)

	session.AssertCalled(t, "Query", mock.AnythingOfType("string"), mock.Anything)
	query.AssertCalled(t, "Exec")
}

func TestNewSessionErrors(t *testing.T) {
	t.Run("NewCluster error", func(t *testing.T) {
		cfg := &config.Configuration{
			Connection: config.Connection{
				TLS: configtls.ClientConfig{
					Config: configtls.Config{
						CAFile: "foobar",
					},
				},
			},
		}
		_, err := NewSession(cfg)
		require.ErrorContains(t, err, "failed to load TLS config")
	})
	t.Run("CreateSession error", func(t *testing.T) {
		cfg := &config.Configuration{}
		_, err := NewSession(cfg)
		require.ErrorContains(t, err, "no hosts provided")
	})
	t.Run("CreateSession error with schema", func(t *testing.T) {
		cfg := &config.Configuration{
			Schema: config.Schema{
				CreateSchema: true,
			},
		}
		_, err := NewSession(cfg)
		require.ErrorContains(t, err, "no hosts provided")
	})
}

func TestInheritSettingsFrom(t *testing.T) {
	primaryFactory := NewFactory()
	primaryFactory.config.Schema.Keyspace = "foo"
	primaryFactory.config.Query.MaxRetryAttempts = 99

	archiveFactory := NewArchiveFactory()
	archiveFactory.config.Schema.Keyspace = "bar"

	archiveFactory.InheritSettingsFrom(primaryFactory)

	require.Equal(t, "bar", archiveFactory.config.Schema.Keyspace)
	require.Equal(t, 99, archiveFactory.config.Query.MaxRetryAttempts)
}

func TestIsArchiveCapable(t *testing.T) {
	tests := []struct {
		name      string
		namespace string
		enabled   bool
		expected  bool
	}{
		{
			name:      "archive capable",
			namespace: "cassandra-archive",
			enabled:   true,
			expected:  true,
		},
		{
			name:      "not capable",
			namespace: "cassandra-archive",
			enabled:   false,
			expected:  false,
		},
		{
			name:      "capable + wrong namespace",
			namespace: "cassandra",
			enabled:   true,
			expected:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			factory := &Factory{
				Options: &Options{
					NamespaceConfig: NamespaceConfig{
						namespace: test.namespace,
						Enabled:   test.enabled,
					},
				},
			}
			result := factory.IsArchiveCapable()
			require.Equal(t, test.expected, result)
		})
	}
}
